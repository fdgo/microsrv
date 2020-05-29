package main

import (
	"context"
	"fmt"

	proto "github.com/wangmhgo/microservice-project/go-micro-learn/examples/basic-practices/micro-api/rpc/proto"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/util/log"
)

type Example struct{}

func (e *Example) Call(ctx context.Context, req *proto.CallRequest, rsp *proto.CallResponse) error {
	log.Log("Example.Call接口收到请求，返回错误")
	return fmt.Errorf("[ERR] Call Error")
}

func main() {
	service := micro.NewService(
		micro.Name("go.micro.retry.example"),
	)

	service.Init()

	// 注册 example handler
	proto.RegisterExampleHandler(service.Server(), new(Example))

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
