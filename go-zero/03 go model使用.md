```
goctl model mysql 指令
goctl model mysql 指令用于生成基于 MySQL 的 model 代码，支持生成带缓存和不带缓存的代码。MySQL 代码生成支持从 sql 文件，数据库连接两个来源生成代码。
```
```
datasource方式
```
```
goctl model mysql datasource -url="root:123456@tcp(10.17.220.211:3306)/xieyuxaun" -table="user" -dir="./model"
```
```
ddl方式
```
```
goctl model mysql ddl
```
```
错误信息提示 "Mysql" is not set，原因是 YAML 配置文件中的键名与 Go 结构体字段名的大小写不匹配。
在 go-zero 的配置解析中，默认使用小写键名，因此你需要将 YAML 中的 Mysql 改为小写 mysql
```
```
如何引入mysql
Name: user-api
Host: 0.0.0.0
Port: 8888

mysql:
  DataSource: root:123456@tcp(10.17.220.211:3306)/xieyuxaun?parseTime=true&loc=Asia%2FShanghai  //loc=Local

package config

import "github.com/zeromicro/go-zero/rest"

type Config struct {
	rest.RestConf
	Mysql struct { //数据库配置，除mysql外，还可能有mongo等其他数据库
		DataSource string //mysql链接地址，满足$user:$password@tcp($ip:端口)/数据库名称
	}
}

type ServiceContext struct {
	Config config.Config

	UserModel model.UserModel //加入表增删改查model
}

func NewServiceContext(c config.Config) *ServiceContext {
	// UserModel是接口类型
	// defaultUserModel实现了接口
	// 调用构造函数得到 *model.defaultUserModel
	// NewUserModel(conn sqlx.SqlConn)
	sqlxConn := sqlx.NewMysql(c.Mysql.DataSource)//新加
	return &ServiceContext{
		Config:    c,
		UserModel: model.NewUserModel(sqlxConn),//新加
	}
}
```