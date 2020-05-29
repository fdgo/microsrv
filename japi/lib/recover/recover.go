package recover

import (
	"microservice/jzapi/basic/time_ex"
	"github.com/gin-gonic/gin"
	"runtime/debug"
	"strings"
	rtnpkg "microservice/jzapi/basic/return_pkg"
)

func Recover( name string ) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {

				DebugStack := ""
				for _, v := range strings.Split(string(debug.Stack()), "\n") {
					DebugStack += v + "<br>"
				}
				str := name + time_ex.GetCurrentTime() +"|"+ c.Request.Host+"|"+c.Request.RequestURI+"|"+c.Request.Method +"|"+  DebugStack +"|"+ c.Request.UserAgent()
				rtnpkg.GinResponse(500, 500,"系统异常，请联系管理员！","系统异常，请联系管理员！",str, c)
			}
		}()
		c.Next()
	}
}
