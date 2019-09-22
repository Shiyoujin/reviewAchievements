package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"reviewAchievements/model"
	"reviewAchievements/utils"
	"strconv"
)

func UpdateLevelTime(c *gin.Context) {
	token := c.Request.Header.Get("token")
	redId, _, _ := utils.GetTokenValue(token)
	step := c.Request.FormValue("step")
	timeStr := c.Request.FormValue("time")
	time, _ := strconv.Atoi(timeStr)
	usrTime, isRecord := model.UpdateLevelTime(step, time, redId)
	c.JSON(http.StatusOK,gin.H{
		"time": usrTime,
		"isRecord": isRecord,
	})
}
