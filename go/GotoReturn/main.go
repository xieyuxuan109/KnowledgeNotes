package main

import "fmt"

// goto和return演示
func main() {
	//尽量不要用goto
	fmt.Println("ok1")
	goto label1
	fmt.Println("ok2")
	fmt.Println("ok3")
	fmt.Println("ok4")
label1:
	fmt.Println("ok5")
	fmt.Println("ok6")
	//return 使用在函数里面，表示结束
}
