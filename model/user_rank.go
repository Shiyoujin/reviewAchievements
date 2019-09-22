package model

import (
	"log"
)

type UserRank struct {
	RedId string `gorm:"column:redId"`
	Total int
	Rank  int
}

// 获取用户闯关后排名
func InsertTotal(total int, redId string) (int, int) {
	var count int = -1
	var recordTotal UserRank
	err := DB.Table("users").Select("total").Where("redId = ? ", redId).First(&recordTotal).Error
	if err != nil {
		log.Println("fail to get total time", err)
	}

	if total > 0 {
		if total < recordTotal.Total || recordTotal.Total == -1 {
			err = DB.Table("users").Where("redId= ?", redId).Update("total", total).Error
			if err != nil {
				log.Println("insert record error:", err)
				return -1, -1
			}
		}
	}

	err = DB.Raw("select count(*) from users u where u.total < ?", total).Row().Scan(&count)
	if err != nil {
		log.Println("fail to get user rank",err)
	}

	return total, count + 1
}

// 获取用户最佳排行
func GetUserRank(redId string) (int, int) {
	var userRank UserRank
	var userRanks []UserRank
	rows, err := DB.Raw("SELECT * FROM ( SELECT redId, total, @curRank := @curRank + 1 AS rank FROM users u, (SELECT @curRank := 0) r ORDER BY total ) a WHERE a.redId = ?", redId).Rows()
	if err != nil {
		log.Println("fail to get user rank: ", err)
	}
	defer rows.Close()

	for rows.Next() {
		DB.ScanRows(rows, &userRank)
		userRanks = append(userRanks, userRank)
	}
	return userRanks[0].Total, userRank.Rank
}

//func GetUserRank(redId string) (float32, int) {
//	var total Total
//	var count int = -1
//	err := DB.Table("users").Select("total").Where("redId = ?", redId).First(&total).Error
//	if err != nil {
//		log.Println("fail to get user rank:",err)
//		return -1, -1
//	}
//	if total.Total > -1 {
//		DB.Raw("select count(*) from users u where u.total < ?", total.Total).Row().Scan(&count)
//	}
//	return total.Total, count + 1
//}

//func CopyInfoFromDB() {
//
//	//获取redis连接池
//	rConn := rPool.RedisPool().Get()
//	defer rConn.Close()
//	for {
//		log.Println("schedule start")
//		rows, err := DB.Raw("select redId, nickName, total from users").Rows()
//		if err != nil {
//			return
//		}
//
//		userInfos := make(map[int]User)
//		i := 0
//		for rows.Next() {
//			var userInfo User
//			DB.ScanRows(rows, &userInfo)
//			fmt.Println(userInfo.Total)
//			userInfos[i] = userInfo
//			i++
//		}
//		rows.Close()
//
//		for i, _ := range userInfos {
//
//			rConn.Do("ZADD", "red", userInfos[i].Total, []string{userInfos[i].RedId, userInfos[i].NickName}, )
//		}
//
//		now := time.Now()
//		next := now.Add(time.Hour * 24)
//		//每天凌晨2点更新缓存
//		//next = time.Date(next.Year(), next.Month(), next.Day(), 2, 0, 0, 0, next.Location())
//		//计算当前时间到次日凌晨2点的差值
//		t := time.NewTimer(next.Sub(now))
//
//		<-t.C
//	}
//
//}
//
////从redis获取用户排名信息
//func GetAllRank() []UserInfo {
//	rConn := rPool.RedisPool().Get()
//	defer rConn.Close()
//
//	dataMap, err := redis.StringMap(rConn.Do("ZRANGE", "red", 0, 100, "withscores"))
//	if err != nil {
//		log.Println("failed to get redis: ", err)
//	}
//
//	var userInfo UserInfo
//	var userInfos []UserInfo
//
//	for key := range dataMap {
//
//		userInfo.time, _ = strconv.Atoi(dataMap[key])
//		key = strings.TrimPrefix(key, "[")
//		key = strings.TrimSuffix(key, "]")
//		keys := strings.Split(key, " ")
//
//		if len(keys) > 1 {
//			userInfo.redId = keys[0]
//			userInfo.nickName = keys[1]
//			userInfos = append(userInfos, userInfo)
//		}
//	}
//	return userInfos
//
//}
//
//func B2S(bs []uint8) string {
//	ba := []byte{}
//	for _, b := range bs {
//		ba = append(ba, byte(b))
//	}
//	return string(ba)
//}
