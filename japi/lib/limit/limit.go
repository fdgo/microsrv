package limit

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

var (
	rl = New(2, time.Second)
)
// 中间件，用令牌桶限制请求频率
func LimitHandler(c *gin.Context) {
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
