package logic

import (
	"context"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"inquora/application/applet/internal/code"
	"inquora/application/user/rpc/user"
	"inquora/pkg/jwt"
	"inquora/pkg/xcode"
	"strings"

	"inquora/application/applet/internal/svc"
	"inquora/application/applet/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginRequest) (resp *types.LoginResponse, err error) {
	req.Mobile = strings.TrimSpace(req.Mobile)
	if len(req.Mobile) == 0 {
		return nil, code.LoginMobileEmpty
	}
	req.VerificationCode = strings.TrimSpace(req.VerificationCode)
	if len(req.VerificationCode) == 0 {
		return nil, code.VerificationCodeEmpty
	}

	if err = checkVerificationCode(l.svcCtx.BizRedis, req.Mobile, req.VerificationCode); err != nil {
		return nil, err
	}
	u, err := l.svcCtx.UserRpc.FindByMobile(l.ctx, &user.FindByMobileRequest{Mobile: req.Mobile})
	if err != nil {
		logx.Errorf("FindByMobile error: %v", err)
		return nil, err
	}
	if u.UserId == 0 {
		return nil, xcode.AccessDenied
	}
	token, err := jwt.BuildTokens(jwt.TokenOptions{
		AccessSecret: l.svcCtx.Config.Auth.AccessSecret,
		AccessExpire: l.svcCtx.Config.Auth.AccessExpire,
		Fields: map[string]interface{}{
			types.UserIdKey: u.UserId,
		},
	})
	if err != nil {
		logx.Errorf("BuildTokens error: %v", err)
		return nil, err
	}

	if err = delActivationCache(req.Mobile, l.svcCtx.BizRedis); err != nil {
		logx.Errorf("delActivationCache error: %v", err)
		return nil, err
	}
	return &types.LoginResponse{
		UserId: u.UserId,
		Token: types.Token{
			AccessToken: token.AccessToken,
		},
	}, nil
}

func checkVerificationCode(rds *redis.Redis, mobile, c string) error {
	cacheCode, err := getActivationCache(mobile, rds)
	if err != nil {
		return err
	}
	if cacheCode == "" {
		return code.VerificationCodeExpire
	}
	if cacheCode != c {
		return code.VerificationCodeInvalid
	}

	return nil
}
