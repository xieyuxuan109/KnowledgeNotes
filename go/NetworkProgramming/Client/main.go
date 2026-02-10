package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	conn, err := net.Dial("tcp", "192.168.1.100:8888")
	if err != nil {
		fmt.Println("client dial err=", err)
		return
	}
	//1.客户端可以发送单行数据，然后退出
	reader := bufio.NewReader(os.Stdin) //os.Stdin代表标准输入[终端] 一般指键盘
	//从终端读取一行，并准备发送给服务器
	for {
		content, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("readstring err=", err)
		}
		content = strings.TrimSpace(content)

		if content == "exit" {
			fmt.Println("客户端退出...")
			break
		}
		//再将line 发送给 服务器
		_, err = conn.Write([]byte(content + "\n")) //传入切片 返回n代表写了多少字节
		if err != nil {
			fmt.Println("readstring err=", err)
		}
	}
}
