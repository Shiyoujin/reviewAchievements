package model

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

func InitDB() (*gorm.DB, error) {
	db, err := gorm.Open("mysql","root:@/reviewachievements?charset=utf8&parseTime=True")
	if err != nil {
		return nil, err
	}
	DB = db
	db.AutoMigrate(&User{})
	return DB, nil
}
