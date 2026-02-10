package main

import "fmt"

// init函数演示
var age = test()

func test() int {
	fmt.Println("test()...")
	return 90
}
func main() {
	//每一个源文件都可以包含多个init函数，该函数会在main函数执行前被Go运行框架调用
	//如果一个函数同时包含全局变量定义，init函数和main函数，则执行流程是变量定义->init函数->main函数
	fmt.Println("age=", age)
	//匿名函数就是没有名字的函数
	//方式一：直接定义，匿名函数只能用一次，因为没有名字，也没有赋值给变量
	a := func(a int, b int) int {
		return a + b
	}(1, 2)
	fmt.Println(a)
	//方式二：赋值给一个变量使用,如果在全局定义，那是全局匿名函数
	b := func(a int, b int) int {
		return a + b
	}

	fmt.Println(b(10, 20))
}
