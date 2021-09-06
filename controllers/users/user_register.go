package users

import (
	"github.com/gin-gonic/gin" // 引入需要的gin包
	"net/http"
	"log"
	"ginPro/utils"
	"ginPro/models/users"
	"ginPro/common"
)

func Register(ctx *gin.Context) {
	DB = common.GetDB()
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
	if isTelExist(DB, tel) {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "message": "手机号已存在"})
		return
	}
	// 创建用户
	newUser := User{
		Name: name,
		Tel: tel,
		Pwd: pwd,
	}
	DB.Create(&newUser)
	
	// 返回结果
	ctx.JSON(200, gin.H{
		"message": "注册成功",
	})
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