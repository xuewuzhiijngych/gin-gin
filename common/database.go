package common

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"ych.hub/model"
)

var DB *gorm.DB

// InitDb 初始化数据库连接池
func InitDb() (db *gorm.DB) {
	dirverName := "mysql"
	host := "127.0.0.1"
	port := "3306"
	database := "gintest"
	username := "root"
	password := "123456"
	charset := "utf8mb4"

	args := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true",
		username,
		password,
		host,
		port,
		database,
		charset,
	)

	db, err := gorm.Open(dirverName, args)
	if err != nil {
		panic("数据库连接失败" + err.Error())
	}
	db.AutoMigrate(&model.User{})
	DB = db
	return
}

// 获取 db实例
func GetDB() *gorm.DB {
	return DB
}
