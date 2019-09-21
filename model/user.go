package model

type User struct {
	ID int `gorm:"primary_key"`
	RedId int   `gorm:"column:redId"`
	NickName string  `gorm:"column:nickName"`
	HeadImgUrl string `gorm:"column:headimgurl"`
	One string
	Two string
	Three string
	Four string
	Five string
	Total string
}
