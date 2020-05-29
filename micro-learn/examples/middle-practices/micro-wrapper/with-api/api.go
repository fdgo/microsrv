package main

import (
	"context"
	"encoding/json"
	"strings"

	proto "github.com/wangmhgo/microservice-project/go-micro-learn/examples/basic-practices/micro-api/api/proto"
	"github.com/micro/go-micro"
	api "github.com/micro/go-micro/api/proto"
	"github.com/micro/go-micro/errors"
	"github.com/micro/go-micro/server"
	"github.com/micro/go-micro/util/log"
)

type Example struct{}

type Foo struct{}

// Example.Call 通过API向外暴露为/example/call，接收http请求
// 即：/example/call请求会调用go.micro.api.example服务的Example.Call方法
func (e *Example) Call(ctx context.Context, req *api.Request, rsp *api.Response) error {
	log.Logf("Example.Call接口收到请求")

	name, ok := req.Get["name"]

	if !ok || len(name.Values) == 0 {
		return errors.BadRequest("go.micro.api.example", "参数不正确")
	}

	rsp.StatusCode = 200

	b, _ := json.Marshal(map[string]string{
		"message": "我们已经收到你的请求，" + strings.Join(name.Values, " "),
	})

	// 设置返回值
	rsp.Body = string(b)

	return nil
}

// logWrapper 包装HandlerFunc类型的接口
func logWrapper(fn server.HandlerFunc) server.HandlerFunc {
	return func(ctx context.Context, req server.Request, rsp interface{}) error {
		log.Logf("[logWrapper] 收到请求，请求接口：%s", req.Endpoint())
		if reqStruct, ok := req.Body().(*api.Request); ok {
			log.Logf("[logWrapper] 请求参数名：%s ", reqStruct.Get["name"].Key)
			err := fn(ctx, req, rsp)
			return err
		}

		return nil
	}
}

// rspHeaderWrapper 处理Response头部
func rspHeaderWrapper(fn server.HandlerFunc) server.HandlerFunc {
	return func(ctx context.Context, req server.Request, rsp interface{}) error {
		if rspT, ok := rsp.(*api.Response); ok {
			if rspT.Header == nil {
				rspT.Header = map[string]*api.Pair{}
			}

			rspT.Header["NEW-HEADER-ADDED"] = &api.Pair{Key: "NEW-HEADER-ADDED", Values: []string{"NEW-HEADER-ADDED-VALUE"}}
			if reqStruct, ok := req.Body().(*api.Request); ok {
				for k, v := range reqStruct.Header {
					log.Logf("[rspHeaderWrapper] 原头：%s:%s", k, v)
					rspT.Header[k] = v
				}
			}
		}
		err := fn(ctx, req, rsp)
		return err
	}
}

func main() {
	service := micro.NewService(
		micro.Name("go.micro.api.example"),
		micro.WrapHandler(logWrapper, rspHeaderWrapper),
	)

	service.Init()

	// 注册 example handler
	proto.RegisterExampleHandler(service.Server(), new(Example))

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
