package main

import "fmt"

// 演示键盘输入语句
func main() {
	var name string
	var age byte
	var sal float32
	var isPass bool
	//fmt.Scanln用法
	fmt.Println("请输入姓名 ")
	fmt.Scanln(&name)
	fmt.Println("请输入年龄 ")
	fmt.Scanln(&age)
	fmt.Println("请输入薪水 ")
	fmt.Scanln(&sal)
	fmt.Println("请输入是否通过考试 ")
	fmt.Scanln(&isPass)
	fmt.Printf("姓名是%v\n年龄是%v\n薪水是%v\n是否通过考试是%v\n", name, age, sal, isPass)
	//fmt.Scanf用法
	fmt.Printf("请输入你的姓名，年龄，薪水，是否通过考试，使用空格隔开")
	fmt.Scanf("%v %v %v %v %v", &name, &age, &sal, &isPass)
	fmt.Printf("姓名是%v\n年龄是%v\n薪水是%v\n是否通过考试是%v\n", name, age, sal, isPass)
}
