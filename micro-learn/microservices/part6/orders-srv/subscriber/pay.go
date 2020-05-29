package subscriber

import (
	"context"

	"github.com/wangmhgo/microservice-project/go-micro-learn/microservices/part6/orders-srv/model/orders"
	payS "github.com/wangmhgo/microservice-project/go-micro-learn/microservices/part6/payment-srv/proto/payment"
	"github.com/micro/go-micro/util/log"
)

var (
	ordersService orders.Service
)

// Init 初始化handler
func Init() {
	ordersService, _ = orders.GetService()
}

// PayOrder 订单支付消息
func PayOrder(ctx context.Context, event *payS.PayEvent) (err error) {
	log.Logf("[PayOrder] 收到支付订单通知，%d，%d", event.OrderId, event.State)

	err = ordersService.UpdateOrderState(event.OrderId, int(event.State))
	if err != nil {
		log.Logf("[PayOrder] 收到支付订单通知，更新状态异常，%s", err)
		return
	}
	return
}
