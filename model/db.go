package model

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

func InitDB() (*gorm.DB, error) {
	db, err := gorm.Open("mysql", "root:123456@/reviewachievements?charset=utf8&parseTime=True")
	if err != nil {
		return nil, err
	}

	DB = db

	db.AutoMigrate(&User{})
	db.Model(&User{}).AddUniqueIndex("idx_redid","redId")
	db.Model(&User{}).AddIndex("idx_total","total")

	return DB, nil
}
