package model

import (
	"log"
)

type UserRank struct {
	OpenId string `gorm:"column:openId"`
	Total  float64
	Rank   int
}

//type Total struct {
//	Totals int64
//}

// 获取用户闯关后排名
func InsertTotal(openId string) (float64, int) {
	var recordTotal float64
	var userInfo UserRank
	err := DB.Raw("SELECT (one*1.333+two*1.265+three*1.125+four*0.999+five*0.75) as totals FROM users u where u.openId = ?", openId).Row().Scan(&recordTotal)
	if err != nil {
		log.Println("fail to insert total", err)
	}
	DB.Table("users").Where("openId = ?", openId).Update("total", recordTotal)

	//_ = DB.Raw("select count(*) from users u where u.total < ?", recordTotal.Totals).Row().Scan(&count)
	rows, rankErr := DB.Raw("SELECT * FROM ( SELECT openId, total, @curRank := @curRank + 1 AS rank FROM users u, (SELECT @curRank := 0) r where total > 0 ORDER BY total ) a WHERE a.openId = ?", openId).Rows()

	if rankErr != nil {
		log.Println("fail to get user rank: ", err)
	}
	defer rows.Close()
	for rows.Next() {
		DB.ScanRows(rows, &userInfo)
	}
	return recordTotal, userInfo.Rank

}

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
