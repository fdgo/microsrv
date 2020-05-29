package login

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/util/log"
	errcode "microservice/jzapi/basic/error_code"
	resp "microservice/jzapi/basic/model_response"
	"microservice/jzapi/basic/regexp"
	repkg "microservice/jzapi/basic/return_pkg"
	"microservice/jzapi/lib/wrapper/tracer/opentracing/gin2micro"
	baseproto "microservice/jzapi/proto/base"
	userproto "microservice/jzapi/proto/user"
	hdler "microservice/jzapi/user_web/srvclient"
	"net/http"
)

func SingleLoginGuest(c *gin.Context) {
	ctx, ok := gin2micro.ContextWithSpan(c)
	if ok == false {
		log.Log("get context err")
	}
	type Input struct {
		Deviceid string `json:"device_id" binding:"required"`
		Channel  uint   `json:"channel"`
	}
	type Token struct {
		Token string `json:"token"`
	}
	var outrsp Token
	var in Input
	if err := c.ShouldBindJSON(&in); err != nil {
		repkg.GinResponse(http.StatusBadRequest, errcode.ERROR_COMMON_PARAM, errcode.ERROR_COMMON_PARAM_MSG+err.Error(), errcode.ERROR_COMMON_PARAM_MSG+err.Error(), outrsp, c)
		return
	}
	out, _ := hdler.UserSer.SingleLoginGuest(ctx, &userproto.SingleLoginGuestInput{
		DeviceId: in.Deviceid,
		Channel:  int32(in.Channel),
	})
	if out == nil {
		repkg.GinResponse(http.StatusInternalServerError, errcode.ERROR_USERSRV_BASE, "服务器正忙，稍后重试！", "服务器正忙，稍后重试！", outrsp, c)
		return
	}
	outrsp.Token = string(out.Data)
	repkg.GinResponse(out.HttpCode, out.Code, out.Msg, out.Detail, outrsp, c)
}
func SingleLoginMobile(c *gin.Context) {
	ctx, ok := gin2micro.ContextWithSpan(c)
	if ok == false {
		//log.Log("get context err")
	}
	type input struct {
		Mobile   string `json:"mobile" binding:"required"`
		Channel  int32  `json:"channel" binding:"required"`
		Code     string `json:"code" binding:"required"`
		TimeDiff int32  `json:"timediff" binding:"required"`
	}
	type Token struct {
		Token string `json:"token"`
	}
	var outrsp Token
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
	out, _ := hdler.UserSer.SingleLoginMobile(ctx, &userproto.SingleLoginMobileInput{
		Channel: in.Channel,
		Mobile:  in.Mobile,
		Code:    in.Code,
	})
	if out == nil {
		repkg.GinResponse(http.StatusInternalServerError, errcode.ERROR_USERSRV_BASE, "服务器正忙，稍后重试！", "服务器正忙，稍后重试！", outrsp, c)
		return
	}
	outrsp.Token = string(out.Data)
	repkg.GinResponse(out.HttpCode, out.Code, out.Msg, out.Detail, outrsp, c)
}
func SingleLoginAccount(c *gin.Context) {
	ctx, ok := gin2micro.ContextWithSpan(c)
	if ok == false {
		log.Log("get context err")
	}
	type Input struct {
		Uuid     string `json:"uuid"`
		Password string `json:"password"`
		Channel  uint   `json:"channel"`
	}
	type Token struct {
		Token string `json:"token"`
	}
	var outrsp Token
	var in Input
	if err := c.ShouldBindJSON(&in); err != nil {
		repkg.GinResponse(http.StatusBadRequest, errcode.ERROR_COMMON_PARAM, errcode.ERROR_COMMON_PARAM_MSG+err.Error(), errcode.ERROR_COMMON_PARAM_MSG+err.Error(), outrsp, c)
		return
	}
	out, _ := hdler.UserSer.SingleLoginAccount(ctx, &userproto.SingleLoginAccountInput{
		Channel:   int32(in.Channel),
		AccountId: in.Uuid,
		Pwd:       in.Password,
	})
	if out == nil {
		repkg.GinResponse(http.StatusInternalServerError, errcode.ERROR_USERSRV_BASE, "服务器正忙，稍后重试！", "服务器正忙，稍后重试！", outrsp, c)
		return
	}
	outrsp.Token = string(out.Data)
	repkg.GinResponse(out.HttpCode, out.Code, out.Msg, out.Detail, outrsp, c)
}

