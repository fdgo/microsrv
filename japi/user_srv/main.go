package main

import (
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry/consul"
	"microservice/jzapi/lib/log/comlog"
	"microservice/jzapi/user_srv/initial"
	"microservice/jzapi/lib/wrapper/tracer/jaeger"
	"github.com/micro/cli"
	userproto "microservice/jzapi/proto/user"
	opentrace "github.com/opentracing/opentracing-go"
	plgtrac "github.com/micro/go-plugins/wrapper/trace/opentracing"
	"microservice/jzapi/user_srv/srvclient"
	"microservice/jzapi/user_srv/services"
	"time"
)

func main()  {
	micReg := consul.NewRegistry(initial.RegistryOptions)

	t, io, err := tracer.NewTracer(initial.GetSrvCfg().Name, "localhost:6831")
	if err != nil {
		comlog.Logger.Fatal(err.Error())
		return
	}
	defer io.Close()
	opentrace.SetGlobalTracer(t)


	service := micro.NewService(
		micro.Name(initial.GetSrvCfg().Name),
		micro.RegisterTTL(time.Second*15),
		micro.RegisterInterval(time.Second*10),
		micro.Registry(micReg),
		micro.WrapHandler(plgtrac.NewHandlerWrapper(t)),
	)
	service.Init(
		micro.Action(func(c *cli.Context) {
			srvclient.Init()
		}),
	)
	userproto.RegisterUserHandler(service.Server(), new(services.Service))
	// 启动服务
	if err := service.Run(); err != nil {
		comlog.Logger.Fatal(err.Error())
		return
	}
}
