package main

import (
	"fmt"
	"github.com/micro/cli"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/registry/consul"
	"github.com/micro/go-micro/util/log"
	"github.com/micro/go-plugins/config/source/grpc"
	plgtrac "github.com/micro/go-plugins/wrapper/trace/opentracing"
	opentrace "github.com/opentracing/opentracing-go"
	glocfg "microservice/jzapi/basic/cfg"
	"microservice/jzapi/basic/cfg/common"
	"microservice/jzapi/basic/cfg/config"
	"microservice/jzapi/basic/const_value"
	"microservice/jzapi/lib/wrapper/tracer/jaeger"
	adminproto "microservice/jzapi/proto/admin"
	"microservice/jzapi/admin_srv/handler"
	start "microservice/jzapi/admin_srv/init"
	"time"
)
var (
	srvName = "admin_srv"
	srvCfg     = &SrvCfg{}
)
type SrvCfg struct {
	common.SrvCfg
}
func main() {
	// 初始化配置、数据库等信息
	Init()
	// 使用consul注册
	micReg := consul.NewRegistry(registryOptions)

	t, io, err := tracer.NewTracer(srvCfg.Name, "localhost:6831")
	if err != nil {
		log.Fatal(err)
	}
	defer io.Close()
	opentrace.SetGlobalTracer(t)
	// 新建服务
	service := micro.NewService(
		micro.Name(srvCfg.Name),
		//micro.Broker(nsq.NewBroker(
		//	broker.Addrs([]string{"127.0.0.1:4150"}...),
		//)),
		micro.RegisterTTL(time.Second*15),
		micro.RegisterInterval(time.Second*10),
		micro.Registry(micReg),
		micro.Version("latest"),
		micro.WrapHandler(plgtrac.NewHandlerWrapper(t)),
	)
	// 服务初始化
	service.Init(
		micro.Action(func(c *cli.Context) {
			// 初始化模型层
			start.InitSrvs()
			// 初始化handler
			handler.Init()
		}),
	)
	//sOpts := broker.NewSubscribeOptions(
	//	nsq.WithMaxInFlight(5),
	//)
	//
	//err =micro.RegisterSubscriber("go.micro.broker.topic.nsq", service.Server(), &handler.NsqService{}, server.SubscriberContext(sOpts.Context))
	//if err !=nil{
	//	log.Fatal(err)
	//}
	// 注册服务
	adminproto.RegisterAdminHandler(service.Server(), new(handler.Service))
	// 启动服务
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}

func registryOptions(ops *registry.Options) {
	consulCfg := &common.Consul{}
	err := config.C().App("consul", consulCfg)
	if err != nil {
		panic(err)
	}
	//fmt.Println(consulCfg.Host, consulCfg.Port)
	ops.Addrs = []string{fmt.Sprintf("%s:%d", consulCfg.Host, consulCfg.Port)}
}

func Init() {
	source := grpc.NewSource(
		grpc.WithAddress(const_value.CONFIG_ADDRESS),
		grpc.WithPath("micro"),
	)
	glocfg.Init(config.WithSource(source))

	err := config.C().App(srvName, srvCfg)
	if err != nil {
		panic(err)
	}
}