//4.快速注册（新接口）
func SingleRegistQuick(c *gin.Context) {
	ctx, ok := gin2micro.ContextWithSpan(c)
	if ok == false {
		log.Log("get context err")
	}
	type Input struct {
		Deviceid string `json:"device_id" binding:"required"`
		Ver      int32  `json:"ver" binding:"required"`
		Channel  uint   `json:"channel"`
		Type     string `json:"type"`
		Pid      string `json:"pid" binding:"required"`
	}
	var outrsp resp.User_RegistQuick
	var in Input
	if err := c.ShouldBindJSON(&in); err != nil {
		repkg.GinResponse(http.StatusBadRequest, errcode.ERROR_COMMON_PARAM, errcode.ERROR_COMMON_PARAM_MSG+err.Error(), errcode.ERROR_COMMON_PARAM_MSG+err.Error(), outrsp, c)
		return
	}
	out, _ := hdler.UserSer.SingleRegistQuick(ctx, &userproto.SingleRegistQuickInput{
		DeviceId: in.Deviceid,
		Ver:      in.Ver,
		Type:     in.Type,
		Channel:  int32(in.Channel),
		Pid:      in.Pid,
	})
	if out == nil {
		repkg.GinResponse(http.StatusInternalServerError, errcode.ERROR_USERSRV_BASE, "服务器正忙，稍后重试！", "服务器正忙，稍后重试！", outrsp, c)
		return
	}
	json.Unmarshal(out.Data, &outrsp)
	repkg.GinResponse(out.HttpCode, out.Code, out.Msg, out.Detail, outrsp, c)
}

//2.手机号注册
func SingleRegistMobile(c *gin.Context) {
	ctx, ok := gin2micro.ContextWithSpan(c)
	if ok == false {
		log.Log("get context err")
	}
	type Input struct {
		Deviceid string `json:"device_id" binding:"required"`
		Mobile   string `json:"mobile" binding:"required"`
		Code     string `json:"code" binding:"required"`
		TimeDiff int32  `json:"timediff" binding:"required"`
		Ver      int32  `json:"ver" binding:"required"`
		Channel  int32  `json:"channel"`
		Type     string `json:"type"`
		Pid      string `json:"pid" binding:"required"`
		Password string `json:"password" binding:"required"`
	}
	//var outrsp mdl.User_Member
	var in Input
	if err := c.ShouldBindJSON(&in); err != nil {
		repkg.GinResponse(http.StatusBadRequest, errcode.ERROR_COMMON_PARAM, errcode.ERROR_COMMON_PARAM_MSG+err.Error(), errcode.ERROR_COMMON_PARAM_MSG+err.Error(), nil, c)
		return
	}
	rep := regexp.Veryfy{}
	if !rep.VerifyMobile(in.Mobile) || !rep.VerifyPasswd(in.Password) {
		//repkg.RespErrParams("请检查手机号或者密码[格式]是否正确!", c)
		return
	}
	out, _ := hdler.UserSer.SingleRegistMobile(ctx, &userproto.SingleRegistMobileInput{
		DeviceId: in.Deviceid,
		Mobile:   in.Mobile,
		Code:     in.Code,
		Timediff: in.TimeDiff,
		Ver:      in.Ver,
		Channel:  in.Channel,
		Type:     in.Type,
		Pid:      in.Pid,
		Password: in.Password,
	})
	if out == nil {
		repkg.GinResponse(http.StatusInternalServerError, errcode.ERROR_USERSRV_BASE, "服务器正忙，稍后重试！", "服务器正忙，稍后重试！", nil, c)
		return
	}
	//json.Unmarshal([]byte(out.Data), &outrsp)
	repkg.GinResponse(out.HttpCode, out.Code, out.Msg, out.Detail, nil, c)
}

