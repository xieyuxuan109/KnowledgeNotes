package main

import "fmt"

func main() {
	var age byte
	fmt.Printf("请输入你的年龄：")
	fmt.Scanf("%v", &age)
	if age > 18 {
		fmt.Printf("你的年龄大于18，你要对自己的行为负责！\n")
	} else {
		fmt.Printf("你还没满18岁\n")
	}
	//注意go语言ifelse后面必须有大括号
	//go支持再if中直接定义一个变量
	//注意else必须紧贴if的大括号
	if age1 := 20; age1 > 18 {
		fmt.Printf("你的年龄大于18，你要对自己的行为负责！\n")
	} else {
		fmt.Printf("你还没满18岁\n")
	}
}
