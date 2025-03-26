package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"sensor/internal/logic"
	"sensor/internal/svc"
)

func DeleteSensorHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewDeleteSensorLogic(r.Context(), svcCtx)
		err := l.DeleteSensor()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.Ok(w)
		}
	}
}
