package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/registry/consul"
	"github.com/micro/go-micro/web"
	"github.com/micro/go-plugins/config/source/grpc"
	"github.com/opentracing/opentracing-go"
	hdl "microservice/jzapi/admin_web/handler"
	hdlogin "microservice/jzapi/admin_web/handler/login"
	hdlmgr "microservice/jzapi/admin_web/handler/manage"
	glocfg "microservice/jzapi/basic/cfg"
	"microservice/jzapi/basic/cfg/common"
	"microservice/jzapi/basic/cfg/config"
	conval "microservice/jzapi/basic/const_value"
	"microservice/jzapi/lib/log/comlog"
	ginlog "microservice/jzapi/lib/log/ginlog"
	"microservice/jzapi/lib/recover"
	"microservice/jzapi/lib/wrapper/tracer/jaeger"
	"microservice/jzapi/lib/wrapper/tracer/opentracing/gin2micro"
	"time"
)

var (
	srvName = "admin_web"
	srvCfg  = &SrvCfg{}
)

type SrvCfg struct {
	common.SrvCfg
}

func main() {
	// 初始化配置
	Init()
	// 使用consul注册
	micReg := consul.NewRegistry(registryOptions)
	t, io, err := tracer.NewTracer(srvCfg.Name, "")
	if err != nil {
		comlog.Logger.Fatal(err.Error())
		return
	}
	defer io.Close()
	opentracing.SetGlobalTracer(t)
	// 创建新服务
	service := web.NewService(
		web.Name(srvCfg.Name),
		web.Version(srvCfg.Version),
		web.RegisterTTL(time.Second*15),
		web.RegisterInterval(time.Second*5),
		web.Registry(micReg),
		web.Address(srvCfg.Addr()),
	)
	// 初始化服务
	service.Init()
	hdl.NewServices(service.Options().Service.Client())
	gin2micro.SetSamplingFrequency(50)
	rfather := gin.New()
	rfather.Use(recover.Recover("admin_web"))
	radmin := rfather.Group("/admin")
	{
		radmin.POST("/users/actions/login", hdlogin.BackLogin)
		radmin.GET("/users", hdlmgr.UserList)
		radmin.POST("/users/actions/add", hdlmgr.AddUser)
		radmin.DELETE("/users/actions/delete", hdlmgr.DelUser)
		radmin.GET("/channels", hdlmgr.ChannelList)
		radmin.POST("/accesses/actions/add", hdlmgr.AddAccess)
		radmin.DELETE("/accesses/actions/delete", hdlmgr.DelAccess)
		radmin.GET("/accesses", hdlmgr.AccessList)
		radmin.GET("/roles", hdlmgr.RoleList)
		radmin.POST("/roles/actions/add", hdlmgr.AddRole)
		radmin.PUT("/roles/actions/edit", hdlmgr.EditRole)
		radmin.DELETE("/roles/actions/delete", hdlmgr.DeleteRole)
	}
	service.Handle("/", rfather)
	if err := service.Run(); err != nil {
		comlog.Logger.Fatal(err.Error())
		return
	}
}
func registryOptions(ops *registry.Options) {
	consulCfg := &common.Consul{}
	err := config.C().App("consul", consulCfg)
	if err != nil {
		comlog.Logger.Fatal(err.Error())
	}
	ops.Addrs = []string{fmt.Sprintf("%s:%d", consulCfg.Host, consulCfg.Port)}
}

func Init() {
	source := grpc.NewSource(
		grpc.WithAddress(conval.CONFIG_ADDRESS),
		grpc.WithPath("micro"),
	)
	glocfg.Init(config.WithSource(source))
	err := config.C().App(srvName, srvCfg)
	if err != nil {
		comlog.Logger.Fatal(err.Error())
		return
	}
	ginlog.Init(0, "admin_web")
}
