package main

import (
	"fmt"
	"unsafe"
)

// 演示go常见数据类型及占用字节数 占用在结束用unsafe包下的Sizeof()查看
func main() {
	//整型 包括byte(uint8),int16,rune(int36),int64 int类型与系统有关 64位系统为int64 32位系统为int32
	var a int = 2
	fmt.Printf("a的数据类型为%T，数据占用字节数为%d\n", a, unsafe.Sizeof(a))
	//浮点型 包括float32,float64(默认) 注意go中无float double类型
	var b float32 = 2.0
	fmt.Printf("b的数据类型为%T，数据占用字节数为%d\n", b, unsafe.Sizeof(b))

	//浮点型数据注意精度损失
	var c float32 = -123.0000901 //打印结果为-123.00009
	var d float64 = -123.0000901 //打印结果为-123.000091
	e := .123                    //特殊写法 结果为0.123
	f := 5.1234e2                //科学记数法 512.34 中间E可大写 5.1234E2 类型为浮点数
	fmt.Println("c=", c)
	fmt.Println("d=", d)
	fmt.Println("e=", e)
	fmt.Println("f=", f)

	//字符型 go语言中没有专门的字符类型 一般采用byte(uint8)，rune(int32)来表示
	var g byte = 'a'        //注意因为byte范围有限，不能储存中文字符，如果要用中文，用rune
	fmt.Println("g=", g)    //输出对应的码值
	fmt.Printf("g=%c\n", g) //输出对应的字符%c
	var h rune = '中'
	fmt.Printf("h=%c", h) //输出h=中

	//布尔型 go中bool类型只能是true或者false 占用一个字节
	//不可以用0或非0代替
	var i bool = true
	fmt.Printf("a的数据类型为%T，数据占用字节数为%d\n", i, unsafe.Sizeof(i)) //bool 1

	//字符串型 传统字符串是由字符组成的char[]，但是go的字符串不同，是由字节组成的[]byte
	//go语言使用UTF-8进行处理字符串的
	//go语言中有两个表示形式""(双引号)或者``(反引号)
	var j string = "北京长城number1" //注意 字符串一但赋值了，Go中的字符串就不可变了
	fmt.Println(j)
	//双引号会识别转义字符，输出特殊字符不方便，单引号是以字符串的原生形式输出
	str1 := "abc\nabc"
	fmt.Println(str1)
	str2 := `
	fmt.Println("c=", c)
	fmt.Println("d=", d)
	fmt.Println("e=", e)
	fmt.Println("f=", f)
	`
	fmt.Println(str2) //把``中内容当成文本输出
	//字符串拼接方式
	str3 := str1 + str2
	fmt.Println(str3)
	//当拼接操作很长时候，可以换行写但是+必须留在上一行末尾
	str4 := "hello" + "hello" + "hello" + "hello" + "hello" + "hello" + "hello" +
		"hello" + "hello" + "hello" + "hello" + "hello" + "hello"
	fmt.Println(str4) //hellohellohellohellohellohellohellohellohellohellohellohellohello

	//基本数据类型默认值，也叫零值 %v：按照变量的原始值输出
	/*
		int=0
		float32=0
		float64=0 浮点型数据使用默认输出类型时候会省略小数点
		bool=false
		string=""
	*/
}
