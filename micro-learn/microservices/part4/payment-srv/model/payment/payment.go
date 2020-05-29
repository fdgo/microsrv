package payment

import (
	"fmt"
	"sync"

	"github.com/wangmhgo/microservice-project/go-micro-learn/microservices/part4/basic/common"
	invS "github.com/micro-in-cn/tutorials/microservices-in-micro/part4/inventory-srv/proto/inventory"
	ordS "github.com/micro-in-cn/tutorials/microservices-in-micro/part4/orders-srv/proto/orders"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/client"
)

var (
	s            *service
	invClient    invS.InventoryService
	ordSClient   ordS.OrdersService
	m            sync.RWMutex
	payPublisher micro.Publisher
)

// service 服务
type service struct {
}

// Service 服务类
type Service interface {
	// PayOrder 支付订单
	PayOrder(orderId int64) (err error)
}

// GetService 获取服务类
func GetService() (Service, error) {
	if s == nil {
		return nil, fmt.Errorf("[GetService] GetService 未初始化")
	}
	return s, nil
}

// Init 初始化库存服务层
func Init() {
	m.Lock()
	defer m.Unlock()

	if s != nil {
		return
	}

	invClient = invS.NewInventoryService("fred.micro.shop.srv.inventory", client.DefaultClient)
	ordSClient = ordS.NewOrdersService("fred.micro.shop.srv.orders", client.DefaultClient)
	payPublisher = micro.NewPublisher(common.TopicPaymentDone, client.DefaultClient)
	s = &service{}
}
