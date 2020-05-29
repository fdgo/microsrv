package orders

import (
	"fmt"
	invS "github.com/wangmhgo/microservice-project/go-micro-learn/microservices/part5/inventory-srv/proto/inventory"
	proto "github.com/wangmhgo/microservice-project/go-micro-learn/microservices/part5/orders-srv/proto/orders"
	"github.com/micro/go-micro/client"
	"sync"
)

var (
	s         *service
	invClient invS.InventoryService
	m         sync.RWMutex
)

// service 服务
type service struct {
}

// Service 订单服务类
type Service interface {
	// New 下单
	New(bookId, userId int64) (orderId int64, err error)

	// GetOrder 获取订单
	GetOrder(orderId int64) (order *proto.Order, err error)

	// UpdateOrderState 更新订单状态
	UpdateOrderState(orderId int64, state int) (err error)
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
	s = &service{}
}
