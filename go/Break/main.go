package main

import (
	"fmt"
)

// break示例演示
func main() {
	//为了生成一个真正的随机数，先设置一个随机的种子，但是1.20过后就不需要了
	//生成0-100中的一个随机整数,不含100,含0

	// var count int = 0
	// for count <= 99 {
	// 	n := rand.Intn(100) + 1
	// 	count++
	// 	fmt.Println(n)
	// }
	//特殊用法
	//当break出现到多层嵌套语句中 可以通过label来指明最终跳出哪一层循环
	//label名称随意

label1:
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			if j == 2 {
				break label1
			}
			fmt.Println("j=", j) //j= 0 j= 1
		}
	}

}
