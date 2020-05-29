package main

import (
	"fmt"
	"time"
	"github.com/micro/cli"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/registry/consul"
	"github.com/micro/go-micro/util/log"
	"github.com/micro/go-plugins/config/source/grpc"
	"github.com/opentracing/opentracing-go"
	"microservice/jzapi/basic/cfg/common"
	"microservice/jzapi/basic/cfg/config"
	"microservice/jzapi/basic/const_value"
	"microservice/jzapi/lib/wrapper/tracer/jaeger"
	baseproto "microservice/jzapi/proto/base"
	glocfg "microservice/jzapi/basic/cfg"
	hdler "microservice/jzapi/base_srv/handler"
	start "microservice/jzapi/base_srv/init"
	plgtrac "github.com/micro/go-plugins/wrapper/trace/opentracing"
)
var (
	srvName = "base_srv"
	srvCfg     = &SrvCfg{}
)
type SrvCfg struct {
	common.SrvCfg
}
func main() {
	// 初始化配置、数据库等信息
	initCfg()
	// 使用consul注册
	micReg := consul.NewRegistry(registryOptions)

	t, io, err := tracer.NewTracer(srvCfg.Name, "")
	if err != nil {
		log.Fatal(err)
	}
	defer io.Close()
	opentracing.SetGlobalTracer(t)
	// 新建服务
	service := micro.NewService(
		micro.Name(srvCfg.Name),
		//micro.Broker(nsq.NewBroker(
		//	broker.Addrs([]string{"127.0.0.1:4150"}...),
		//)),
		micro.RegisterTTL(time.Second*15),
		micro.RegisterInterval(time.Second*10),
		micro.Registry(micReg),
		micro.Version(srvCfg.Version),
		micro.Address(srvCfg.Addr()),
		micro.WrapHandler(plgtrac.NewHandlerWrapper(t)),
	)
	// 服务初始化
	service.Init(
		micro.Action(func(c *cli.Context) {
			// 初始化handler
			start.InitSrvs()
			// 初始化handler
			hdler.Init()
		}),
	)
	// 注册服务
	baseproto.RegisterBaseHandler(service.Server(), new(hdler.Service))
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
	ops.Addrs = []string{fmt.Sprintf("%s:%d", consulCfg.Host, consulCfg.Port)}
}
func initCfg() {
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