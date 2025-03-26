package logic

import (
	"context"
	"sensor/model"
	"sensor/internal/svc"
	"sensor/internal/types"
	"strconv"
	"strings"
	"github.com/zeromicro/go-zero/core/logx"
)

type CreateSensorLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateSensorLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateSensorLogic {
	return &CreateSensorLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateSensorLogic) CreateSensor(req *types.SensorDataRequest) (resp *types.SensorDataResponse, err error) {
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
		// 假设 SensorDataRequest 和 SensorData 的字段一一对应
		Num:          number,
		Name:        req.Name,
		Value:       req.Value,
	}
	// 调用模型层的方法来添加传感器
	num, err := model.AddSensor(sensor)
	if err != nil {
		return nil, err
	}

	// 设置返回的传感器ID
	sensor.Num = int(num)

	// 构建响应
	resp = &types.SensorDataResponse{
		Num:          sensor.Num,
		Name:        sensor.Name,
		Value:		 sensor.Value,
	}

	return resp, nil
}
