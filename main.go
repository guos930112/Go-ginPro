package main

import(
	"fmt"
	"github.com/gin-gonic/gin" // 引入需要的gin包
	"github.com/jinzhu/gorm"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
	"math/rand"
	"log"
	"time"
)

// 创建用户数据表的结构体
type User struct {
	gorm.Model
	Name string `gorm:"type:varchar(50);not null"`
	Tel  string `gorm:"varchar(11);not null;unique"`
	Pwd  string `gorm:"size:255;not null"`
}

// gin 框架学习
func main() {
	db := InitDB()
	defer db.Close()

	r := gin.Default()
	// 根目录下的接口返回
	// 测试接口
	r.GET("/ping", func(c *gin.Context){
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	// 增加用户注册的接口
	r.POST("/api/auth/register", func(ctx *gin.Context){
		// 获取参数
		name := ctx.PostForm("name")
		tel := ctx.PostForm("tel")
		pwd := ctx.PostForm("pwd")
		// 验证参数
		if len(tel) != 11 {
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "message": "手机号必须为11位"})
			return 
		}
		if len(pwd) < 6 {
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "message": "密码不能少于6位"})
			return 
		}
		// 如果名称没有传，给一个10位的随机字符串
		if len(name) == 0 {
			name = RandomString(10)
		}
		log.Println("name:", name, "tel=", tel, "pwd=", pwd)
 		// 判断手机号是否存在 需要查库，这里使用gorm来连接数据库 需要下载 go get -u github.com/jinzhu/gorm
		// 同时还需要安装 mysql驱动 go get github.com/go-sql-driver/mysql 
		if isTelExist(db, tel) {
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "message": "手机号已存在"})
			return
		}
		// 创建用户
		newUser := User{
			Name: name,
			Tel: tel,
			Pwd: pwd,
		}
		db.Create(&newUser)
		// 返回结果
		ctx.JSON(200, gin.H{
			"message": "注册成功",
		})
	})
	fmt.Println("start studying go web framework Gin....")
	panic(r.Run())
}


// 定义一个根据数值返回一个随机长度相等于数值的字符串
func RandomString(n int) string {
	var letters = []byte("abcdefghijklmnopqrstuvwsyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	res := make([]byte, n)
	rand.Seed(time.Now().Unix())
	for i := range res {
		res[i] = letters[rand.Intn(len(letters))]
	}
	return string(res)
}

// 连接数据库
func InitDB() *gorm.DB {
	driverName := "mysql"
	host := "localhost"
	port := "3306"
	database := "ginPro_db"
	username := "root"
	password := "root"
	charset := "utf8"
	args := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true", 
		username,
		password,
		host,
		port,
		database,
		charset)

	db, err := gorm.Open(driverName, args)
	if err != nil {
		panic("failed to connect database err:" + err.Error())
	}
	// 可以自动创建数据表
	db.AutoMigrate(&User{})

	return db
}

// 判断手机号是否存在 因为是唯一约束
func isTelExist(db *gorm.DB, tel string) bool {
	var user User 
	db.Where("tel = ?", tel).First(&user)
	if user.ID != 0 {
		return true 
	}
	return false
}