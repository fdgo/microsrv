package services

import (
	"context"
	//"encoding/json"
	"microservice/jzapi/basic/cfg/common"
	//mresp "microservice/jzapi/basic/model_response"
	//"microservice/jzapi/basic/randstring"
	//"microservice/jzapi/lib/db"
	userproto "microservice/jzapi/proto/user"
)

var (
	jwtkey = "jwt"
	jwtcfg = &keyCfg{}
)

type keyCfg struct {
	common.Jwt
}

//完成
func  (s *Service)MultiLoginMobile(ctx context.Context, in *userproto.MultiLoginMobileInput, out *userproto.CommonOutput) error {
	//output := new(userproto.CommonOutput)
	//var outresp mresp.User_MultiLoginMobile
	//tx := db.GetDB().Begin()
	//tmp_avatar := "http://101.132.65.222/upload/avatar/" + randstring.Rand1NumString() + ".jpg"
	//tmp_nickname := nickName()
	//tmp_uuid := randstring.GetRandAccntPwd()
	//tmp_realname := randstring.GetRandAccntPwd()
	//tmp_password := randstring.GetRandAccntPwd()
	//tmp_salt := randstring.GetRandAccntPwd()
	//_, err := dao.NewDaoBasic(tx).IsMobileExist(in.Mobile)
	//if err != nil { //不存在该手机号对应的账号，随机生成
	//	tmp_basic := &mdl.User_Basic{
	//		UuId:     tmp_uuid,
	//		RealName: tmp_realname,
	//		Gender:   1,
	//		Age:      25,
	//		IdCard:   "1234567890",
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
	//		Channel:  1,
	//		IsDelete: 100,
	//	}
	//	if err, _ := dao.NewDaoBasic(tx).CreateUser(tmp_basic); err != nil {
	//		output.HttpCode = http.StatusNotAcceptable
	//		output.Code = http.StatusNotAcceptable
	//		output.Msg = "手机注册失败!"
	//		output.Detail = "手机注册失败!"
	//		tx.Rollback()
	//		return nil
	//	}
	//	if err, _ := dao.NewDaoMember(tx).CreateUser(tmp_member); err != nil {
	//		output.HttpCode = http.StatusNotAcceptable
	//		output.Code = http.StatusNotAcceptable
	//		output.Msg = "手机注册失败!"
	//		output.Detail = "手机注册失败!"
	//		tx.Rollback()
	//		return nil
	//	}
	//	outresp.TbAcct.Uuid = tmp_uuid
	//	outresp.TbAcct.Level = "0"
	//	outresp.TbAcct.Exp = "0"
	//	outresp.TbAcct.NickName = tmp_nickname
	//	outresp.TbAcct.Avatar = tmp_avatar
	//	outresp.TbAcct.Vip = "0"
	//	outresp.TbAcct.VipValidityTime = "0"
	//	outresp.TbAcct.MyCoin = "0"
	//	outresp.TbAcct.BindPhone = in.Mobile
	//	outresp.TbAcct.BindEmail = "0"
	//	outresp.TbAcct.Stage = 1
	//	outresp.TbAcct.Series = 1
	//	tmpoutrsp, _ := json.Marshal(outresp)
	//	output.Data = tmpoutrsp
	//	output.HttpCode = http.StatusCreated
	//	output.Code = http.StatusCreated
	//	output.Msg = "手机注册成功!"
	//	output.Detail = "手机注册成功!"
	//	tx.Commit()
	//	return nil
	//} //存在该手机号对应的账号，登陆即可
	//loginrcd := &mdl.User_Login_Record{
	//	UUid:      tmp_uuid,
	//	Mobile:    in.Mobile,
	//	ChannelId: 1,
	//}
	//err = dao.NewDaoLoginRecord(tx).AddLoginRecord(loginrcd) //
	//if err != nil {
	//	output.HttpCode = http.StatusNotAcceptable
	//	output.Code = http.StatusNotAcceptable
	//	output.Msg = "创建登陆记录失败!"
	//	output.Detail = "创建登陆记录失败!"
	//	resp := &mresp.User_MultiLoginMobile{}
	//	buf, _ := json.Marshal(resp)
	//	output.Data = buf
	//	tx.Rollback()
	//	return nil
	//}
	//output.HttpCode = http.StatusOK
	//output.Code = http.StatusOK
	//output.Msg = "登陆成功!"
	//output.Detail = "登陆成功!"
	//tmp := &mresp.User_MultiLoginMobile{
	//	TbAcct: &mresp.BindAccount{
	//		Uuid:            "NULL",
	//		Level:           "0",
	//		Exp:             "0",
	//		NickName:        "NULL",
	//		Avatar:          "NULL",
	//		Vip:             "0",
	//		VipValidityTime: "NULL",
	//		MyCoin:          "0",
	//		BindPhone:       "NULL",
	//		BindEmail:       "",
	//		Stage:           0,
	//		Series:          0,
	//	},
	//	Token: CreateToken(in.Mobile, "NULL", "NULL", "multilogin"),
	//}
	//buf, _ := json.Marshal(tmp)
	//output.Data = buf
	//tx.Commit()
	return nil
}
