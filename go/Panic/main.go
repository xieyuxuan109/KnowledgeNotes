package main

import (
	"errors"
	"fmt"
)

// 演示go语言的错误处理机制
// defer recover panic
// recover()，panic()是一个内置函数，可以捕获到异常
// 示例一：defer+recover
// 示例二：自定义错误 errors.New panic
// errors.New可以创建一个error类型的值，表示一个错误
// panic内置函数，接受一个interface{}类型的值作为参数，可以接受error类型的变量，输出错误信息，并退出程序
func readConf(name string) (err error) {
	if name == "config.ini" {
		return nil
	} else {
		return errors.New("读取文件错误")
	}
}
func test02() {
	err := readConf("config.in")
	if err != nil {
		panic(err)
	}
	fmt.Printf("test02后面的代码继续执行")
}
func test() {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("res=", err)
			//这里可以将信息发送给管理员
		}
	}()
	num1 := 10
	num2 := 0
	res := num1 / num2
	fmt.Println("res=", res)
}

func main() {
	// test()
	// fmt.Println("main()下面的代码...")
	test02()
}
