package main

import (
	"fmt"
)

// for循环代码实现
func main() {
	//方式一：类似C语言实现方式
	for i := 1; i < 10; i++ {
		fmt.Println("你好，世界！")
	}
	//方式二：只写循环条件 类似C语言while实现方式
	var j int = 1
	for j <= 10 {
		fmt.Println("你好，我是谢宇轩！", j)
		j++
	}
	//方式三：死循环+break语句配合使用
	var k int = 1
	for {
		fmt.Println("你好")
		k++
		if k > 10 {
			break
		}
	}
	//方式四：for-range遍历 可以遍历中文
	//传统是一个字节一个字节遍历的，但是中文再UTF-8中是三个字节
	//for-range+[]rune就可以遍历中文
	//按字符遍历，也可以不用转化，但是会跳字节
	var str string = "abcd北京"
	var slice []rune = []rune(str)
	for index, val := range slice {
		fmt.Printf("index=%d val=%c\n", index, val)
	}

	//注意：go语言没有while和do-while语言 可以用for语言实现类似效果
	//打印九九乘法表
	for i := 1; i < 10; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%v*%v=%v ", j, i, i*j)
		}
		fmt.Printf("\n")
	}
}
