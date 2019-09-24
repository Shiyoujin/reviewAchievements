package Filter

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Authorize() gin.HandlerFunc {
	return func(c *gin.Context) {

		//从 header中取出token
		token := c.Request.Header.Get("token") // 访问令牌

		if token != "" {
			// 验证通过，会继续访问下一个中间件
			//先测试接口，后面会增加 token的校验
			c.Next()
		} else {
			// 验证不通过，不再调用后续的函数处理
			c.Abort()
			c.JSON(http.StatusUnauthorized, gin.H{"message": "请求的token为空"})
			// return可省略, 只要前面执行Abort()就可以让后面的handler函数不再执行
			return
		}
	}
}
