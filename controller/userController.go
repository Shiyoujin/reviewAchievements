package controller

import (
	"encoding/base64"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/tidwall/gjson"
	"log"
	"reviewAchievements/model"
	"strings"
)

func GetToken(c *gin.Context) {

	token := c.Query("token")

	fmt.Println(token)
	//用 . 分成 Payload和Signature 两部分
	tokenSlice := strings.Split(token, ".")
	//取 Payload并替换
	tokenPayload := strings.Replace(tokenSlice[0], " ", "+", -1)
	fmt.Println(tokenPayload)
	//base64解码
	resultJson, err := base64.StdEncoding.DecodeString(tokenPayload)

	if err != nil {
		log.Println(err)
		fmt.Println("base64解码出错")
	}

	//json解析
	jsonObject := gjson.Parse((string(resultJson)))
	fmt.Println(jsonObject)
	redId := jsonObject.Get("redId").String()
	openId := jsonObject.Get("openid").String()
	nickName := jsonObject.Get("nickname").String()
	headImgUrl := jsonObject.Get("headImgUrl").String()

	headImgUrlSlice := strings.Split(headImgUrl, ":")
	head := headImgUrlSlice[0] + "s:"
	fmt.Println("------", headImgUrlSlice)
	head = head + headImgUrlSlice[1]

	fmt.Println(redId, nickName, head, openId)


	user := new(model.User)
	model.DB.Where("openId = ?", openId).First(user)

	fmt.Println("openId + " + user.OpenId)
	//如果没有查询到之前登录的数据
	if user.OpenId == "" {

		err := model.DB.Create(&model.User{RedId: redId, OpenId: openId, NickName: nickName, HeadImgUrl: headImgUrl}).Error
		if err != nil {
			log.Println(err)
		}

	}

	//301重定向，永久跳转 http.StatusMovedPermanently
	//但这里是 302
	c.Redirect(302, "https://upred.atowerlight.cn/tesgs/#/?token="+tokenSlice[0])
	fmt.Println("getToken运行结束")

}

//func PersonAchievements(context *gin.Context) {
//
//	//从 header中取出token
//	token := context.Request.Header.Get("token")
//
//	redId, _, _ := utils.GetTokenValue(token)
//
//	fmt.Println(redId)
//	user := new(model.User)
//
//	model.DB.Where("redId = ?", redId).First(user)
//
//	context.JSON(200, gin.H{
//		"redId": user.RedId,
//		"one":   user.One,
//		"two":   user.Two,
//		"three": user.Three,
//		"four":  user.Four,
//		"five":  user.Five,
//		"total": user.Total,
//	})
//}
