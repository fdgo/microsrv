package handler

import (
	//"fmt"
	"context"
	pb "ds_server/proto/user"
	srv "ds_server/services/user/service"
)
type UserHandler struct {
	Usersrv *srv.UserService
}
func (userdle *UserHandler) Regist(c context.Context, req *pb.RegistIn, rsp *pb.CommonOut) error {
	return userdle.Usersrv.Regist(c,req,rsp)
}

func (userdle *UserHandler) Login(c context.Context, req *pb.LoginIn, rsp *pb.CommonOut) error {
	return userdle.Usersrv.Login(c,req,rsp)
}

func (userdle *UserHandler) GetMemberUserAgent(c context.Context, req *pb.GetMemberUserAgentIn, rsp *pb.CommonOut) error {
	return userdle.Usersrv.GetMemberUserAgent(c,req,rsp)
}

func (userdle *UserHandler) ConnectUs(c context.Context, req *pb.ConnectUsIn, rsp *pb.CommonOut) error {
	return userdle.Usersrv.ConnectUs(c,req,rsp)
}

func (userdle *UserHandler) ServerStream(c context.Context, req *pb.WsIn, rsp pb.User_ServerStreamStream) error {
	return userdle.Usersrv.ServerStream(c,req,rsp)
}

func (userdle *UserHandler) Stream(c context.Context,  rsp pb.User_StreamStream) error {
	return userdle.Usersrv.Stream(c,rsp)
}

func (userdle *UserHandler) ExchangeRate(c context.Context,  req *pb.ExchangeRateIn, rsp *pb.CommonOut) error {
	return userdle.Usersrv.ExchangeRate(c,req,rsp)
}

func (userdle *UserHandler) MemberDeposit(c context.Context,  req *pb.MemberDepositIn, rsp *pb.CommonOut) error {
	return userdle.Usersrv.MemberDeposit(c,req,rsp)
}

func (userdle *UserHandler) OnlinePay(c context.Context,  req *pb.OnlinePayIn, rsp *pb.CommonOut) error {
	return userdle.Usersrv.OnlinePay(c,req,rsp)
}
func (userdle *UserHandler) MemberDepositLog(c context.Context,  req *pb.MemberDepositLogIn, rsp *pb.CommonOut) error {
	return userdle.Usersrv.MemberDepositLog(c,req,rsp)
}

func (userdle *UserHandler) MemerClassSet(c context.Context,  req *pb.MemerClassSetIn, rsp *pb.CommonOut) error {
	return userdle.Usersrv.MemerClassSet(c,req,rsp)
}

func (userdle *UserHandler) ModifyBasicPwd(c context.Context, req *pb.ModifyBasicPwdIn, rsp *pb.CommonOut) error {
	return userdle.Usersrv.ModifyBasicPwd(c,req,rsp)
}

func (userdle *UserHandler) SetPaypwd (c context.Context, req *pb.SetPaypwdIn, rsp *pb.CommonOut) error {
	return userdle.Usersrv.SetPaypwd(c,req,rsp)
}

func (userdle *UserHandler) ModifyPayPwd(c context.Context,  req *pb.ModifyPayPwdIn, rsp *pb.CommonOut) error {
	return userdle.Usersrv.ModifyPayPwd(c,req,rsp)
}

func (userdle *UserHandler) AgentClassSet(c context.Context, req *pb.AgentClassSetIn, rsp *pb.CommonOut) error {
	return userdle.Usersrv.AgentClassSet(c,req,rsp)
}
func (userdle *UserHandler) MemberUsdtRecharge(c context.Context, req *pb.MemberUsdtRechargeIn, rsp *pb.CommonOut) error {
	return userdle.Usersrv.MemberUsdtRecharge(c,req,rsp)
}

func (userdle *UserHandler) ModifyLoginPwd(c context.Context, req *pb.ModifyLoginPwdIn, rsp *pb.CommonOut) error {
	return userdle.Usersrv.ModifyLoginPwd(c,req,rsp)
}

