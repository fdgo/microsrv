package manage

import (
	"encoding/json"
	//"encoding/json"
	"github.com/gin-gonic/gin"
	hdler "microservice/jzapi/admin_web/handler"
	errcode "microservice/jzapi/basic/error_code"
	repkg "microservice/jzapi/basic/return_pkg"
	//"microservice/jzapi/basic/time_ex"
	"microservice/jzapi/lib/wrapper/tracer/opentracing/gin2micro"
	userproto "microservice/jzapi/proto/user"
	"net/http"
)

func ChannelList(c *gin.Context) {
	ctx, ok := gin2micro.ContextWithSpan(c)
	if ok == false {
		//log.Log("get context err")
	}
	type outresp struct {
		ChannelId   uint    `json:"channel_id"`
		ChannelName string  `json:"channel_name"`
		CreatedAt   string  `json:"created_at"`
	}
	sliresp := make([]outresp, 0)
	out, _ := hdler.UserSer.GetChannelInfo(ctx, &userproto.GetChannelInfoInput{})
	if out == nil {
		repkg.GinResponse(http.StatusInternalServerError, errcode.ERROR_ADMINSRV_BASE, "服务器正忙，稍后重试！", "服务器正忙，稍后重试！", sliresp, c)
		return
	}
	json.Unmarshal(out.Data, &sliresp)
	repkg.GinResponse(out.HttpCode, out.Code, out.Msg, out.Detail, sliresp, c)
}