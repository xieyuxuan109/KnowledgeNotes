package main

import "fmt"

// 演示排序算法和查找算法
func Sort(num *[10]int) {
	for i := 0; i < len(num); i++ {
		for j := 0; j < len(num)-i-1; j++ {
			if num[j] > num[j+1] {
				temp := num[j+1]
				num[j+1] = num[j]
				num[j] = temp
			}
		}
	}
}
func Search() {
	var str = []string{"谢宇轩", "石云娜", "完"}
	var str2 string
	fmt.Scanln(&str2)
	var index int = -1
	for i := 0; i < len(str); i++ {
		if str2 == str[i] {
			index = i
			fmt.Printf("%d,%s", index, str[index])
		}
	}
	if index == -1 {
		fmt.Println("没找到")
	}
}

// 也可以递归实现
func MinSearch(arr [10]int, num int) {
	left := 0
	right := len(arr) - 1
	for right >= left {
		mid := (right - left) / 2
		if arr[mid] == num {
			fmt.Println("找到了")
			fmt.Println(mid)
			return
		}
		if arr[mid] > num {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
}
func main() {
	//排序分类
	//内部排序法：
	//指的是将需要处理的所有数据都加载到内部贮存器中进行排序(交换式排序法，选择式排序法，插入式排序法)
	//外部排序法：
	//指的是数据量过大无法全部加载到内存中，需要借助外部贮存进行排序（合并排序法和直接合并排序法）

	//交换排序 冒泡 快速排序
	var num = [10]int{10, 18, 60, 34, 20, 10, 48, 28, 87, 63}
	Sort(&num)
	fmt.Println(num)

	//查找算法
	//顺序查找
	//Search()
	MinSearch(num, 10)
}
