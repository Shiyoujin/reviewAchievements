package main

import (
	"log"
	"reviewAchievements/model"
)

func main() {
	db, err := model.InitDB()
	if err != nil {
		log.Println("err open databases", err)
		return
	}
	defer db.Close()
	model.Query(7)
}
