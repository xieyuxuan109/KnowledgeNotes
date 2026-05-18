## metadata定义
```
Metadata是指在RPC调用中，与业务数据分离的辅助信息，通常用于服务治理、链路追踪、认证授权等场景
```
### Metadata vs 请求参数
|对比维度|	Metadata|	请求参数|
|---|---|---|
|作用|	描述"如何调用"|	描述"调什么"|
|业务逻辑|	不参与核心业务逻辑|	业务逻辑的核心输入|
|典型数据|	Token、TraceID、UserID|	用户名、订单ID、查询条件|
|传输位置|	gRPC Header 中|	Protobuf Body 中|
|类型安全|	字符串键值对|	Protobuf 强类型|
### Metadata存哪些数据
```
应该放 Metadata 的数据：
├── 认证信息：Token、API Key、SessionID
├── 链路追踪：TraceID、SpanID、ParentID
├── 用户上下文：UserID、Role、TenantID
├── 客户端信息：ClientIP、User-Agent、AppVersion
├── 灰度标识：Canary、Version
├── 国际化：Language、Timezone
└── 请求控制：Timeout、RetryCount

不应该放 Metadata 的数据：
├── 业务核心字段：商品ID、订单金额、用户姓名
├── 复杂数据结构：整个用户对象、列表数据
└── 敏感信息：密码、信用卡号（建议走专门的安全通道）
```
### 示例
```
// 正确的 Metadata 使用
ctx := context.Background()
md := metadata.Pairs(
    "authorization", "Bearer xxxxxxxx",  // 认证
    "trace-id", "abc123def456",           // 链路追踪
    "user-id", "10086",                   // 用户标识（非业务参数）
    "x-request-id", "req-001",            // 请求追踪
)
ctx = metadata.NewOutgoingContext(ctx, md)

// 错误：业务参数不应该放 Metadata
// md := metadata.Pairs("order-id", "12345")  // 这是业务参数，应该放 Protobuf
```
## 拦截器
### 添加client拦截器
### 添加server拦截器