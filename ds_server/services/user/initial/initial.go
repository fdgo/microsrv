package initial

import (
	dao "ds_server/services/user/dao"
	"ds_server/services/user/handler"
	"ds_server/services/user/plugin"
	usrv "ds_server/services/user/service"
	msql "ds_server/support/lib/mysqlex"
	rds "ds_server/support/lib/redisex"
	cfg "ds_server/support/utils/cfg"
	"ds_server/support/utils/cfg/config"
	"ds_server/support/utils/constex"
	"ds_server/support/utils/hystrix"
	"ds_server/support/utils/trace"
	"fmt"
	"time"

	"github.com/micro/go-log"
	"github.com/micro/go-micro"
	srv "github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-plugins/config/source/grpc"
	"github.com/micro/go-plugins/registry/consul"
	ocplugin "github.com/micro/go-plugins/wrapper/trace/opentracing"
	"github.com/opentracing/opentracing-go"
)

func NewUserSrv() (srv.Service, handler.UserHandler) {
	initConfig()
	initHystrix()
	initTracer()
	db := initDb()
	srv := micro.NewService(
		micro.Name(constex.SRV_USER),
		micro.Registry(initConsul()),
		micro.Version("latest"),
		micro.RegisterTTL(time.Second*30),
		micro.RegisterInterval(time.Second*15),
		micro.WrapHandler(plugin.LogWrapper),
		micro.WrapHandler(ocplugin.NewHandlerWrapper(opentracing.GlobalTracer())),
		micro.WrapClient(hystrix.NewClientWrapper()),
		micro.Address(":8029"),
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
	err = config.C().App("rediscluster", constex.RedisCluterCfg)
	err = config.C().App("jwt", constex.JwtCfg)
	err = config.C().App("memberclass", constex.MemberClassCfg)
	err = config.C().App("agentclass", constex.AgentClassCfg)
	err = config.C().App("exchangerate", constex.ExchangeRateCfg)
	fmt.Println("consul:", constex.ConsulCfg, " mysql:", constex.MysqlCfg, " redis:", constex.RedisCfg, " rediscluster:", constex.RedisCluterCfg,
		" jwt:", constex.JwtCfg, "memberclass:", constex.MemberClassCfg, "agentclass", constex.AgentClassCfg, "exchangerate:", constex.ExchangeRateCfg)
	if err != nil {
		return
	}
}
func initHystrix() {
	hystrix.Configure([]string{
		constex.SRV_USER + "UserHandler.GetAll",
		constex.SRV_USER + "UserHandler.XXX",
	})
}
func initTracer() {
	trace.SetSamplingFrequency(50)
	t, io, err := trace.NewTracer(constex.SRV_USER, "localhost:6831")
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
		op.Timeout =  5 * time.Second
	})
}
func initDb() handler.UserHandler {
	msql := msql.MysqlInstanceg()
	rds := rds.RedisInstanceg()
	usersrv := &usrv.UserService{
		RedisCache:                    rds,
		DsUserBasicinfoDao:            dao.NewDsUserBasicinfoMgr(msql),
		DsSysInfoDao:                  dao.NewDsSysInfoMgr(msql),
		DsUserMemberDepositHistoryDao: dao.NewDsUserMemberDepositHistoryMgr(msql),
		DsUserMemberAgentDao:          dao.NewDsUserMemberAgentMgr(msql),
		DsUserMemberAccountDao:        dao.NewDsUserMemberAccountMgr(msql),
		DsUserMemberClassDao:          dao.NewDsUserMemberClassMgr(msql),
		DsUserAgentClassDao:           dao.NewDsUserAgentClassMgr(msql),
	}
	return handler.UserHandler{Usersrv: usersrv}
}
