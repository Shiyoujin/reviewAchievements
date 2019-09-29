package model

import "log"

type UserInfo struct {
	OpenId   string `gorm:"column:openId"`
	NickName string `gorm:"column:nickName"`
	Total    float64
}

// 获取前100名
func GetAllRank() ([]*UserInfo){
	var userInfos []*UserInfo

	rows, err := DB.Raw("select openId, nickName, total from users where total > 0 order by total limit 100").Rows()
	if err != nil {
		log.Println(err)
	}

	defer rows.Close()

	for rows.Next() {
		var userInfo UserInfo
		DB.ScanRows(rows,&userInfo)
		userInfos = append(userInfos,&userInfo)
	}

	return userInfos
}