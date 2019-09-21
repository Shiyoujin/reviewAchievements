package controller

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"reviewAchievements/model"
	"reviewAchievements/utils"
	"strconv"
)

func GetAllRank(c *gin.Context) {
	bankInfos := model.GetAllRank()
	bankInfos_Json, err := json.Marshal(bankInfos)
	if err != nil {
		fmt.Println("Marshal json error: ", err)
	}
	c.Data(http.StatusOK, "application/json", bankInfos_Json)
}

func InsertController(c *gin.Context) {
	token := c.Request.Header.Get("token")
	redId, _, _ := utils.GetTokenValue(token)
	totalTimeStr := c.Request.FormValue("total")
	totalTime, _ := strconv.Atoi(totalTimeStr)
	total, rank := model.InsertTotal(totalTime, redId)
	c.JSON(http.StatusOK, gin.H{
		"totalTime": total,
		"rank": rank,
	})
}

func GetUserRank(c *gin.Context) {
	token := c.Request.Header.Get("token")
	redId, _, _ := utils.GetTokenValue(token)
	total, rank := model.GetUserRank(redId)
	c.JSON(http.StatusOK, gin.H{
		"totalTime": total,
		"rank": rank,
	})
}

//定时将mysql数据写入redis
func CopyInfo(c *gin.Context)  {
	model.CopyInfoFromDB()
	c.JSON(http.StatusOK, gin.H{
		"message": "success",
	})
}
