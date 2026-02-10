package main

import (
	"fmt"

	"github.com/gomodule/redigo/redis"
)

func main() {
	//通过go写入数据和读取数据
	//1.连接到redis
	conn, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		fmt.Println("redis err=", err)
		return
	}
	defer conn.Close()
	_, err = conn.Do("set", "name", "xieyuxuan") //会返回结果，err
	if err != nil {
		fmt.Println("redis err=", err)
		return
	}
	//返回的是interface{}类型 但是不能使用类型断言转换
	r, err := redis.String(conn.Do("get", "name"))
	if err != nil {
		fmt.Println("redis err=", err)
		return
	}
	r1, err := redis.Strings(conn.Do("hmget", "1", "xie", "yu", "12"))
	if err != nil {
		fmt.Println("redis err=", err)
		return
	}
	for i, v := range r1 {
		fmt.Println(i, v)
	}
	fmt.Println("操作成功！", r)
}
