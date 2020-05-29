package service

import (
	"context"
	"encoding/json"
	mb "microservice/jzapi/base_srv/mobilecode"
	r6 "microservice/jzapi/basic/randstring"
	tr "microservice/jzapi/basic/time_ex"
	baseproto "microservice/jzapi/proto/base"
)

func (e *service)  SendMobileCode(ctx context.Context, in *baseproto.SendMobileCodeIn) *baseproto.CommonOutput {
	out := new(baseproto.CommonOutput)
	mct := &mb.MobileCodeTime{
		Times:0,
		Mobile:in.Mobile,
		Code:r6.Rand6NumString(),
		Time:tr.GetCurrentTimeStamp(),
	}
	err := mb.SendMobileCode(Ca,   mct, int64(in.Times) ,in.Timediff)
	if err!=nil{
		out.HttpCode = 400
		out.Code = 400
		out.Msg = err.Error()
		out.Detail=err.Error()
		out.Data = []byte("false")
		return out
	}
	out.HttpCode = 200
	out.Code = 200
	out.Msg = "验证码发送成功！"
	out.Detail = "验证码发送成功！"
	out.Data = []byte("true")
	return out
}
func (e *service)  IsMobileCodeOk (ctx context.Context, in *baseproto.IsMobileCodeOkIn) *baseproto.CommonOutput {
	out := new(baseproto.CommonOutput)
	err,data := mb.GetCodeFromRedis(Ca,mb.CodeMobileKeyPrefix+in.Mobile)
	if err!=nil{
		out.HttpCode = 400
		out.Code = 400
		out.Msg = "验证码错误！"
		out.Detail = "验证码错误！"
		out.Data = []byte("false")
		return out
	}
	var tmp mb.MobileCodeTime
	json.Unmarshal([]byte(data),&tmp)
	if  in.Code != tmp.Code{
		out.HttpCode = 400
		out.Code = 400
		out.Msg = "验证码错误！"
		out.Detail = "验证码错误！"
		out.Data =[]byte("false")
		return out
	}
	if tr.GetCurrentTimeStamp() - tmp.Time > int64(in.Timediff){
		out.Detail = "验证码已经过期!"
		out.Msg = "验证码已经过期!"
		out.HttpCode = 400
		out.Code = 400
		out.Data = []byte("false")
		return out
	}
	out.Detail = "验证【验证码】成功！"
	out.Msg = "验证【验证码】成功！"
	out.HttpCode = 200
	out.Code = 200
	out.Data =[]byte("true")
	return out
}
