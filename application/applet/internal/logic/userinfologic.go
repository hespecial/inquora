package logic

import (
	"context"
	"inquora/application/user/rpc/user"
	"inquora/pkg/xcode"

	"inquora/application/applet/internal/svc"
	"inquora/application/applet/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserInfoLogic {
	return &UserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserInfoLogic) UserInfo() (resp *types.UserInfoResponse, err error) {
	userId, ok := l.ctx.Value(types.UserIdKey).(int64)
	if !ok {
		logx.Errorf("assert user id fail")
		return nil, xcode.AccessDenied
	}
	u, err := l.svcCtx.UserRpc.FindById(l.ctx, &user.FindByIdRequest{UserId: userId})
	if err != nil {
		logx.Errorf("FindById userId: %d error: %v", userId, err)
		return nil, err
	}
	return &types.UserInfoResponse{
		UserId:   userId,
		Username: u.Username,
		Avatar:   u.Avatar,
	}, nil
}
