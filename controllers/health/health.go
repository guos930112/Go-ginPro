package health

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin" // 引入需要的gin包
)

func HealthTest(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func IndexRedirect(c *gin.Context) {
	c.Redirect(http.StatusMovedPermanently, "https://www.baidu.com")
}

func LongAsync(c *gin.Context) {
	// 异步执行 需要用的是上下文的副本
	copyContext := c.Copy()
	// 异步处理
	go func() {
		time.Sleep(3 * time.Second)
		log.Println("异步执行: " + copyContext.Request.URL.Path)
	}()
}

// 定义中间件
func MiddleWare() gin.HandleFunc {
	return func(c *gin.Context) {
		t := time.Now()
		log.Panicln("中间件开始执行了...")
		// 设置变量到 Context的key中，可以通过Get() 获取
		c.Set("request", "中间件")
		status := c.Writer.Status()
		log.Println("中间件执行完毕", status)
		t2 := time.Since(t)
		log.Println("time keep:", t2)
	}
}