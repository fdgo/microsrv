package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/yangwenmai/ratelimit/simpleratelimit"
	"net/http"
	"time"
)

var (
	rl = simpleratelimit.New(10000, time.Minute)
)

// 中间件，用令牌桶限制请求频率
func LimitHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		if rl.Limit() {
			c.JSON(
				http.StatusOK,
				gin.H{
					"code":    9999,
					"message": "请求频率太高",
				},
			)
			c.Abort()
			return
		}
		c.Next()
	}
}
