package handler

import (
	"context"
	"github.com/micro/go-micro/util/log"
	baseservice "microservice/jzapi/base_srv/service"
	baseproto "microservice/jzapi/proto/base"
)

var (
	baseServeice baseservice.Service
)

// Init 初始化handler
func Init() {
	var err error
	baseServeice, err = baseservice.GetService()
	if err != nil {
		log.Fatal("[Init] 初始化Handler错误，%s", err)
		return
	}
}

type Service struct{}

func (s *Service) IsMobileCodeOk(ctx context.Context, in *baseproto.IsMobileCodeOkIn, out *baseproto.CommonOutput) error {
	tmpout := baseServeice.IsMobileCodeOk(ctx, in)
	(*out).Data = (*tmpout).Data
	(*out).Msg = (*tmpout).Msg
	(*out).Code = (*tmpout).Code
	(*out).HttpCode = (*tmpout).HttpCode
	(*out).Detail = (*tmpout).Detail
	return nil
}
func (s *Service) SendMobileCode(ctx context.Context, in *baseproto.SendMobileCodeIn, out *baseproto.CommonOutput) error {
	tmpout := baseServeice.SendMobileCode(ctx, in)
	(*out).Data = (*tmpout).Data
	(*out).Msg = (*tmpout).Msg
	(*out).Code = (*tmpout).Code
	(*out).HttpCode = (*tmpout).HttpCode
	(*out).Detail = (*tmpout).Detail
	return nil
}
