package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/xieyuxuan109/project01/Gorm/controller"
	"github.com/xieyuxuan109/project01/Gorm/midwares"
)

func AdminRoutersInit(r *gin.Engine) {
	url := r.Group("/admin", midwares.InitMidwares, midwares.GetInfo)
	{
		url.GET("/user", controller.UserController{}.SearchAll)
		url.GET("/add", controller.UserController{}.Add)
		url.GET("/edit", controller.UserController{}.Edit)
		url.GET("/delete", controller.UserController{}.Delete)
	}
}
