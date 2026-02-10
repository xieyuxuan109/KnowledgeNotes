package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// 演示文件的操作
func main() {
	//流：数据在数据源(文件)和程序(内存)之间经历的路径
	//输入流：数据从数据源(文件)到程序(内存)的路径
	//输出流：数据从程序(内存)到数据源(文件)的路径
	//1.file叫file 对象
	//2.file叫file文件句柄
	//3.file叫file指针
	file, err := os.Open("C:/test01/test.txt") // 推荐使用正斜杠/分割路径，因为/不用转义
	if err != nil {
		fmt.Println("Open file err=", err)
	}
	defer file.Close()
	//创建一个*Reader 带缓冲4096个字节 好处多多
	reader := bufio.NewReader(file)
	for {
		str, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		fmt.Print(str)
	}
	fmt.Printf("文件读取结束")

	// 第二种读取文件方式
	//优点：可以不用打开和关闭文件，已经封装好了
	//缺点：不适用较大文件
	file1 := "c:/test01/test.txt"
	content, err := os.ReadFile(file1)
	if err != nil {
		fmt.Printf("read file err=%v\n", err)
	}
	fmt.Printf("%v", string(content))
	
}
