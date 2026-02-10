package main

import "fmt"

// 函数演示
func Cal(n1 float64, n2 float64, operator byte) float64 {
	var res float64
	switch operator {
	case '+':
		res = n1 + n2
	case '-':
		res = n1 - n2
	case '*':
		res = n1 * n2
	case '/':
		res = n1 / n2
	default:
		fmt.Printf("输入错误")
	}
	return res
}
func getSum(a int, b int) int {
	return a + b
}

type myfunc func(int, int) int

func fun(myfun myfunc, n1 int, n2 int) int {
	return myfun(n1, n2)
}

func f(n1 int, n2 int) (sum int, sub int) {
	sum = n1 + n2
	sub = n1 - n2
	return
}
func Sum(n1 int, args ...int) (sum1 int) {
	sum1 = n1
	for i := 0; i < len(args); i++ {
		sum1 += args[i]
	}
	return
}
func main() {
	var n1 float64
	var n2 float64
	var ch byte
	fmt.Printf("请输入一个算式：")
	fmt.Scanf("%f%c%f", &n1, &ch, &n2)
	res := Cal(n1, n2, ch)
	fmt.Println(res)

	//注意函数参数传递分为值传递和地址传递

	//go函数可以返回多个值，如果想忽略某个返回值，用_忽略
	//go函数返回一个之的时候返回值列表一般不加()，多个值必须加()
	//1一个值加变量名也需要()
	//函数递归调用 类似C语言
	//go语言不支持重载

	//go中函数也属于一种数据类型，可以赋值给一个变量，变量为函数类型
	//go语言中变量类型主要分为四类 值类型 引用类型 接口类型 函数类型
	a := getSum
	fmt.Printf("a的数据类型为%T getSum()的数据类型为%T\n", a, getSum) //a的数据类型为func(int, int) int getSum()的数据类型为func(int, int) int
	ans := a(10, 20)
	fmt.Printf("10+20=%v\n", ans) //30

	//函数作为参数传递
	ans2 := fun(getSum, 10, 20)
	fmt.Println("10+20=", ans2)

	//为了简化数据类型定义，go语言支持自定义数据类型
	//但是go语言从语法上认为他们是两种类型
	//但实际上他们都是一种类型
	//一般用于简化函数类型，但是必须注意定义顺序，type定义在函数之前
	type myInt int
	var num1 myInt = 1
	var num2 int = int(num1) //必须强转，因为go语言认为他们不一样
	fmt.Println(num2)        //1
	//函数也可以直接先对函数返回值命名，这样后就可以不写了，直接写return就行
	//并且相当于将返回值帮你定义了，你不用重复定义
	sum, sub := f(1, 2)
	fmt.Println(sum, sub) //3 -1

	//go语言支持可变参数 args(参数名字，随便取)...int(函数类型)
	//args是一个切片，通过args[index]
	//一般会使用for
	//注意：可变参数一定要放在形参列表最后
	sum = Sum(1, 2, 3, 4, 5, 6, 7, 8)
	fmt.Println("1+2+3+4+5+6+7+8=", sum)
}
