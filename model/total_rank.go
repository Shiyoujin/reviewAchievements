package model

import "log"

type UserInfo struct {

	RedId string   `gorm:"column:redId"`
	NickName string  `gorm:"column:nickName"`
	Total int
}

// 获取前100名
func GetAllRank() ([]*UserInfo){
	var userInfos []*UserInfo

	rows, err :=DB.Raw("select redId, nickName, total from users order by total limit 100").Rows()
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