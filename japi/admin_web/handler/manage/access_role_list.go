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

func AddAccess(c *gin.Context) {
	ctx, ok := gin2micro.ContextWithSpan(c)
	if ok == false {
		//log.Log("get context err")
	}
	type input struct {
		AcsName string `json:"acs_name" binding:"required"`
		AcsTag  string `json:"acs_tag" binding:"required"`
	}
	var in input
	if err := c.ShouldBindJSON(&in); err != nil {
		repkg.GinResponse(http.StatusBadRequest, errcode.ERROR_COMMON_PARAM, errcode.ERROR_COMMON_PARAM_MSG+err.Error(), errcode.ERROR_COMMON_PARAM_MSG+err.Error(), []byte{}, c)
		return
	}
	out, _ := hdler.UserSer.AddAccess(ctx, &userproto.AddAccessInput{AcsName: in.AcsName,AcsTag:in.AcsTag})
	if out == nil {
		repkg.GinResponse(http.StatusInternalServerError, errcode.ERROR_USERSRV_BASE, "服务器正忙，稍后重试！", "服务器正忙，稍后重试！", []byte{}, c)
		return
	}
	repkg.GinResponse(out.HttpCode, out.Code, out.Msg, out.Detail, []byte{}, c)
}
func DelAccess(c *gin.Context) {
	ctx, ok := gin2micro.ContextWithSpan(c)
	if ok == false {
		//log.Log("get context err")
	}
	type input struct {
		AcsName string `json:"acs_name" binding:"required"`
		AcsTag  string `json:"acs_tag" binding:"required"`
	}
	var in input
	if err := c.ShouldBindJSON(&in); err != nil {
		repkg.GinResponse(http.StatusBadRequest, errcode.ERROR_COMMON_PARAM, errcode.ERROR_COMMON_PARAM_MSG+err.Error(), errcode.ERROR_COMMON_PARAM_MSG+err.Error(), []byte{}, c)
		return
	}
	out, _ := hdler.UserSer.DeleteAccess(ctx, &userproto.DeleteAccessInput{AcsName: in.AcsName,AcsTag:in.AcsTag})
	if out == nil {
		repkg.GinResponse(http.StatusInternalServerError, errcode.ERROR_USERSRV_BASE, "服务器正忙，稍后重试！", "服务器正忙，稍后重试！", []byte{}, c)
		return
	}
	repkg.GinResponse(out.HttpCode, out.Code, out.Msg, out.Detail, []byte{}, c)
}
func AccessList(c *gin.Context) {
	ctx, ok := gin2micro.ContextWithSpan(c)
	if ok == false {
		//log.Log("get context err")
	}
	type src struct {
		AcsId   string `json:"acs_id"`
		AcsName string `json:"acs_name"`
		AcsTag  string `json:"acs_tag"`
	}
	outrsp := make([]src, 0)
	out, _ := hdler.UserSer.GetAccessList(ctx, &userproto.GetAccessListInput{})
	if out == nil {
		repkg.GinResponse(http.StatusInternalServerError, errcode.ERROR_USERSRV_BASE, "服务器正忙，稍后重试！", "服务器正忙，稍后重试！", outrsp, c)
		return
	}
	json.Unmarshal(out.Data, &outrsp)

	repkg.GinResponse(out.HttpCode, out.Code, out.Msg, out.Detail, outrsp, c)
}

func RoleList(c *gin.Context) {
	ctx, ok := gin2micro.ContextWithSpan(c)
	if ok == false {
		//log.Log("get context err")
	}
	type src struct {
		RoleId   string `json:"role_id"`
		RoleName string `json:"role_name"`
	}
	outrsp := make([]src, 0)
	out, _ := hdler.UserSer.GetRoleList(ctx, &userproto.GetRoleListInput{})
	if out == nil {
		repkg.GinResponse(http.StatusInternalServerError, errcode.ERROR_USERSRV_BASE, "服务器正忙，稍后重试！", "服务器正忙，稍后重试！", outrsp, c)
		return
	}
	json.Unmarshal(out.Data, &outrsp)

	repkg.GinResponse(out.HttpCode, out.Code, out.Msg, out.Detail, outrsp, c)
}

func AddRole(c *gin.Context) {
	ctx, ok := gin2micro.ContextWithSpan(c)
	if ok == false {
		//log.Log("get context err")
	}
	type input struct {
		RoleName string `json:"role_name" binding:"required"`
	}
	var in input
	if err := c.ShouldBindJSON(&in); err != nil {
		repkg.GinResponse(http.StatusBadRequest, errcode.ERROR_COMMON_PARAM, errcode.ERROR_COMMON_PARAM_MSG+err.Error(), errcode.ERROR_COMMON_PARAM_MSG+err.Error(), []byte{}, c)
		return
	}
	out, _ := hdler.UserSer.AddRole(ctx, &userproto.AddRoleInput{RoleName: in.RoleName})
	if out == nil {
		repkg.GinResponse(http.StatusInternalServerError, errcode.ERROR_USERSRV_BASE, "服务器正忙，稍后重试！", "服务器正忙，稍后重试！", []byte{}, c)
		return
	}
	repkg.GinResponse(out.HttpCode, out.Code, out.Msg, out.Detail, []byte{}, c)
}

func EditRole(c *gin.Context) {
	ctx, ok := gin2micro.ContextWithSpan(c)
	if ok == false {
		//log.Log("get context err")
	}
	type input struct {
		RoleNameOld string `json:"rolename_old" binding:"required"`
		RoleNameNew string `json:"rolename_new" binding:"required"`
	}
	var in input
	if err := c.ShouldBindJSON(&in); err != nil {
		repkg.GinResponse(http.StatusBadRequest, errcode.ERROR_COMMON_PARAM, errcode.ERROR_COMMON_PARAM_MSG+err.Error(), errcode.ERROR_COMMON_PARAM_MSG+err.Error(), []byte{}, c)
		return
	}
	out, _ := hdler.UserSer.EditRole(ctx, &userproto.EditRoleInput{RolenameOld: in.RoleNameOld, RolenameNew: in.RoleNameNew})
	if out == nil {
		repkg.GinResponse(http.StatusInternalServerError, errcode.ERROR_USERSRV_BASE, "服务器正忙，稍后重试！", "服务器正忙，稍后重试！", []byte{}, c)
		return
	}
	repkg.GinResponse(out.HttpCode, out.Code, out.Msg, out.Detail, []byte{}, c)
}

func DeleteRole(c *gin.Context) {
	ctx, ok := gin2micro.ContextWithSpan(c)
	if ok == false {
		//log.Log("get context err")
	}
	type input struct {
		RoleName string `json:"role_name" binding:"required"`
	}
	var in input
	if err := c.ShouldBindJSON(&in); err != nil {
		repkg.GinResponse(http.StatusBadRequest, errcode.ERROR_COMMON_PARAM, errcode.ERROR_COMMON_PARAM_MSG+err.Error(), errcode.ERROR_COMMON_PARAM_MSG+err.Error(), []byte{}, c)
		return
	}
	out, _ := hdler.UserSer.DeleteRole(ctx, &userproto.DeleteRoleInput{Rolename: in.RoleName})
	if out == nil {
		repkg.GinResponse(http.StatusInternalServerError, errcode.ERROR_USERSRV_BASE, "服务器正忙，稍后重试！", "服务器正忙，稍后重试！", []byte{}, c)
		return
	}
	repkg.GinResponse(out.HttpCode, out.Code, out.Msg, out.Detail, []byte{}, c)
}
