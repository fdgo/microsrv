package base

import (
	"ds_server/client"
	reqmdl "ds_server/models/user/gin_req"
	rspmdl "ds_server/models/user/gin_rsp"
	baseproto "ds_server/proto/base"
	"ds_server/support/utils/param"
	"ds_server/support/utils/rsp"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
)

func VfCode(c *gin.Context) {
	var ret_resp rspmdl.MobileCode_rsp
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	var in reqmdl.MobileCode_req
	if err := c.ShouldBindJSON(&in); err != nil {
		rsp.RespGin(400, 400, "输入有误,请重写输入!", "参数有误", ret_resp, c)
		return
	}
	isok, _ := param.IsParam(in)
	if !isok {
		rsp.RespGin(400, 400, "输入有误,请重写输入!", "参数有误", ret_resp, c)
		return
	}
	var rtin baseproto.VfCodeIn
	rtin.Mobile = in.Mobile

	ret, err := client.BaseClient.VfCode(c, &rtin)
	if err != nil {
		rsp.RespGin(400, 400, "注册失败!", err.Error(), ret_resp, c)
		return
	}
	json.Unmarshal((*ret).Data, &ret_resp)
	rsp.RespGin((*ret).Httpcode, (*ret).Code, (*ret).Msg, (*ret).Innermsg, ret_resp, c)
}
