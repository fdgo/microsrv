package srvclient

import (
	"github.com/micro/go-micro/client"
	baseproto "microservice/jzapi/proto/base"
)


var (
	baseClient baseproto.BaseService
)

func Init() {
	baseClient = baseproto.NewBaseService("jz.micro.jzapi.srv.base", client.DefaultClient)
}
