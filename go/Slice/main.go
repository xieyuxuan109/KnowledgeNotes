package main

import (
	"fmt"
)

// 演示切片
func main() {
	//切片可以简单的理解为动态的数组，长度可变
	//切片是数组的一个引用
	//切片定义语法
	var a []int
	fmt.Println(a)
	//基本使用
	//方式一：基于数组创建
	var intArr [5]int = [...]int{1, 2, 3, 4, 5}
	slice := intArr[1:3]
	fmt.Println(slice)
	fmt.Println("切片长度", len(slice))
	fmt.Println("切片容量", cap(slice)) //切片容量动态变化cap 表示底层数组的长度
	fmt.Printf("%T\n", slice)

	//切片内存形式
	//数据结构本质就是一个结构体，第一个寻访第一个元素地址，第二个存放slice长度，第三个存放cap
	//切片是数组的引用，改变切片相当于改变数组
	fmt.Println(slice[1]) //3
	intArr[2] = 10
	fmt.Println(slice[1]) //10

	//切片使用的第二种方式make创建，类型，大小，容量(可选) 创建后有默认值0
	var slice2 []int = make([]int, 4, 10) //可以用来创建切片 映射 通道
	//注意：cap>=len,通过make创建的切片对应的数组是由make底层维护，对外部可见
	fmt.Println(slice2)

	//第三种方式：定义一个切片，直接指向具体数组，使用原理类似make
	arrSlice := []string{"hello", "nihao", "谢谢"}
	fmt.Println(arrSlice)

	//方式一和方式二区别
	//一个底层数组是可见的，一个是不可见的
	//切片可以继续切片

	//切片遍历
	var slice3 []int = []int{1, 2, 3}
	for key, val := range slice3 {
		fmt.Println(key, val)
	}

	//切片使用注意事项
	//append内置函数，追加元素，实际上就是新开一个数组，然后将新的数据加在里面，然后重新引用
	slice4 := append(slice3, 400, 500, 600) //追加单个元素
	fmt.Println(slice4)
	slice3 = append(slice3, slice4...) //三个点必须有，追加切片
	fmt.Println(slice3)

	//切片拷贝
	slice5 := make([]int, 10)
	copy(slice5, slice4) //将slice4的内容拷贝进slice5，覆盖性拷贝,两个必须都是slice类型才可以拷贝
	fmt.Println(slice5)

	//特殊情况
	//引用类型调用函数时侯传的是地址
	slice6 := make([]int, 1)
	copy(slice6, slice4) //只会copy一个函数
	fmt.Println(slice6)

	//string和切片区联系 string本质上就是[]byte数组
	str := "helloxieyuxuan"
	slice1 := str[:6]
	fmt.Println(slice1)

	//string是不可变的，也就是说不能通过a[0]='a'来改变
	//但是可以通过string->[]byte或[]rune->string来改变
}
