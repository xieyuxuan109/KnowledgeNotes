package main

import (
	"encoding/json"
	"fmt"
)

// 演示结构体的使用
type Stu struct {
	Name  string `json:"name"`
	ID    int    `json:"id"`
	Score []int  `json:"score"`
}
type Stu1 struct {
	Name string
	ID   int
}
type Pupil struct {
	Name  string
	Age   int
	Score int
}

func (pupil *Pupil) ShowInfo() {
	fmt.Printf("姓名:%v 年龄:%v 成绩:%v\n", pupil.Name, pupil.Age, pupil.Score)
}
func (pupil *Pupil) SetScore(score int) {
	pupil.Score = score
}
func (pupil *Pupil) GetScore() int {
	return pupil.Score
}

type Collage struct {
	Pupil //嵌套匿名结构体
	//也可以嵌套结构体指针
}

func main() {
	//结构体是自定义的数据类型
	//结构体是一个完整的个体
	//由字段组成
	//结构体字段如果有引用类型需要先make过后才能使用
	//结构体是值类型
	//结构体使用方式一
	var stu Stu
	stu.Score = make([]int, 10)
	stu.Score[1] = 1
	stu.Score[2] = 4
	stu.Score[3] = 7
	stu.Score[4] = 9
	stu.Score[5] = 0
	stu.Score[6] = 1
	stu.Score[7] = 1
	fmt.Println(stu.Score)

	//方式二
	//类似map 边创建边赋值
	stu1 := Stu1{"xieyuxuan", 1}
	fmt.Println(stu1)

	//也可以使用new，但是返回的是指针，方式三
	p3 := new(Stu)
	p3.ID = 1 //等价于(*p3).ID，go设计者专门设计的,但是标准还是用(*p3).ID

	//方式四
	var stu2 *Stu1 = &Stu1{}
	stu2.Name = "Jasper"
	stu2.ID = 2
	fmt.Println(stu2)
	//或者
	var stu3 *Stu1 = &Stu1{"jasper", 1}
	fmt.Println(stu3)

	//结构体使用注意事项
	//结构体字段内存连续分布
	//结构体之间也相互转换的话，结构体必须每个字段名完全一样(名字，个数，类型)
	//结构体名字被type重新定义后，golang认为新的数据类型，但是相互间可以强转
	//序列化和反序列化 json tag
	stu4 := Stu1{"xieyuxaun", 25}
	jsonStr, _ := json.Marshal(stu4) //使用到了反射，序列化
	fmt.Println(jsonStr)
	fmt.Println("jsonStr", string(jsonStr)) //反序列化

	//面向对象编程的三大特性
	//继承，封装，多态
	//封装：封装指的是把抽象出的字段和对字段的操作封装在一起，数据被保护在内部。程序其他包只有通过授权的操作(方法)，才能对字段进行操作
	//封装好处
	//可以隐藏实现细节，可以对数据进行验证，保证安全合理Get Set方法

	//如何体现封装
	//对结构体属性进行封装
	//通过方法，包实现封装

	// 继承
	//使用嵌套匿名结构体实现继承
	//嵌入匿名结构体后使用方法会发生变化
	//提高代码复用性
	//继承后被继承的结构体大写小写字段名和方法名也可以使用
	col := &Collage{}
	col.Pupil.Name = "123" //可以简化 col.Name
	col.Pupil.Age = 19     //但是当继承的结构体和自己本身字段名或方法一样时候，需要具体指明，不然会就近原则
	col.Pupil.Score = 100
	fmt.Println(col.Pupil.GetScore())
	//但是当一个结构体嵌入两个结构体时候，两个结构体含有相同的字段和方法但是结构体本身没有一样名字的字段和方法
	//访问时候必须指定匿名结构体名字
	//继承结构体可以边定义边赋值，但是用{}来赋值
	col2 := &Collage{Pupil{"123", 1, 1}}
	fmt.Println(col2)
	//组合
	//嵌套有名结构体就是组合
	//组合必须使用具体路径，不能简写，带上有名结构体名字

}
