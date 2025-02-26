package logic

import (
	"context"
	"encoding/json"
	"inquora/application/applet/internal/code"
	"inquora/application/article/rpc/pb"

	"inquora/application/applet/internal/svc"
	"inquora/application/applet/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

const (
	minContentLen = 80
)

type PublishLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPublishLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PublishLogic {
	return &PublishLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PublishLogic) Publish(req *types.PublishRequest) (resp *types.PublishResponse, err error) {
	if len(req.Title) == 0 {
		return nil, code.ArtitleTitleEmpty
	}
	if len(req.Content) < minContentLen {
		return nil, code.ArticleContentTooFewWords
	}
	if len(req.Cover) == 0 {
		return nil, code.ArticleCoverEmpty
	}

	userId, err := l.ctx.Value(types.UserIdKey).(json.Number).Int64()
	if err != nil {
		return nil, err
	}

	res, err := l.svcCtx.ArticleRpc.Publish(l.ctx, &pb.PublishRequest{
		UserId:      userId,
		Title:       req.Title,
		Content:     req.Content,
		Description: req.Description,
		Cover:       req.Cover,
	})
	if err != nil {
		l.Logger.Errorf("ArticleRpc.Publish err:%v", err)
		return nil, err
	}

	return &types.PublishResponse{ArticleId: res.ArticleId}, nil
}
