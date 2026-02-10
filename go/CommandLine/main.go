package main

import (
	"flag"
	"fmt"
	"os"
)

// 演示如何获取命令行参数
func main() {
	//使用os.Args()
	fmt.Println("命令行参数有", len(os.Args))
	for i, v := range os.Args {
		fmt.Printf("args[%v]=%v\n", i, v)
	}

	//下面展示如何指定参数进行传值
	var user string
	var pwd string
	var host string
	var port int //后面这个类型跟你需要接受的类型有关
	//&user就是接受用户命令行中输入的-u后面的参数值，必须传引用
	//"u"，就是-u指定参数
	//""，默认值
	//"用户名，默认为空" 文字说明
	flag.StringVar(&user, "u", "", "用户名，默认为空")
	flag.StringVar(&pwd, "pwd", "", "密码，默认为空")
	flag.StringVar(&host, "h", "localhost", "主机名，默认为localhost")
	flag.IntVar(&port, "p", 3306, "端口号，默认为3306")
	flag.Parse() //这是一个非常重要的操作，转换，必须调用该方法
	//输出结果
	fmt.Printf("user=%v,pwd=%v,host=%v,port=%v", user, pwd, host, port)
}
