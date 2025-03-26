package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"sensor/internal/logic"
	"sensor/internal/svc"
	"sensor/internal/types"
)

func CreateSensorHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SensorDataRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewCreateSensorLogic(r.Context(), svcCtx)
		resp, err := l.CreateSensor(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
