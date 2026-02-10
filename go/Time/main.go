package main

import (
	"fmt"
	"strconv"
	"time"
)

// 关于时间的函数演示
func main() {
	//获取当前时间 类型time.Time
	now := time.Now()
	fmt.Printf("%v %T\n", now, now) //2025-12-30 17:41:02.7220361 +0800 CST m=+0.000528501 time.Time

	//常见处理
	fmt.Printf("年=%v\n", now.Year())
	fmt.Printf("月=%v\n", now.Month())
	fmt.Printf("日=%v\n", int(now.Month())) //注意这个强转不太一样，"五"->5，对应里面的枚举类型
	fmt.Printf("时=%v\n", now.Day())
	fmt.Printf("分=%v\n", now.Minute())
	fmt.Printf("秒=%v\n", now.Second())

	//格式化日期时间 第一种
	fmt.Printf("当前年月日 %d-%02d-%02d %02d:%02d:%02d \n", now.Year(),
		now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second())

	//格式化日期时间 第二种
	fmt.Printf(now.Format("2006-01-02 15:04:05"))
	fmt.Println()
	fmt.Printf(now.Format("15:04:05"))
	fmt.Println()
	fmt.Printf(now.Format("2006-01-02"))
	fmt.Println()

	//时间戳
	fmt.Println(now.Unix())

	// 程序运行的时间
	str := ""
	begin := time.Now().Unix()
	for i := 0; i < 100000; i++ {
		str += strconv.Itoa(i)
	}
	end := time.Now().Unix()
	fmt.Printf("程序运行的时间是%v", end-begin)

	//使程序休眠
	time.Sleep(10 * time.Second)
	end2 := time.Now().Unix()
	fmt.Printf("程序运行的时间是%v", end2-end)
}
