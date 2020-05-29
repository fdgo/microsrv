package manage

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	hdler "microservice/jzapi/admin_web/handler"
	errcode "microservice/jzapi/basic/error_code"
	repkg "microservice/jzapi/basic/return_pkg"
	"microservice/jzapi/lib/wrapper/tracer/opentracing/gin2micro"
	userproto "microservice/jzapi/proto/user"
	"net/http"
)

func UserList(c *gin.Context) {
	ctx, ok := gin2micro.ContextWithSpan(c)
	if ok == false {
		//log.Log("get context err")
	}
	type src struct {
		Uuid      string `json:"uuid"`
		UserName  string `json:"user_name"`
		RealName  string `json:"real_name"`
		RoleId    string `json:"role_id"`
		RoleName  string `json:"role_name"`
		Mobile    string `json:"mobile"`
		Address   string `json:"address"`
		LastLogin string `json:"last_login"`
		ChannelId uint   `json:"channel_id"`
	}
	outrsp := make([]src, 0)
	out, _ := hdler.UserSer.GetBackUerInfo(ctx, &userproto.GetBackUerInfoInput{})
	if out == nil {
		repkg.GinResponse(http.StatusInternalServerError, errcode.ERROR_ADMINSRV_BASE, "服务器正忙，稍后重试！", "服务器正忙，稍后重试！", outrsp, c)
		return
	}
	json.Unmarshal(out.Data, &outrsp)

	repkg.GinResponse(out.HttpCode, out.Code, out.Msg, out.Detail, outrsp, c)
}

func DelUser(c *gin.Context) {
	ctx, ok := gin2micro.ContextWithSpan(c)
	if ok == false {
		//log.Log("get context err")
	}
	type Input struct {
		RealName string `json:"real_name"`
		Mobile   string `json:"mobile"`
		UserName string `json:"user_name"`
		NickName string `json:"nick_name"`
		Status   int32  `json:"status" `
	}
	var in Input
	if err := c.ShouldBindJSON(&in); err != nil {
		repkg.GinResponse(http.StatusBadRequest, errcode.ERROR_COMMON_PARAM, errcode.ERROR_COMMON_PARAM_MSG+err.Error(), errcode.ERROR_COMMON_PARAM_MSG+err.Error(), "", c)
		return
	}
	out, _ := hdler.UserSer.DelUser(ctx, &userproto.DelUserInput{
		RealName: in.RealName,
		Mobile:   in.Mobile,
		UserName: in.UserName,
		NickName: in.NickName,
		Status:   in.Status,
	})
	if out == nil {
		repkg.GinResponse(http.StatusInternalServerError, errcode.ERROR_USERSRV_BASE, "服务器正忙，稍后重试！", "服务器正忙，稍后重试！", "", c)
		return
	}
	repkg.GinResponse(out.HttpCode, out.Code, out.Msg, out.Detail, out.Data, c)
}

func AddUser(c *gin.Context) {
	ctx, ok := gin2micro.ContextWithSpan(c)
	if ok == false {
		//log.Log("get context err")
	}
	type Input struct {
		RealName  string `json:"real_name"`
		Gender    uint   `json:"gender"`
		Age       uint   `json:"age"`
		IdCard    string `json:"id_card"`
		Mobile    string `json:"mobile" binding:"required"`
		Address   string `json:"address" binding:"required"`
		RoleName  string `json:"role_name" binding:"required"`
		UserName  string `json:"user_name" binding:"required"`
		PassWord  string `json:"password" binding:"required"`
		NickName  string `json:"nick_name" `
		Avatar    string `json:"avatar" `
		ChannelId uint   `json:"channel_id" `
		LastLogin string `json:"last_login" `
		Status    uint   `json:"status" `
	}
	var in Input
	if err := c.ShouldBindJSON(&in); err != nil {
		repkg.GinResponse(http.StatusBadRequest, errcode.ERROR_COMMON_PARAM, errcode.ERROR_COMMON_PARAM_MSG+err.Error(), errcode.ERROR_COMMON_PARAM_MSG+err.Error(), "", c)
		return
	}
	out, _ := hdler.UserSer.AddUser(ctx, &userproto.AddUserInput{
		RealName:  in.RealName,
		Gender:    int32(in.Gender),
		Age:       int32(in.Age),
		IdCard:    in.IdCard,
		Mobile:    in.Mobile,
		Address:   in.Address,
		RoleName:  in.RoleName,
		UserName:  in.UserName,
		Password:  in.PassWord,
		NickName:  in.NickName,
		Avatar:    in.Avatar,
		ChannelId: int32(in.ChannelId),
		LastLogin: in.LastLogin,
		Status:    int32(in.Status),
	})
	if out == nil {
		repkg.GinResponse(http.StatusInternalServerError, errcode.ERROR_USERSRV_BASE, "服务器正忙，稍后重试！", "服务器正忙，稍后重试！", "", c)
		return
	}
	repkg.GinResponse(out.HttpCode, out.Code, out.Msg, out.Detail, out.Data, c)
}
