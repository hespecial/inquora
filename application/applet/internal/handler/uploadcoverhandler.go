package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"inquora/application/applet/internal/logic"
	"inquora/application/applet/internal/svc"
)

func UploadCoverHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewUploadCoverLogic(r.Context(), svcCtx)
		resp, err := l.UploadCover(r)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
