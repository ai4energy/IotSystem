package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"github.com/eclipse/paho.mqtt.golang"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
	"sensor/internal/config"
	"sensor/internal/logic"
	"sensor/internal/handler"
	"sensor/internal/svc"
	"sensor/model"
	"sensor/internal/types"
)

var configFile = flag.String("f", "etc/sensor-api.yaml", "the config file")

func init() {
	model.InitDB()
}

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	// 初始化服务上下文
	ctx := svc.NewServiceContext(c)

	// 注册 HTTP 处理器
	handler.RegisterHandlers(server, ctx)

	// MQTT 配置
	mqttOpts := mqtt.NewClientOptions().AddBroker("tcp://127.0.0.1:1883")
	mqttOpts.SetClientID("gozero_mqtt_client")

	client := mqtt.NewClient(mqttOpts)
	token := client.Connect()
	token.Wait()
	if token.Error() != nil {
		log.Fatalf("Error connecting to MQTT broker: %v", token.Error())
	}

	token = client.Subscribe("device", 0, func(client mqtt.Client, msg mqtt.Message) {
		fmt.Printf("Received message: %s from topic: %s\n", msg.Payload(), msg.Topic())
		var data SensorData
		err := json.Unmarshal(msg.Payload(), &data)
		if err != nil {
			log.Printf("Failed to parse MQTT message: %v", err)
			return
		}

		// 调用逻辑层方法来更新传感器数据
		updateSensorLogic := logic.NewUpdateSensorLogic(context.Background(), ctx)
		req := &types.SensorDataRequest{
			Name:  data.Name,
			Value: data.Value,
		}
		_, err = updateSensorLogic.UpdateSensor(req)
		if err != nil {
			log.Printf("Failed to update sensor data: %v", err)
		}
	})
	token.Wait()
	if token.Error() != nil {
		log.Fatalf("Error subscribing to MQTT topic: %v", token.Error())
	}

	log.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}

type SensorData struct {
	Name  string  `json:"name"`
	Value int `json:"value"`
}