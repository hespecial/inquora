package logic

import (
	"context"

	"inquora/application/user/rpc/internal/svc"
	"inquora/application/user/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindByMobileLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindByMobileLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindByMobileLogic {
	return &FindByMobileLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FindByMobileLogic) FindByMobile(in *pb.FindByMobileRequest) (*pb.FindByMobileResponse, error) {
	user, err := l.svcCtx.UserModel.FindOneByMobile(l.ctx, in.Mobile)
	if err != nil {
		logx.Errorf("FindById userId: %s error: %v", in.Mobile, err)
		return nil, err
	}
	return &pb.FindByMobileResponse{
		UserId:   user.Id,
		Username: user.Username,
		Mobile:   user.Mobile,
		Avatar:   user.Avatar,
	}, nil
}
