### 编写pb文件
```
syntax="proto3";
package user;
option go_package = "./user";
message getUserReq{
    int64 userID=1;
}
message getUserResp{
    int64 userID=1;
    string username=2;
    int64 gender=3;
}
service user{
    rpc getUser(getUserReq) returns (getUserResp);
}
```
### 生成命令
```
# 单个 rpc 服务生成示例指令
$ goctl rpc protoc greet.proto --go_out=./pb --go-grpc_out=./pb --zrpc_out=. --client=true
# 多个 rpc 服务生成示例指令
$ goctl rpc protoc greet.proto --go_out=./pb --go-grpc_out=./pb --zrpc_out=. --client=true -m
```
### 修改api/etc的yaml配置
```
Name: user.rpc//服务名称
Mode: dev
ListenOn: 0.0.0.0:8080
Etcd:
  Hosts:
  - 10.17.220.221:2379
  Key: user.rpc
Mysql:
  DataSource: root:root@tcp(10.17.220.221:3306)/xieyuxuan?parseTime=true&loc=Local
CacheRedis:
  - Host: 10.17.220.221:6379
```
### 修改api/internal/config配置
```
type Config struct {
	zrpc.RpcServerConf
	Mysql struct { //数据库配置 使用匿名结构体 Mysql是字段名不是结构体名字
		DataSource string //mysql连接地址
	}
	CacheRedis cache.CacheConf
}
```
### 修改api/internal/svc
```
type ServiceContext struct {
	Config    config.Config
	UserModel model.UserModel
}
func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config:    c,
		UserModel: model.NewUserModel(conn, c.CacheRedis),
	}
}
```
### 编写grpc逻辑
```
func (l *GetUserLogic) GetUser(in *user.GetUserReq) (*user.GetUserResp, error) {
	// todo: add your logic here and delete this line
	// 根据userID查询数据库返回用户信息
	u, err := l.svcCtx.UserModel.FindOneByUserId(l.ctx, in.UserID)
	// 两种错 一种连接数据库错误 二是查询不到用户
	if err != nil && err != sqlx.ErrNotFound {
		logx.Errorw("FindOneByUserId failed", logx.Field("err", err))
		return nil, errors.New("内部错误")
	}
	if err != nil {
		return nil, errors.New("查询失败")
	}
	return &user.GetUserResp{
		UserID:   u.UserId,
		Username: u.Username,
		Gender:   u.Gender,
	}, nil
}
```
### 启动etcd命令
```
//默认启动需要https才能访问
docker run -d \
  --name my-etcd \
  --network host \
  -e ALLOW_NONE_AUTHENTICATION=yes \
  nexus.dennis.com.cn:8082/google_containers/etcd:3.5.24-0 \
  /usr/local/bin/etcd \
  --listen-client-urls http://0.0.0.0:2379,http://0.0.0.0:2379 \
  --advertise-client-urls http://127.0.0.1:2379
```
### 安装grpc测试工具 只有dev或者test模式才能进行测试
```bash
//grpc ui
go install github.com/fullstorydev/grpcui/cmd/grpcui@latest
//确保$GOPATH/bin添加到环境变量里面
```
```bash
//使用 端口换成自己都端口
grpcui -plaintext localhost:8080
```