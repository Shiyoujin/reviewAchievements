package route

import (
	"github.com/gin-gonic/gin"
	"reviewAchievements/controller"
)

func LOAD(router *gin.Engine, mw ...gin.HandlerFunc) *gin.Engine {


	//router.POST("/userRank",controller.GetUserRank)
	router.POST("/allRank",controller.GetAllRank)
    //router.POST("/copyInfo",controller.CopyInfo)
	router.POST("/addTotal",controller.AddGrades)
	router.POST("/pass",controller.UpdateLevelTime)

	return router
}