//3.账号密码注册（新接口）
func SingleRegistAccount(c *gin.Context) {
	ctx, ok := gin2micro.ContextWithSpan(c)
	if ok == false {
		log.Log("get context err")
	}
	type Input struct {
		Deviceid string `json:"device_id" binding:"required"`
		Account  string `json:"account" binding:"required"`
		Password string `json:"password" binding:"required"`
		Ver      int32  `json:"ver" binding:"required"`
		Channel  uint   `json:"channel"`
		Type     string `json:"type"`
		Pid      string `json:"pid" binding:"required"`
	}
	//var outrsp mdl.User_Member
	var in Input
	if err := c.ShouldBindJSON(&in); err != nil {
		repkg.GinResponse(http.StatusBadRequest, errcode.ERROR_COMMON_PARAM, errcode.ERROR_COMMON_PARAM_MSG+err.Error(), errcode.ERROR_COMMON_PARAM_MSG+err.Error(), nil, c)
		return
	}
	out, _ := hdler.UserSer.SingleRegistAccount(ctx, &userproto.SingleRegistAccountInput{
		DeviceId: in.Deviceid,
		Account:  in.Account,
		Password: in.Password,
		Ver:      in.Ver,
		Type:     in.Type,
		Channel:  int32(in.Channel),
		Pid:      in.Pid,
	})
	if out == nil {
		repkg.GinResponse(http.StatusInternalServerError, errcode.ERROR_USERSRV_BASE, "服务器正忙，稍后重试！", "服务器正忙，稍后重试！", nil, c)
		return
	}
	//var ret mdl.User_Member
	//json.Unmarshal([]byte(out.Data), &ret)
	repkg.GinResponse(out.HttpCode, out.Code, out.Msg, out.Detail, nil, c)
}

//5.账号密码登录
//func DeviceLogin(c *gin.Context) {
//	ctx, ok := gin2micro.ContextWithSpan(c)
//	if ok == false {
//		log.Log("get context err")
//	}
//	type Input struct {
//		Deviceid string `json:"device_id" binding:"required"`
//		Account  string `json:"account" binding:"required"`
//		Password string `json:"password" binding:"required"`
//		Ver      int32  `json:"ver" binding:"required"`
//		Channel  int32  `json:"channel"`
//		Type     string `json:"type"`
//		Pid      string `json:"pid" binding:"required"`
//	}
//	var in Input
//	if err := c.ShouldBindJSON(&in); err != nil {
//		repkg.GinResponse(http.StatusBadRequest, errcode.ERROR_COMMON_PARAM, errcode.ERROR_COMMON_PARAM_MSG+err.Error(), nil, c)
//		return
//	}
//	rep := regexp.Veryfy{}
//	if !rep.VerifyUserAccount(in.Account) || !rep.VerifyPasswd(in.Password) {
//		//repkg.RespErrParams("请检查账号或者密码[格式]是否正确!", c)
//		return
//	}
//	_, err := hdler.UserSer.DeviceLogin(ctx, &userproto.DeviceLoginInput{
//		DeviceId: in.Deviceid,
//		Account:  in.Account,
//		Password: in.Password,
//		Ver:      in.Ver,
//		Type:     in.Type,
//		Channel:  in.Channel,
//		Pid:      in.Pid,
//	})
//	if err != nil {
//		//repkg.RespErrEachServer(errcode.USER_DEVICELOGIN_MSG+err.Error(), errcode.USER_DEVICELOGIN, c)
//		return
//	}
//	//签发token并缓存redis
//	//rsp, err := hdler.BaseSer.MakeAccessToken(ctx, &authproto.AccessTokenInput{
//	//	DeviceId:in.Deviceid,
//	//	Account:in.Account,
//	//})
//	//token := rsp.Token
//	//out.Ret.Token = token
//	//repkg.RespSuccess("账号密码登录成功 ！", out, c)
//}


