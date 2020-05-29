package service

import (
	"context"
	pb "ds_server/proto/base"
	"ds_server/services/base/dao"
)

type BaseService struct {
	BaseDao *dao.BaseDao
}

func (basesrv *BaseService) VfCode(c context.Context, req *pb.VfCodeIn, rsp *pb.CommonOut) ( error) {
	basesrv.BaseDao.VfCode(c,req,rsp)
	return nil
}
