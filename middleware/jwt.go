package middleware

import (
	"Termbin/util"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// JWTAuth jwt 鉴权
func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")

		// token为空
		if token == "" {
			c.Next()
			return
		}

		claims, err := util.ParseToken(token)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"err": "invalid token",
			})
			c.Abort()
			return
		}
		if claims.ExpiresAt < time.Now().Unix() {
			c.JSON(http.StatusOK, gin.H{
				"err": "expired token",
			})
			c.Abort()
			return
		}

		// jwt 验证通过的处理
		c.Set("UserID", claims.UserID)
		c.Set("UserEmail", claims.UserEmail)

		c.Next()
	}
}
