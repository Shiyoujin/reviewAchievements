package model

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"net/url"
	"os"
)

var DB *gorm.DB


func InitDB() (*gorm.DB, error) {
	//db, err := gorm.Open("mysql", "root:@/reviewachievements?charset=utf8&parseTime=True")

	var err error
	host := url.QueryEscape(os.Getenv("REDROCK_CORE_DB_HOST"))
	port := url.QueryEscape(os.Getenv("REDROCK_CORE_DB_PORT"))
	username := url.QueryEscape(os.Getenv("REDROCK_CORE_DB_USERNAME"))
	password := url.QueryEscape(os.Getenv("REDROCK_CORE_DB_PASSWORD"))
	dbname := url.QueryEscape(os.Getenv("REDROCK_CORE_DB_DBNAME"))
	dburl := username + ":" + password + "@(" + host + ":" + port + ")/" + dbname + "?charset=utf8mb4&parseTime=True"
	//MD, err = gorm.Open("mysql", "root: @tcp(127.0.0.1:3306)/province_question?parseTime=true&charset=utf8&loc=Local")
	DB, err = gorm.Open("mysql", dburl)

	if err != nil {
		return nil, err
	}

	DB.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&User{})
	DB.Model(&User{}).AddUniqueIndex("idx_openid", "openId")
	DB.Model(&User{}).AddIndex("idx_total", "total")

	return DB, nil
}
