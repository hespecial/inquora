package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"inquora/application/applet/internal/logic"
	"inquora/application/applet/internal/svc"
	"inquora/application/applet/internal/types"
)

func ArticlesListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ArticleListRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewArticlesListLogic(r.Context(), svcCtx)
		resp, err := l.ArticlesList(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
