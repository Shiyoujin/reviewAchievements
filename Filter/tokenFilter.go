package Filter

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//token校验和跨域
func Authorize() gin.HandlerFunc {
	return func(c *gin.Context) {

		//从 header中取出token
		token := c.Request.Header.Get("token") // 访问令牌

		method := c.Request.Method
		// 放行所有OPTIONS方法，因为有的模板是要请求两次的
		//OPTIONS 方法不验证token
		if method == "OPTIONS" {
			//将预检请求的结果缓存40分钟
			//下一次检查跨域的时间
			c.Header("Access-Control-Max-Age", "2400")
			c.Header("Access-Control-Allow-Origin", "*")
			c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token")
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, PATCH, DELETE")
			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
			c.Header("Access-Control-Allow-Credentials", "true")

			c.AbortWithStatus(http.StatusNoContent)
			c.Next()
		} else if token != "" {
			//将预检请求的结果缓存40分钟
			//下一次检查跨域的时间
			c.Header("Access-Control-Max-Age", "2400")
			c.Header("Access-Control-Allow-Origin", "*")
			// 验证通过，会继续访问下一个中间件
			//先测试接口，后面会增加 token的校验
			c.Header("Access-Control-Allow-Origin", "*")
			c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token")
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, PATCH, DELETE")
			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
			c.Header("Access-Control-Allow-Credentials", "true")
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
