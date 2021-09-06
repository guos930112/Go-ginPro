package health

import (
	"github.com/gin-gonic/gin" // 引入需要的gin包
)

func HealthTest(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}