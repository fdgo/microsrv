package login

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	errcode "microservice/jzapi/basic/error_code"
	mresp "microservice/jzapi/basic/model_response"
	repkg "microservice/jzapi/basic/return_pkg"
	"microservice/jzapi/lib/wrapper/tracer/opentracing/gin2micro"
	baseproto "microservice/jzapi/proto/base"
	userproto "microservice/jzapi/proto/user"
	hdler "microservice/jzapi/user_web/srvclient"
	"net/http"
)

func MultiLoginMobile(c *gin.Context) {
	ctx, ok := gin2micro.ContextWithSpan(c)
	if ok == false {
		//log.Log("get context err")
	}
	type input struct {
		Mobile   string `json:"mobile" binding:"required"`
		Code     string `json:"code" binding:"required"`
		TimeDiff int32  `json:"timediff" binding:"required"`
	}
	var outrsp mresp.User_MultiLoginMobile
	var in input
	if err := c.ShouldBindJSON(&in); err != nil {
		repkg.GinResponse(http.StatusBadRequest, errcode.ERROR_COMMON_PARAM, errcode.ERROR_COMMON_PARAM_MSG+err.Error(), errcode.ERROR_COMMON_PARAM_MSG+err.Error(), outrsp, c)
		return
	}
	outcode, _ := hdler.BaseSer.IsMobileCodeOk(ctx, &baseproto.IsMobileCodeOkIn{
		Mobile:   in.Mobile,
		Code:     in.Code,
		Timediff: in.TimeDiff,
	})
	if string(outcode.Data) == "false" {
		repkg.GinResponse(http.StatusBadRequest, http.StatusBadRequest, outcode.Msg, outcode.Detail, outrsp, c)
		return
	}
	out, _ := hdler.UserSer.MultiLoginMobile(ctx, &userproto.MultiLoginMobileInput{
		Mobile: in.Mobile,
		Code:   in.Code,
	})
	if out == nil {
		repkg.GinResponse(http.StatusInternalServerError, errcode.ERROR_USERSRV_BASE, "服务器正忙，稍后重试！", "服务器正忙，稍后重试！", outrsp, c)
		return
	}
	json.Unmarshal([]byte(out.Data), &outrsp)
	repkg.GinResponse(out.HttpCode, out.Code, out.Msg, out.Detail, outrsp, c)
}
