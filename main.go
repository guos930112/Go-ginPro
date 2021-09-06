package main

import(
	"fmt"
	"github.com/gin-gonic/gin" // 引入需要的gin包
	. "ginPro/common"
	. "ginPro/routers"
	_ "github.com/go-sql-driver/mysql"
	
)

// gin 框架学习
func main() {
	db := common.InitDB()
	defer db.Close()

	r := gin.Default()
	r = CollectRoute(r)

	fmt.Println("start studying go web framework Gin....")
	panic(r.Run())
}
