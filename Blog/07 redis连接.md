## redis连接
```
func InitRedis() *redis.Client {
	r := global.Config.Redis
	redisDB := redis.NewClient(&redis.Options{
		Addr:     r.Addr,     // 不写默认就是这个
		Password: r.Password, // 密码
		DB:       r.DB,       // 默认是0
	})
	_, err := redisDB.Ping().Result()
	if err != nil {
		logrus.Fatalf("redis连接失败 %s", err)
	}
	logrus.Info("redis连接成功")
	return redisDB
}
```
## 黑名单应用
```
type BlackType int8

const (
	UserBlackType   BlackType = 1
	AdminBlackType  BlackType = 2
	DeviceBlackType BlackType = 3
)

func ParseBlackType(val string) BlackType {
	switch val {
	case "1":
		return UserBlackType
	case "2":
		return AdminBlackType
	case "3":
		return DeviceBlackType
	default:
		return UserBlackType
	}
}
func TokenBlack(token string, value BlackType) {
	key := fmt.Sprintf("token_black_%s", token)
	claims, err := jwts.ParseToken(token, false)
	if err != nil {
		logrus.Errorf("token解析失败 %s", err)
		return
	}
	second := claims.RegisteredClaims.ExpiresAt.Unix() - time.Now().Unix()                               //新写法
	_, err = global.Redis.Set(key, fmt.Sprintf("%d", value), time.Duration(second)*time.Second).Result() //key value 过期时间
	if err != nil {
		logrus.Errorf("redis 添加黑名单失败 %s", err)
		return
	}
}

func IsTokenBlack(token string) (BlackType, bool) {
	key := fmt.Sprintf("token_black_%s", token)
	value, err := global.Redis.Get(key).Result()
	if err != nil {
		return 0, false
	}
	blk := ParseBlackType(value)
	return blk, true
}
```