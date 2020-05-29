package handler

import (
	"context"
	"github.com/micro/go-micro/util/log"
	adminservice "microservice/jzapi/admin_srv/service"
	adminproto "microservice/jzapi/proto/admin"
)

type Service struct{}

var (
	userService adminservice.ServiceInterface
)

func Init() {
	var err error
	userService, err = adminservice.GetService()
	if err != nil {
		log.Fatal("[Init] 初始化Handler错误")
		return
	}
}

func (s *Service) GetAccessList(ctx context.Context, in *adminproto.GetAccessListInput, out *adminproto.CommonOutput) error {
	tmpout := userService.GetAccessList(ctx, in)
	(*out).HttpCode = (*tmpout).HttpCode
	(*out).Code = (*tmpout).Code
	(*out).Msg = (*tmpout).Msg
	(*out).Detail = (*tmpout).Detail
	(*out).Data = (*tmpout).Data
	return nil
}
func (s *Service) GetRoleList(ctx context.Context, in *adminproto.GetRoleListInput, out *adminproto.CommonOutput) error {
	tmpout := userService.GetRoleList(ctx, in)
	(*out).HttpCode = (*tmpout).HttpCode
	(*out).Code = (*tmpout).Code
	(*out).Msg = (*tmpout).Msg
	(*out).Detail = (*tmpout).Detail
	(*out).Data = (*tmpout).Data
	return nil
}
func (s *Service) AddRole(ctx context.Context, in *adminproto.AddRoleInput, out *adminproto.CommonOutput) error {
	tmpout := userService.AddRole(ctx, in)
	(*out).HttpCode = (*tmpout).HttpCode
	(*out).Code = (*tmpout).Code
	(*out).Msg = (*tmpout).Msg
	(*out).Detail = (*tmpout).Detail
	(*out).Data = (*tmpout).Data
	return nil
}
func (s *Service) EditRole(ctx context.Context, in *adminproto.EditRoleInput, out *adminproto.CommonOutput) error {
	tmpout := userService.EditRole(ctx, in)
	(*out).HttpCode = (*tmpout).HttpCode
	(*out).Code = (*tmpout).Code
	(*out).Msg = (*tmpout).Msg
	(*out).Detail = (*tmpout).Detail
	(*out).Data = (*tmpout).Data
	return nil
}
func (s *Service) DeleteRole(ctx context.Context, in *adminproto.DeleteRoleInput, out *adminproto.CommonOutput) error {
	tmpout := userService.DeleteRole(ctx, in)
	(*out).HttpCode = (*tmpout).HttpCode
	(*out).Code = (*tmpout).Code
	(*out).Msg = (*tmpout).Msg
	(*out).Detail = (*tmpout).Detail
	(*out).Data = (*tmpout).Data
	return nil
}
