package logic

import (
	"context"
	"sensor/internal/svc"
	"sensor/internal/types"
	"sensor/model"
	"strconv"
	"strings"
	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateSensorLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateSensorLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateSensorLogic {
	return &UpdateSensorLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateSensorLogic) UpdateSensor(req *types.SensorDataRequest) (resp *types.SensorDataResponse, err error) {
	// 假设 SensorDataRequest 包含 num 字段
	// 原始字符串
	str := req.Name
	// 去掉非数字字符，提取数字部分
	numberStr := strings.TrimLeft(str, "x") // 去掉前缀 "x"
	// 将提取的数字字符串转换为整数
	number, err := strconv.Atoi(numberStr)
	if err != nil {
		return
	}

	// 将请求中的数据转换为模型数据
	sensor := model.SensorData{
		Num:    number,
		Name:  req.Name,
		Value: req.Value,
		// 其他字段...
	}

	// 调用模型层的方法来更新传感器
	err = model.UpdateSensor(req.Name, req.Value)
	if err != nil {
		return nil, err
	}

	// 构建响应
	resp = &types.SensorDataResponse{
		Num:    sensor.Num,
		Name:  sensor.Name,
		Value: sensor.Value,
		// 其他字段...
	}

	return resp, nil
}
