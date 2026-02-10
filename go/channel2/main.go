
package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func Put(PutNum chan int) {
	for i := 2; i < 1000000; i++ { // 从2开始，0和1不是素数
		PutNum <- i
	}
	close(PutNum)
	wg.Done()
}

func Prime(PrimeNum chan int, PutNum chan int, ExitNum chan bool) {
	for v := range PutNum {
		isPrime := true

		// 优化：只需要检查到平方根
		for i := 2; i*i <= v; i++ {
			if v%i == 0 {
				isPrime = false
				break
			}
		}

		if isPrime {
			PrimeNum <- v
		}
	}
	ExitNum <- true
	wg.Done()
}

func PrintNum(Prime chan int) {
	for v := range Prime {
		fmt.Println(v)
	}
	wg.Done()
}

func main() {
	PutNum := make(chan int, 100)
	PrimeNum := make(chan int, 100)
	ExitNum := make(chan bool, 16)

	// 启动数据生产者
	wg.Add(1)
	go Put(PutNum)

	// 启动16个素数判断协程
	for i := 0; i < 16; i++ {
		wg.Add(1)
		go Prime(PrimeNum, PutNum, ExitNum)
	}

	// 启动结果打印协程
	wg.Add(1)
	go PrintNum(PrimeNum)

	// 监控所有判断协程完成
	go func() {
		for i := 0; i < 16; i++ { // 必须是16，与启动的协程数一致
			<-ExitNum
		}
		close(PrimeNum) // 必须关闭，否则PrintNum会永远等待
	}()

	// 等待所有协程完成
	wg.Wait()
}
