package model

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var DB *gorm.DB


func InitDB() (*gorm.DB, error) {
	DB, err := gorm.Open("mysql", "root:zhy123@/file-server?charset=utf8&parseTime=True")
	if err != nil {
		return nil, err
	}
	DB.AutoMigrate(&User{})
	return DB, nil
}