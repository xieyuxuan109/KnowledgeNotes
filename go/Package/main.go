package main

import (
	"fmt"

	util "github.com/xieyuxuan109/project01/ValueTypesAndReferenceTypes/model" //util为自己给包取的别名
	//去完别名之前的名字就不可用了
)

// package演示
func main() {
	//package旨在解决分类，避免臃肿和冲突
	//包的本质就是一个文件夹
	//go的每个文件都是属于一个包 go通过包来管理文件项目
	//打包方法 package + 包名
	//引入方法 import + 路径
	//大写函数和变量可以跨包使用，小写只能本包使用
	res := util.Cal(1, 99, '+')
	fmt.Println(res)
}
