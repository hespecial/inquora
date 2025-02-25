package logic

import (
	"context"

	"inquora/application/like/rpc/internal/svc"
	"inquora/application/like/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type IsThumbUpLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewIsThumbUpLogic(ctx context.Context, svcCtx *svc.ServiceContext) *IsThumbUpLogic {
	return &IsThumbUpLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *IsThumbUpLogic) IsThumbUp(in *pb.IsThumbUpRequest) (*pb.IsThumbUpResponse, error) {
	// todo: add your logic here and delete this line

	return &pb.IsThumbUpResponse{}, nil
}
