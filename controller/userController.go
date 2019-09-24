package controller

import (
	"encoding/base64"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/tidwall/gjson"
	"log"
	"net/http"
	"reviewAchievements/model"
	"reviewAchievements/utils"
	"strings"
)

func GetToken(c *gin.Context) {
	token := c.Query("token")

	//用 . 分成 Payload和Signature 两部分
	tokenSlice := strings.Split(token, ".")
	//取 Payload并替换
	tokenPayload := strings.Replace(tokenSlice[0], "%20", "+", -1)
	//base64解码
	resultJson, err := base64.StdEncoding.DecodeString(tokenPayload)

	if err != nil {
		log.Fatalln(err)
	}

	//json解析
	jsonObject := gjson.Parse((string(resultJson)))
	redId := jsonObject.Get("redId").String()
	nickName := jsonObject.Get("nickname").String()
	headImgUrl := jsonObject.Get("headImgUrl").String()

	headImgUrlSlice := strings.Split(headImgUrl, ":")
	head := headImgUrlSlice[0] + "s:"
	head = head + headImgUrlSlice[1]

	//初始化建表
	//db,err :=model.InitDB()
	//defer db.Close()
	//if err !=nil {
	//	log.Println(err)
	//}

	user := new(model.User)
	model.DB.First(user, 1)

	//如果没有查询到之前登录的数据
	if user.RedId == "" {

		model.DB.Create(&model.User{RedId: redId, NickName: nickName, HeadImgUrl: headImgUrl})
		log.Fatalln("新插入用户信息：" + nickName)

	}

	//301重定向
	c.Redirect(http.StatusMovedPermanently, "http://localhost:3000/#/?token="+tokenSlice[0])
	fmt.Println("getToken运行结束")

}

func PersonAchievements(context *gin.Context) {

	//从 header中取出token
	token := context.Request.Header.Get("token")

	redId, _, _ := utils.GetTokenValue(token)

	fmt.Println(redId)
	user := new(model.User)

	model.DB.Where("redId = ?", redId).First(user)

	context.JSON(200, gin.H{
		"redId": user.RedId,
		"one":   user.One,
		"two":   user.Two,
		"three": user.Three,
		"four":  user.Four,
		"five":  user.Five,
		"total": user.Total,
	})

}
