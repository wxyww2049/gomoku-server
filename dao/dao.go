package dao

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Setup() {
	db, err := gorm.Open(mysql.Open(fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local",
		"majiangqun", "damajiang", "101.42.20.21", 3306, "five_son_chess")), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	DB = db
	//db.AutoMigrate(&po.User{})
}
