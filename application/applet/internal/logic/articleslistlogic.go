package logic

import (
	"context"

	"inquora/application/applet/internal/svc"
	"inquora/application/applet/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ArticlesListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewArticlesListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ArticlesListLogic {
	return &ArticlesListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ArticlesListLogic) ArticlesList(req *types.ArticleListRequest) (resp *types.ArticleListResponse, err error) {

	return
}
