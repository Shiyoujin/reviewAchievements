package utils

import (
	"encoding/base64"
	"github.com/tidwall/gjson"
	"log"
	"strings"
)

func GetTokenValue(token string) (string, string, string) {

	//用 . 分成 Payload和Signature 两部分
	tokenSlice := strings.Split(token, ".")

	//base64解码
	resultJson, err := base64.StdEncoding.DecodeString(tokenSlice[0])
	if err != nil {
		log.Println(err)
	}
	//json解析
	jsonObject := gjson.Parse((string(resultJson)))
	redId := jsonObject.Get("redId").String()
	nickName := jsonObject.Get("nickname").String()
	headImgUrl := jsonObject.Get("headImgUrl").String()

	return redId, nickName, headImgUrl
}
