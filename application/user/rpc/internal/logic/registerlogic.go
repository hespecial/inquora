package logic

import (
	"context"
	"inquora/application/user/rpc/internal/code"
	"inquora/application/user/rpc/internal/model"
	"time"

	"inquora/application/user/rpc/internal/svc"
	"inquora/application/user/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RegisterLogic) Register(in *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	if in.Username == "" {
		return nil, code.RegisterNameEmpty
	}

	user := model.User{
		Username:   in.Username,
		Mobile:     in.Mobile,
		CreateTime: time.Now(),
		UpdateTime: time.Now(),
	}
	ret, err := l.svcCtx.UserModel.Insert(l.ctx, &user)
	if err != nil {
		logx.Errorf("Register req: %v error: %v", in, err)
		return nil, err
	}
	user.Id, err = ret.LastInsertId()
	if err != nil {
		logx.Errorf("LastInsertId error: %v", err)
		return nil, err
	}

	return &pb.RegisterResponse{
		UserId: user.Id,
	}, nil
}
