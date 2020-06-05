package trace

import (
	token "ds_server/support/utils/auth"
	"ds_server/support/utils/cors"
	"ds_server/support/utils/limit"
	"ds_server/support/utils/logex"
	rsp "ds_server/support/utils/rsp"
	time_ex "ds_server/support/utils/timex"
	"ds_server/support/utils/trace"
	"github.com/gin-gonic/gin"
	"net/http"
	"runtime/debug"
	"strings"
	"time"
)

func Log() gin.HandlerFunc {
	return logex.GinLogger()
}

func TracerWrapper(c *gin.Context) {
	trace.TracerWrapper(c)
}

func Cors() gin.HandlerFunc {
	return cors.Cors()
}

func NoRoute() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if ctx.Request.Method == "OPTIONS" {
			ctx.JSON(200, nil)
		}
	}
}
func Recover(name string) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {

				DebugStack := ""
				for _, v := range strings.Split(string(debug.Stack()), "\n") {
					DebugStack += v + "<br>"
				}
				str := name + time_ex.GetCurrentTime() + "|" + c.Request.Host + "|" + c.Request.RequestURI + "|" + c.Request.Method + "|" + DebugStack + "|" + c.Request.UserAgent()
				rtnpkg.GinResponse(500, 500, "系统异常，请联系管理员！", "系统异常，请联系管理员！", str, c)
			}
		}()
		c.Next()
	}
}

// 中间件，用令牌桶限制请求频率
func Limit(c *gin.Context) {
	limit.LimitHandler(c)
}

func Auth(token *token.JwtToken) gin.HandlerFunc {
	return func(c *gin.Context) {
		istokenok, msg, sub := token.Decode(c.Request, c.Writer)
		if !istokenok {
			rsp.RespGin(400, 400, " 请先登录!", msg, "nil", c)
			c.Abort()
			return
		}
		c.Request.Header.Set("X-Head-Uuid", sub.Uuid)
		c.Request.Header.Set("X-Head-Mobile", sub.Mobile)
		c.Request.Header.Set("X-Head-UserName", sub.UserName)
		c.Request.Header.Set("X-Head-InvCodeAgent", sub.InvCodeAgent)
		c.Request.Header.Set("X-Head-InvCodeSelf", sub.InvCodeSelf)
		c.Request.Header.Set("X-Head-TimeStamp", time_ex.TimeStampToTimeStr(sub.ExpiresAt))
		c.Next()
	}
}
