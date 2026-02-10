package main

import (
	"github.com/gin-gonic/gin"
	"github.com/xieyuxuan109/project01/Gorm/routers"
)

func main() {
	r := gin.Default()
	routers.AdminRoutersInit(r)
	r.Run(":8080")
}
