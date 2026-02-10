package main

import (
	"fmt"
	"sync"
)

// 使用sync包来实现主线程和协程的同步
var MyMap = make(map[int]int, 10)
var wg sync.WaitGroup
var lock sync.Mutex // 🔴 修正1：改为具体类型，不是接口

func test(n int) {
	defer wg.Done() // ✅ 最佳实践：确保Done()被调用

	sum := 1                  // 🔴 修正2：阶乘必须从1开始（0!=1）
	for i := 1; i <= n; i++ { // 🔴 修正3：计算n的阶乘，不是固定200
		sum *= i
	}

	lock.Lock()
	MyMap[n] = sum
	lock.Unlock()
}

func main() {
	for i := 1; i <= 20; i++ { // ✅ 建议：不要用200，阶乘会溢出
		wg.Add(1)
		go test(i)
	}
	wg.Wait()

	// 读取时也需要加锁（安全考虑）
	lock.Lock()
	for i, v := range MyMap {
		fmt.Printf("%d! = %d\n", i, v)
	}
	lock.Unlock()
}
