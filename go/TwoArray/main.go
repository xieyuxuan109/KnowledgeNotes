package main

import (
	"fmt"
)

// 演示二维数组
func main() {
	//定义
	var arr [4][6]int //4行6列
	//其他定义方式和一维数组类似，但列下标不能省略
	//赋值
	arr[0][1] = 1
	for i := 0; i < 4; i++ {
		for j := 0; j < 6; j++ {
			fmt.Print(arr[i][j]) //print就是比println少一个换行
		}
		fmt.Println()
	}

	//遍历 for-range遍历
	for _, v := range arr {
		for _, v2 := range v {
			fmt.Print(v2)
		}
		fmt.Println()
	}
}
