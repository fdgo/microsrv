package initial

import (
	"fmt"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-plugins/config/source/grpc"
	glocfg "microservice/jzapi/basic/cfg"
	"microservice/jzapi/basic/cfg/common"
	"microservice/jzapi/basic/cfg/config"
	"microservice/jzapi/basic/const_value"
	"microservice/jzapi/lib/log/comlog"
)
type SrvCfg struct {
	common.SrvCfg
}
type AbdCfg struct {
	common.Abd
}
var (
	srvName = "user_web"
	srvCfg     = &SrvCfg{}

	abdName = "abd"
	abdCfg = &AbdCfg{}
)
func init()  {
	conAddressSrv()
}

func conAddressSrv()  {
	source := grpc.NewSource(
		grpc.WithAddress(const_value.CONFIG_ADDRESS),
		grpc.WithPath("micro"),
	)
	glocfg.Init(config.WithSource(source))
	err := config.C().App(srvName, srvCfg)
	if err != nil {
		comlog.Logger.Fatal(err.Error())
		return
	}

	err = config.C().App(abdName, abdCfg)
	if err != nil {
		comlog.Logger.Fatal(err.Error())
		return
	}
	fmt.Println("456546abd:",abdCfg,"8888888",srvCfg)
}

func RegistryOptions(ops *registry.Options) {
	consulCfg := &common.Consul{}
	err := config.C().App("consul", consulCfg)
	if err != nil {
		comlog.Logger.Fatal(err.Error())
		return
	}
	ops.Addrs = []string{fmt.Sprintf("%s:%d", consulCfg.Host, consulCfg.Port)}
	fmt.Println("ererewr",ops.Addrs)
}
func GetSrvCfg() *SrvCfg  {
	return srvCfg
}
