package main

import (
	"fmt"

	"github.com/wangmhgo/microservice-project/go-micro-learn/micro-benchmark/micro/internal"
	"github.com/wangmhgo/microservice-project/go-micro-learn/micro-benchmark/pb"
	"github.com/micro/go-micro"
	"github.com/micro/go-plugins/transport/utp"
)

func main() {
	service := micro.NewService(
		micro.Name("go.micro.benchmark.hello.utp_transport"),
		micro.Version("latest"),
		micro.Transport(utp.NewTransport()),
	)

	service.Init()

	pb.RegisterHelloHandler(service.Server(), &internal.HelloS{})

	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}
