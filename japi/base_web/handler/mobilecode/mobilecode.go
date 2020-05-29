package mobilecode

import (
	"github.com/gin-gonic/gin"
	hdler "microservice/jzapi/base_web/handler"
	errcode "microservice/jzapi/basic/error_code"
	repkg "microservice/jzapi/basic/return_pkg"
	"microservice/jzapi/lib/wrapper/tracer/opentracing/gin2micro"
	baseproto "microservice/jzapi/proto/base"
	"net/http"
	"strconv"
)

//1.发送验证码
func SendMobileCode(c *gin.Context) {
	ctx, ok := gin2micro.ContextWithSpan(c)
	if ok == false {
		//log.Log("get context err")
	}
	type input struct {
		Mobile   string `json:"mobile" binding:"required"`
		Times    int32  `json:"times" binding:"required"`
		Timediff int64  `json:"timediff" binding:"required"`
	}
	var in input
	if err := c.ShouldBindJSON(&in); err != nil {
		repkg.GinResponse(http.StatusBadRequest, errcode.ERROR_COMMON_PARAM, errcode.ERROR_COMMON_PARAM_MSG+err.Error(), errcode.ERROR_COMMON_PARAM_MSG+err.Error(), "NULL", c)
		return
	}
	out, _ := hdler.BaseSer.SendMobileCode(ctx, &baseproto.SendMobileCodeIn{
		Mobile:   in.Mobile,
		Times:    in.Times,
		Timediff: in.Timediff,
	})
	if out == nil {
		repkg.GinResponse(http.StatusInternalServerError, errcode.ERROR_BASESRV_BASE, "服务器正忙，稍后重试！", "服务器正忙，稍后重试！",  string(out.Data), c)
		return
	}
	repkg.GinResponse(out.HttpCode, out.Code, out.Msg, out.Detail,  string(out.Data), c)
}

func IsMobileCodeOk(c *gin.Context) {
	//fmt.Println(c.Request.Header["X-Head-Username"])
	//fmt.Println(c.Request.Header["X-Head-Accountid"])
	//fmt.Println(c.Request.Header["X-Head-Timestamp"])
	ctx, ok := gin2micro.ContextWithSpan(c)
	if ok == false {
		//log.Log("get context err")
	}
	code := c.Query("code")
	tf,_ := strconv.Atoi(c.Query("timediff"))
	mobile := c.Query("mobile")
	out, _ := hdler.BaseSer.IsMobileCodeOk(ctx, &baseproto.IsMobileCodeOkIn{
		Mobile:  mobile ,
		Timediff: int32(tf),
		Code:code,
	})
	if out == nil {
		repkg.GinResponse(http.StatusInternalServerError, errcode.ERROR_BASESRV_BASE, "服务器正忙，稍后重试！", "服务器正忙，稍后重试！", string(out.Data), c)
		return
	}
	repkg.GinResponse(out.HttpCode, out.Code, out.Msg, out.Detail, string(out.Data), c)
}
