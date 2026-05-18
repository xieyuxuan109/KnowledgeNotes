### 编写另一个微服务api文件
```
syntax = "v1"
info (
	title:   "mall-order"
	desc:    "go-zero学习项目"
	author:  "xieyuxuan"
	email:   "xieyuxuan.com"
	version: "1.0"
)
type searchRequest {
	OrderID string `form:"orderID"`
}
type searchResponse {
	OrderID  string `json:"orderID"`
	Status   int    `json:"status"`
	Username string `json:"username"`
}
@server (
	prefix: api
)
service order-api {
	@handler SearchHandler
	get /order/search (searchRequest) returns (searchResponse)
}
```
### 生成代码
```
goctl api go -api user.api -dir . -style=goZero
```
### 修改api/etc的yaml配置 
```
Name: order-api
Host: 0.0.0.0
Port: 8888
UserRPC:
  Etcd:
    Hosts:
      - 10.17.220.221:2379
    # 确保和注册时一样
    Key: user.rpc
```
### 修改api/internal/config配置
```
type Config struct {
	rest.RestConf
	UserRPC zrpc.RpcClientConf //因为可能连接多个客户端所以最好起一个名字
}
```
### 修改api/internal/svc配置
```
//添加在service/order/api/go.mod 一种不太好的解决方法
replace user-rpc => ../../../user/rpc
```
```
type ServiceContext struct {
	Config  config.Config
	UserRPC userclient.User
}
func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:  c,
		UserRPC: userclient.NewUser(zrpc.MustNewClient(c.UserRPC)),
	}
}
```
### 编写逻辑代码
```
func (l *SearchLogic) Search(req *types.SearchRequest) (resp *types.SearchResponse, err error) {
	// todo: add your logic here and delete this line
	UserResp, err := l.svcCtx.UserRPC.GetUser(l.ctx, &userclient.GetUserReq{UserID: 111})
	if err != nil {
		logx.Errorw("UserPRC.GetUser failed", logx.Field("err=", err))
		return nil, err
	}
	return &types.SearchResponse{
		OrderID:  "212121212121212",
		Status:   1000,
		Username: UserResp.GetUsername(), //rpc调用rpc得来的
	}, nil
}
```