package service

import (
	"context"
	"fmt"
	"github.com/micro/go-micro/client"
	adminproto "microservice/jzapi/proto/admin"
	baseproto "microservice/jzapi/proto/base"
	userproto "microservice/jzapi/proto/user"
	"sync"
)

var (
	srv        *service
	rw         sync.RWMutex
	baseClient baseproto.BaseService
	userClient userproto.UserService
)

type service struct{}

type ServiceInterface interface {
	GetAccessList(ctx context.Context, in *adminproto.GetAccessListInput) *adminproto.CommonOutput
	GetRoleList(ctx context.Context, in *adminproto.GetRoleListInput) *adminproto.CommonOutput
	AddRole(ctx context.Context, in *adminproto.AddRoleInput) *adminproto.CommonOutput
	EditRole(ctx context.Context, in *adminproto.EditRoleInput) *adminproto.CommonOutput
	DeleteRole(ctx context.Context, in *adminproto.DeleteRoleInput) *adminproto.CommonOutput
}

func GetService() (ServiceInterface, error) {
	if srv == nil {
		return nil, fmt.Errorf("[GetService] GetService 未初始化")
	}
	return srv, nil
}
func Init() {
	rw.Lock()
	defer rw.Unlock()

	if srv != nil {
		return
	}
	baseClient = baseproto.NewBaseService("jz.micro.jzapi.srv.base", client.DefaultClient)
	userClient = userproto.NewUserService("jz.micro.jzapi.srv.user", client.DefaultClient)
	srv = &service{}
}
