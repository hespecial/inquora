package logic

import (
	"context"
	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/core/service"

	"github.com/zeromicro/go-zero/core/logx"
	"inquora/application/like/mq/internal/svc"
)

type ThumbUpLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewThumbUpLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ThumbUpLogic {
	return &ThumbUpLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ThumbUpLogic) Consume(_ context.Context, key, val string) error {
	l.Logger.Infof("get key: %s val: %s", key, val)
	return nil
}

func Consumers(ctx context.Context, svcCtx *svc.ServiceContext) []service.Service {
	return []service.Service{
		kq.MustNewQueue(svcCtx.Config.KqConsumerConf, NewThumbUpLogic(ctx, svcCtx)),
	}
}
