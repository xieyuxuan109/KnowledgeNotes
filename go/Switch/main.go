package main

import "fmt"

// switch演示
func main() {
	//go中的switch分支自带break
	//go中case可以跟多个值
	var ch byte
	fmt.Printf("请输入一个字符：")
	fmt.Scanf("%c", &ch)
	//switch后面也可以不接表达式，类似ifelse使用
	//case后面也可以接范围，例如 case a>60：
	//如果需要穿透，需要添加fallthrough关键字，默认只穿透一层
	//还可以用作type-switch，用来判断interface类型
	switch ch {
	case 'a':
		fmt.Printf("星期一")
	case 'b':
		fmt.Printf("星期二")
	case 'c':
		fmt.Printf("星期三")
	case 'd':
		fmt.Printf("星期四")
	case 'e':
		fmt.Printf("星期五")
	case 'f':
		fmt.Printf("星期六")
	case 'g':
		fmt.Printf("星期天")
	default:
		fmt.Println("输入错误")
	}
}
