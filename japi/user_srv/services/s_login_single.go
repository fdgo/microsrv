package services

import (
	"context"
	"fmt"

	//"encoding/json"
	//"microservice/jzapi/user_srv_ex/dao"

	//resp "microservice/jzapi/basic/model_response"
	//"microservice/jzapi/basic/randstring"
	//"microservice/jzapi/lib/db"
	//baseproto "microservice/jzapi/proto/base"
	userproto "microservice/jzapi/proto/user"
	//"microservice/jzapi/user_srv/dao"
	//mdl "microservice/jzapi/user_srv/model"
	//"net/http"
)
func (s *Service) SingleRegistQuick(ctx context.Context, in *userproto.SingleRegistQuickInput, out *userproto.CommonOutput) error {
	fmt.Println("hello, single-regist-quick...")
	//out = new(userproto.CommonOutput)
	//tx := db.GetDB().Begin()
	//var outrsp resp.User_RegistQuick
	//_, err := dao.NewDaoGuest(tx).IsDeviceExist(in.DeviceId, in.Channel)
	//if err != nil {
	//	tmp_avatar := "http://101.132.65.222/upload/avatar/" + randstring.Rand1NumString() + ".jpg"
	//	tmp_nickname := nickName()
	//	tmp_uuid := randstring.GetUuidStr()
	//	tmp_realname := randstring.GetRandAccntPwd()
	//	tmp_password := randstring.GetRandAccntPwd()
	//	tmp_salt := randstring.GetRandAccntPwd()
	//	tmp_basic := &mdl.User_Basic{
	//		UuId:     tmp_uuid,
	//		RealName: tmp_realname,
	//		Gender:   1,
	//		Age:      25,
	//		IdCard:   randstring.GetUuidStr(),
	//		Mobile:   randstring.GetUuidStr(),
	//		Address:  "beijing",
	//	}
	//	tmp_guest := &mdl.User_Guest{
	//		UuId:       tmp_uuid,
	//		Status:     1,
	//		DeviceId:   in.DeviceId,
	//		Password:   tmp_password,
	//		Salt:       tmp_salt,
	//		NickName:   tmp_nickname,
	//		Avatar:     tmp_avatar,
	//		Channel_Id: uint(in.Channel),
	//		IsDelete:   100,
	//	}
	//	if err, _ := dao.NewDaoBasic(tx).CreateUser(tmp_basic); err != nil {
	//		out.HttpCode = http.StatusNotAcceptable
	//		out.Code = http.StatusNotAcceptable
	//		out.Msg = "快速注册失败!"
	//		out.Detail = "快速注册失败!"
	//		tx.Rollback()
	//		return nil
	//	}
	//	if err, _ := dao.NewDaoGuest(tx).CreateUser(tmp_guest); err != nil {
	//		out.HttpCode = http.StatusNotAcceptable
	//		out.Code = http.StatusNotAcceptable
	//		out.Msg = "快速注册失败!"
	//		out.Detail = "快速注册失败!"
	//		tx.Rollback()
	//		return nil
	//	}
	//	outrsp.Uuid = tmp_uuid
	//	outrsp.BindAccount.Uuid = tmp_uuid
	//	outrsp.BindAccount.DeviceId = in.DeviceId
	//	outrsp.BindAccount.Level = "0"
	//	outrsp.BindAccount.Exp = "0"
	//	outrsp.BindAccount.NickName = tmp_nickname
	//	outrsp.BindAccount.Avatar = tmp_avatar
	//	outrsp.BindAccount.Vip = "0"
	//	outrsp.BindAccount.VipValidityTime = "0"
	//	outrsp.BindAccount.MyCoin = "0"
	//	outrsp.BindAccount.BindPhone = randstring.GetUuidStr()
	//	outrsp.BindAccount.BindEmail = randstring.GetUuidStr()
	//	outrsp.BindAccount.Stage = 0
	//	outrsp.BindAccount.Series = 0
	//	outrsp.Password = tmp_password
	//	//tmpoutrsp, _ := json.Marshal(outrsp)
	//	//output.Data = string(tmpoutrsp)
	//	out.HttpCode = http.StatusCreated
	//	out.Code = http.StatusCreated
	//	out.Msg = "设备注册成功!"
	//	out.Detail = "设备注册成功!"
	//	tx.Commit()
		//return nil
	//} else {
		//json.Marshal(outrsp)
		//out.Data = []byte{}//string(tmpoutrsp)
		//out.HttpCode = http.StatusConflict
		//out.Code = http.StatusConflict
		//out.Msg = "该设备已经注册!"
		//out.Detail = "该设备已经注册!"
		//tx.Commit()
		return nil
	//}
}

