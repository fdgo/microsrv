package srvclient

import (
	"github.com/micro/go-micro/client"
	baseSer "microservice/jzapi/proto/base"
	userSer "microservice/jzapi/proto/user"
)


var (
	UserSer userSer.UserService
	BaseSer baseSer.BaseService
)

func NewServices(client client.Client) {
	UserSer = userSer.NewUserService("jz.micro.jzapi.srv.user", client)
	BaseSer = baseSer.NewBaseService("jz.micro.jzapi.srv.base", client)
}
