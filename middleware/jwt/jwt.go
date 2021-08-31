package jwt

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-gin-example/pkg/e"
	"go-gin-example/pkg/util"
	"net/http"
	"time"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		var data interface{}

		code = e.SUCCESS
		token := c.Query("token")
		if token == "" {
			code = e.INVALID_PARAMS
		}else{
			claims,err := util.PraseToken(token)
			if err != nil {
				fmt.Printf("鉴权失败:%s",err)
				code = e.ERROR_AUTH_CHECK_TOKEN_FAIL
			}else if time.Now().Unix() > claims.ExpiresAt {
				code = e.ERROR_AUTH_CHECK_TOKEN_TIMEOUT
			}
		}
		if code != e.SUCCESS {
			c.JSON(http.StatusUnauthorized,gin.H{
				"code":code,
				"msg":e.GetMsg(code),
				"data":data,
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
