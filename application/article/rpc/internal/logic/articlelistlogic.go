package logic

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/mr"
	"github.com/zeromicro/go-zero/core/threading"
	"inquora/application/article/rpc/internal/code"
	"inquora/application/article/rpc/internal/model"
	"inquora/application/article/rpc/internal/types"
	"strconv"
	"time"

	"inquora/application/article/rpc/internal/svc"
	"inquora/application/article/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

const (
	prefixArticles = "biz#articles#%d#%d"
	articlesExpire = 3600 * 24 * 2
)

type ArticleListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewArticleListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ArticleListLogic {
	return &ArticleListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ArticleListLogic) ArticleList(in *pb.ArticleListRequest) (*pb.ArticleListResponse, error) {
	if in.SortType != types.SortPublishTime && in.SortType != types.SortLikeCount {
		return nil, code.SortTypeInvalid
	}
	if in.UserId <= 0 {
		return nil, code.UserIdInvalid
	}
	if in.PageSize <= 0 || in.PageSize > types.DefaultLimit {
		in.PageSize = types.DefaultPageSize
	}
	if in.Cursor <= 0 {
		if in.SortType == types.SortPublishTime {
			in.Cursor = time.Now().Unix()
		} else {
			in.Cursor = types.DefaultSortLikeCursor
		}
	}

	var (
		sortField       string
		sortLikeNum     int64
		sortPublishTime string
	)
	if in.SortType == types.SortLikeCount {
		sortField = "like_num"
		sortLikeNum = in.Cursor
	} else {
		sortField = "publish_time"
		sortPublishTime = time.Unix(in.Cursor, 0).Format("2006-01-02 15:04:05")
	}

	var articles []*model.Article
	var err error
	var isCache, isEnd bool
	articleIds, err := l.getUserArticleIdsFromCache(l.ctx, in.UserId, in.Cursor, in.PageSize, in.SortType)
	if err != nil {
		l.Logger.Errorf("get user article id from cache err: %v", err)
	}
	if len(articleIds) > 0 {
		isCache = true
		if articleIds[len(articleIds)-1] == -1 {
			isEnd = true
		}

		articles, err = l.getArticleListByIds(l.ctx, articleIds)
		if err != nil {
			l.Logger.Errorf("get article list by ids err: %v", err)
			return nil, err
		}
	} else {
		isEnd = true
		v, err, _ := l.svcCtx.SingleFlightGroup.Do(fmt.Sprintf("UserArticles:%d:%d", in.UserId, in.SortType), func() (interface{}, error) {
			return l.svcCtx.ArticleModel.GetUserArticles(l.ctx, in.UserId, types.ArticleStatusVisible, sortLikeNum, sortPublishTime, sortField, types.DefaultLimit)
		})
		if err != nil {
			l.Logger.Errorf("get user articles err: %v", err)
			return nil, err
		}
		if v == nil {
			return &pb.ArticleListResponse{IsEnd: isEnd}, nil
		}
		articles = v.([]*model.Article)
		if len(articles) >= int(in.PageSize) {
			isEnd = false
			articles = articles[:in.PageSize]
		}
	}

	var curPage []*pb.ArticleItem
	for _, article := range articles {
		curPage = append(curPage, &pb.ArticleItem{
			Id:           article.Id,
			Title:        article.Title,
			Content:      article.Content,
			LikeCount:    article.LikeNum,
			CommentCount: article.CommentNum,
			PublishTime:  article.PublishTime.Unix(),
		})
	}

	var lastId, cursor int64
	if len(curPage) > 0 {
		pageLast := curPage[len(curPage)-1]
		lastId = pageLast.Id
		cursor = pageLast.PublishTime
		if in.SortType == types.SortLikeCount {
			cursor = pageLast.LikeCount
		}
		for k, article := range curPage {
			if in.SortType == types.SortPublishTime {
				if article.PublishTime == in.Cursor && article.Id == in.ArticleId {
					curPage = curPage[k:]
					break
				}
			} else {
				if article.LikeCount == in.Cursor && article.Id == in.ArticleId {
					curPage = curPage[k:]
					break
				}
			}
		}
	}

	ret := &pb.ArticleListResponse{
		IsEnd:     isEnd,
		Cursor:    cursor,
		ArticleId: lastId,
		Articles:  curPage,
	}

	if !isCache {
		threading.GoSafe(func() {
			if len(articles) < types.DefaultLimit && len(articles) > 0 {
				articles = append(articles, &model.Article{Id: -1})
			}
			err = l.addCacheArticles(context.Background(), articles, in.UserId, in.SortType)
			if err != nil {
				logx.Errorf("addCacheArticles error: %v", err)
			}
		})
	}

	return ret, nil
}

func (l *ArticleListLogic) getUserArticleIdsFromCache(ctx context.Context, uid, cursor, ps int64, sortType int32) ([]int64, error) {
	key := fmt.Sprintf(prefixArticles, uid, sortType)
	exist, err := l.svcCtx.BizRedis.ExistsCtx(ctx, key)
	if err != nil {
		l.Logger.Errorf("ExistsCtx key: %s error: %v", key, err)
		return nil, err
	}
	if !exist {
		return nil, nil
	}

	// 热点数据续期
	if err = l.svcCtx.BizRedis.ExpireCtx(ctx, key, articlesExpire); err != nil {
		l.Logger.Errorf("ExpireCtx key: %s error: %v", key, err)
	}

	pairs, err := l.svcCtx.BizRedis.ZrevrangebyscoreWithScoresAndLimitCtx(ctx, key, 0, cursor, 0, int(ps))
	if err != nil {
		l.Logger.Errorf("ZrevrangebyscoreWithScoresAndLimit key: %s error: %v", key, err)
		return nil, err
	}

	var ids []int64
	for _, pair := range pairs {
		id, err := strconv.ParseInt(pair.Key, 10, 64)
		if err != nil {
			l.Logger.Errorf("strconv.ParseInt key: %s error: %v", pair.Key, err)
			return nil, err
		}
		ids = append(ids, id)
	}

	return ids, nil
}

func (l *ArticleListLogic) getArticleListByIds(ctx context.Context, ids []int64) ([]*model.Article, error) {
	return mr.MapReduce[int64, *model.Article, []*model.Article](
		func(source chan<- int64) {
			for _, id := range ids {
				if id == -1 {
					continue
				}
				source <- id
			}
		},
		func(id int64, writer mr.Writer[*model.Article], cancel func(error)) {
			article, err := l.svcCtx.ArticleModel.FindOne(ctx, id)
			if err != nil {
				cancel(err)
				return
			}
			writer.Write(article)
		},
		func(pipe <-chan *model.Article, writer mr.Writer[[]*model.Article], cancel func(error)) {
			articles := make([]*model.Article, 0)
			for article := range pipe {
				articles = append(articles, article)
			}
			writer.Write(articles)
		},
	)
}

func (l *ArticleListLogic) addCacheArticles(ctx context.Context, articles []*model.Article, userId int64, sortType int32) error {
	if len(articles) == 0 {
		return nil
	}
	key := fmt.Sprintf(prefixArticles, userId, sortType)
	for _, article := range articles {
		var score int64
		if sortType == types.SortLikeCount {
			score = article.LikeNum
		} else if sortType == types.SortPublishTime && article.Id != -1 {
			score = article.PublishTime.Local().Unix()
		}
		_, err := l.svcCtx.BizRedis.ZaddCtx(ctx, key, score, strconv.Itoa(int(article.Id)))
		if err != nil {
			return err
		}
	}

	return l.svcCtx.BizRedis.ExpireCtx(ctx, key, articlesExpire)
}
