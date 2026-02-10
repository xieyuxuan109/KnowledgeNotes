package main

import (
	"fmt"

	"github.com/xieyuxuan109/project01/FactoryPattern/model"
)

// 工厂模式演示
func main() {
	//想使用其他包里面小写的结构体时，使用工厂模式调用
	stu := model.NewStu("xieyuxuan", 90.9)
	fmt.Println(*stu)
	fmt.Println(stu.Name)       //注意指针也可以这样访问
	fmt.Println(stu.GetScore()) //小写字段名可以这样访问
}
