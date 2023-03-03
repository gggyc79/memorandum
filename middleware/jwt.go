package middleware

import (
	"beiwanglu/pkg/util"
	"time"

	"github.com/gin-gonic/gin"
)

// JWT token验证中间件
func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		code = 200
		token := c.GetHeader("Authorization")
		if token == "" {
			code = 404
		} else {
			claims, err := util.ParseToken(token)
			if err != nil {
				code = 400
			} else if time.Now().Unix() > claims.ExpiresAt {
				code = 400
			}
		}
		if code != 200 {
			c.JSON(400, gin.H{
				"status": code,
				"msg":    "失败",
				"data":   "可能是身份过期了，请重新登录",
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
