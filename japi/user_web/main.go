package main

import (
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/web"
	"github.com/micro/go-plugins/registry/etcdv3"
	"github.com/opentracing/opentracing-go"
	"microservice/jzapi/lib/log/comlog"
	"microservice/jzapi/lib/recover"
	"microservice/jzapi/lib/wrapper/tracer/jaeger"
	"microservice/jzapi/user_web/handler/login"
	"microservice/jzapi/user_web/initial"
	hdl "microservice/jzapi/user_web/srvclient"
	"time"
)

func main() {
	micReg := etcdv3.NewRegistry(initial.RegistryOptions)
	t, io, err := tracer.NewTracer(initial.GetSrvCfg().Name, "localhost:6831")
	if err != nil {
		comlog.Logger.Fatal(err.Error())
		return
	}
	defer io.Close()
	opentracing.SetGlobalTracer(t)
	service := web.NewService(
		web.Name(initial.GetSrvCfg().Name),
		//web.Version(initial.GetSrvCfg().Version),
		web.RegisterTTL(time.Second*15),
		web.RegisterInterval(time.Second*5),
		web.Registry(micReg),
		web.Address(initial.GetSrvCfg().Addr()),
	)
	service.Init()
	//--------------------------------------------------

	hdl.NewServices(service.Options().Service.Client())
	rfather := gin.New()
	rfather.Use(gin.Logger())
	rfather.Use(recover.Recover("user_web"))
	rlogin := rfather.Group("/user")
	rlogin.POST("/multi/login", login.MultiLoginMobile) //合集注册登陆二合一

	rlogin.POST("/regist/single/quick", login.SingleRegistQuick)     //单包快速注册
	rlogin.POST("/regist/single/mobile", login.SingleRegistMobile)   //单包手机号注册
	rlogin.POST("/regist/single/account", login.SingleRegistAccount) //单包账号注册

	rlogin.POST("/login/single/guest", login.SingleLoginGuest)     //单包游客登陆
	rlogin.POST("/login/single/mobile", login.SingleLoginMobile)   //单包手机号登陆
	rlogin.POST("/login/single/account", login.SingleLoginAccount) //单包账号登陆

	service.Handle("/", rfather)
	if err := service.Run(); err != nil {
		comlog.Logger.Fatal(err.Error())
		return
	}
}
