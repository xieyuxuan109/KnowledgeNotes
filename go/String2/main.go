package main

import (
	"fmt"
	"strings"
)

// 字符串常用函数
// 注意：字符串转换时候本身没有改变，只是返回的值变了
func main() {
	//12.返回字符串中最后一次出现子字符串的下标,如果没有返回-1
	a := strings.LastIndex("123abc", "abc") //3
	fmt.Println(a)

	//13.将指定字符串替换成另外一个字串,最后一个参数表示替换几个，-1表示全部替换
	b := strings.Replace("go go go hello", "go", "go语言", -1) //go语言 go语言 go语言 hello
	fmt.Println(b)

	//14.按照指定的某个字符为分割标识，将一个字符串拆分成字符串数组
	c := strings.Split("hello,world", ",")
	fmt.Println(c[0]) //hello
	fmt.Println(c[1]) //world

	//15.进行字符串字母大小写转换
	str := "Hello World"
	str1 := strings.ToLower(str)
	fmt.Println(str1) //hello world
	str2 := strings.ToUpper(str)
	fmt.Println(str2) //HELLO WORLD

	//16.将字符串左右两边的空格全部去掉
	str3 := "   xieyuxuan    "
	str3 = strings.TrimSpace(str3)
	fmt.Println(str3) //xieyuxuan

	//17.将字符串左右两边指定的字符去掉
	str4 := "!!hello!"
	str4 = strings.Trim(str4, "!")
	fmt.Println(str4) //hello

	//18.将字符串左边指定的字符去掉
	str5 := "!!hello!"
	str5 = strings.TrimLeft(str5, "!")
	fmt.Println(str5) //hello!

	//19.将字符串右边指定的字符去掉
	str6 := "!!hello!"
	str6 = strings.TrimRight(str6, "!")
	fmt.Println(str6) //!!hello

	//20.判断字符串是否以指定字符串开头
	h := strings.HasPrefix("ftp://192.168.10.1", "ftp") //true
	fmt.Println(h)

	//21.判断字符串是否以指定的字符串结束
	i := strings.HasSuffix("NLT_abc.hpg", "abc") //false
	fmt.Println(i)
}
