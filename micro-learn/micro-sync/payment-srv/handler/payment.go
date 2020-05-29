package handler

import (
	"context"

	"github.com/wangmhgo/microservice-project/go-micro-learn/micro-sync/payment-srv/model"
	proto "github.com/wangmhgo/microservice-project/go-micro-learn/micro-sync/payment-srv/proto/payment"
	"github.com/micro/go-micro/util/log"
)

var (
	paymentService model.Service
)

type Service struct {
}

// Init 初始化handler
func Init() {
	paymentService, _ = model.GetService()
}

// New 新增订单
func (e *Service) PayOrder(ctx context.Context, req *proto.Request, rsp *proto.Response) (err error) {
	log.Log("[PayOrder] 收到支付请求")

	err = paymentService.PayOrder(req.OrderId)
	if err != nil {
		rsp.Success = false
		rsp.Error = &proto.Error{
			Detail: err.Error(),
		}
		return
	}

	rsp.Success = true
	return
}
