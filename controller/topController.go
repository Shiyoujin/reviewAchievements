package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"reviewAchievements/model"
	"reviewAchievements/utils"
)

func GetAllRank(c *gin.Context) {
	rankInfos := model.GetAllRank()
	c.JSON(http.StatusOK,rankInfos)
}

func AddGrades(c *gin.Context) {
	token := c.Request.Header.Get("token")
	redId, _, _ := utils.GetTokenValue(token)
	//totalTimeStr := c.Request.FormValue("total")
	//totalTime, _ := strconv.Atoi(totalTimeStr)
	total, rank := model.InsertTotal(redId)
	c.JSON(http.StatusOK, gin.H{
		"totalTime": total,
		"rank": rank,
	})
}

//func GetUserRank(c *gin.Context) {
//	token := c.Request.Header.Get("token")
//	redId, _, _ := utils.GetTokenValue(token)
//	total, rank := model.GetUserRank(redId)
//	c.JSON(http.StatusOK, gin.H{
//		"totalTime": total,
//		"rank": rank,
//	})
//}

////定时将mysql数据写入redis
//func CopyInfo(c *gin.Context)  {
//	model.CopyInfoFromDB()
//	c.JSON(http.StatusOK, gin.H{
//		"message": "success",
//	})
//}
