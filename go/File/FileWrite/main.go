package main

import (
	"bufio"
	"fmt"
	"os"
)
func PathExists(path string)(bool,error){
	_,err:=os.Stat(path)
	if err==nil{
		return true ,nil
	}
	if os.IsNotExist(err){
		return false,nil
	}
	return false,err
}
// 演示文件写入
func main() {
	//创建5句子新文件，写入五句hello world
	filePath := "C:/test01/abc.txt" //写入路径
	//判断文件路径是否存在os.Stat()
	//如果函数返回错误为nil 说明文件夹存在
	//如果返回错误类型使用os.IsNotExist()判断为true 说明文件或文件夹不存在
	//如果返回其他错误类型，则不确定是否存在


	// func OpenFile(name string, flag int, perm FileMode) (*File, error)
	// name：要打开的文件的路径（字符串类型）。
	// flag：指定文件打开方式的标志。这些标志在 os 包中定义，常见的包括：
	// os.O_RDONLY：只读
	// os.O_WRONLY：只写
	// os.O_RDWR：读写
	// os.O_CREATE：如果文件不存在则创建
	// os.O_APPEND：追加写入（在文件末尾写入）
	// os.O_TRUNC：打开文件时清空文件（截断为0）
	// 等等。这些标志可以通过按位或（|）组合使用。
	// perm：文件模式，用于指定文件的权限（如果创建了文件）。类型为 os.FileMode，通常用八进制数表示，例如 0644 表示文件所有者有读写权限，组用户和其他用户有读权限。但是windows下无效，在linux下才有用
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("err=", err)
	} //打开文件一般都会先判断打开是否成功
	defer file.Close() //defer好用
	str := "helloworld\n"
	writer := bufio.NewWriter(file)
	for i := 0; i < 5; i++ {
		writer.WriteString(str)
	}

	writer.Flush() //必须刷新，不然数据无法写入磁盘
}