//1.发送验证码
//func SendYzm(c *gin.Context) {
//	_, ok := gin2micro.ContextWithSpan(c)
//	if ok == false {
//		log.Log("get context err")
//	}
//	type input struct {
//		Mobile   string `json:"mobile" binding:"required"`
//		Times    int32  `json:"times" binding:"required"`
//		Timediff int64  `json:"timediff" binding:"required"`
//	}
//	var in input
//	if err := c.ShouldBindJSON(&in); err != nil {
//		repkg.GinResponse(http.StatusBadRequest, errcode.ERROR_COMMON_PARAM, errcode.ERROR_COMMON_PARAM_MSG+err.Error(), nil, c)
//		return
//	}
//	//out, _ := hdler.BaseSer.SendMobileCode(ctx, &baseproto.SendMobileCodeIn{
//	//	Mobile:   in.Mobile,
//	//	Times:    in.Times,
//	//	Timediff: in.Timediff,
//	//})
//	//repkg.GinResponse(out.StatusCode, out.StatusMsg, out.Data, c)
//}
//
////10.身份证绑定
//func BindIdcard(c *gin.Context) { //10.身份证绑定
//	_, ok := gin2micro.ContextWithSpan(c)
//	if ok == false {
//		log.Log("get context err")
//	}
//	type Input struct {
//		Deviceid  string `json:"deviceid" binding:"required"`
//		Accountid string `json:"account_id" binding:"required"`
//		Name      string `json:"name" binding:"required"`
//		Idcard    string `json:"idcard" binding:"required"`
//	}
//	var in Input
//	if err := c.ShouldBindJSON(&in); err != nil {
//		repkg.GinResponse(http.StatusBadRequest, errcode.ERROR_COMMON_PARAM, errcode.ERROR_COMMON_PARAM_MSG+err.Error(), nil, c)
//		return
//	}
//	rep := regexp.Veryfy{}
//	if !rep.VerifyIDcard(in.Idcard) || !rep.VerifyUserAccount(in.Accountid) {
//		//repkg.RespErrParams("请检查账号或者身份证号或者密码[格式]是否正确!", c)
//		return
//	}
//	//out, _ := hdler.UserSer.BindIDcard(ctx, &userproto.BindIDcardInput{
//	//	Name:      in.Name,
//	//	Idcard:    in.Idcard,
//	//	Deviceid:  in.Deviceid,
//	//	AccountId: in.Accountid,
//	//})
//	//repkg.GinResponse(out.StatusCode, out.StatusMsg, out.Data, c)
//}
//
////12. 绑定新手机号
//func BindMobile(c *gin.Context) {
//	_, ok := gin2micro.ContextWithSpan(c)
//	if ok == false {
//		log.Log("get context err")
//	}
//	type Input struct {
//		Deviceid  string `json:"deviceid" binding:"required"`
//		Accountid string `json:"account_id" binding:"required"`
//		Mobile    string `json:"mobile" binding:"required"`
//		Code      string `json:"code" binding:"required"`
//	}
//	var in Input
//	if err := c.ShouldBindJSON(&in); err != nil {
//		repkg.GinResponse(http.StatusBadRequest, errcode.ERROR_COMMON_PARAM, errcode.ERROR_COMMON_PARAM_MSG+err.Error(), nil, c)
//		return
//	}
//	rep := regexp.Veryfy{}
//	if !rep.VerifyUserAccount(in.Accountid) || !rep.VerifyMobile(in.Mobile) {
//		//repkg.RespErrParams("请检查账号或者手机号[格式]是否正确!", c)
//		return
//	}
//	//out, _ := hdler.UserSer.BindMobile(ctx, &userproto.BindMobileInput{
//	//	Code:      in.Code,
//	//	Mobile:    in.Mobile,
//	//	AccountId: in.Accountid,
//	//	Deviceid:  in.Deviceid,
//	//})
//	//repkg.GinResponse(out.StatusCode, out.StatusMsg, out.Data, c)
//}

//7.手机号找回密码
//func ForgetMobile(c *gin.Context) {
//	_, ok := gin2micro.ContextWithSpan(c)
//	if ok == false {
//		log.Log("get context err")
//	}
//	type Input struct {
//		Deviceid string `json:"deviceid" binding:"required"`
//		Mobile   string `json:"mobile" binding:"required"`
//		Password string `json:"password" binding:"required"`
//		Code     string `json:"code" binding:"required"`
//	}
//	var in Input
//	if err := c.ShouldBindJSON(&in); err != nil {
//		repkg.GinResponse(http.StatusBadRequest, errcode.ERROR_COMMON_PARAM, errcode.ERROR_COMMON_PARAM_MSG+err.Error(), nil, c)
//		return
//	}
//
//	rep := regexp.Veryfy{}
//	if !rep.VerifyMobile(in.Mobile) || !rep.VerifyPasswd(in.Password) {
//		//repkg.RespErrParams("请检查手机号或者密码[格式]是否正确!", c)
//		return
//	}
//	//_, err := hdler.BaseSer.GetMobileCode(ctx, &baseproto.GetMobileCodeIn{
//	//	Mobile: in.Mobile,
//	//})
//	//if err != nil {
//	//	//repkg.RespErrEachServer(errcode.USER_GETMOBILECODE_MSG+"验证码错误 !", errcode.USER_GETMOBILECODE, c)
//	//	return
//	//}
//	//if rsp.Code != in.Code {
//	//	repkg.RespErrEachServer(errcode.USER_GETMOBILECODE_MSG+"验证码错误 !", errcode.USER_GETMOBILECODE, c)
//	//	return
//	//}
//	//out, err := hdler.UserSer.UpdatePwdByMobile(ctx, &userproto.UpdatePwdByMobileInput{
//	//	Password: in.Password,
//	//	Mobile:   in.Mobile,
//	//	Code:     in.Code,
//	//})
//	//repkg.GinResponse(out.StatusCode, out.StatusMsg, out.Data, c)
//}

