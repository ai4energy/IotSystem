import paho.mqtt.client as mqtt
import time
import random
import json

# MQTT 服务器信息
MQTT_BROKER = "127.0.0.1"  # 公共 MQTT 测试服务器
MQTT_PORT = 1883
MQTT_TOPIC = "device"

# 创建 MQTT 客户端
client = mqtt.Client()


# 连接到 MQTT 服务器
def connect_mqtt():
    try:
        print("正在连接到 MQTT 服务器...")
        client.connect(MQTT_BROKER, MQTT_PORT)
        print("成功连接到 MQTT 服务器！")
    except Exception as e:
        print(f"连接失败: {e}")
        exit(1)


# 发送消息
def send_message():
    while True:
        # 构造随机消息内容（模拟传感器数据）
        message_data = {
            "name": "x" + str(random.randint(1, 4)),
            "value": random.randint(1, 100),
        }
        message_json = json.dumps(message_data)  # 转换为 JSON 字符串

        print(f"发送消息: {message_json}")

        # 发布消息到指定主题
        client.publish(MQTT_TOPIC, message_json)

        # 等待 1 秒
        time.sleep(1)


# 主函数
if __name__ == "__main__":
    connect_mqtt()  # 连接 MQTT 服务器
    try:
        send_message()  # 开始发送消息
    except KeyboardInterrupt:
        print("程序已终止")
    finally:
        client.disconnect()  # 断开连接
