package main

import (
	"fmt"
	"time"

	"github.com/wangmhgo/microservice-project/go-micro-learn/microservices/part7/basic"
	"github.com/wangmhgo/microservice-project/go-micro-learn/microservices/part7/basic/common"
	"github.com/wangmhgo/microservice-project/go-micro-learn/microservices/part7/basic/config"
	"github.com/wangmhgo/microservice-project/go-micro-learn/microservices/part7/payment-srv/handler"
	"github.com/wangmhgo/microservice-project/go-micro-learn/microservices/part7/payment-srv/model"
	s "github.com/wangmhgo/microservice-project/go-micro-learn/microservices/part7/payment-srv/proto/payment"
	"github.com/wangmhgo/microservice-project/go-micro-learn/microservices/part7/plugins/tracer/jaeger"
	"github.com/micro/cli"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/registry/consul"
	"github.com/micro/go-micro/util/log"
	"github.com/micro/go-plugins/config/source/grpc"
	openTrace "github.com/micro/go-plugins/wrapper/trace/opentracing"
	"github.com/opentracing/opentracing-go"
)

var (
	appName = "payment_srv"
	cfg     = &appCfg{}
)

type appCfg struct {
	common.AppCfg
}

func main() {
	// 初始化配置、数据库等信息
	initCfg()

	// 使用consul注册
	micReg := consul.NewRegistry(registryOptions)

	t, io, err := tracer.NewTracer(cfg.Name, "")
	if err != nil {
		log.Fatal(err)
	}
	defer io.Close()

	opentracing.SetGlobalTracer(t)
	// 新建服务
	service := micro.NewService(
		micro.Name(cfg.Name),
		micro.Version(cfg.Version),
		micro.RegisterTTL(time.Second*15),
		micro.RegisterInterval(time.Second*10),
		micro.Registry(micReg),
		micro.Address(cfg.Addr()),
		micro.WrapHandler(openTrace.NewHandlerWrapper(opentracing.GlobalTracer())),
	)

	// 服务初始化
	service.Init(
		micro.Action(func(c *cli.Context) {
			// 初始化模型层
			model.Init()
			// 初始化handler
			handler.Init()
		}),
	)

	// 注册服务
	s.RegisterPaymentHandler(service.Server(), new(handler.Service))

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
		grpc.WithAddress("127.0.0.1:9600"),
		grpc.WithPath("micro"),
	)

	basic.Init(config.WithSource(source))

	err := config.C().App(appName, cfg)
	if err != nil {
		panic(err)
	}

	log.Logf("[initCfg] 配置，cfg：%v", cfg)

	return
}
