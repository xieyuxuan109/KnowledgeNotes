package main

import (
	"fmt"
	"strconv"
)

// 演示基本数据类型之间的转换
func main() {
	//go中数据类型只能显示转换，不能自动转换，转换方式类似于python
	//int->float32
	var a int32 = 100
	var b float32 = float32(a)
	var c int8 = int8(a)
	var d float64 = float64(a)              //低精度到高精度也需要强制转化 高到低有可能会存在溢出现象
	fmt.Printf("b=%v c=%v d=%v\n", b, c, d) //发生溢出时候编译不会报错 但是结果会和预期不一样
	fmt.Printf("a type is %T\n", a)         //int32 被转化的是变量贮存的数据，变量本身的数据类型没有发生变化

	//基本数据类型和string的转换
	//方式1 fmt包中的Sprintf("%参数",表达式)推荐 生成一个格式化的字符串
	var num1 int = 99
	var num2 float32 = 23.456
	var num3 bool = true
	var num4 byte = 'h'
	var str string
	str = fmt.Sprintf("%d", num1)
	//%q 将对应的字面值加上引号输出
	fmt.Printf("str type %T str=%q\n", str, str) //str type string str="99"
	str = fmt.Sprintf("%f", num2)
	fmt.Printf("str type %T str=%q\n", str, str) //str type string str="23.455999"
	str = fmt.Sprintf("%t", num3)
	fmt.Printf("str type %T str=%q\n", str, str) //str type string str="true"
	str = fmt.Sprintf("%c", num4)
	fmt.Printf("str type %T str=%q\n", str, str) //str type string str="h"

	//方式2 strconv包中的函数
	//例如strconv.FormatInt(i int64,base int)string
	//i是需要转的整数，base表示转成什么进制
	str = strconv.FormatInt(int64(num1), 10)
	fmt.Printf("str type %T str=%q\n", str, str) //str type string str="99"
	//strconv.FormatFloat(f float64,fmt byte,prex,bitSize int)
	//fmt是格式 prex表示小数保留几位 bitSize表示转化过后是float32还是float64类型
	str = strconv.FormatFloat(float64(num2), 'f', 10, 64)
	fmt.Printf("str type %T str=%q\n", str, str) //str type string str="23.4559993744"
	//strconv.FormatBool(b bool)
	str = strconv.FormatBool(num3)
	fmt.Printf("str type %T str=%q\n", str, str) //str type string str="true"

	//strconv.Itoa(num int)string
	var num5 int = 4567
	str = strconv.Itoa(num5)
	fmt.Printf("str type %T str=%q\n", str, str) //str type string str="4567"

	//字符串转为其他类型
	var x bool
	var str5 string = "true"
	var str6 = "123456"
	var n1 int64
	var n2 int
	x, _ = strconv.ParseBool(str5)
	fmt.Printf("x type %T x=%v\n", x, x)
	n1, _ = strconv.ParseInt(str6, 10, 64) //注意返回值为int64
	n2 = int(n1)                           // bitsize旨在帮你先检查一遍转化为相应位数会不会溢出
	fmt.Printf("n1 type %T n1=%v\n", n1, n1)
	fmt.Printf("n2 type %T n2=%v\n", n2, n2)
	var str7 string = "123.456"
	var f1 float64
	f1, _ = strconv.ParseFloat(str7, 64)
	fmt.Printf("f1 type %T f1=%v\n", f1, f1)
	//注意：String类型数据转成其他类型时候需要确保String可以转成有效的数据
	//比如"hello"转成整数 会转成数据类型默认值
}
