// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"io"
// 	"os"
// )

// // 完整试验文件读取写入操作
// func main() {
// 	filePath := "C:/test01/a.txt"
// 	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY, 0644)
// 	if err != nil {
// 		fmt.Println("err=", err)
// 	}

//		str := "王慧琳你比蒋涵文明多了\n"
//		writer := bufio.NewWriter(file)
//		for i := 0; i < 10000; i++ {
//			writer.WriteString(str)
//		}
//		writer.Flush()
//		file.Close()
//		file1, err1 := os.Open(filePath)
//		if err1 != nil {
//			fmt.Println("err=", err1)
//		}
//		defer file1.Close()
//		reader := bufio.NewReader(file1)
//		for {
//			str, err := reader.ReadString('\n')
//			if err == io.EOF {
//				break
//			}
//			fmt.Print(str)
//		}
//	}
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

// 完整试验文件读取写入操作
func main() {
	filePath := "C:/Jasper_xie/a.txt"

	// 获取目录路径
	dir := filepath.Dir(filePath)

	// 创建目录（如果不存在），包括所有父目录
	err := os.MkdirAll(dir, 0755)
	if err != nil {
		fmt.Printf("创建目录失败: %v\n", err)
		return
	}

	// 创建并打开文件
	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Printf("创建文件失败: %v\n", err)
		return
	}
	defer file.Close()

	str := "蒋涵、王慧琳\n"
	writer := bufio.NewWriter(file)
	for i := 0; i < 10000; i++ {
		writer.WriteString(str)
	}
	writer.Flush()

	// 重新打开文件进行读取
	file1, err1 := os.Open(filePath)
	if err1 != nil {
		fmt.Printf("打开文件失败: %v\n", err1)
		return
	}
	defer file1.Close()

	reader := bufio.NewReader(file1)
	for {
		str, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Printf("读取文件失败: %v\n", err)
			break
		}
		fmt.Print(str)
	}
}
