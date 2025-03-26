package handler

import (
	"net/http"
	"github.com/zeromicro/go-zero/rest/httpx"
	"sensor/internal/logic"
	"sensor/internal/svc"
)

func ListSensorsHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		l := logic.NewListSensorsLogic(r.Context(), svcCtx)
		resp, err := l.ListSensors()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
