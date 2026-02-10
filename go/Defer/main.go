package main

import "fmt"

func sayHello() {

	defer fmt.Println("hello...defer")
	fmt.Println("hello...say")
}
func main() {
	//当go遇到defer时候，不会立即执行defer后面的语句，同时将defer要用到的资源和值拷贝入栈
	//当函数调用结束时，从defer栈中依次取出执行，遵循先入后出
	fmt.Println("hello...main1")
	sayHello()
	fmt.Println("hello...main2")
	// hello...main1
	// hello...say
	// hello...defer
	// hello...main2
}
