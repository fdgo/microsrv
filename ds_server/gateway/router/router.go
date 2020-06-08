package router

import (
	midware "ds_server/gateway/middleware"
	"ds_server/gateway/service/base"
	"ds_server/gateway/service/user"
	token "ds_server/support/utils/auth"
	"ds_server/support/utils/constex"
	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	router.Use(gin.Recovery())
	router.Use(midware.Log())
	router.Use(midware.TracerWrapper)
	router.Use(midware.Cors())
	router.Use(midware.NoRoute())
	router.Use(midware.Limit)
	router.Use(midware.Recover("gateway"))

	rf := router.Group("/api/v1")
	rc := rf.Group("base")
	{
		rc.POST("/mobilecode", base.VfCode)
	}
	/************************************************************************/
	ru := rf.Group("/user")
	{
		ru.GET("/ws", user.WebsocketMsg)
		ru.POST("/regist", user.Regist) //ContextWithSpan
		rus := ru.Use(midware.Auth(&token.JwtToken{
			SigningKey: []byte(constex.JwtCfg.SecretKey),
		}))
		rus.GET("/info", user.GetMemberUserAgent)
	}
	return router
}
