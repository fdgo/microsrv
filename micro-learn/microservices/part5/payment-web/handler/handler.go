package handler

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	auth "github.com/wangmhgo/microservice-project/go-micro-learn/microservices/part5/auth/proto/auth"
	payS "github.com/wangmhgo/microservice-project/go-micro-learn/microservices/part5/payment-srv/proto/payment"
	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/util/log"
)

var (
	serviceClient payS.PaymentService
	authClient    auth.Service
)

// Error 错误结构体
type Error struct {
	Code   string `json:"code"`
	Detail string `json:"detail"`
}

func Init() {
	serviceClient = payS.NewPaymentService("fred.micro.shop.srv.payment", client.DefaultClient)
	authClient = auth.NewService("fred.micro.shop.srv.auth", client.DefaultClient)
}

// PayOrder 支付订单
func PayOrder(w http.ResponseWriter, r *http.Request) {

	// 只接受POST请求
	if r.Method != "POST" {
		log.Logf("非法请求")
		http.Error(w, "非法请求", 400)
		return
	}

	r.ParseForm()

	orderId, _ := strconv.ParseInt(r.Form.Get("orderId"), 10, 10)

	// 调用后台服务
	_, err := serviceClient.PayOrder(context.TODO(), &payS.Request{
		OrderId: orderId,
	})

	// 返回结果
	response := map[string]interface{}{}

	// 返回结果
	response["ref"] = time.Now().UnixNano()
	if err != nil {
		response["success"] = false
		response["error"] = Error{
			Detail: err.Error(),
		}
	} else {
		response["success"] = true
	}

	w.Header().Add("Content-Type", "application/json; charset=utf-8")

	// 返回JSON结构
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}
