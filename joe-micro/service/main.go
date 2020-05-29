package main

import (
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/registry/consul"
	ocplugin "github.com/micro/go-plugins/wrapper/trace/opentracing"
	"github.com/opentracing/opentracing-go"
	"github.com/wangmhgo/microservice-project/joe-micro/lib/log"
	"github.com/wangmhgo/microservice-project/joe-micro/lib/queue"
	"github.com/wangmhgo/microservice-project/joe-micro/lib/trace"
	"github.com/wangmhgo/microservice-project/joe-micro/service/handler"
	srv "github.com/wangmhgo/microservice-project/joe-micro/service/proto/service"
	"github.com/wangmhgo/microservice-project/joe-micro/service/subscriber"
	"time"
)

func main() {
	name := "go.micro.srv.service"

	/************************************/
	/********** 服务发现  cousul   ********/
	/************************************/
	reg := consul.NewRegistry(func(op *registry.Options) {
		op.Addrs = []string{
			"127.0.0.1:8500",
		}
		op.Timeout =  5 * time.Second
	})




	/************************************/
	/********** 链路追踪  trace   ********/
	/************************************/
	trace.SetSamplingFrequency(50)
	t, io, err := trace.NewTracer(name, "localhost:6831")
	if err != nil {
		log.Fatal(err)
	}
	defer io.Close()
	opentracing.SetGlobalTracer(t)

	/************************************/
	/********** 消息队列  queue   ********/
	/************************************/
	queue.Init("",nil,1,true)
	subscriber.Registersubscriber()  //注册消费者


	// New Service
	service := micro.NewService(
		micro.Name(name),
		micro.Registry(reg),
		micro.RegisterTTL(time.Second*15),        //重新注册时间
		micro.RegisterInterval(time.Second*10),   //注册过期时间
		micro.Version("latest"),
		micro.WrapHandler(ocplugin.NewHandlerWrapper(opentracing.GlobalTracer())),

	)

	// Initialise service
	service.Init()

	// Register Handler
	err=srv.RegisterServiceHandler(service.Server(), new(handler.Service))
	if err != nil {
		log.Fatal(err)
	}




	// Run service

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
