package main

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type MyClaims struct {
	UserName string `json:"username"`

	jwt.StandardClaims
}

func main() {
	//加密key
	mySigningKey := []byte("nihaoxieyuxuan")
	//MyClaims map[string]interface实现了
	//StardardClaims struct
	//加密方式有很多，可以仔细了解
	//工作中一般使用rs256非对称加密
	//hs256对称加密
	c := MyClaims{
		UserName: "qimiao",
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix() - 60,
			ExpiresAt: time.Now().Unix() + 60*60*2,
			Issuer:    "谢宇轩",
		},
	}
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, c) //第一个参数是加密方法 第二个是实现Claims接口的结构体
	s, e := t.SignedString(mySigningKey)              //加密token
	if e != nil {
		fmt.Println(e)
	}
	fmt.Println(s)
	//解密
	token, err := jwt.ParseWithClaims(s, &MyClaims{}, func(t *jwt.Token) (interface{}, error) {
		return mySigningKey, nil
	}) //第一个参数是加密token，第二个是结构体模板，第三个是写一个func
	fmt.Println(token.Claims.(*MyClaims)) //记住一定得断言
	fmt.Println(err)
}

//map类型类似
