package main

import (
	"fmt"
	"sort"
)

// 映射演示
func main() {
	//map是一种无序key-value的数据结构，又称为字段或者关联数组，类似其他编程语言的集合
	//map声明
	//key可以是很多类型string bool 结构体等等 但是不能有slice map function，因为他们不能用==来判断
	//key通常是int string作为key的类型
	//value通常是数字(整数，浮点数),string,map,struct
	//注意：引用类型声明后不会分配内存，需要make后才可用
	a := make(map[string]string, 10) //make申请内存，最多可以放10个key-value，可以不写
	//赋值时候key重复会覆盖之前的value
	a["xieyuaun"] = "123"
	a["siyunna1"] = "123"
	a["jianghan1"] = "123"
	a["1"] = "123"
	a["2"] = "123"
	a["3"] = "123"
	a["4"] = "123"
	a["5"] = "123"
	a["6"] = "123"
	a["7"] = "123"
	a["8"] = "123"
	//其他使用方式和切片类似 先声明后make 便声明边make 声明make赋值
	fmt.Println(a)

	//案例
	stu := make(map[string]map[string]string)
	stu["stu01"] = make(map[string]string)
	stu["stu01"]["姓名"] = "谢宇轩"
	stu["stu01"]["性别"] = "男"

	//修改直接用之前的key修改
	//没有的key就是增加
	//删除 内置函数delete,当指定的key不存在的时候不进行任何操作也不报错
	delete(a, "1")
	fmt.Println(a)

	//go语言没有能一次性删除所有的key，只能遍历然后一步一步删除
	//也可以map=make()一个新空间，让之前的空间变成垃圾，被回收

	//查询
	val, findness := a["1"] //会返回两个值，如果存在会findness为true，否则为false
	if findness {
		fmt.Println(val)
	} else {
		fmt.Println("没有这个值")
	}
	//map遍历只能用for-range
	for k, v := range a {
		fmt.Println(k, v)
	}
	fmt.Println(len(a))

	//map切片 能够让map动态增加
	monster := make([]map[string]string, 10)
	newMonster := map[string]string{
		"1": "2",
		"2": "3",
	}
	monster = append(monster, newMonster)
	fmt.Println(monster)

	//map排序 将key存入切片
	map1 := make(map[int]int, 10)
	map1[10] = 100
	map1[2] = 13
	map1[4] = 56
	map1[8] = 90
	var keys []int
	for k, _ := range map1 {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	fmt.Println(keys) //将key排好序
	//将key排好序后将key按序号输出
	for _, v := range keys {
		fmt.Println(map1[v])
	}
	//注意事项
	//map是引用类型
	//map可以自动扩容
	//map也经常使用struct而不是嵌套map
}
