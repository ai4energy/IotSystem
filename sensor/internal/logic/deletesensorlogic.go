package logic

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"
	"sensor/internal/svc"
)

type DeleteSensorLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteSensorLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteSensorLogic {
	return &DeleteSensorLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteSensorLogic) DeleteSensor() error {
	// 调用模型层的方法来删除传感器

	return nil
}
