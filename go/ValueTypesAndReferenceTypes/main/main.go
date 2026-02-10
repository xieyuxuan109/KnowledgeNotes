package main

import (
	"fmt"

	"github.com/xieyuxuan109/project01/ValueTypesAndReferenceTypes/model"
)

func main() {
	//值类型总结
	//int float bool string 数组 结构体

	//引用类型
	//指针 slice切片 map chan管道 interface

	// 1. 值类型 (Value Types)
	// 变量直接存储值
	// 赋值和传递时拷贝整个值
	// 内存通常在栈上分配（除非逃逸分析确定需要在堆上）
	// 变量在声明后会分配内存并初始化为零值

	// 2. 引用类型 (Reference Types)
	// 变量存储的是内存地址（引用/指针）
	// 赋值和传递时拷贝地址，不拷贝底层数据
	// 内存通常在堆上分配
	// 变量在声明时为nil，需要显式初始化（如使用make）
	fmt.Println(model.Arr) //调用成功 10
	var res float64
	res = model.Cal(1.2, 1.2, '+') //专业术语 该函数可导出
	fmt.Println(res)
}
