package initial

import (
	"ds_server/support/utils/cfg"
	"ds_server/support/utils/cfg/config"
	"ds_server/support/utils/constex"
	"ds_server/support/utils/trace"
	"fmt"
	"time"

	"github.com/micro/go-log"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/web"
	"github.com/micro/go-plugins/config/source/grpc"
	"github.com/micro/go-plugins/registry/consul"
	"github.com/opentracing/opentracing-go"
)

func initConfig() {
	source := grpc.NewSource(
		grpc.WithAddress(constex.CONFIG_ADDRESS),
		grpc.WithPath("micro"),
	)
	cfg.Init(config.WithSource(source))
	err := config.C().App("consul", constex.ConsulCfg)
	err = config.C().App("mysql", constex.MysqlCfg)
	err = config.C().App("redis", constex.RedisCfg)
	err = config.C().App("aioss", constex.AiossCfg)
	err = config.C().App("jwt", constex.JwtCfg)
	err = config.C().App("memberclass", constex.MemberClassCfg)
	err = config.C().App("agentclass", constex.AgentClassCfg)
	fmt.Println("============constex.JwtCfg=======", constex.AiossCfg.AccessKeyId)
	if err != nil {
		return
	}
}
func NewGateWaysrv() web.Service {
	initConfig()
	reg := consul.NewRegistry(func(op *registry.Options) {
		op.Addrs = []string{
			constex.ConsulCfg.Host,
		}
		op.Timeout =  5 * time.Second
	})
	trace.SetSamplingFrequency(50)
	t, io, err := trace.NewTracer(constex.SRV_GATAWAY, "localhost:6831")
	if err != nil {
		log.Fatal(err)
	}
	defer io.Close()
	opentracing.SetGlobalTracer(t)

	return web.NewService(
		web.Name(constex.SRV_GATAWAY),
		web.Registry(reg),
		web.Address(":8030"),
	)
}
