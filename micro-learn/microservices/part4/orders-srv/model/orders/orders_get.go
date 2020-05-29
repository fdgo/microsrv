package orders

import (
	proto "github.com/micro-in-cn/tutorials/microservices-in-micro/part4/orders-srv/proto/orders"
	"github.com/wangmhgo/microservice-project/go-micro-learn/microservices/part4/plugins/db"
	"github.com/micro/go-micro/util/log"
)

// GetOrder 获取订单
func (s *service) GetOrder(orderId int64) (order *proto.Order, err error) {
	order = &proto.Order{}

	// 获取数据库
	o := db.GetDB()
	// 查询
	err = o.QueryRow("SELECT id, user_id, book_id, inv_his_id, state FROM orders WHERE id = ?", orderId).Scan(
		&order.Id, &order.UserId, &order.BookId, &order.InvHistoryId, &order.State)
	if err != nil {
		log.Logf("[GetOrder] 查询数据失败，err：%s", err)
		return
	}

	return
}
