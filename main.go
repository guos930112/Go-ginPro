package main

import (
	"fmt"
	"ginPro/routers"

	"ginPro/common"
	"github.com/gin-gonic/gin" // 引入需要的gin包
	_ "github.com/go-sql-driver/mysql"
)


// gin 框架学习
func main() {
	db := common.InitDB()
	defer db.Close()

	// 1.创建路由
	// 默认 使用了2个中间件 Logger() Recovery()
	r := gin.Default()
	// 注册中间件
	// r.Use(MiddleWare())
	// r.Use(LoggerByTime()) // 日志
	// r.Use(Cors())  // 跨域
	r = routers.CollectRoute(r)

	fmt.Println("start studying go web framework Gin....")

	// 3.监听端口，默认是8080
	panic(r.Run(":7077"))
}
