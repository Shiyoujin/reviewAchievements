package model

type User struct {
	ID int `gorm:"primary_key"`
	RedId string   `gorm:"column:redId"`
	NickName string  `gorm:"column:nickName"`
	HeadImgUrl string `gorm:"column:headimgurl"`
	One int  `gorm:"default:'-1'"`
	Two int  `gorm:"default:'-1'"`
	Three int  `gorm:"default:'-1'"`
	Four int  `gorm:"default:'-1'"`
	Five int  `gorm:"default:'-1'"`
	Total int  `gorm:"default:'-1'"`
}
