package main

//演示统计一个文件里面各种字符分别有多少
import (
	"bufio"
	"fmt"
	"io"
	"os"
)

type Count struct {
	Num   int64
	Alpha int64
	Space int64
}

func main() {
	filePath := "C:/test01/a.txt"
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("err=", err)
	}
	defer file.Close()
	var count Count
	reader := bufio.NewReader(file)
	for {
		str, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		for _, v := range str {
			//swith语句的特殊使用，switch后面可以什么都不接，相当于if语句但是更加直观
			switch {
			case v >= 'a' && v <= 'z':
				fallthrough
			case v >= 'A' && v <= 'Z':
				count.Alpha++
			case v <= '9' && v >= '0':
				count.Num++
			case v == ' ' || v == '\t':
				count.Space++
			}
		}
	}
	fmt.Println(count.Alpha, count.Space, count.Num)
}
