package models

// 结构体名称和首字母需大写 名称需对应 默认操作结构体的复数的数据库users
// 不然要自定义TableName函数指定操作数据库
type Article struct {
	Id     int    `json:"id"`
	Title  string `json:"title"` //注意如果是驼峰默认数据库中字段名有下划线_
	CateId int    `json:"cateid"`
	State  string `json:"state"` //对应email
}

// 绑定一格tablename函数，指定操作数据表
func (Article) TableName() string {
	return "user"
}
