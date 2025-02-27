package model

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ ArticleModel = (*customArticleModel)(nil)

type (
	// ArticleModel is an interface to be customized, add more methods here,
	// and implement the added methods in customArticleModel.
	ArticleModel interface {
		articleModel
		GetUserArticles(ctx context.Context, userId int64, status int, likeNum int64, pubTime, sortField string, limit int) ([]*Article, error)
	}

	customArticleModel struct {
		*defaultArticleModel
	}
)

// NewArticleModel returns a model for the database table.
func NewArticleModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) ArticleModel {
	return &customArticleModel{
		defaultArticleModel: newArticleModel(conn, c, opts...),
	}
}

func (m *customArticleModel) GetUserArticles(ctx context.Context, userId int64, status int, likeNum int64, pubTime, sortField string, limit int) ([]*Article, error) {
	var articles []*Article
	var fieldVal any = pubTime
	if sortField == "like_num" {
		fieldVal = likeNum
	}
	query := fmt.Sprintf("select %s from %s where `author_id` = %d and `status` = %d and `%s` <= ? order by `%s` desc limit %d",
		articleRows, m.table, userId, status, sortField, sortField, limit)
	if err := m.QueryRowsNoCacheCtx(ctx, &articles, query, fieldVal); err != nil {
		return nil, err
	}
	return articles, nil
}
