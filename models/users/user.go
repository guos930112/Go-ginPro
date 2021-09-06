package users

import (
	"github.com/jinzhu/gorm"
)

// 创建用户数据表的结构体
type User struct {
	gorm.Model
	Name string `gorm:"type:varchar(50);not null"`
	Tel  string `gorm:"varchar(11);not null;unique"`
	Pwd  string `gorm:"size:255;not null"`
}