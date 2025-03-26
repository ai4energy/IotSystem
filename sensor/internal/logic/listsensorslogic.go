package logic

import (
	"context"
	"sensor/internal/svc"
	"sensor/internal/types"
	"sensor/model"
	"github.com/zeromicro/go-zero/core/logx"
)

type ListSensorsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListSensorsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListSensorsLogic {
	return &ListSensorsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListSensorsLogic) ListSensors() (resp []types.SensorDataResponse, err error) {
	// 调用模型层的方法来获取所有传感器数据
	sensors, err := model.GetSensors()
	if err != nil {
		l.Errorf("failed to list sensors: %v", err)
		return nil, err
	}

	// 将模型数据转换为响应数据
	for _, sensor := range sensors {
		resp = append(resp, types.SensorDataResponse{
			Num:          sensor.Num,
			Name:        sensor.Name,
			Value:       sensor.Value,
		})
	}

	return resp, nil
}