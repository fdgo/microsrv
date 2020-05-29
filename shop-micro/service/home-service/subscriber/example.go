package subscriber

import (
	"context"
	"github.com/micro/go-micro/util/log"

	example "github.com/wangmhgo/microservice-project/shop-micro/service/home-service/proto"
)

type Example struct{}

func (e *Example) Handle(ctx context.Context, msg *example.Message) error {
	log.Log("Handler Received message: ", msg.Say)
	return nil
}

func Handler(ctx context.Context, msg *example.Message) error {
	log.Log("Function Received message: ", msg.Say)
	return nil
}
