### 修改相关api文件
```
@server (
	jwt:    Auth //添加jwt鉴权
	prefix: api
)
service user-api {
	@handler DetailHandler
	get /user/detail (DetailRequest) returns (DetailResponse)
}
```
### 重新生成
```
goctl api go -api user.api -dir . -style=goZero
```
### 修改api/etc的yaml配置
```
Name: user-api
Host: 0.0.0.0
Port: 8888
Log:
  ServiceName: "user-srv"  # 服务名称，用于区分不同服务（默认："" 空字符串）
  Mode: file             # 日志输出模式：console/file/volume（默认：console）
  Encoding: json            # 日志编码格式：json/plain（默认：json）
  Path: logs                # 日志文件路径（默认：logs）
  Level: info               # 日志级别：debug/info/error/severe（默认：info） 输出>=info日志
  State: true               # 是否启用状态日志（默认：true）
Auth:
  AccessSecret: "夏天夏天悄悄过去"
  AccessExpire: 100000000
Mysql:
  DataSource: root:root@tcp(10.17.220.221:3306)/xieyuxuan?parseTime=true&loc=Local
CacheRedis:
  - Host: 10.17.220.221:6379
```
### 修改api/internal/config配置
```
type Config struct {
	rest.RestConf
	Auth struct { //jwt权限配置
		AccessSecret string //jwt密钥
		AccessExpire int64  //有效期，单位秒
	}
	Mysql struct { //数据库配置 使用匿名结构体 Mysql是字段名不是结构体名字
		DataSource string //mysql连接地址
	}
	CacheRedis cache.CacheConf
}
```
### 添加jwt相关函数
```
//生成函数
func (l *LoginLogic) getJwtToken(secretKey string, iat, seconds, userid int64) (string, error) {
	claims := make(jwt.MapClaims)
	claims["exp"] = iat + seconds
	claims["iat"] = iat
	claims["userid"] = userid
	claims["author"] = "q1mi" // 添加一些自定义的
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(secretKey))
}
```
```
go-zero中内置jwt解析函数，不用额外配置
```
### 调用jwt函数
```
func (l *LoginLogic) Login(req *types.LoginRequest) (resp *types.LoginResponse, err error) {
	// todo: add your logic here and delete this line
    ······
	now := time.Now().Unix()
	expire := l.svcCtx.Config.Auth.AccessExpire
	token, err := l.getJwtToken(l.svcCtx.Config.Auth.AccessSecret, now, expire, user.UserId)//调用函数
	fmt.Printf("JWT author:%v\n", l.ctx.Value("author"))//取出token里面参数
	if err != nil {
		logx.Errorw("l.generateAccessToken failed", logx.Field("err", err))
		return nil, errors.New("内部错误")
	}
    ······
	return &types.LoginResponse{
		Message:      "success",
		AccessToken:  token,
		AccessExpire: int(now + expire),
		RefreshAfter: int(now + expire/2),
	}, nil
}
```

