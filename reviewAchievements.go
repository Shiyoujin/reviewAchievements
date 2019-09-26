package main

import (
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"os"
	"reviewAchievements/model"
	"reviewAchievements/route"
)

func main() {
	db, err := model.InitDB()
	if err != nil {
		log.Println("err open databases", err)
		return
	}
	defer db.Close()

	r := gin.New()
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
	r.Use(gin.Logger())

	r.Use(gin.Recovery())

	// 使用跨域中间件
	route.LOAD(r)
	r.Run(":8085")

}
