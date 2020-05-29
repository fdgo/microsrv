package initial

import (
	brep "ds_server/services/base/dao"
	"ds_server/services/base/handler"
	"ds_server/services/base/plugin"
	bsrv "ds_server/services/base/service"
	cfg "ds_server/support/utils/cfg"
	"ds_server/support/utils/cfg/config"
	"ds_server/support/utils/constex"
	"ds_server/support/utils/hystrix"
	"ds_server/support/utils/trace"
	"fmt"
	"github.com/micro/go-log"
	"github.com/micro/go-micro"
	srv "github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-plugins/config/source/grpc"
	"github.com/micro/go-plugins/registry/consul"
	ocplugin "github.com/micro/go-plugins/wrapper/trace/opentracing"
	"github.com/opentracing/opentracing-go"
	"time"
)

func NewBaseSrv() (srv.Service, handler.BaseHandler) {
	initConfig()
	initHystrix()
	initTracer()
	db := initDb()
	srv := micro.NewService(
		micro.Name(constex.SRV_BASE),
		micro.Registry(initConsul()),
		micro.Version("latest"),
		micro.RegisterTTL(time.Second*30),
		micro.RegisterInterval(time.Second*15),
		micro.WrapHandler(plugin.LogWrapper),
		micro.WrapHandler(ocplugin.NewHandlerWrapper(opentracing.GlobalTracer())),
		micro.WrapClient(hystrix.NewClientWrapper()),
	)
	srv.Init()
	return srv, db
}
func initConfig() {
	source := grpc.NewSource(
		grpc.WithAddress(constex.CONFIG_ADDRESS),
		grpc.WithPath("micro"),
	)
	cfg.Init(config.WithSource(source))
	err := config.C().App("consul", constex.ConsulCfg)
	err = config.C().App("mysql", constex.MysqlCfg)
	err = config.C().App("redis", constex.RedisCfg)
	fmt.Println("consul:",constex.ConsulCfg," mysql:",constex.MysqlCfg," redis:",constex.RedisCfg)
	if err != nil {
		return
	}
}
func initHystrix() {
	hystrix.Configure([]string{
		constex.SRV_BASE + "BaseHandler.VfCode",
		constex.SRV_BASE + "BaseHandler.XXX",
	})
}
func initTracer() {
	trace.SetSamplingFrequency(50)
	t, io, err := trace.NewTracer(constex.SRV_BASE, "localhost:6831")
	if err != nil {
		log.Fatal(err)
	}
	defer io.Close()
	opentracing.SetGlobalTracer(t)
}
func initConsul() registry.Registry {
	return consul.NewRegistry(func(op *registry.Options) {
		op.Addrs = []string{
			constex.ConsulCfg.Host,
		}
	})
}
func initDb() handler.BaseHandler {
	basesrv := &bsrv.BaseService{BaseDao:brep.NewBaseDao()}
	return handler.BaseHandler{BaseSvr: basesrv}
}
