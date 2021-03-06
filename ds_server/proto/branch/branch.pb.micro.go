// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: branch.proto

package ds_srv_branch

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

// Client API for Branch service

type BranchService interface {
	GetBranchByID(ctx context.Context, in *IdRequest, opts ...client.CallOption) (*Response, error)
	GetBranchDynamicByID(ctx context.Context, in *IdRequest, opts ...client.CallOption) (*Response, error)
	SelectBranch(ctx context.Context, in *SelectBranchRequest, opts ...client.CallOption) (*Response, error)
	CreateBranch(ctx context.Context, in *CreateRequest, opts ...client.CallOption) (*Response, error)
	CreateBranchDynamic(ctx context.Context, in *CreateBranchDynamicRequest, opts ...client.CallOption) (*Response, error)
	CreateBranchUrl(ctx context.Context, in *UrlRequest, opts ...client.CallOption) (*Response, error)
	GetBranchID(ctx context.Context, in *IdRequest, opts ...client.CallOption) (*IdResponse, error)
	SelectBranchDynamics(ctx context.Context, in *SelectBranchDynamicsRequest, opts ...client.CallOption) (*Response, error)
	GetAddress(ctx context.Context, in *IdRequest, opts ...client.CallOption) (*AddResponse, error)
}

type branchService struct {
	c    client.Client
	name string
}

func NewBranchService(name string, c client.Client) BranchService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "ds.srv.branch"
	}
	return &branchService{
		c:    c,
		name: name,
	}
}

func (c *branchService) GetBranchByID(ctx context.Context, in *IdRequest, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "Branch.GetBranchByID", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *branchService) GetBranchDynamicByID(ctx context.Context, in *IdRequest, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "Branch.GetBranchDynamicByID", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *branchService) SelectBranch(ctx context.Context, in *SelectBranchRequest, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "Branch.SelectBranch", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *branchService) CreateBranch(ctx context.Context, in *CreateRequest, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "Branch.CreateBranch", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *branchService) CreateBranchDynamic(ctx context.Context, in *CreateBranchDynamicRequest, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "Branch.CreateBranchDynamic", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *branchService) CreateBranchUrl(ctx context.Context, in *UrlRequest, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "Branch.CreateBranchUrl", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *branchService) GetBranchID(ctx context.Context, in *IdRequest, opts ...client.CallOption) (*IdResponse, error) {
	req := c.c.NewRequest(c.name, "Branch.GetBranchID", in)
	out := new(IdResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *branchService) SelectBranchDynamics(ctx context.Context, in *SelectBranchDynamicsRequest, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "Branch.SelectBranchDynamics", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *branchService) GetAddress(ctx context.Context, in *IdRequest, opts ...client.CallOption) (*AddResponse, error) {
	req := c.c.NewRequest(c.name, "Branch.GetAddress", in)
	out := new(AddResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Branch service

type BranchHandler interface {
	GetBranchByID(context.Context, *IdRequest, *Response) error
	GetBranchDynamicByID(context.Context, *IdRequest, *Response) error
	SelectBranch(context.Context, *SelectBranchRequest, *Response) error
	CreateBranch(context.Context, *CreateRequest, *Response) error
	CreateBranchDynamic(context.Context, *CreateBranchDynamicRequest, *Response) error
	CreateBranchUrl(context.Context, *UrlRequest, *Response) error
	GetBranchID(context.Context, *IdRequest, *IdResponse) error
	SelectBranchDynamics(context.Context, *SelectBranchDynamicsRequest, *Response) error
	GetAddress(context.Context, *IdRequest, *AddResponse) error
}

func RegisterBranchHandler(s server.Server, hdlr BranchHandler, opts ...server.HandlerOption) error {
	type branch interface {
		GetBranchByID(ctx context.Context, in *IdRequest, out *Response) error
		GetBranchDynamicByID(ctx context.Context, in *IdRequest, out *Response) error
		SelectBranch(ctx context.Context, in *SelectBranchRequest, out *Response) error
		CreateBranch(ctx context.Context, in *CreateRequest, out *Response) error
		CreateBranchDynamic(ctx context.Context, in *CreateBranchDynamicRequest, out *Response) error
		CreateBranchUrl(ctx context.Context, in *UrlRequest, out *Response) error
		GetBranchID(ctx context.Context, in *IdRequest, out *IdResponse) error
		SelectBranchDynamics(ctx context.Context, in *SelectBranchDynamicsRequest, out *Response) error
		GetAddress(ctx context.Context, in *IdRequest, out *AddResponse) error
	}
	type Branch struct {
		branch
	}
	h := &branchHandler{hdlr}
	return s.Handle(s.NewHandler(&Branch{h}, opts...))
}

type branchHandler struct {
	BranchHandler
}

func (h *branchHandler) GetBranchByID(ctx context.Context, in *IdRequest, out *Response) error {
	return h.BranchHandler.GetBranchByID(ctx, in, out)
}

func (h *branchHandler) GetBranchDynamicByID(ctx context.Context, in *IdRequest, out *Response) error {
	return h.BranchHandler.GetBranchDynamicByID(ctx, in, out)
}

func (h *branchHandler) SelectBranch(ctx context.Context, in *SelectBranchRequest, out *Response) error {
	return h.BranchHandler.SelectBranch(ctx, in, out)
}

func (h *branchHandler) CreateBranch(ctx context.Context, in *CreateRequest, out *Response) error {
	return h.BranchHandler.CreateBranch(ctx, in, out)
}

func (h *branchHandler) CreateBranchDynamic(ctx context.Context, in *CreateBranchDynamicRequest, out *Response) error {
	return h.BranchHandler.CreateBranchDynamic(ctx, in, out)
}

func (h *branchHandler) CreateBranchUrl(ctx context.Context, in *UrlRequest, out *Response) error {
	return h.BranchHandler.CreateBranchUrl(ctx, in, out)
}

func (h *branchHandler) GetBranchID(ctx context.Context, in *IdRequest, out *IdResponse) error {
	return h.BranchHandler.GetBranchID(ctx, in, out)
}

func (h *branchHandler) SelectBranchDynamics(ctx context.Context, in *SelectBranchDynamicsRequest, out *Response) error {
	return h.BranchHandler.SelectBranchDynamics(ctx, in, out)
}

func (h *branchHandler) GetAddress(ctx context.Context, in *IdRequest, out *AddResponse) error {
	return h.BranchHandler.GetAddress(ctx, in, out)
}
