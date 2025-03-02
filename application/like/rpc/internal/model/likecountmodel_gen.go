// Code generated by goctl. DO NOT EDIT.
// versions:
//  goctl version: 1.7.6

package model

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	likeCountFieldNames          = builder.RawFieldNames(&LikeCount{})
	likeCountRows                = strings.Join(likeCountFieldNames, ",")
	likeCountRowsExpectAutoSet   = strings.Join(stringx.Remove(likeCountFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	likeCountRowsWithPlaceHolder = strings.Join(stringx.Remove(likeCountFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"

	cacheInquoraLikeLikeCountIdPrefix         = "cache:inquoraLike:likeCount:id:"
	cacheInquoraLikeLikeCountBizIdObjIdPrefix = "cache:inquoraLike:likeCount:bizId:objId:"
)

type (
	likeCountModel interface {
		Insert(ctx context.Context, data *LikeCount) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*LikeCount, error)
		FindOneByBizIdObjId(ctx context.Context, bizId string, objId int64) (*LikeCount, error)
		Update(ctx context.Context, data *LikeCount) error
		Delete(ctx context.Context, id int64) error
	}

	defaultLikeCountModel struct {
		sqlc.CachedConn
		table string
	}

	LikeCount struct {
		Id         int64     `db:"id"`          // 主键ID
		BizId      string    `db:"biz_id"`      // 业务ID
		ObjId      int64     `db:"obj_id"`      // 点赞对象id
		LikeNum    int64     `db:"like_num"`    // 点赞数
		DislikeNum int64     `db:"dislike_num"` // 点踩数
		CreateTime time.Time `db:"create_time"` // 创建时间
		UpdateTime time.Time `db:"update_time"` // 最后修改时间
	}
)

func newLikeCountModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) *defaultLikeCountModel {
	return &defaultLikeCountModel{
		CachedConn: sqlc.NewConn(conn, c, opts...),
		table:      "`like_count`",
	}
}

func (m *defaultLikeCountModel) Delete(ctx context.Context, id int64) error {
	data, err := m.FindOne(ctx, id)
	if err != nil {
		return err
	}

	inquoraLikeLikeCountBizIdObjIdKey := fmt.Sprintf("%s%v:%v", cacheInquoraLikeLikeCountBizIdObjIdPrefix, data.BizId, data.ObjId)
	inquoraLikeLikeCountIdKey := fmt.Sprintf("%s%v", cacheInquoraLikeLikeCountIdPrefix, id)
	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, inquoraLikeLikeCountBizIdObjIdKey, inquoraLikeLikeCountIdKey)
	return err
}

func (m *defaultLikeCountModel) FindOne(ctx context.Context, id int64) (*LikeCount, error) {
	inquoraLikeLikeCountIdKey := fmt.Sprintf("%s%v", cacheInquoraLikeLikeCountIdPrefix, id)
	var resp LikeCount
	err := m.QueryRowCtx(ctx, &resp, inquoraLikeLikeCountIdKey, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", likeCountRows, m.table)
		return conn.QueryRowCtx(ctx, v, query, id)
	})
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultLikeCountModel) FindOneByBizIdObjId(ctx context.Context, bizId string, objId int64) (*LikeCount, error) {
	inquoraLikeLikeCountBizIdObjIdKey := fmt.Sprintf("%s%v:%v", cacheInquoraLikeLikeCountBizIdObjIdPrefix, bizId, objId)
	var resp LikeCount
	err := m.QueryRowIndexCtx(ctx, &resp, inquoraLikeLikeCountBizIdObjIdKey, m.formatPrimary, func(ctx context.Context, conn sqlx.SqlConn, v any) (i any, e error) {
		query := fmt.Sprintf("select %s from %s where `biz_id` = ? and `obj_id` = ? limit 1", likeCountRows, m.table)
		if err := conn.QueryRowCtx(ctx, &resp, query, bizId, objId); err != nil {
			return nil, err
		}
		return resp.Id, nil
	}, m.queryPrimary)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultLikeCountModel) Insert(ctx context.Context, data *LikeCount) (sql.Result, error) {
	inquoraLikeLikeCountBizIdObjIdKey := fmt.Sprintf("%s%v:%v", cacheInquoraLikeLikeCountBizIdObjIdPrefix, data.BizId, data.ObjId)
	inquoraLikeLikeCountIdKey := fmt.Sprintf("%s%v", cacheInquoraLikeLikeCountIdPrefix, data.Id)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?)", m.table, likeCountRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.BizId, data.ObjId, data.LikeNum, data.DislikeNum)
	}, inquoraLikeLikeCountBizIdObjIdKey, inquoraLikeLikeCountIdKey)
	return ret, err
}

func (m *defaultLikeCountModel) Update(ctx context.Context, newData *LikeCount) error {
	data, err := m.FindOne(ctx, newData.Id)
	if err != nil {
		return err
	}

	inquoraLikeLikeCountBizIdObjIdKey := fmt.Sprintf("%s%v:%v", cacheInquoraLikeLikeCountBizIdObjIdPrefix, data.BizId, data.ObjId)
	inquoraLikeLikeCountIdKey := fmt.Sprintf("%s%v", cacheInquoraLikeLikeCountIdPrefix, data.Id)
	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, likeCountRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, newData.BizId, newData.ObjId, newData.LikeNum, newData.DislikeNum, newData.Id)
	}, inquoraLikeLikeCountBizIdObjIdKey, inquoraLikeLikeCountIdKey)
	return err
}

func (m *defaultLikeCountModel) formatPrimary(primary any) string {
	return fmt.Sprintf("%s%v", cacheInquoraLikeLikeCountIdPrefix, primary)
}

func (m *defaultLikeCountModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary any) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", likeCountRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultLikeCountModel) tableName() string {
	return m.table
}
