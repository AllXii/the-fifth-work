package dao

import "github.com/jinzhu/gorm"

var DB *gorm.DB

func InitMySQL()(err error){
	dsn := "root:123456@(localhost:3306)/db0?charset=utf8mb4"
	DB, err = gorm.Open("mysql", dsn)
	if err != nil {
		return
	}
	return DB.DB().Ping()
}