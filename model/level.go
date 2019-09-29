package model

import (
	"fmt"
	"log"
)

type Level struct {
	One int
	Two int
	Three int
	Four int
	Five int
}

//更新每关的时间
func UpdateLevelTime(step string, time int, openId string) (int, bool) {
	var level Level
	var recordTime int
	err := DB.Table("users").Select(step).Where("openId = ? ", openId).First(&level).Error
	if err != nil{
		log.Println(err)
	}

	//判断是第几关
	switch step {
	case "one":
		fmt.Println("one")
		recordTime = level.One
	case "two":
		fmt.Println("two")
		recordTime = level.Two
	case "three":
		fmt.Println("three")
		recordTime = level.Three
	case "four":
		fmt.Println("four")
		recordTime = level.Four
	case "five":
		fmt.Println("five")
		recordTime = level.Five
	default:
		recordTime = 0
	}

	//如果是新纪录，则修改数据库
	if time > 0 {
		if time < recordTime || recordTime == -1 {
			err := DB.Table("users").Where("openId = ? ", openId).Update(step, time).Error
			if err != nil {
				log.Println(err)
			}
			return time, true
		}
	}

	if recordTime == -1 {
		return recordTime, false
	}

	return time,false

}