package models

// 结构体名称和首字母需大写 名称需对应 默认操作结构体的复数的数据库users
// 不然要自定义TableName函数指定操作数据库
type User struct {
	Id       int    `json:"id"`
	Username string `json:"username"` //注意如果是驼峰默认数据库中字段名有下划线_
	Age      int    `json:"age"`
	Email    string `json:"email"`   //对应email
	AddTime  int    `json:"addtime"` //add_time对应驼峰式AddTime
}

// 绑定一格tablename函数，指定操作数据表
func (User) TableName() string {
	return "user"
}

//结果
// {
//   "result": [
//     {
//       "id": 1,
//       "username": "xieyuxuan",
//       "age": 1,
//       "email": "qqq@163",
//       "addtime": "123"
//     },
//     {
//       "id": 2,
//       "username": "chenjieyan",
//       "age": 2,
//       "email": "aaa@163",
//       "addtime": "123"
//     },
//     {
//       "id": 3,
//       "username": "wangyuhan",
//       "age": 3,
//       "email": "bbb@163",
//       "addtime": "123"
//     }
//   ]
// }
