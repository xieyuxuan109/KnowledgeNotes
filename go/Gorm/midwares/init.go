package midwares

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func InitMidwares(c *gin.Context) {
	fmt.Println(time.Now())
	c.Set("name", "zhangsan1")
	fmt.Println(c.Request.URL)
}
func GetInfo(c *gin.Context) {
	con, _ := c.Get("name")
	fmt.Println(con)
}
