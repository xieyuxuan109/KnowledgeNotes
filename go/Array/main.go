package main

import (
	"fmt"
	"math/rand/v2"
)

// 数组演示
func main() {
	//类似于C语言数组
	//数组定义，如果不赋值，默认会有零值
	var arr [5]int
	//函数
	arr[0] = 10
	fmt.Println(arr[0])
	//内存布局
	//类似C语言
	//数组名为第一个元素的地址

	//四种初始化数组的方式
	var numArr01 [3]int = [3]int{1, 2, 3}
	var numArr02 = [3]int{1, 2, 3}
	var numArr03 = [...]int{1, 2, 3}
	var numArr04 = [...]int{1: 1, 0: 2, 5: 3} //[2 1 0 0 0 3]指定下标
	fmt.Println(numArr01)
	fmt.Println(numArr02)
	fmt.Println(numArr03)
	fmt.Println(numArr04)

	//注意事项
	//数组只能存放相同类型的数据
	//数组一但声明，长度是固定不变的
	//var arr []int是切片不是数组
	//数组中的元素可以是任何数据类型，包括值类型和引用类型，但是不能混用
	//数组创建后如果不赋值，会有默认值
	//go语言数组属于值类型

	//数组的反转
	var numArr05 [5]int
	for i := 0; i < 5; i++ {
		numArr05[i] = rand.IntN(100) //0-100 大于等于0小于100
	}
	fmt.Println(numArr05)
	for i := 0; i < len(numArr05)/2; i++ {
		temp := numArr05[i]
		numArr05[i] = numArr05[len(numArr05)-1-i]
		numArr05[len(numArr05)-1-i] = temp
	}
	fmt.Println(numArr05)
}
