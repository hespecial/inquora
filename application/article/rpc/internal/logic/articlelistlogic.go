package logic

import (
	"context"
	"inquora/application/article/rpc/internal/code"
	"inquora/application/article/rpc/internal/types"
	"time"

	"inquora/application/article/rpc/internal/svc"
	"inquora/application/article/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
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
	if in.PageSize <= 0 {
		in.PageSize = types.DefaultPageSize
	}
	if in.Cursor <= 0 {
		if in.SortType == types.SortPublishTime {
			in.Cursor = time.Now().Unix()
		} else {
			in.Cursor = types.DefaultSortLikeCursor
		}
	}

	return &pb.ArticleListResponse{}, nil
}
