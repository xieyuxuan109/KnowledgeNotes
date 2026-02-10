package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Open("C:/test01/test.txt")
	if err != nil {
		fmt.Println("err!=", err)
		return
	}
	defer file.Close()
	reader := bufio.NewReader(file) //bufio 缓冲io 适合文件较大的方式读取
	for {
		str, err := reader.ReadString('\n') //会将换行读取
		if err == io.EOF {
			break
		}
		fmt.Print(str)
	}
}
