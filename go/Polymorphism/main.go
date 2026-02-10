package main

import "fmt"

// 多态演示
func main() {
	//利用接口编程
	//本质就是利用接口的特性，可以存储多个不同的数据类型或者接受多个不同的数据类型
	a := []interface{}{
		1, "123", 6.6, true,
	} //多态的体现
	//类型断言
	//大多数应用到接口上，判断接口指向的是什么类型
	b, ok := a[0].(int) //断言
	if ok {
		fmt.Println(b)
	}
	c := a[1].(string)
	fmt.Println(c)

}
