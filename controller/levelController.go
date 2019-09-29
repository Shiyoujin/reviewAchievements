package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"reviewAchievements/model"
	"reviewAchievements/utils"
	"strconv"
)

func UpdateLevelTime(c *gin.Context) {
	fmt.Println("enen")
	token := c.Request.Header.Get("token")
	openId, _, _ := utils.GetTokenValue(token)
	step := c.Request.FormValue("step")
	timeStr := c.Request.FormValue("time")
	time, _ := strconv.Atoi(timeStr)
	usrTime, isRecord := model.UpdateLevelTime(step, time, openId)
	c.JSON(http.StatusOK,gin.H{
		"time": usrTime,
		"isRecord": isRecord,
	})
}
