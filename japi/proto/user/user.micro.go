// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: user.proto

package user

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

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for User service

type UserService interface {
	MultiLoginMobile(ctx context.Context, in *MultiLoginMobileInput, opts ...client.CallOption) (*CommonOutput, error)
	SingleRegistQuick(ctx context.Context, in *SingleRegistQuickInput, opts ...client.CallOption) (*CommonOutput, error)
	SingleRegistAccount(ctx context.Context, in *SingleRegistAccountInput, opts ...client.CallOption) (*CommonOutput, error)
	SingleRegistMobile(ctx context.Context, in *SingleRegistMobileInput, opts ...client.CallOption) (*CommonOutput, error)
	SingleLoginGuest(ctx context.Context, in *SingleLoginGuestInput, opts ...client.CallOption) (*CommonOutput, error)
	SingleLoginMobile(ctx context.Context, in *SingleLoginMobileInput, opts ...client.CallOption) (*CommonOutput, error)
	SingleLoginAccount(ctx context.Context, in *SingleLoginAccountInput, opts ...client.CallOption) (*CommonOutput, error)
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
		name = "user"
	}
	return &userService{
		c:    c,
		name: name,
	}
}

func (c *userService) MultiLoginMobile(ctx context.Context, in *MultiLoginMobileInput, opts ...client.CallOption) (*CommonOutput, error) {
	req := c.c.NewRequest(c.name, "User.MultiLoginMobile", in)
	out := new(CommonOutput)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) SingleRegistQuick(ctx context.Context, in *SingleRegistQuickInput, opts ...client.CallOption) (*CommonOutput, error) {
	req := c.c.NewRequest(c.name, "User.SingleRegistQuick", in)
	out := new(CommonOutput)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) SingleRegistAccount(ctx context.Context, in *SingleRegistAccountInput, opts ...client.CallOption) (*CommonOutput, error) {
	req := c.c.NewRequest(c.name, "User.SingleRegistAccount", in)
	out := new(CommonOutput)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) SingleRegistMobile(ctx context.Context, in *SingleRegistMobileInput, opts ...client.CallOption) (*CommonOutput, error) {
	req := c.c.NewRequest(c.name, "User.SingleRegistMobile", in)
	out := new(CommonOutput)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) SingleLoginGuest(ctx context.Context, in *SingleLoginGuestInput, opts ...client.CallOption) (*CommonOutput, error) {
	req := c.c.NewRequest(c.name, "User.SingleLoginGuest", in)
	out := new(CommonOutput)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) SingleLoginMobile(ctx context.Context, in *SingleLoginMobileInput, opts ...client.CallOption) (*CommonOutput, error) {
	req := c.c.NewRequest(c.name, "User.SingleLoginMobile", in)
	out := new(CommonOutput)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) SingleLoginAccount(ctx context.Context, in *SingleLoginAccountInput, opts ...client.CallOption) (*CommonOutput, error) {
	req := c.c.NewRequest(c.name, "User.SingleLoginAccount", in)
	out := new(CommonOutput)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for User service

type UserHandler interface {
	MultiLoginMobile(context.Context, *MultiLoginMobileInput, *CommonOutput) error
	SingleRegistQuick(context.Context, *SingleRegistQuickInput, *CommonOutput) error
	SingleRegistAccount(context.Context, *SingleRegistAccountInput, *CommonOutput) error
	SingleRegistMobile(context.Context, *SingleRegistMobileInput, *CommonOutput) error
	SingleLoginGuest(context.Context, *SingleLoginGuestInput, *CommonOutput) error
	SingleLoginMobile(context.Context, *SingleLoginMobileInput, *CommonOutput) error
	SingleLoginAccount(context.Context, *SingleLoginAccountInput, *CommonOutput) error
}

func RegisterUserHandler(s server.Server, hdlr UserHandler, opts ...server.HandlerOption) error {
	type user interface {
		MultiLoginMobile(ctx context.Context, in *MultiLoginMobileInput, out *CommonOutput) error
		SingleRegistQuick(ctx context.Context, in *SingleRegistQuickInput, out *CommonOutput) error
		SingleRegistAccount(ctx context.Context, in *SingleRegistAccountInput, out *CommonOutput) error
		SingleRegistMobile(ctx context.Context, in *SingleRegistMobileInput, out *CommonOutput) error
		SingleLoginGuest(ctx context.Context, in *SingleLoginGuestInput, out *CommonOutput) error
		SingleLoginMobile(ctx context.Context, in *SingleLoginMobileInput, out *CommonOutput) error
		SingleLoginAccount(ctx context.Context, in *SingleLoginAccountInput, out *CommonOutput) error
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

func (h *userHandler) MultiLoginMobile(ctx context.Context, in *MultiLoginMobileInput, out *CommonOutput) error {
	return h.UserHandler.MultiLoginMobile(ctx, in, out)
}

func (h *userHandler) SingleRegistQuick(ctx context.Context, in *SingleRegistQuickInput, out *CommonOutput) error {
	return h.UserHandler.SingleRegistQuick(ctx, in, out)
}

func (h *userHandler) SingleRegistAccount(ctx context.Context, in *SingleRegistAccountInput, out *CommonOutput) error {
	return h.UserHandler.SingleRegistAccount(ctx, in, out)
}

func (h *userHandler) SingleRegistMobile(ctx context.Context, in *SingleRegistMobileInput, out *CommonOutput) error {
	return h.UserHandler.SingleRegistMobile(ctx, in, out)
}

func (h *userHandler) SingleLoginGuest(ctx context.Context, in *SingleLoginGuestInput, out *CommonOutput) error {
	return h.UserHandler.SingleLoginGuest(ctx, in, out)
}

func (h *userHandler) SingleLoginMobile(ctx context.Context, in *SingleLoginMobileInput, out *CommonOutput) error {
	return h.UserHandler.SingleLoginMobile(ctx, in, out)
}

func (h *userHandler) SingleLoginAccount(ctx context.Context, in *SingleLoginAccountInput, out *CommonOutput) error {
	return h.UserHandler.SingleLoginAccount(ctx, in, out)
}