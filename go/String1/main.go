package main

import (
	"fmt"
	"strconv"
	"strings"
)

// 演示常见字符串处理函数
func main() {
	//1.len 对于字符串按字节算的 对于数组 切片是元素个数
	var a string = "12345"
	b := "谢宇轩"
	fmt.Println(len(a)) //5
	fmt.Println(len(b)) //9 在go语言中一个汉字占三个字节

	//2.字符串遍历 如果有中文需要先转化成[]rune
	r := []rune(b)
	for i := 0; i < len(r); i++ {
		fmt.Printf("%c", r[i])
	}

	//3.整数转字符串
	c := strconv.Itoa(12345)
	fmt.Println(c)

	//4.字符串转整数,注意会有两个返回值，一个string，一个err
	d, _ := strconv.Atoi("12345")
	fmt.Println(d)

	//5.整数转2，8，16进制的字符串
	e := strconv.FormatInt(100, 2) //1100100
	fmt.Println(e)                 //1100100

	//6.查找一个字符串是否含有特定的子串
	f := strings.Contains("123谢宇轩", "谢")
	fmt.Println(f) //true

	//7.统计一个字符串中含有几个指定的子串
	g := strings.Count("谢宇轩谢谢", "谢")
	fmt.Println(g) //3

	//8.字符串的比较
	h := "11"
	i := "11"
	fmt.Println(h == i)                          //区分大小写比较true
	fmt.Println(strings.EqualFold("abc", "ABC")) //不区分大小写比较true

	//9.返回子串在字符串中第一次出现的index值
	j := strings.Index("谢宇轩", "宇")
	fmt.Println(j) //3

	//10.byte转字符串
	k := string([]byte{97, 98, 99})
	fmt.Println(k)

	//11.字符串转byte
	bytes := []byte("hello go")
	fmt.Println(bytes)
}
