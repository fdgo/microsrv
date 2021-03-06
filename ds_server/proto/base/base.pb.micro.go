// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: base.proto

package ds_srv_base

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

// Client API for Base service

type BaseService interface {
	VfCode(ctx context.Context, in *VfCodeIn, opts ...client.CallOption) (*CommonOut, error)
}

type baseService struct {
	c    client.Client
	name string
}

func NewBaseService(name string, c client.Client) BaseService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "ds.srv.base"
	}
	return &baseService{
		c:    c,
		name: name,
	}
}

func (c *baseService) VfCode(ctx context.Context, in *VfCodeIn, opts ...client.CallOption) (*CommonOut, error) {
	req := c.c.NewRequest(c.name, "Base.VfCode", in)
	out := new(CommonOut)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Base service

type BaseHandler interface {
	VfCode(context.Context, *VfCodeIn, *CommonOut) error
}

func RegisterBaseHandler(s server.Server, hdlr BaseHandler, opts ...server.HandlerOption) error {
	type base interface {
		VfCode(ctx context.Context, in *VfCodeIn, out *CommonOut) error
	}
	type Base struct {
		base
	}
	h := &baseHandler{hdlr}
	return s.Handle(s.NewHandler(&Base{h}, opts...))
}

type baseHandler struct {
	BaseHandler
}

func (h *baseHandler) VfCode(ctx context.Context, in *VfCodeIn, out *CommonOut) error {
	return h.BaseHandler.VfCode(ctx, in, out)
}