func (s *Service) SingleRegistMobile(ctx context.Context, in *userproto.SingleRegistMobileInput, out *userproto.CommonOutput) error {
	//out = new(userproto.CommonOutput)
	//var outrsp mdl.User_Member
	//tx := db.GetDB().Begin()
	//if in.Code != "111111" {
	//	retcode, _ := baseClient.IsMobileCodeOk(ctx, &baseproto.IsMobileCodeOkIn{
	//		Mobile:   in.Mobile,
	//		Code:     in.Code,
	//		Timediff: in.Timediff,
	//	})
	//	if string(retcode.Data) == "false" {
	//		out.HttpCode = http.StatusBadRequest
	//		out.Code = http.StatusBadRequest
	//		out.Msg = retcode.Msg
	//		out.Detail = retcode.Msg
	//		buf, _ := json.Marshal(outrsp)
	//		out.Data = buf
	//		return nil
	//	}
	//}
	//_, err := dao.NewDaoBasic(tx).IsMobileExist(in.Mobile)
	//if err != nil { //不存在该手机号对应的账号，随机生成
	//	tmp_avatar := "http://101.132.65.222/upload/avatar/" + randstring.Rand1NumString() + ".jpg"
	//	tmp_nickname := nickName()
	//	tmp_uuid := randstring.GetUuidStr()
	//	tmp_realname := randstring.GetRandAccntPwd()
	//	tmp_password := randstring.GetRandAccntPwd()
	//	tmp_salt := randstring.GetRandAccntPwd()
	//	tmp_basic := &mdl.User_Basic{
	//		UuId:     tmp_uuid,
	//		RealName: tmp_realname,
	//		Gender:   1,
	//		Age:      25,
	//		IdCard:   tmp_uuid,
	//		Mobile:   in.Mobile,
	//		Address:  "beijing",
	//	}
	//	tmp_member := &mdl.User_Member{
	//		UuId:     tmp_uuid,
	//		Status:   1,
	//		Password: tmp_password,
	//		Salt:     tmp_salt,
	//		NickName: tmp_nickname,
	//		Avatar:   tmp_avatar,
	//		Channel:  uint(in.Channel),
	//		IsDelete: 100,
	//	}
	//	if err, _ := dao.NewDaoBasic(tx).CreateUser(tmp_basic); err != nil {
	//		out.HttpCode = http.StatusNotAcceptable
	//		out.Code = http.StatusNotAcceptable
	//		out.Msg = "手机注册失败!"
	//		out.Detail = "手机注册失败!"
	//		tx.Rollback()
	//		return nil
	//	}
	//	if err, _ := dao.NewDaoMember(tx).CreateUser(tmp_member); err != nil {
	//		out.HttpCode = http.StatusNotAcceptable
	//		out.Code = http.StatusNotAcceptable
	//		out.Msg = "手机注册失败!"
	//		out.Detail = "手机注册失败!"
	//		tx.Rollback()
	//		return nil
	//	}
	//	outrsp.Status = 1
	//	outrsp.UuId = tmp_uuid
	//	outrsp.Password = tmp_password
	//	outrsp.Avatar = tmp_avatar
	//	outrsp.Channel = uint(in.Channel)
	//	outrsp.IsDelete = 100
	//	outrsp.NickName = tmp_nickname
	//	outrsp.Salt = tmp_salt
	//	outrsp.UserName = tmp_realname
	//	tmpoutrsp, _ := json.Marshal(outrsp)
	//	out.Data = tmpoutrsp
	//	out.HttpCode = http.StatusCreated
	//	out.Code = http.StatusCreated
	//	out.Msg = "手机注册成功!"
	//	out.Detail = "手机注册成功!"
	//	tx.Commit()
	//	return nil
	//} //存在该手机号对应的账号，登陆即可
	//tmpoutrsp, _ := json.Marshal(outrsp)
	//out.Data = tmpoutrsp
	//out.HttpCode = http.StatusConflict
	//out.Code = http.StatusConflict
	//out.Msg = "该手机号号已经注册!"
	//out.Detail = "该手机号号已经注册!"
	//tx.Commit()
	return nil
}