//8.身份证找回密码
//func ForgetIdcard(c *gin.Context) { //8.身份证找回密码
//
//	_, ok := gin2micro.ContextWithSpan(c)
//	if ok == false {
//		log.Log("get context err")
//	}
//	type Input struct {
//		Deviceid  string `json:"deviceid" binding:"required"`
//		Accountid string `json:"account_id" binding:"required"`
//		Password  string `json:"password" binding:"required"`
//		Name      string `json:"name" binding:"required"`
//		Idcard    string `json:"idcard" binding:"required"`
//	}
//	var in Input
//	if err := c.ShouldBindJSON(&in); err != nil {
//		repkg.GinResponse(http.StatusBadRequest, errcode.ERROR_COMMON_PARAM, errcode.ERROR_COMMON_PARAM_MSG+err.Error(), nil, c)
//		return
//	}
//	rep := regexp.Veryfy{}
//	if !rep.VerifyIDcard(in.Idcard) || !rep.VerifyPasswd(in.Password) || !rep.VerifyUserAccount(in.Accountid) {
//		//repkg.RespErrParams("请检查账号或者身份证号或者密码[格式]是否正确!", c)
//		return
//	}
//	//out, _ := hdler.UserSer.UpdatePwdByIDcard(ctx, &userproto.UpdatePwdByIDcardInput{
//	//	Deviceid:  in.Deviceid,
//	//	Password:  in.Password,
//	//	AccountId: in.Accountid,
//	//	Idcard:    in.Idcard,
//	//	Name:      in.Name,
//	//})
//	//repkg.GinResponse(out.StatusCode, out.StatusMsg, out.Data, c)
//}

////9.身份证验证
//func VerifyIdcard(c *gin.Context) { //9. 身份证验证
//	_, ok := gin2micro.ContextWithSpan(c)
//	if ok == false {
//		log.Log("get context err")
//	}
//	type Input struct {
//		Deviceid  string `json:"deviceid" binding:"required"`
//		Accountid string `json:"account" binding:"required"`
//		Idcard    string `json:"idcard" binding:"required"`
//	}
//	var in Input
//	if err := c.ShouldBindJSON(&in); err != nil {
//		repkg.GinResponse(http.StatusBadRequest, errcode.ERROR_COMMON_PARAM, errcode.ERROR_COMMON_PARAM_MSG+err.Error(), nil, c)
//		return
//	}
//	rep := regexp.Veryfy{}
//	if !rep.VerifyIDcard(in.Idcard) || !rep.VerifyUserAccount(in.Accountid) {
//		//repkg.RespErrParams("请检查账号或者身份证号或者密码[格式]是否正确!", c)
//		return
//	}
//	//out, _ := hdler.UserSer.VerifyIdcard(ctx, &userproto.VerifyIdcardInput{
//	//	Idcard:    in.Idcard,
//	//	AccountId: in.Accountid,
//	//	Deviceid:  in.Deviceid,
//	//})
//	//repkg.GinResponse(out.StatusCode, out.StatusMsg, out.Data, c)
//}

