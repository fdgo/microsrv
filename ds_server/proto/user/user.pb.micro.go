// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: user.proto

package ds_srv_user

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
)

// Reference imports to suppress errex if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Reference imports to suppress errex if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for User service

type UserService interface {
	Regist(ctx context.Context, in *RegistIn, opts ...client.CallOption) (*CommonOut, error)
	Login(ctx context.Context, in *LoginIn, opts ...client.CallOption) (*CommonOut, error)
	ModifyBasicPwd(ctx context.Context, in *ModifyBasicPwdIn, opts ...client.CallOption) (*CommonOut, error)
	ModifyLoginPwd(ctx context.Context, in *ModifyLoginPwdIn, opts ...client.CallOption) (*CommonOut, error)
	ModifyPayPwd(ctx context.Context, in *ModifyPayPwdIn, opts ...client.CallOption) (*CommonOut, error)
	SetPaypwd(ctx context.Context, in *SetPaypwdIn, opts ...client.CallOption) (*CommonOut, error)
	GetMemberUserAgent(ctx context.Context, in *GetMemberUserAgentIn, opts ...client.CallOption) (*CommonOut, error)
	ConnectUs(ctx context.Context, in *ConnectUsIn, opts ...client.CallOption) (*CommonOut, error)
	Stream(ctx context.Context, opts ...client.CallOption) (User_StreamService, error)
	ServerStream(ctx context.Context, in *WsIn, opts ...client.CallOption) (User_ServerStreamService, error)
	ExchangeRate(ctx context.Context, in *ExchangeRateIn, opts ...client.CallOption) (*CommonOut, error)
	MemberDeposit(ctx context.Context, in *MemberDepositIn, opts ...client.CallOption) (*CommonOut, error)
	OnlinePay(ctx context.Context, in *OnlinePayIn, opts ...client.CallOption) (*CommonOut, error)
	MemberDepositLog(ctx context.Context, in *MemberDepositLogIn, opts ...client.CallOption) (*CommonOut, error)
	MemerClassSet(ctx context.Context, in *MemerClassSetIn, opts ...client.CallOption) (*CommonOut, error)
	AgentClassSet(ctx context.Context, in *AgentClassSetIn, opts ...client.CallOption) (*CommonOut, error)
	MemberUsdtRecharge(ctx context.Context, in *MemberUsdtRechargeIn, opts ...client.CallOption) (*CommonOut, error)
}

type userService struct {
	c    client.Client
	name string
}

func NewUserService(name string, c client.Client) UserService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "ds.srv.user"
	}
	return &userService{
		c:    c,
		name: name,
	}
}

func (c *userService) Regist(ctx context.Context, in *RegistIn, opts ...client.CallOption) (*CommonOut, error) {
	req := c.c.NewRequest(c.name, "User.Regist", in)
	out := new(CommonOut)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) Login(ctx context.Context, in *LoginIn, opts ...client.CallOption) (*CommonOut, error) {
	req := c.c.NewRequest(c.name, "User.Login", in)
	out := new(CommonOut)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) ModifyBasicPwd(ctx context.Context, in *ModifyBasicPwdIn, opts ...client.CallOption) (*CommonOut, error) {
	req := c.c.NewRequest(c.name, "User.ModifyBasicPwd", in)
	out := new(CommonOut)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) ModifyLoginPwd(ctx context.Context, in *ModifyLoginPwdIn, opts ...client.CallOption) (*CommonOut, error) {
	req := c.c.NewRequest(c.name, "User.ModifyLoginPwd", in)
	out := new(CommonOut)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) ModifyPayPwd(ctx context.Context, in *ModifyPayPwdIn, opts ...client.CallOption) (*CommonOut, error) {
	req := c.c.NewRequest(c.name, "User.ModifyPayPwd", in)
	out := new(CommonOut)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) SetPaypwd(ctx context.Context, in *SetPaypwdIn, opts ...client.CallOption) (*CommonOut, error) {
	req := c.c.NewRequest(c.name, "User.SetPaypwd", in)
	out := new(CommonOut)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) GetMemberUserAgent(ctx context.Context, in *GetMemberUserAgentIn, opts ...client.CallOption) (*CommonOut, error) {
	req := c.c.NewRequest(c.name, "User.GetMemberUserAgent", in)
	out := new(CommonOut)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) ConnectUs(ctx context.Context, in *ConnectUsIn, opts ...client.CallOption) (*CommonOut, error) {
	req := c.c.NewRequest(c.name, "User.ConnectUs", in)
	out := new(CommonOut)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) Stream(ctx context.Context, opts ...client.CallOption) (User_StreamService, error) {
	req := c.c.NewRequest(c.name, "User.Stream", &WsIn{})
	stream, err := c.c.Stream(ctx, req, opts...)
	if err != nil {
		return nil, err
	}
	return &userServiceStream{stream}, nil
}