func (s *Service) SingleRegistAccount(ctx context.Context, in *userproto.SingleRegistAccountInput, out *userproto.CommonOutput) error {
	//out = new(userproto.CommonOutput)
	//var outrsp mdl.User_Member
	//tx := db.GetDB().Begin()
	//_, err := dao.NewDaoMember(tx).IsAccountExist(in.Account, in.Password)
	//if err != nil {
	//	tmp_avatar := "http://101.132.65.222/upload/avatar/" + randstring.Rand1NumString() + ".jpg"
	//	tmp_nickname := nickName()
	//	tmp_realname := randstring.GetRandAccntPwd()
	//	tmp_salt := randstring.GetRandAccntPwd()
	//	tmp_basic := &mdl.User_Basic{
	//		UuId:     in.Account,
	//		RealName: tmp_realname,
	//		Gender:   1,
	//		Age:      25,
	//		IdCard:   in.Account,
	//		Mobile:   randstring.GetUuidStr(),
	//		Address:  "beijing",
	//	}
	//	tmp_member := &mdl.User_Member{
	//		UuId:     in.Account,
	//		Status:   1,
	//		Password: in.Password,
	//		Salt:     tmp_salt,
	//		NickName: tmp_nickname,
	//		Avatar:   tmp_avatar,
	//		Channel:  uint(in.Channel),
	//		IsDelete: 100,
	//	}
	//	if err, _ := dao.NewDaoBasic(tx).CreateUser(tmp_basic); err != nil {
	//		out.HttpCode = http.StatusNotAcceptable
	//		out.Code = http.StatusNotAcceptable
	//		out.Msg = "账号注册失败!" + err.Error()
	//		out.Detail = "账号注册失败!"
	//		tx.Rollback()
	//		return nil
	//	}
	//	if err, _ := dao.NewDaoMember(tx).CreateUser(tmp_member); err != nil {
	//		out.HttpCode = http.StatusNotAcceptable
	//		out.Code = http.StatusNotAcceptable
	//		out.Msg = "账号注册失败!" + err.Error()
	//		out.Detail = "账号注册失败!"
	//		tx.Rollback()
	//		return nil
	//	}
	//	outrsp.Status = 1
	//	outrsp.UuId = in.Account
	//	outrsp.Password = in.Password
	//	outrsp.Avatar = tmp_avatar
	//	outrsp.Channel = uint(in.Channel)
	//	outrsp.IsDelete = 100
	//	outrsp.NickName = tmp_nickname
	//	outrsp.Salt = tmp_salt
	//	outrsp.UserName = tmp_realname
	//	tmpoutrsp, _ := json.Marshal(outrsp)
	//	out.Data = tmpoutrsp
	//	out.HttpCode = http.StatusCreated
	//	out.Code = http.StatusCreated
	//	out.Msg = "账号注册成功!"
	//	out.Detail = "账号注册成功!"
	//	tx.Commit()
	//	return nil
	//}
	//tmpoutrsp, _ := json.Marshal(outrsp)
	//out.Data = tmpoutrsp
	//out.HttpCode = http.StatusConflict
	//out.Code = http.StatusConflict
	//out.Msg = "该账号已经注册!"
	//out.Detail = "该账号已经注册!"
	//tx.Commit()
	return nil
}
func (s *Service) SingleLoginGuest(ctx context.Context, in *userproto.SingleLoginGuestInput, out *userproto.CommonOutput) error {
	//out = new(userproto.CommonOutput)
	//tx := db.GetDB().Begin()
	//meminfo, err := dao.NewDaoGuest(tx).IsDeviceExist(in.DeviceId, in.Channel)
	//if err != nil {
	//	out.HttpCode = http.StatusBadRequest
	//	out.Code = http.StatusBadRequest
	//	out.Msg = "不存在与该设备对应的账号或者渠道！"
	//	out.Detail = "不存在与该设备对应的账号或者渠道！"
	//	out.Data = []byte("false")
	//	tx.Rollback()
	//	return nil
	//}
	//loginrcd := &mdl.User_Login_Record{
	//	UUid:      meminfo.UuId,
	//	ChannelId: meminfo.Channel_Id,
	//	DeviceId:  meminfo.DeviceId,
	//}
	//err = dao.NewDaoLoginRecord(tx).AddLoginRecord(loginrcd)
	//if err != nil {
	//	out.HttpCode = http.StatusNotAcceptable
	//	out.Code = http.StatusNotAcceptable
	//	out.Msg = "创建登陆记录失败!"
	//	out.Detail = "创建登陆记录失败!"
	//	out.Data = []byte("false")
	//	tx.Rollback()
	//	return nil
	//}
	//out.HttpCode = http.StatusOK
	//out.Code = http.StatusOK
	//out.Msg = "登陆成功！"
	//out.Detail = "登陆成功！"
	//out.Data =[]byte( CreateToken("", "", meminfo.DeviceId, "singleloginmobile") )
	//tx.Commit()
	return nil
}
func (s *Service) SingleLoginAccount(ctx context.Context, in *userproto.SingleLoginAccountInput, out *userproto.CommonOutput) error {
	//out = new(userproto.CommonOutput)
	//tx := db.GetDB().Begin()
	//meminfo, err := dao.NewDaoMember(tx).IsAccountExist(in.AccountId, in.Pwd)
	//if err != nil { //不存该账号
	//	out.HttpCode = http.StatusBadRequest
	//	out.Code = http.StatusBadRequest
	//	out.Msg = "不存在该账号或者该账号密码错误！"
	//	out.Detail = "不存在该账号或者该账号密码错误！"
	//	out.Data = []byte("false")
	//	tx.Rollback()
	//	return nil
	//}
	//mem, err := dao.NewDaoMember(tx).GetMemberUerInfo(meminfo.UuId)
	//if err != nil {
	//	out.HttpCode = http.StatusBadRequest
	//	out.Code = http.StatusBadRequest
	//	out.Msg = "该账号不存在！"
	//	out.Detail = "该账号不存在！"
	//	out.Data = []byte("false")
	//	tx.Rollback()
	//	return nil
	//}
	//loginrcd := &mdl.User_Login_Record{
	//	UUid:      mem.UuId,
	//	ChannelId: meminfo.Channel,
	//}
	//err = dao.NewDaoLoginRecord(db.GetDB()).AddLoginRecord(loginrcd)
	//if err != nil {
	//	out.HttpCode = http.StatusNotAcceptable
	//	out.Code = http.StatusNotAcceptable
	//	out.Msg = "创建登陆记录失败!"
	//	out.Detail = "创建登陆记录失败!"
	//	out.Data = []byte("false")
	//	tx.Rollback()
	//	return nil
	//}
	//out.HttpCode = http.StatusOK
	//out.Code = http.StatusOK
	//out.Msg = "登陆成功！"
	//out.Detail = "登陆成功！"
	//out.Data = []byte(CreateToken("", mem.UuId, "NULL", "singleloginmobile"))
	//tx.Commit()
	return nil
}
func (s *Service) SingleLoginMobile(ctx context.Context, in *userproto.SingleLoginMobileInput, out *userproto.CommonOutput) error {
	//out = new(userproto.CommonOutput)
	//tx := db.GetDB().Begin()
	//meminfo, err := dao.NewDaoBasic(tx).IsMobileExist(in.Mobile)
	//if err != nil { //不存在该手机号对应的账号，随机生成
	//	out.HttpCode = http.StatusBadRequest
	//	out.Code = http.StatusBadRequest
	//	out.Msg = "该手机号不存在！"
	//	out.Detail = "该手机号不存在！"
	//	out.Data = []byte("false")
	//	tx.Rollback()
	//	return nil
	//}
	//mem, err := dao.NewDaoMember(tx).GetMemberUerInfo(meminfo.UuId)
	//if err != nil {
	//	out.HttpCode = http.StatusBadRequest
	//	out.Code = http.StatusBadRequest
	//	out.Msg = "该手机号不存在！"
	//	out.Detail = "该手机号不存在！"
	//	out.Data = []byte("false")
	//	tx.Rollback()
	//	return nil
	//}
	//loginrcd := &mdl.User_Login_Record{
	//	UUid:      mem.UuId,
	//	Mobile:    in.Mobile,
	//	ChannelId: mem.Channel,
	//}
	//err = dao.NewDaoLoginRecord(tx).AddLoginRecord(loginrcd)
	//if err != nil {
	//	out.HttpCode = http.StatusNotAcceptable
	//	out.Code = http.StatusNotAcceptable
	//	out.Msg = "创建登陆记录失败!"
	//	out.Detail = "创建登陆记录失败!"
	//	out.Data = []byte("false")
	//	tx.Rollback()
	//	return nil
	//}
	//out.HttpCode = http.StatusOK
	//out.Code = http.StatusOK
	//out.Msg = "登陆成功！"
	//out.Detail = "登陆成功！"
	//out.Data = []byte(CreateToken(meminfo.Mobile, meminfo.UuId, "NULL", "singleloginmobile"))
	//tx.Commit()
	return nil
}