//11.手机号验证
//func VerifyByMobile(c *gin.Context) { //11.手机号验证
//ctx, ok := gin2micro.ContextWithSpan(c)
//if ok == false {
//	log.Log("get context err")
//}
//type Input struct {
//	Deviceid string `json:"deviceid" binding:"required"`
//	Mobile   string `json:"mobile" binding:"required"`
//	Code     string `json:"code" binding:"required"`
//}
//var in Input
//if err := c.ShouldBindJSON(&in); err != nil {
//	repkg.RespErrParams(err.Error(), c)
//	return
//}
//rep := regexp.Veryfy{}
//if !rep.VerifyMobile(in.Mobile) {
//	repkg.RespErrParams("请检查手机号[格式]是否正确!", c)
//	return
//}
//out, err := hdler.UserSer.IsMobileExist(ctx, &userproto.IsMobileExistInput{
//	DeviceId: in.Deviceid,
//	Moblie:   in.Mobile,
//})
//if err != nil {
//	repkg.RespErrEachServer(errcode.USER_ISMOBILEEXIST_MSG+err.Error(), errcode.USER_ISMOBILEEXIST, c)
//	return
//}
//if out.StatusCode == 1 {
//	repkg.RespErrEachServer(errcode.USER_ISMOBILEEXIST_MSG+err.Error(), errcode.USER_ISMOBILEEXIST, c)
//	return
//}
//_, err = hdler.BaseSer.GetMobileCode(ctx, &baseproto.GetMobileCodeIn{
//	Mobile: in.Mobile,
//})
//if err != nil {
//	repkg.RespErrEachServer(errcode.USER_GETMOBILECODE_MSG+"验证码错误 !", errcode.USER_GETMOBILECODE, c)
//	return
//}
//if rsp.Code != in.Code {
//	repkg.RespErrEachServer(errcode.USER_GETMOBILECODE_MSG+"验证码错误 !", errcode.USER_GETMOBILECODE, c)
//	return
//}
//repkg.RespSuccess(errcode.SUCCESS_MSG+"手机号验证成功 !", errcode.SUCCESS_MSG, c)
//}
//func GetCustomer(c *gin.Context) {
//	_, ok := gin2micro.ContextWithSpan(c)
//	if ok == false {
//		log.Log("get context err")
//	}
//	type Input struct {
//		Deviceid string `json:"deviceid" binding:"required"`
//		GameID   string `json:"game_id" binding:"required"`
//	}
//	var in Input
//	if err := c.ShouldBindJSON(&in); err != nil {
//		repkg.GinResponse(http.StatusBadRequest, errcode.ERROR_COMMON_PARAM, errcode.ERROR_COMMON_PARAM_MSG+err.Error(), nil, c)
//		return
//	}
//	//out, _ := hdler.UserSer.GetCustomer(ctx, &userproto.GetCustomerInput{
//	//	Deviceid: in.Deviceid,
//	//	GameId:   in.GameID,
//	//})
//	//repkg.GinResponse(out.StatusCode, out.StatusMsg, out.Data, c)
//}
//func PayLog(c *gin.Context) {
//	_, ok := gin2micro.ContextWithSpan(c)
//	if ok == false {
//		log.Log("get context err")
//	}
//	type Input struct {
//		Deviceid string `json:"deviceid" binding:"required"`
//		Token    string `json:"token" binding:"required"`
//		Devicesn string `json:"devicesn" binding:"required"`
//	}
//	var in Input
//	if err := c.ShouldBindJSON(&in); err != nil {
//		repkg.GinResponse(http.StatusBadRequest, errcode.ERROR_COMMON_PARAM, errcode.ERROR_COMMON_PARAM_MSG+err.Error(), nil, c)
//		return
//	}
//	//out, _ := hdler.UserSer.PayLog(ctx, &userproto.PayLogInput{
//	//	Deviceid: in.Deviceid,
//	//	Token:    in.Token,
//	//	Devicesn: in.Devicesn,
//	//})
//	//repkg.GinResponse(out.StatusCode, out.StatusMsg, out.Data, c)
//}
//
//func CdKey(c *gin.Context) {
//	_, ok := gin2micro.ContextWithSpan(c)
//	if ok == false {
//		log.Log("get context err")
//	}
//	type Input struct {
//		Deviceid string `json:"deviceid" binding:"required"`
//		Token    string `json:"token" binding:"required"`
//		Devicesn string `json:"devicesn" binding:"required"`
//		UserId   string `json:"user_id" binding:"required"`
//		GameId   string `json:"game_id" binding:"required"`
//		Channel  string `json:"channel" binding:"required"`
//	}
//	var in Input
//	if err := c.ShouldBindJSON(&in); err != nil {
//		repkg.GinResponse(http.StatusBadRequest, errcode.ERROR_COMMON_PARAM, errcode.ERROR_COMMON_PARAM_MSG+err.Error(), nil, c)
//		return
//	}
//	//out, _ := hdler.UserSer.CdKeyList(ctx, &userproto.CdKeyListInput{
//	//	Deviceid: in.Deviceid,
//	//	Token:    in.Token,
//	//	Devicesn: in.Devicesn,
//	//	UserId:   in.UserId,
//	//	GameId:   in.GameId,
//	//	Channel:  in.Channel,
//	//})
//	//repkg.GinResponse(out.StatusCode, out.StatusMsg, out.Data, c)
//}
//
//func CdKeyInfo(c *gin.Context) {
//	_, ok := gin2micro.ContextWithSpan(c)
//	if ok == false {
//		log.Log("get context err")
//	}
//	type Input struct {
//		Deviceid string `json:"deviceid" binding:"required"`
//		Token    string `json:"token" binding:"required"`
//		Devicesn string `json:"devicesn" binding:"required"`
//		UserId   string `json:"user_id" binding:"required"`
//		GameId   string `json:"game_id" binding:"required"`
//		CdkeyId  string `json:"cdkey_id" binding:"required"`
//	}
//	var in Input
//	if err := c.ShouldBindJSON(&in); err != nil {
//		repkg.GinResponse(http.StatusBadRequest, errcode.ERROR_COMMON_PARAM, errcode.ERROR_COMMON_PARAM_MSG+err.Error(), nil, c)
//		return
//	}
//	//out, _ := hdler.UserSer.CdKeyInfo(ctx, &userproto.CdKeyInfoInput{
//	//	Deviceid: in.Deviceid,
//	//	Token:    in.Token,
//	//	Devicesn: in.Devicesn,
//	//	UserId:   in.UserId,
//	//	GameId:   in.GameId,
//	//	CdkeyId:  in.CdkeyId,
//	//})
//	//repkg.GinResponse(out.StatusCode, out.StatusMsg, out.Data, c)
//}
//func CdKeyMy(c *gin.Context) {
//	_, ok := gin2micro.ContextWithSpan(c)
//	if ok == false {
//		log.Log("get context err")
//	}
//	type Input struct {
//		Deviceid string `json:"deviceid" binding:"required"`
//		Token    string `json:"token" binding:"required"`
//		Devicesn string `json:"devicesn"`
//		UserId   string `json:"user_id" binding:"required"`
//		GameId   string `json:"game_id" binding:"required"`
//		Channel  string `json:"channel"`
//	}
//	var in Input
//	if err := c.ShouldBindJSON(&in); err != nil {
//		repkg.GinResponse(http.StatusBadRequest, errcode.ERROR_COMMON_PARAM, errcode.ERROR_COMMON_PARAM_MSG+err.Error(), nil, c)
//		return
//	}
//	//out, _ := hdler.UserSer.CdKeyMy(ctx, &userproto.CdKeyMyInput{
//	//	Deviceid: in.Deviceid,
//	//	Token:    in.Token,
//	//	Devicesn: in.Devicesn,
//	//	UserId:   in.UserId,
//	//	GameId:   in.GameId,
//	//	Channel:  in.Channel,
//	//})
//	//repkg.GinResponse(out.StatusCode, out.StatusMsg, out.Data, c)
//}
//func BaseConfig(c *gin.Context) {
//	_, ok := gin2micro.ContextWithSpan(c)
//	if ok == false {
//		log.Log("get context err")
//	}
//	type Input struct {
//		Deviceid string `json:"deviceid" binding:"required"`
//		Forced   int    `json:"forced" binding:"required"`
//		Pid      string `json:"pid" binding:"required"`
//		GameId   string `json:"game_id"`
//	}
//	var in Input
//	if err := c.ShouldBindJSON(&in); err != nil {
//		repkg.GinResponse(http.StatusBadRequest, errcode.ERROR_COMMON_PARAM, errcode.ERROR_COMMON_PARAM_MSG+err.Error(), nil, c)
//		return
//	}
//	//out, err := hdler.UserSer.CdKeyMy(ctx, &userproto.CdKeyMyInput{
//	//	Deviceid: in.Deviceid,
//	//	Token:    in.Token,
//	//	Devicesn: in.Devicesn,
//	//	UserId:   in.UserId,
//	//	GameId:   in.GameId,
//	//	Channel:  in.Channel,
//	//})
//	//if err != nil {
//	//	repkg.RespErrEachServer(errcode.USER_CDKEYMY_MSG+"查看我的礼包错误 ！", errcode.USER_CDKEYMY, c)
//	//	return
//	//}
//	//repkg.RespSuccess(errcode.SUCCESS_MSG+"查看我的礼包成功   !", out, c)
//}
