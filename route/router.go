package route

import (
	"github.com/gin-gonic/gin"
	"reviewAchievements/Filter"
	"reviewAchievements/controller"
)

func LOAD(router *gin.Engine, mw ...gin.HandlerFunc) *gin.Engine {

	router.GET("/token", controller.GetToken)
	//以下的接口，都使用Authorize()中间件身份验证
	router.Use(Filter.Authorize())

	//router.POST("/userRank",controller.GetUserRank)
	router.GET("/allRank", controller.GetAllRank)
	//router.POST("/copyInfo",controller.CopyInfo)
	router.GET("/addTotal", controller.AddGrades)
	router.POST("/pass", controller.UpdateLevelTime)
	router.GET("/getToken", controller.PersonAchievements)
	router.GET("/allTime", controller.PassTime)
	return router
}
