package dao

import (
	"context"
	pb "ds_server/proto/base"
	"ds_server/support/lib/redisex"
	"ds_server/support/utils/constex"
	"ds_server/support/utils/httpex"
	string_ex "ds_server/support/utils/stringex"
	"encoding/json"
	"errors"
	"net/url"
	"ds_server/models/user/gin_rsp"
)

type BaseDao struct {
}

func NewBaseDao() *BaseDao {
	return &BaseDao{}
}

// DsUserBasicinfoMgr open func
func (basedao *BaseDao) VfCode(c context.Context, req *pb.VfCodeIn, rsp *pb.CommonOut) error {
	rand_code := string_ex.Rand6NumString()
	tplValue := url.QueryEscape("#code#=" + rand_code)
	url := "httpex://v.juhe.cn/vercodesms/send.php?key=a9d9d80988b76855a1c0198d0148d1e3&tplId=66115&tplValue=" + tplValue + "&mobile=" + req.Mobile
	_, msgcode, _ := httpex.Get(url)
	var resp gin_rsp.MobileCode_rsp
	json.Unmarshal([]byte(msgcode),&resp)
	if resp.ErrorCode !=0{
		return errors.New("短信发送失败！")
	}
	redisex.RedisInstanceg().Set(constex.REDIS_USER_VFCODE+req.Mobile,rand_code,0)
	(*rsp).Httpcode = 200
	(*rsp).Msg = "success"
	(*rsp).Code = 200
	(*rsp).Data = []byte(msgcode)
	return nil
}
