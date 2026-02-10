package controller

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/xieyuxuan109/project01/Gorm/models"
)

type UserController struct{} //一般里面会有basecontroller

func (con UserController) SearchAll(c *gin.Context) {
	//查询数据库
	UserList := []models.User{} //定义成切片，因为不止一个
	models.DB.Find(&UserList)   //find可以查询数据，里面传入结构体切片指针
	c.JSON(http.StatusOK, gin.H{
		"result": UserList,
	})
}
func (con UserController) Add(c *gin.Context) {
	//定义数据
	User := models.User{
		Id:       5,
		Username: "itying Gorm",
		Age:      22,
		Email:    "222@qq.com",
		AddTime:  int(time.Now().Unix()), // 类型转换
	}
	//传入数据 增加数据
	models.DB.Create(&User)
	fmt.Println(User)
	c.String(200, "增加数据成功")
}
func (con UserController) Edit(c *gin.Context) {
	//保存所有字段
	user := models.User{} //里面可以传入条件字段
	models.DB.Find(&user).Where("id=?", 6)
	//修改字段 更新数据
	user.Age = 100
	user.Username = "jianghan"
	user.AddTime = int(time.Now().Unix())
	models.DB.Save(&user) //保存修改
	fmt.Println(user)
	c.String(200, "修改数据成功")
}
func (con UserController) Delete(c *gin.Context) {
	//保存所有字段
	user := models.User{Id: 5} //里面可以传入条件字段
	//删除数据
	models.DB.Delete(&user) //也可以后面加上条件
	c.String(200, "删除数据成功")
}

// 查询条件
// "id in (?)",[]int{3,5,6}
//"id>? AND id<?",a,b
//"id like ?","%会%"
//id= ? or id=?,a,b
//select("名字","好处")可以指定字段名
//可以指定order("id desc").order()让id降序排序 可以写多个 如果相等 里面按第二个order排序
//limit() 获取前几条,里面传整数
//offset() 跳过前n条数据，从第n+1条开始获取
//count(&num)统计数据条数

//同时gorm支持使用原生sql语句进行操作
//models.DB.Exec("语句")