type User_StreamService interface {
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*WsIn) error
	Recv() (*CommonOut, error)
}

type userServiceStream struct {
	stream client.Stream
}

func (x *userServiceStream) Close() error {
	return x.stream.Close()
}

func (x *userServiceStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *userServiceStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *userServiceStream) Send(m *WsIn) error {
	return x.stream.Send(m)
}

func (x *userServiceStream) Recv() (*CommonOut, error) {
	m := new(CommonOut)
	err := x.stream.Recv(m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

func (c *userService) ServerStream(ctx context.Context, in *WsIn, opts ...client.CallOption) (User_ServerStreamService, error) {
	req := c.c.NewRequest(c.name, "User.ServerStream", &WsIn{})
	stream, err := c.c.Stream(ctx, req, opts...)
	if err != nil {
		return nil, err
	}
	if err := stream.Send(in); err != nil {
		return nil, err
	}
	return &userServiceServerStream{stream}, nil
}

type User_ServerStreamService interface {
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Recv() (*CommonOut, error)
}

type userServiceServerStream struct {
	stream client.Stream
}

func (x *userServiceServerStream) Close() error {
	return x.stream.Close()
}

func (x *userServiceServerStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *userServiceServerStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *userServiceServerStream) Recv() (*CommonOut, error) {
	m := new(CommonOut)
	err := x.stream.Recv(m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

func (c *userService) ExchangeRate(ctx context.Context, in *ExchangeRateIn, opts ...client.CallOption) (*CommonOut, error) {
	req := c.c.NewRequest(c.name, "User.ExchangeRate", in)
	out := new(CommonOut)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) MemberDeposit(ctx context.Context, in *MemberDepositIn, opts ...client.CallOption) (*CommonOut, error) {
	req := c.c.NewRequest(c.name, "User.MemberDeposit", in)
	out := new(CommonOut)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) OnlinePay(ctx context.Context, in *OnlinePayIn, opts ...client.CallOption) (*CommonOut, error) {
	req := c.c.NewRequest(c.name, "User.OnlinePay", in)
	out := new(CommonOut)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) MemberDepositLog(ctx context.Context, in *MemberDepositLogIn, opts ...client.CallOption) (*CommonOut, error) {
	req := c.c.NewRequest(c.name, "User.MemberDepositLog", in)
	out := new(CommonOut)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) MemerClassSet(ctx context.Context, in *MemerClassSetIn, opts ...client.CallOption) (*CommonOut, error) {
	req := c.c.NewRequest(c.name, "User.MemerClassSet", in)
	out := new(CommonOut)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) AgentClassSet(ctx context.Context, in *AgentClassSetIn, opts ...client.CallOption) (*CommonOut, error) {
	req := c.c.NewRequest(c.name, "User.AgentClassSet", in)
	out := new(CommonOut)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) MemberUsdtRecharge(ctx context.Context, in *MemberUsdtRechargeIn, opts ...client.CallOption) (*CommonOut, error) {
	req := c.c.NewRequest(c.name, "User.MemberUsdtRecharge", in)
	out := new(CommonOut)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for User service

type UserHandler interface {
	Regist(context.Context, *RegistIn, *CommonOut) error
	Login(context.Context, *LoginIn, *CommonOut) error
	ModifyBasicPwd(context.Context, *ModifyBasicPwdIn, *CommonOut) error
	ModifyLoginPwd(context.Context, *ModifyLoginPwdIn, *CommonOut) error
	ModifyPayPwd(context.Context, *ModifyPayPwdIn, *CommonOut) error
	SetPaypwd(context.Context, *SetPaypwdIn, *CommonOut) error
	GetMemberUserAgent(context.Context, *GetMemberUserAgentIn, *CommonOut) error
	ConnectUs(context.Context, *ConnectUsIn, *CommonOut) error
	Stream(context.Context, User_StreamStream) error
	ServerStream(context.Context, *WsIn, User_ServerStreamStream) error
	ExchangeRate(context.Context, *ExchangeRateIn, *CommonOut) error
	MemberDeposit(context.Context, *MemberDepositIn, *CommonOut) error
	OnlinePay(context.Context, *OnlinePayIn, *CommonOut) error
	MemberDepositLog(context.Context, *MemberDepositLogIn, *CommonOut) error
	MemerClassSet(context.Context, *MemerClassSetIn, *CommonOut) error
	AgentClassSet(context.Context, *AgentClassSetIn, *CommonOut) error
	MemberUsdtRecharge(context.Context, *MemberUsdtRechargeIn, *CommonOut) error
}

func RegisterUserHandler(s server.Server, hdlr UserHandler, opts ...server.HandlerOption) error {
	type user interface {
		Regist(ctx context.Context, in *RegistIn, out *CommonOut) error
		Login(ctx context.Context, in *LoginIn, out *CommonOut) error
		ModifyBasicPwd(ctx context.Context, in *ModifyBasicPwdIn, out *CommonOut) error
		ModifyLoginPwd(ctx context.Context, in *ModifyLoginPwdIn, out *CommonOut) error
		ModifyPayPwd(ctx context.Context, in *ModifyPayPwdIn, out *CommonOut) error
		SetPaypwd(ctx context.Context, in *SetPaypwdIn, out *CommonOut) error
		GetMemberUserAgent(ctx context.Context, in *GetMemberUserAgentIn, out *CommonOut) error
		ConnectUs(ctx context.Context, in *ConnectUsIn, out *CommonOut) error
		Stream(ctx context.Context, stream server.Stream) error
		ServerStream(ctx context.Context, stream server.Stream) error
		ExchangeRate(ctx context.Context, in *ExchangeRateIn, out *CommonOut) error
		MemberDeposit(ctx context.Context, in *MemberDepositIn, out *CommonOut) error
		OnlinePay(ctx context.Context, in *OnlinePayIn, out *CommonOut) error
		MemberDepositLog(ctx context.Context, in *MemberDepositLogIn, out *CommonOut) error
		MemerClassSet(ctx context.Context, in *MemerClassSetIn, out *CommonOut) error
		AgentClassSet(ctx context.Context, in *AgentClassSetIn, out *CommonOut) error
		MemberUsdtRecharge(ctx context.Context, in *MemberUsdtRechargeIn, out *CommonOut) error
	}
	type User struct {
		user
	}
	h := &userHandler{hdlr}
	return s.Handle(s.NewHandler(&User{h}, opts...))
}

type userHandler struct {
	UserHandler
}

func (h *userHandler) Regist(ctx context.Context, in *RegistIn, out *CommonOut) error {
	return h.UserHandler.Regist(ctx, in, out)
}

func (h *userHandler) Login(ctx context.Context, in *LoginIn, out *CommonOut) error {
	return h.UserHandler.Login(ctx, in, out)
}

func (h *userHandler) ModifyBasicPwd(ctx context.Context, in *ModifyBasicPwdIn, out *CommonOut) error {
	return h.UserHandler.ModifyBasicPwd(ctx, in, out)
}

func (h *userHandler) ModifyLoginPwd(ctx context.Context, in *ModifyLoginPwdIn, out *CommonOut) error {
	return h.UserHandler.ModifyLoginPwd(ctx, in, out)
}

func (h *userHandler) ModifyPayPwd(ctx context.Context, in *ModifyPayPwdIn, out *CommonOut) error {
	return h.UserHandler.ModifyPayPwd(ctx, in, out)
}

func (h *userHandler) SetPaypwd(ctx context.Context, in *SetPaypwdIn, out *CommonOut) error {
	return h.UserHandler.SetPaypwd(ctx, in, out)
}

func (h *userHandler) GetMemberUserAgent(ctx context.Context, in *GetMemberUserAgentIn, out *CommonOut) error {
	return h.UserHandler.GetMemberUserAgent(ctx, in, out)
}

func (h *userHandler) ConnectUs(ctx context.Context, in *ConnectUsIn, out *CommonOut) error {
	return h.UserHandler.ConnectUs(ctx, in, out)
}

func (h *userHandler) Stream(ctx context.Context, stream server.Stream) error {
	return h.UserHandler.Stream(ctx, &userStreamStream{stream})
}

type User_StreamStream interface {
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*CommonOut) error
	Recv() (*WsIn, error)
}

type userStreamStream struct {
	stream server.Stream
}

func (x *userStreamStream) Close() error {
	return x.stream.Close()
}

func (x *userStreamStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *userStreamStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *userStreamStream) Send(m *CommonOut) error {
	return x.stream.Send(m)
}

func (x *userStreamStream) Recv() (*WsIn, error) {
	m := new(WsIn)
	if err := x.stream.Recv(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (h *userHandler) ServerStream(ctx context.Context, stream server.Stream) error {
	m := new(WsIn)
	if err := stream.Recv(m); err != nil {
		return err
	}
	return h.UserHandler.ServerStream(ctx, m, &userServerStreamStream{stream})
}

type User_ServerStreamStream interface {
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*CommonOut) error
}

type userServerStreamStream struct {
	stream server.Stream
}

func (x *userServerStreamStream) Close() error {
	return x.stream.Close()
}

func (x *userServerStreamStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *userServerStreamStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *userServerStreamStream) Send(m *CommonOut) error {
	return x.stream.Send(m)
}

func (h *userHandler) ExchangeRate(ctx context.Context, in *ExchangeRateIn, out *CommonOut) error {
	return h.UserHandler.ExchangeRate(ctx, in, out)
}

func (h *userHandler) MemberDeposit(ctx context.Context, in *MemberDepositIn, out *CommonOut) error {
	return h.UserHandler.MemberDeposit(ctx, in, out)
}

func (h *userHandler) OnlinePay(ctx context.Context, in *OnlinePayIn, out *CommonOut) error {
	return h.UserHandler.OnlinePay(ctx, in, out)
}

func (h *userHandler) MemberDepositLog(ctx context.Context, in *MemberDepositLogIn, out *CommonOut) error {
	return h.UserHandler.MemberDepositLog(ctx, in, out)
}

func (h *userHandler) MemerClassSet(ctx context.Context, in *MemerClassSetIn, out *CommonOut) error {
	return h.UserHandler.MemerClassSet(ctx, in, out)
}

func (h *userHandler) AgentClassSet(ctx context.Context, in *AgentClassSetIn, out *CommonOut) error {
	return h.UserHandler.AgentClassSet(ctx, in, out)
}

func (h *userHandler) MemberUsdtRecharge(ctx context.Context, in *MemberUsdtRechargeIn, out *CommonOut) error {
	return h.UserHandler.MemberUsdtRecharge(ctx, in, out)
}
