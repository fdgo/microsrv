package handler

import (
	"context"
	pb "ds_server/proto/base"
	srv "ds_server/services/base/service"
)

type BaseHandler struct {
	BaseSvr *srv.BaseService
}
func (basehdl *BaseHandler) VfCode(c context.Context, req *pb.VfCodeIn, rsp *pb.CommonOut) error {
	return basehdl.BaseSvr.VfCode(c,req,rsp)
}

