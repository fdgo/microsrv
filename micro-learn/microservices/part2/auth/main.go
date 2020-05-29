package main

import (
	"fmt"

	"github.com/wangmhgo/microservice-project/go-micro-learn/microservices/part2/auth/handler"
	"github.com/wangmhgo/microservice-project/go-micro-learn/microservices/part2/auth/model"
	s "github.com/wangmhgo/microservice-project/go-micro-learn/microservices/part2/auth/proto/auth"
	"github.com/wangmhgo/microservice-project/go-micro-learn/microservices/part2/basic"
	"github.com/wangmhgo/microservice-project/go-micro-learn/microservices/part2/basic/config"
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
		micro.Name("fred.micro.shop.srv.auth"),
		micro.Registry(micReg),
		micro.Version("latest"),
	)

	// 服务初始化
	service.Init(
		micro.Action(func(c *cli.Context) {
			// 初始化handler
			model.Init()
			// 初始化handler
			handler.Init()
		}),
	)

	// 注册服务
	s.RegisterServiceHandler(service.Server(), new(handler.Service))

	// 启动服务
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}

func registryOptions(ops *registry.Options) {
	consulCfg := config.GetConsulConfig()
	ops.Addrs = []string{fmt.Sprintf("%s:%d", consulCfg.GetHost(), consulCfg.GetPort())}
}