//func (s *Service) VistorLogin(ctx context.Context, in *userproto.LoginInput, out *userproto.CommonOutput) error {
//	//msg,isok := userService.VistorLogin(c,in)
//	//if isok {
//	//	out.Code = errorcode.SUCCESS
//	//}else {
//	//	out.Code = errorcode.ERROR
//	//}
//	//out.Msg = msg
//	//out.Date = isok
//
//	return nil
//}
//func (s *Service) MobileLogin(ctx context.Context, in *userproto.LoginInput, out *userproto.CommonOutput) error {
//	//msg,isok :=  userService.MobileLogin(c,in)
//	//if isok {
//	//	out.Code = errorcode.SUCCESS
//	//}else {
//	//	out.Code = errorcode.ERROR
//	//}
//	//out.Msg = msg
//	//out.Date = isok
//	return nil
//}
//func (s *Service) AccountLogin(ctx context.Context, in *userproto.LoginInput, out *userproto.CommonOutput) error {
//	//msg,isok,acct := userService.AccountLogin(c,in)
//	//if isok {
//	//	out.Code = errorcode.SUCCESS
//	//	out.Msg = acct
//	//}else {
//	//	out.Code = errorcode.ERROR
//	//	out.Msg = msg
//	//}
//	//out.Date = isok
//	return nil
//}
//func (s *Service) ActiveAccount(ctx context.Context, in *userproto.ActiveAccountInput, out *userproto.CommonOutput) error {
//	userService.ActiveAccount(ctx,in)
//	return nil
//}
//
//func (s *Service) DeviceId(ctx context.Context, in *userproto.DeviceIdInput, out *userproto.CommonOutput) error {
//	//deviceid,isok := userService.DeviceId(c,in)
//	//if isok{
//	//	out.Code = errorcode.SUCCESS
//	//	out.Msg = deviceid
//	//}else {
//	//	out.Code = errorcode.ERROR
//	//	out.Msg = deviceid
//	//}
//	//out.Date = isok
//	return nil
//}
//func (s *Service)DeviceLogin(ctx context.Context, in *userproto.DeviceLoginInput,  out *userproto.CommonOutput) error{
//	//err := userService.DeviceLogin(c, in,out)
//	return nil
//}
//func (s *Service) UpdatePwdByMobile(ctx context.Context, in *userproto.UpdatePwdByMobileInput, out *userproto.CommonOutput) error{
//	//err := userService.UpdatePwdByMobile(c, in,out)
//	return nil
//}
//func (s *Service) BindIDcard(ctx context.Context, in *userproto.BindIDcardInput, out *userproto.CommonOutput) ( error) {
//	//err := userService.BindIDcard(c,in,out)
//	return nil
//}
//func (s *Service) UpdatePwdByIDcard(ctx context.Context, in *userproto.UpdatePwdByIDcardInput, out *userproto.CommonOutput) error {
//	//err := userService.UpdatePwdByIDcard(c, in,out)
//	return nil
//}
//func (s *Service) VerifyIdcard(ctx context.Context, in *userproto.VerifyIdcardInput, out *userproto.CommonOutput) error {
//	//err := userService.VerifyIdcard(c, in,out)
//	return nil
//}
//func (s *Service) BindMobile(ctx context.Context, in *userproto.BindMobileInput,out * userproto.CommonOutput) error{
//	//err := userService.BindMobile(c, in,out)
//	return nil
//}
//func (s *Service)  GetCustomer(ctx context.Context, in *userproto.GetCustomerInput, out *userproto.CommonOutput )error {
//	//err := userService.GetCustomer(c, in,out)
//	return nil
//}
//func (s *Service)  PayLog(ctx context.Context, in *userproto.PayLogInput, out *userproto.CommonOutput) error{
//	//err := userService.PayLog(c, in,out)
//	return nil
//}
//func (s *Service)  CdKeyList(ctx context.Context, in *userproto.CdKeyListInput, out *userproto.CommonOutput) error{
//	//err := userService.CdKeyList(c, in,out)
//	return nil
//}
//func (s *Service)  CdKeyInfo(ctx context.Context, in *userproto.CdKeyInfoInput, out *userproto.CommonOutput) error {
//	//err := userService.CdKeyInfo(c,in,out)
//	return nil
//}
//func (s *Service)  CdKeyMy(ctx context.Context, in *userproto.CdKeyMyInput, out *userproto.CommonOutput) error {
//	//err := userService.CdKeyMy(c, in,out)
//	return nil
//}
