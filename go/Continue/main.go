package main

import "fmt"

// continue演示
func main() {
	//和break一样，后面可以加标签，指定跳过那一层循环
here:
	for i := 0; i < 2; i++ {
		for j := 1; j < 4; j++ {
			if j == 2 {
				continue here
			}
			fmt.Println("j=", j) //j=1 j=1
		}
	}
}
