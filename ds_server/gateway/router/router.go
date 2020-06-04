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
		ru.POST("/login", user.Login)
		ru.POST("/memclass/set", user.MemerClassSet)
		ru.POST("/agentclass/set", user.AgentClassSet)
		ru.GET("/con", user.ConnectUs)
		ru.POST("/loginpwd/modify", user.ModifyLoginPwd)//修改登陆密码
		ru.POST("/basicpwd/modify", user.ForgetPwd)//忘记密码
		rus := ru.Use(midware.Auth(&token.JwtToken{
			SigningKey: []byte(constex.JwtCfg.SecretKey),
		}))
		rus.POST("/paypwd/set", user.SetPaypwd)
		rus.POST("/paypwd/modify", user.ModifyPayPwd)
		rus.GET("/info", user.GetMemberUserAgent)
		rus.POST("/usdtRechargeCallback", user.UsdtRechargeCallback)
	}
	rm := rf.Group("/member")
	{
		rms := rm.Use(midware.Auth(&token.JwtToken{
			SigningKey: []byte(constex.JwtCfg.SecretKey),
		}))
		rms.GET("/exrate", user.ExchangeRate)
		rms.POST("/deposit", user.MemberDeposit)
		rms.GET("/depositlog", user.MemberDepositLog)
		rms.POST("/onlinepay", user.OnlinePay)
	}
	/************************************************************************/
	rx := rf.Group("/xxx")
	{
		rxs := rx.Use(midware.Auth(&token.JwtToken{
			SigningKey: []byte(constex.JwtCfg.SecretKey),
		}))
		rxs.GET("/yyy/:id", user.UserInfo)
	}
	return router
}
