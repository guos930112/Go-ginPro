package routers

import (
	"github.com/gin-gonic/gin" // 引入需要的gin包
	"ginPro/controllers/health"
	"ginPro/controllers/users"
)

func CollectRoute(r *gin.Engine) *gin.Engine {
	// 2.绑定路由规则，执行函数
	r.GET("/", health.IndexRedirect)
	r.GET("/ping", health.HealthTest)
	r.GET("/long_async", health.LongAsync)
	r.POST("/api/auth/register", users.Register)

	return r 
}