package main

import (
	"fmt"
	"net"
)

func Process(conn net.Conn) {
	//这里循环介绍客户端发送的数据
	defer conn.Close() //关闭conn
	for {
		//创建一个新的切片
		buf := make([]byte, 1024)
		//如果客户端没有write[发送],那么协程就会阻塞到这里
		n, err := conn.Read(buf) //传入切片 从conn中读取 回阻塞
		if err != nil {
			fmt.Println("err=", err)
			return
		}
		//显示客户端发送的内容到服务器终端
		fmt.Print(string(buf[:n]))
	}
}

func main() {
	fmt.Println("服务器开始监听...")
	//0.0.0.0表示所有本机网络接口监听8888
	// 监听 第一参数表示协议 第二个表示ip+端口 net.Println("tcp","自己ip地址:8888")
	//只能监听一会
	listen, err := net.Listen("tcp", "0.0.0.0:8888")
	if err != nil {
		fmt.Println("listen err=", err)
		return
	}
	defer listen.Close()
	//循环等待客户端连接
	for {
		//等待客户端连接
		fmt.Println("等待客户端连接...")
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("Accept() err=", err)
		} else {
			fmt.Printf("Accept() suc con=%v 客户端ip=%v\n", conn, conn.RemoteAddr().String())
		}
		//这里准备起一个协程，为客户端服务
		//telnet可以测试一个端口是否连接
		go Process(conn)
	}
	//fmt.Printf("%v", listen)

}
