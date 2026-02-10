package main

import "fmt"

// 演示指针的使用
func main() {
	//类似C语言中指针
	var i int = 1
	fmt.Println("i的地址=", &i)
	var ptr *int = &i
	fmt.Printf("ptr=%v\n", ptr)
	fmt.Printf("ptr的地址=%v\n", &ptr)
	fmt.Printf("ptr指向的值=%v\n", *ptr)
}
