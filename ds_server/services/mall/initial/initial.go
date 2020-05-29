package initial

import (
	"ds_server/services/mall/handler"
	"ds_server/services/mall/model"
	"ds_server/services/mall/plugin"
	usrv "ds_server/services/mall/service"
	db "ds_server/support/lib/mysqlex"
	cfg "ds_server/support/utils/cfg"
	"ds_server/support/utils/cfg/config"
	"ds_server/support/utils/constex"
	"ds_server/support/utils/hystrix"
	"ds_server/support/utils/logger"
	"ds_server/support/utils/trace"
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/micro/go-log"
	"github.com/micro/go-micro"
	srv "github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-plugins/config/source/grpc"
	"github.com/micro/go-plugins/registry/consul"
	ocplugin "github.com/micro/go-plugins/wrapper/trace/opentracing"
	"github.com/opentracing/opentracing-go"
)

func NewMallSrv() (srv.Service, handler.MallHandler) {
	initConfig()
	initHystrix()
	initTracer()
	inittable()
	db := initDb()
	srv := micro.NewService(
		micro.Name(constex.SRV_MALL),
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
	fmt.Println("consul:", constex.ConsulCfg, " mysql:", constex.MysqlCfg, " redis:", constex.RedisCfg)
	if err != nil {
		return
	}
}

//容错
func initHystrix() {
	hystrix.Configure([]string{
		constex.SRV_USER + "MallHandler.GetMall",
		constex.SRV_USER + "MallHandler.XXX",
	})
}

func initTracer() {
	trace.SetSamplingFrequency(50)
	t, io, err := trace.NewTracer(constex.SRV_MALL, "localhost:6831")
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

func initDb() handler.MallHandler {
	//db := db.MysqlInstance()
	messagesrv := &usrv.MallService{}
	return handler.MallHandler{Mall: messagesrv}
}

type TSQLLogger struct{}

func (slog TSQLLogger) Print(values ...interface{}) {
	vals := gorm.LogFormatter(values...)
	logger.Log.SqlDebug(vals...)
}

var d = []interface{}{
	&model.DsOrder{},
	&model.DsOrderParameter{},
	&model.DsProduct{},
	&model.DsProductUrl{},
	&model.DsProductDetail{},
	&model.DsProductStore{},
	&model.DsReceivingAddress{},
	&model.DsProductStoreChange{},
}

func inittable() {
	db := db.MysqlInstanceg()
	db.LogMode(true)
	db.SetLogger(TSQLLogger{})
	// 如果有注册models, 则进行建表同步
	if len(d) > 0 {
		for _, m := range d {
			if !db.HasTable(m) {
				err := db.CreateTable(m).Error
				if err != nil {
					log.Fatal(err)
				}
			}
		}
		db.AutoMigrate(d...)
	}
}
