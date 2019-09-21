package model

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"log"
	rPool "reviewAchievements/cache/redis"
	"strconv"
	"strings"
	"time"
)

type UserInfo struct {
	time     int
	redId    string
	nickName string
}

func InsertTotal(total int, redId string) (int, int) {
	var count int = -1

	err := DB.Table("users").Where("redid= ?", redId).Update("total", total).Error
	if err != nil {
		fmt.Println("insert error line[24]: ", err)
		return -1, -1
	}
	DB.Raw("select count(*) from users u where u.total < ?", total).Row().Scan(&count)

	return total, count + 1
}

func GetUserRank(redId string) (int, int) {
	var userInfo User
	var count int = -1
	err := DB.Table("users").Where("redId = ?", redId).First(&userInfo).Error
	if err != nil {
		fmt.Println("failed to get userRank line[37]: ", err)
		return -1, -1
	}
	if userInfo.Total > -1 {
		DB.Raw("select count(*) from users u where u.total < ?", userInfo.Total).Row().Scan(&count)
	}
	return userInfo.Total, count + 1
}

func CopyInfoFromDB() {
	go func() {
		//获取redis连接池
		rConn := rPool.RedisPool().Get()
		defer rConn.Close()
		for {
			rows, err := DB.Raw("select redId, nickName, time from users").Rows()
			if err != nil {
				return
			}

			userInfos := make(map[int]UserInfo)
			i := 0
			for rows.Next() {
				var userInfo UserInfo
				DB.ScanRows(rows, &userInfo)
				userInfos[i] = userInfo
				i++
			}
			rows.Close()

			for i, _ := range userInfos {

				rConn.Do("ZADD", "turn_key", userInfos[i].time, []string{userInfos[i].redId, userInfos[i].nickName}, )
			}

			now := time.Now()
			next := now.Add(time.Hour * 24)
			//每天凌晨2点更新缓存
			next = time.Date(next.Year(), next.Month(), next.Day(), 2, 0, 0, 0, next.Location())
			//计算当前时间到次日凌晨2点的差值
			t := time.NewTimer(next.Sub(now))
			<-t.C
		}
	}()
}

//从redis获取用户排名信息
func GetAllRank() []UserInfo {
	rConn := rPool.RedisPool().Get()
	defer rConn.Close()

	dataMap, err := redis.StringMap(rConn.Do("ZRANGE", "turn_key", 0, 100, "withscores"))
	if err != nil {
		log.Println("failed to get redis: ", err)
	}

	var userInfo UserInfo
	var userInfos []UserInfo

	for key := range dataMap {

		userInfo.time, _ = strconv.Atoi(dataMap[key])
		key = strings.TrimPrefix(key, "[")
		key = strings.TrimSuffix(key, "]")
		keys := strings.Split(key, " ")

		if len(keys) > 1 {
			userInfo.redId = keys[0]
			userInfo.nickName = keys[1]
			userInfos = append(userInfos, userInfo)
		}
	}
	return userInfos

}

func B2S(bs []uint8) string {
	ba := []byte{}
	for _, b := range bs {
		ba = append(ba, byte(b))
	}
	return string(ba)
}
