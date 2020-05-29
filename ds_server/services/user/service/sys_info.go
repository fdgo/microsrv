package service

import (
	"context"
	mygormdl "ds_server/models/user/gorm_mysql"
	pb "ds_server/proto/user"
	"ds_server/support/utils/rsp"
	"encoding/json"
	"errors"
)

func (usersrv *UserService) ConnectUs(ctx context.Context, req *pb.ConnectUsIn, rsq *pb.CommonOut) error {
	tmp_resp := mygormdl.DsSysInfo{}
	temp_resp_err, _ := json.Marshal(tmp_resp)
	ret_resp, _ := usersrv.DsSysInfoDao.ConnectUs()
	if len(ret_resp.ConnectUs) == 0 {
		rsp.RespSrv(400, 400, "获取公司链接失败!", "获取公司链接失败!", temp_resp_err, rsq)
		return errors.New("获取公司链接失败!")
	}
	ret_resp_nor, _ := json.Marshal(ret_resp)
	rsp.RespSrv(200, 200, "获取公司链接成功!", "获取公司链接成功!", ret_resp_nor, rsq)
	return nil
}
func (usersrv *UserService) ServerStream(ctx context.Context, req *pb.WsIn, rsq pb.User_ServerStreamStream) error {
	//usersrv.UserDao.ServerStream(ctx,req,rsq)
	//rsp.RespSrv(200, 200, "校验验证码成功!", "校验验证码成功!", []byte("true"), rsq)
	return nil
}
func (usersrv *UserService) Stream(ctx context.Context, rsq pb.User_StreamStream) error {
	//usersrv.UserDao.Stream(ctx,rsq)
	//rsp.RespSrv(200, 200, "校验验证码成功!", "校验验证码成功!", []byte("true"), rsq)
	return nil
}
