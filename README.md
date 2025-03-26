# IotSystem
 
系统架构图：
![image](assets\架构.png)

按照以下步骤，运行项目：

## 1. 运行mqtt后端服务

### 使用docker运行mqtt后端服务

1. 下载docker-desktop，或者从wsl运行docker

2. 目前国内无法拉取mqtt镜像，需要科学上网，或者按照[这个文档](https://www.cnblogs.com/guangdelw/p/18357276)的方法下载emqx的docker镜像

3. 导入压缩包形式的docker镜像：

```
docker load -i xxx.tar.gz
```

4. 运行emqx镜像，可参照[这个文档](https://docs.emqx.com/zh/emqx/latest/getting-started/getting-started.html)

5. 可通过MQTTX客户端连接mqtt服务端，测试是否正常

### 2. 运行python设备模拟器

进入`mqtt`目录，运行（要先安装python的依赖包）：
```
python python-device.py
```

**此时如果开启了MQTTX客户端，则可以看到模拟设备发送的数据**

## 3. 运行Go后端服务

进入`sensor`目录，运行：

```
go mod tidy

go run sensor.go
```

该文件夹中，`.db`，文件为数据库文件，可用开源的数据库管理软件DBeaver查看。`device.csv`文件为随机生成的数据库原始数据，`.db`数据库的原始数据文件由它导入。

## 4. 运行Vue大屏

进入`Screen-Vue3`目录，运行：
```
npm install
npm run dev
```

