package main

import (
	"fmt"

	"github.com/wangmhgo/microservice-project/go-micro-learn/microservices/part3/basic"
	"github.com/wangmhgo/microservice-project/go-micro-learn/microservices/part3/basic/common"
	"github.com/wangmhgo/microservice-project/go-micro-learn/microservices/part3/basic/config"
	"github.com/wangmhgo/microservice-project/go-micro-learn/microservices/part3/orders-srv/handler"
	"github.com/wangmhgo/microservice-project/go-micro-learn/microservices/part3/orders-srv/model"
	proto "github.com/micro-in-cn/tutorials/microservices-in-micro/part3/orders-srv/proto/orders"
	"github.com/wangmhgo/microservice-project/go-micro-learn/microservices/part3/orders-srv/subscriber"
	"github.com/micro/cli"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/registry/consul"
	"github.com/micro/go-micro/util/log"
)

func main() {
	// 初始化配置、数据库等信息
	basic.Init()

	// 使用consul注册
	micReg := consul.NewRegistry(registryOptions)

	// 新建服务
	service := micro.NewService(
		micro.Name("fred.micro.shop.srv.orders"),
		micro.Registry(micReg),
		micro.Version("latest"),
	)

	// 服务初始化
	service.Init(
		micro.Action(func(c *cli.Context) {
			// 初始化模型层
			model.Init()
			// 初始化handler
			handler.Init()
			// 初始化sub
			subscriber.Init()
		}),
	)

	// 侦听订单支付消息
	err := micro.RegisterSubscriber(common.TopicPaymentDone, service.Server(), subscriber.PayOrder)
	if err != nil {
		log.Fatal(err)
	}

	// 注册服务
	err = proto.RegisterOrdersHandler(service.Server(), new(handler.Orders))
	if err != nil {
		log.Fatal(err)
	}

	// 启动服务
	if err = service.Run(); err != nil {
		log.Fatal(err)
	}
}

func registryOptions(ops *registry.Options) {
	consulCfg := config.GetConsulConfig()
	ops.Addrs = []string{fmt.Sprintf("%s:%d", consulCfg.GetHost(), consulCfg.GetPort())}
}
