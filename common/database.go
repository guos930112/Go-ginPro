package common

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

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

	// 给 DB赋值
	DB = db

	return db
}

func GetDB() *gorm.DB{
	return DB
}