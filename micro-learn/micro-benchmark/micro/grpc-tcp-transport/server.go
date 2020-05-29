package main

import (
	"fmt"

	"github.com/wangmhgo/microservice-project/go-micro-learn/micro-benchmark/micro/internal"
	"github.com/wangmhgo/microservice-project/go-micro-learn/micro-benchmark/pb"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/service/grpc"
	"github.com/micro/go-plugins/transport/tcp"
)

func main() {
	service := grpc.NewService(
		micro.Name("go.micro.benchmark.hello.grpc_tcp"),
		micro.Version("latest"),
		micro.Transport(tcp.NewTransport()),
	)

	service.Init()

	pb.RegisterHelloHandler(service.Server(), &internal.HelloS{})

	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}
