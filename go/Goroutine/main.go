package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

var MyMap = make(map[int]int, 10)
var Lock sync.Mutex

func test(n int) {
	sum := 1
	for i := 1; i <= n; i++ {
		sum *= i
	}
	Lock.Lock()
	MyMap[n] = sum
	Lock.Unlock()
}

// 演示go语言goroutine使用
func main() {
	//go主线程(有程序员直接称为线程或者进程)，一个go线程上，可以起多个协程，协程是轻量级的线程
	//go协程的特点
	//有独立的栈空间
	//共享程序堆空间
	//调度有用户控制
	//协程是轻量级的线程

	//goroutine调度模型
	//MPG模式
	//M：操作系统的主线程
	//P：协程执行需要的上下文
	//G：协程

	//golang默认程序运行在多个cpu上
	//runtime包
	//逻辑CPU个数是你的操作系统（如Windows、Linux）看到的、可以调度的“处理器”数量。
	//实际物理CPU个数是指你主板上插着的、实体的CPU芯片数量。
	//runtime.NumCPU()返回逻辑cpu个数
	//runtime.GOMAXPROCS(n int)可以设置同时执行的最大cpu数量 并返回先前的设置
	cpuNum := runtime.NumCPU()
	fmt.Println(cpuNum)
	//go程序默认会运行在所有逻辑cpu上，如果需要空出几个cpu需要手动设置
	num := runtime.GOMAXPROCS(6) //会返回之前设定的值
	fmt.Println(num)

	//goroutine来完成任务存在并发并行安全的问题
	//go build -race 文件名 这个可以用来显示是否存在资源竞争问题
	//此外还存在时间共存的问题，就是有可能主程序已经结束，但协程还未结束，这样协程会被迫终止

	//为什么资源竞争问题需要解决
	//协程A在自己的CPU缓存中修改了数据
	//协程B在自己的CPU缓存中读取的可能是旧值
	//缓存一致性协议（MESI）虽然能解决，但需要时间，协程间的读写顺序无法保证

	//引入全局互斥锁
	//解法一：全局变量加锁同步
	//拥到sync包 Synchronize同步
	//里面有一个mutex包，里面有lock和unlock两个方法 lock就是全局锁 mutex是互斥的意思
	for i := 0; i < 200; i++ {
		go test(i)
	}
	time.Sleep(time.Second * 10) //解决时间不统一问题,但仍不是很好的解决办法
	//但报错fatal error: concurrent map writes 存在资源竞争问题
	//这里也必须加锁，保证不会读写同时操作
	Lock.Lock()
	Lock.Lock()
	for i, v := range MyMap {
		fmt.Println(i, v)
	}
	Lock.Unlock()
	//虽然已经解决上面问题，但是解决方法并不好
	//于是出现了channal
}
