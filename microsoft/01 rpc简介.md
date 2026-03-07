## rpc介绍 
```
rpc是为了解决类似远程、跨内存空间的函数/方法调用，像调用本地方法一样跨服务调用
restfulAPI也可以是实现类似功能，但比较繁琐
```
```
package main
import (
    "fmt"
    "net"
    "net/rpc"
    "time"
)
// 1. 定义参数结构
type HelloArgs struct {
    Name string
}
// 2. 定义回复结构
type HelloReply struct {
    Message string
    Time    string
}
// 3. 定义服务对象
type HelloService struct{}
// 4. 实现 SayHello 方法 (就这一个函数)
func (h *HelloService) SayHello(args *HelloArgs, reply *HelloReply) error {
    reply.Message = fmt.Sprintf("你好，%s！欢迎使用RPC", args.Name)
    reply.Time = time.Now().Format("2006-01-02 15:04:05")
    return nil
}
func main() {
    // 启动服务器
    go startServer()
    // 等待服务器启动
    time.Sleep(time.Second)
    // 启动客户端
    startClient()
}
// 服务器端
func startServer() {
    // 注册服务
    service := new(HelloService)
    rpc.Register(service)
    // 监听端口
    listener, err := net.Listen("tcp", ":12345")
    if err != nil {
        fmt.Println("服务器启动失败:", err)
        return
    }
    defer listener.Close()
    fmt.Println("RPC服务器已启动，端口: 12345")
    // 接受连接
    rpc.Accept(listener)
}
// 客户端
func startClient() {
    // 连接服务器
    client, err := rpc.Dial("tcp", "localhost:12345")
    if err != nil {
        fmt.Println("连接服务器失败:", err)
        return
    }
    defer client.Close()
    // 准备参数
    args := &HelloArgs{Name: "小明"}
    var reply HelloReply
    // 调用远程函数 SayHello
    err = client.Call("HelloService.SayHello", args, &reply)
    if err != nil {
        fmt.Println("调用失败:", err)
        return
    }
    // 打印结果
    fmt.Println("服务器回复:")
    fmt.Println("消息:", reply.Message)
    fmt.Println("时间:", reply.Time)
}
```
