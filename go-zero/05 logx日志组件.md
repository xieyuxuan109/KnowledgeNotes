### 修改api/etc的yaml配置
```
Name: user-api
Host: 0.0.0.0
Port: 8888
Log:
  ServiceName: "user-srv"  # 服务名称，用于区分不同服务（默认："" 空字符串）
  Mode: console             # 日志输出模式：console/file/volume（默认：console）
  Encoding: json            # 日志编码格式：json/plain（默认：json）
  Path: logs                # 日志文件路径（默认：logs）
  Level: info               # 日志级别：debug/info/error/severe（默认：info） 输出>=info日志
  State: true               # 是否启用状态日志（默认：true）
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
### 使用logx记录日志
```
	logx.Debugv(req) //json.Marshal
	logx.Debugf("req:%#v", req)
    logx.Infov(req) //json.Marshal
	logx.Infof("req:%#v", req)
    logx.Errorw("user_signup_UserModel.FindOneByUsername failed", logx.Field("err", err)) //key-value
    logx.Errorf("user_signup_UserModel.Insert failed err:%v", err) 
```
### 也可与第三方日志库集成
### 日志脱敏