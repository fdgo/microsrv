package login

import (
	"github.com/gin-gonic/gin"
	hdler "microservice/jzapi/admin_web/handler"
	errcode "microservice/jzapi/basic/error_code"
	repkg "microservice/jzapi/basic/return_pkg"
	"microservice/jzapi/lib/wrapper/tracer/opentracing/gin2micro"
	userproto "microservice/jzapi/proto/user"
	"net/http"
)

func BackLogin(c *gin.Context) {
	ctx, ok := gin2micro.ContextWithSpan(c)
	if ok == false {
		//log.Log("get context err")
	}
	type input struct {
		BackName string `json:"back_name" binding:"required"`
		Password string `json:"password" binding:"required"`
	}
	var in input
	if err := c.ShouldBindJSON(&in); err != nil {
		repkg.GinResponse(http.StatusBadRequest, errcode.ERROR_COMMON_PARAM, errcode.ERROR_COMMON_PARAM_MSG+err.Error(), errcode.ERROR_COMMON_PARAM_MSG+err.Error(), []byte{}, c)
		return
	}
	out, _ := hdler.UserSer.BackLogin(ctx, &userproto.BackLoginInput{BackName: in.BackName,Password:in.Password})
	if out == nil {
		repkg.GinResponse(http.StatusInternalServerError, errcode.ERROR_USERSRV_BASE, "服务器正忙，稍后重试！", "服务器正忙，稍后重试！", []byte{}, c)
		return
	}
	repkg.GinResponse(out.HttpCode, out.Code, out.Msg, out.Detail, out.Data, c)
}
