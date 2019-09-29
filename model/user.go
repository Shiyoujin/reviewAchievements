package model

import "log"

type User struct {
	ID         int     `gorm:"primary_key"`
	OpenId     string  `gorm:"column:openId"`
	RedId      string  `gorm:"column:redId"`
	NickName   string  `gorm:"column:nickName"`
	HeadImgUrl string  `gorm:"column:headimgurl"`
	One        int     `gorm:"default:'-1'"`
	Two        int     `gorm:"default:'-1'"`
	Three      int     `gorm:"default:'-1'"`
	Four       int     `gorm:"default:'-1'"`
	Five       int     `gorm:"default:'-1'"`
	Total      float64 `gorm:"default:'-1'"`
}

type Pass struct {
	One   int `gorm:"default:'-1'"`
	Two   int `gorm:"default:'-1'"`
	Three int `gorm:"default:'-1'"`
	Four  int `gorm:"default:'-1'"`
	Five  int `gorm:"default:'-1'"`
}

func PassTime(openId string) Pass {
	var userInfo Pass
	err := DB.Raw("select one, two , three, four , five from users where openId = ?", openId).Scan(&userInfo).Error
	if err != nil {
		log.Println("fail to get pass time", err)
	}
	return userInfo
}
