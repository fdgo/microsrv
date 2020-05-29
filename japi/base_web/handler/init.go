package handler
import (
	baseSer "microservice/jzapi/proto/base"
	"github.com/micro/go-micro/client"
)
var (
	BaseSer baseSer.BaseService
)
func NewServices(client client.Client) {
	BaseSer = baseSer.NewBaseService("jz.micro.jzapi.srv.base", client)
}

