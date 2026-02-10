package model

type stu struct { //小写
	Name  string
	score float64 //小写
}

// 因为stu结构体首字母小写，是私有的，只能在model包使用
// 我们通过工厂模式解决
func NewStu(n string, s float64) *stu {
	return &stu{
		n,
		s,
	}
}
func (stu *stu) GetScore() float64 {
	return stu.score
}
