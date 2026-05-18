### 运行命令
```
docker run --rm --name jaeger \
  -e COLLECTOR_ZIPKIN_HOST_PORT=:9411 \
  -p 6831:6831/udp \
  -p 6832:6832/udp \
  -p 5778:5778 \
  -p 16686:16686 \
  -p 4317:4317 \
  -p 4318:4318 \
  -p 14250:14250 \
  -p 14268:14268 \
  -p 14269:14269 \
  -p 9411:9411 \
  jaegertracing/all-in-one:1.55
```
### 端口说明
| 端口 | 协议 | 组件 | 功能 |
| :--- | :--- | :--- | :--- |
| 6831 | UDP | agent | 通过 Thrift-compact 协议接收 jaeger.thrift（大多数 SDK 使用） |
| 6832 | UDP | agent | 通过 Thrift-binary 协议接收 jaeger.thrift（Node.js SDK 使用） |
| 5775 | UDP | agent | (已弃用) 通过 compact Thrift 协议接收 zipkin.thrift（仅限旧版客户端） |
| 5778 | HTTP | agent | 提供配置服务（如采样策略等） |
| 16686| HTTP | query | 提供前端 UI 界面服务 |
| 4317 | HTTP | collector | 通过 gRPC 接收 OpenTelemetry 协议 (OTLP) 数据 |
| 4318 | HTTP | collector | 通过 HTTP 接收 OpenTelemetry 协议 (OTLP) 数据 |
| 14268| HTTP | collector | 直接接收来自客户端的 jaeger.thrift 数据 |
| 14250| HTTP | collector | 接收 model.proto 数据 |
| 9411 | HTTP | collector | 兼容 Zipkin 的接入端点 (可选) |
```
容器启动后，请使用浏览器打开 http://localhost:16686 即可访问 Jaeger UI 界面
```
