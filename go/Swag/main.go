package main

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"          // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger"      // gin-swagger middleware
	_ "github.com/xieyuxuan109/project01/Swag/docs" //swag init初始化文档 最前面的docs是包的别名
)

//常见注释
// @title：文档的标题。
// @version：API的版本。
// @description：对API的详细描述。
// @termsOfService：服务条款的URL。
// @contact.name、@contact.url、@contact.email：联系信息。
// @host：API的主机地址。
// @BasePath：API的基础路径。
// @Router：指定接口的URL路径和HTTP方法。
// @Summary：接口的简要描述。
// @Param：描述接口的参数，例如查询参数、路径参数、请求体等。
// @Success：描述接口的成功响应。
// @Failure：描述接口的失败响应。

// @title 谢宇轩的API文档
//@version 1.0
//@description API文档
//@host 127.0.0.1:8080
//@BasePath:/

// 安装完swag
// 安装其他依赖
// go get -u github.com/swaggo/gin-swagger
// go get -u github.com/swaggo/files
func Index(c *gin.Context) {
	c.JSON(200, gin.H{"msg": 200})
}
func main() {
	r := gin.Default()
	r.GET("/", Index)
	// 路由方法
	// 设置访问api文档
	//url:= ginSwagger.URL("http://127.0.0.1:8080/swagger/doc.json")
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler)) //这一步很重要
	r.Run()
}
