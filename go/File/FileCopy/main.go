package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// 演示文件复制
// 自己编写一个函数实现传入文件路径的复制
func CopyFile(dstFileName string, srcFileName string) (written int64, err error) {
	srcFile, err := os.Open(srcFileName)
	if err != nil {
		fmt.Println("ERR=", err)
	}
	defer srcFile.Close()
	reader := bufio.NewReader(srcFile)
	dstFile, err := os.OpenFile(dstFileName, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("ERR=", err)
	}
	defer dstFile.Close()
	writer := bufio.NewWriter(dstFile)
	return io.Copy(writer, reader)

}
func main() {
	//主要使用io.Copy func Copy(dst Writer,src Reader)(written int64,err error)
	//注意copy函数只能接受writer和reader，不是文件路径进行复制，比较不方便
	//将C:/test01/a.txt拷贝到C:/test02/a.txt
	//注意：除了文本文件其他文件也可以被拷贝，例如图片、音频
	srcFileName := "C:/test01/a.txt"
	dstFileName := "C:/test02/a.txt"
	_, err := CopyFile(dstFileName, srcFileName)
	if err == nil {
		fmt.Println("拷贝完成.....")
	}
}
