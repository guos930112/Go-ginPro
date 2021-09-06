package routers

import (
	"github.com/gin-gonic/gin" // 引入需要的gin包
	"ginPro/controllers/health"
	"ginPro/controllers/users"
)

func CollectRoute(r *gin.Engine) *gin.Engine {
	r.GET("/ping", health.HealthTest)
	r.POST("/api/auth/register", users.Register)
	return r 
}