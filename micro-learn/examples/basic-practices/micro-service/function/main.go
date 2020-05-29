package main

import (
	"context"

	proto "github.com/wangmhgo/microservice-project/go-micro-learn/examples/basic-practices/micro-service/proto"
	"github.com/micro/go-micro"
)

type Greeter struct{}

func (g *Greeter) Hello(ctx context.Context, req *proto.HelloRequest, rsp *proto.HelloResponse) error {
	rsp.Greeting = "你好呀！ " + req.Name
	return nil
}

func main() {
	fnc := micro.NewFunction(
		micro.Name("greeter"),
	)

	fnc.Init()

	fnc.Handle(new(Greeter))

	fnc.Run()
}
