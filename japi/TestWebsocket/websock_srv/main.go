package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	proto "microservice/jzapi/TestWebsocket/websock_srv/proto"
	"github.com/micro/go-micro"
	"github.com/micro/go-plugins/config/source/grpc"
	glocfg "microservice/jzapi/basic/cfg"
	"microservice/jzapi/basic/cfg/config"
	"microservice/jzapi/basic/cfg/common"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/registry/consul"
	"microservice/jzapi/lib/wrapper/tracer/jaeger"
	ocplugin "github.com/micro/go-plugins/wrapper/trace/opentracing"
	openTrace "github.com/opentracing/opentracing-go"
)

type Streamer struct{}

// Server side stream
func (e *Streamer) ServerStream(ctx context.Context, req *proto.Request, stream proto.Streamer_ServerStreamStream) error {
	log.Printf("Got msg %v", req.Count)
	for i := 0; i < int(req.Count); i++ {
		if err := stream.Send(&proto.Response{Count: int64(i)}); err != nil {
			return err
		}
	}
	return nil
}

// Bidirectional stream
func (e *Streamer) Stream(ctx context.Context, stream proto.Streamer_StreamStream) error {
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		log.Printf("Got msg %v", req.Count)
		if err := stream.Send(&proto.Response{Count: req.Count}); err != nil {
			return err
		}
	}
}
var (
	appName = "user_srv"
	cfg     = &userCfg{}
)
type userCfg struct {
	common.AppCfg
}
func main() {
	Init()

	// 使用consul注册
	micReg := consul.NewRegistry(registryOptions)

	t, io, err := tracer.NewTracer(cfg.Name, "localhost:6831")
	if err != nil {
		log.Fatal(err)
	}
	defer io.Close()
	openTrace.SetGlobalTracer(t)
	// new service
	service := micro.NewService(
		micro.Name("jz.micro.websock-srv.stream"),
		micro.RegisterTTL(time.Second*15),
		micro.RegisterInterval(time.Second*10),
		micro.Registry(micReg),
		micro.Version("latest"),
		micro.WrapHandler(ocplugin.NewHandlerWrapper(t)),
	)

	// Init command line
	service.Init()

	// Register Handler
	proto.RegisterStreamerHandler(service.Server(), new(Streamer))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
func registryOptions(ops *registry.Options) {
	consulCfg := &common.Consul{}
	err := config.C().App("consul", consulCfg)
	if err != nil {
		panic(err)
	}

	ops.Addrs = []string{fmt.Sprintf("%s:%d", consulCfg.Host, consulCfg.Port)}
}
func Init() {
	source := grpc.NewSource(
		grpc.WithAddress("127.0.0.1:9600"),
		grpc.WithPath("micro"),
	)


	glocfg.Init(config.WithSource(source))

	err := config.C().App(appName, cfg)
	if err != nil {
		panic(err)
	}
	fmt.Println(cfg)
	//log.Logf("[initCfg] 配置，cfg：%v", cfg)
	//dboper.Init()
	return
}
