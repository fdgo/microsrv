package handler

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	hystrix_go "github.com/afex/hystrix-go/hystrix"
	auth "github.com/wangmhgo/microservice-project/go-micro-learn/microservices/part8/auth/proto/auth"
	invS "github.com/wangmhgo/microservice-project/go-micro-learn/microservices/part8/inventory-srv/proto/inventory"
	orders "github.com/wangmhgo/microservice-project/go-micro-learn/microservices/part8/orders-srv/proto/orders"
	"github.com/wangmhgo/microservice-project/go-micro-learn/microservices/part8/plugins/session"
	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/util/log"
	"github.com/micro/go-plugins/wrapper/breaker/hystrix"
)

var (
	serviceClient orders.OrdersService
	authClient    auth.Service
	invClient     invS.InventoryService
)

// Error 错误结构体
type Error struct {
	Code   string `json:"code"`
	Detail string `json:"detail"`
}

func Init() {
	hystrix_go.DefaultVolumeThreshold = 1
	hystrix_go.DefaultErrorPercentThreshold = 1
	cl := hystrix.NewClientWrapper()(client.DefaultClient)
	cl.Init(
		client.Retries(3),
		client.Retry(func(ctx context.Context, req client.Request, retryCount int, err error) (bool, error) {
			log.Log(req.Method(), retryCount, " client retry")
			return true, nil
		}),
	)
	serviceClient = orders.NewOrdersService("fred.micro.shop.srv.orders", cl)
	authClient = auth.NewService("fred.micro.shop.srv.auth", cl)
}

// New 新增订单入口
func New(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	// 只接受POST请求
	if r.Method != "POST" {
		log.Logf("非法请求")
		http.Error(w, "非法请求", 400)
		return
	}

	r.ParseForm()
	bookId, _ := strconv.ParseInt(r.Form.Get("bookId"), 10, 10)

	// 返回结果
	response := map[string]interface{}{}

	// 调用后台服务
	rsp, err := serviceClient.New(ctx, &orders.Request{
		BookId: bookId,
		UserId: session.GetSession(w, r).Values["userId"].(int64),
	})

	// 返回结果
	response["ref"] = time.Now().UnixNano()
	if err != nil {
		response["success"] = false
		response["error"] = Error{
			Detail: err.Error(),
		}
	} else {
		response["success"] = true
		response["orderId"] = rsp.Order.Id
	}

	w.Header().Add("Content-Type", "application/json; charset=utf-8")

	// 返回JSON结构
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}

//
func Hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello"))
}
