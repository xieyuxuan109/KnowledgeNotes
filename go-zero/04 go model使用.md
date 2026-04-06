### 启动mysql容器命令
```
docker run -d \
  --name my-mysql \
  -p 3306:3306 \
  -e MYSQL_ROOT_PASSWORD=root \
  --restart unless-stopped \
  mysql:8.0
```
### 启动redis容器命令
```
docker run -d --name my-redis -p 6379:6379 redis:6.2
```
## 连接方式

### ddl方式 
```
# 不带缓存
goctl model mysql ddl -src user.sql -dir ./model
# 带 Redis 缓存（生产环境推荐）
goctl model mysql ddl -src user.sql -dir ./model -cache
```
### datasource方式
```
goctl model mysql datasource -url="root:123456@tcp(10.17.220.211:3306)/xieyuxaun" -table="user" -dir="./model"
```
### 引入model层配置
```
cd service/user/model
go mod init go-zero-demo2/service/user/model
go mod edit -replace go-zero-demo2/service/user/model=../model
go mod edit -require go-zero-demo2/service/user/model@v0.0.0
```
### 引入示例
```
"go-zero-demo2/service/user/model"
```
## 如何引入mysql
### 修改api/etc的yaml配置
### 加入loc=Local后注意数据库还是存的UTC时间
```
Name: user-api
Host: 0.0.0.0
Port: 8888
Mysql:
  DataSource: root:123456@tcp(10.17.220.211:3306)/xieyuxaun?parseTime=true&loc=Local
```
### 修改api/internal/config配置
```
package config
import (
	"github.com/zeromicro/go-zero/rest"
)
type Config struct {
	rest.RestConf
	Mysql struct { //数据库配置 使用匿名结构体 Mysql是字段名不是结构体名字
		DataSource string //mysql连接地址
	}
}
```
### 修改api/internal/svc配置
```
type ServiceContext struct {
	Config config.Config
	UserModel model.UserModel //model.数据库表名+Model
}
func NewServiceContext(c config.Config) *ServiceContext {
	// UserModel是接口类型
	// *defaultUserModel实现了接口
	// 调用构造函数得到 *model.defaultUserModel
	// NewUserModel(conn sqlx.SqlConn)
	// 需要传入sqlx.SqlConn连接
	sqlxConn := sqlx.NewMysql(c.Mysql.DataSource)//获得连接
	return &ServiceContext{
		Config:    c,
		UserModel: model.NewUserModel(sqlxConn),//传入连接
	}
}
```
## 带缓存还需执行的操作
### 修改api/etc的yaml配置
```
Name: user-api
Host: 0.0.0.0
Port: 8888
Mysql:
  DataSource: root:root@tcp(10.17.220.221:3306)/xieyuxuan?parseTime=true&loc=Local
CacheRedis:
  - Host: 10.17.220.221:6379
```
### 修改api/internal/config配置
```
type Config struct {
	rest.RestConf
	Mysql struct { //数据库配置 使用匿名结构体 Mysql是字段名不是结构体名字
		DataSource string //mysql连接地址
	}
	CacheRedis cache.CacheConf
}
```
### 修改api/internal/svc配置
```
func NewServiceContext(c config.Config) *ServiceContext {
	sqlxConn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config:    c,
		UserModel: model.NewUserModel(sqlxConn, c.CacheRedis),//添加redis
	}
}
```
### 调用api接口后查看redis是否生效
```
root@jasper:/home/jasper# docker exec -it f3c5dc71664c redis-cli
127.0.0.1:6379> keys cache:user:username:*
1) "cache:user:username:xi"
```