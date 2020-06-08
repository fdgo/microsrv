package service

import (
	"context"
	rspmdl "ds_server/models/user/gin_rsp"
	mygormdl "ds_server/models/user/gorm_mysql"
	pb "ds_server/proto/user"
	"ds_server/support/lib/redisex"
	"ds_server/support/utils/constex"
	"ds_server/support/utils/sign/md5"
	"ds_server/support/utils/stringex"
	"encoding/json"
	errr "errors"
	"fmt"
	"github.com/go-errors/errors"
	"time"
)


func (usersrv *UserService) Regist(ctx context.Context, req *pb.RegistIn, rsq *pb.CommonOut) error {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	vfcode,_ := redisex.RedisInstanceg().Get(constex.REDIS_USER_VFCODE+req.Mobile).Result()
	if req.Verifycode != vfcode{
		return errr.New("短信验证码无效")
	}
	_, errret := usersrv.DsUserMemberAgentDao.IsInvCodeExist(req.Invcodeagent)
	if errret != nil {
		fmt.Println(400, 400, "邀请码无效!", "邀请码无效！", []byte(""), rsq)
		return errr.New("邀请码无效")
	}
	salt := stringex.GetRandomString(16)
	hash := md5.HashForPwd(salt, req.Pwd)
RECYCLE:
	uuid := stringex.GetRandomString(10)
	tx := usersrv.DsUserBasicinfoDao.Begin()
	err := tx.Create(&mygormdl.DsUserBasicinfo{UUID: uuid, Mobile: req.Mobile, Salt: salt, NickName: "DS" + stringex.Rand6NumString(), Birthday: time.Now().Local(),
		CreateTime: time.Now().Local(), UpdateTime: time.Now().Local(), LastLoginTime: time.Now().Local(), LastLoginIP: req.ClientIp, Hash: hash,
	}).Error
	if err != nil {
		if err.Error() == fmt.Sprintf("Error 1062: Duplicate entry '%s' for key 'mobile'", req.Mobile) {
			fmt.Println(400, 400, "手机号已经注册，请用新号码注册!", "手机号已经注册，请用新号码注册!", []byte(""), rsq)
			tx.Rollback()
			return errr.New("手机号已经注册，请用新号码注册!")
		}
		if err.Error() == fmt.Sprintf("Error 1062: Duplicate entry '%s' for key 'PRIMARY'", uuid) {
			tx.Rollback()
			goto RECYCLE
		}
	}
	err = tx.Create(&mygormdl.DsUserMemberAccount{UUID: uuid, Mobile: req.Mobile, Salt: "", Hash: "",
		Balance: 0, CreateTime: time.Now().Local(), UpdateTime: time.Now().Local()}).Error
	if err != nil {
		if err.Error() == fmt.Sprintf("Error 1062: Duplicate entry '%s' for key 'mobile'", req.Mobile) {
			fmt.Println(400, 400, "手机号已经注册，请用新号码注册!", "手机号已经注册，请用新号码注册!", []byte(""), rsq)
			tx.Rollback()
			return errr.New("手机号已经注册，请用新号码注册!")
		}
		if err.Error() == fmt.Sprintf("Error 1062: Duplicate entry '%s' for key 'PRIMARY'", uuid) {
			tx.Rollback()
			goto RECYCLE
		}
	}
	result, errer := usersrv.DsUserMemberAgentDao.GetUserMemAgentEx(tx, req.Invcodeagent)
	if errer != nil {
		fmt.Println(400, 400, "手机号已经注册，请用新号码注册!", "手机号已经注册，请用新号码注册!", []byte(""), rsq)
		tx.Rollback()
		return errr.New("邀请码错误!")
	}
	invcodeself := stringex.Rand6NumString()
	err = tx.Create(&mygormdl.DsUserMemberAgent{UUIDSelf: uuid, MobileSelf: req.Mobile, InvcodeSelf: invcodeself, UUIDAgent: result.UUIDSelf,
		MobileAgent: result.MobileSelf, InvcodeAgent: req.Invcodeagent, CreateTime: time.Now(), UpdateTime: time.Now(),
		MemclassSelf: 0, MemberTag: "", MemberName: "", AgentClass: 0,
		AgentTag: "", AgentName: ""}).Error
	if err != nil {
		if err.Error() == fmt.Sprintf("Error 1062: Duplicate entry '%s' for key 'mobile_self'", req.Mobile) {
			fmt.Println(400, 400, "手机号已经注册，请用新号码注册!", "手机号已经注册，请用新号码注册!", []byte(""), rsq)
			tx.Rollback()
			return errr.New("手机号已经注册，请用新号码注册!")
		}
		if err.Error() == fmt.Sprintf("Error 1062: Duplicate entry '%s' for key 'invcode_self'", invcodeself) {
			tx.Rollback()
			goto RECYCLE
		}
		if err.Error() == fmt.Sprintf("Error 1062: Duplicate entry '%s' for key 'PRIMARY'", uuid) {
			tx.Rollback()
			goto RECYCLE
		}
	}
	tx.Commit()
	fmt.Println(200, 200, "手机号注册成功!", "手机号注册成功!", []byte(""), rsq)
	return nil
}
func (usersrv *UserService) Login(ctx context.Context, req *pb.LoginIn, rsq *pb.CommonOut) error {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	tmp_resp := rspmdl.Login_rsp{}
	tmp_resp_err, _ := json.Marshal(tmp_resp)
	var err error
	if req.Type == "1" {
		userbasic, err := usersrv.DsUserBasicinfoDao.GetHashSalt(req.Mobile)
		if err != nil {
			fmt.Println(400, 400, "该手机号不存在!", "该手机号不存在!", tmp_resp_err, rsq)
			return errors.New("该手机号不存在!")
		}
		if userbasic.Hash != md5.HashForPwd(userbasic.Salt, req.Pwd) {
			fmt.Println(400, 400, "密码错误!", "密码错误!", tmp_resp_err, rsq)
			return errors.New("密码错误!")
		}
	} else if req.Type == "0" {
		mobilecode,_ := redisex.RedisInstanceg().Get(constex.REDIS_USER_VFCODE+req.Mobile).Result()
		if req.Mobidecode!= mobilecode{
			return errr.New("短信验证码无效")
		}
		//手机验证码
		_, err = usersrv.DsUserBasicinfoDao.LoginVfcode(req.Mobile)
		if err != nil {
			fmt.Println(400, 400, "该手机号不存在或者验证码错误!", "该手机号不存在或者验证码错误!", tmp_resp_err, rsq)
			return errors.New("该手机号不存在或者验证码错误!")
		}
	} else {
		fmt.Println(400, 400, "登陆方式有误!", "登陆方式有误!", tmp_resp_err, rsq)
		return err
	}
	tx := usersrv.DsUserBasicinfoDao.Begin()
	basicinfo, errb := usersrv.DsUserBasicinfoDao.GetUserBasicInfoex(tx, req.Mobile)
	if errb != nil {
		fmt.Println(400, 400, "获取会员信息失败!", "获取会员信息失败", tmp_resp_err, rsq)
		tx.Rollback()
		return err
	}
	memagent, errm := usersrv.DsUserMemberAgentDao.GetUserMemAgent(tx, basicinfo.UUID)
	if errm != nil {
		fmt.Println(400, 400, "获取会员信息失败!", "获取会员信息失败", tmp_resp_err, rsq)
		tx.Rollback()
		return err
	}
	memacct, errat := usersrv.DsUserMemberAccountDao.GetSelfMemberAccount(tx, basicinfo.UUID)
	if errat != nil {
		fmt.Println(400, 400, "获取会员信息失败!", "获取会员信息失败", tmp_resp_err, rsq)
		tx.Rollback()
		return err
	}
	tmp_resp.NickName = basicinfo.NickName
	tmp_resp.UUIDSelf = basicinfo.UUID
	tmp_resp.MobileSelf = req.Mobile
	tmp_resp.InvcodeSelf = memagent.InvcodeSelf
	tmp_resp.AgentName = memagent.AgentName
	tmp_resp.AgentTag = memagent.AgentTag
	tmp_resp.AgentClass = memagent.AgentClass
	tmp_resp.MemberName = memagent.MemberName
	tmp_resp.MemberTag = memagent.MemberTag
	tmp_resp.MemclassSelf = memagent.MemclassSelf
	tmp_resp.InvcodeAgent = memagent.InvcodeAgent
	tmp_resp.MobileAgent = memagent.MobileAgent
	tmp_resp.UUIDAgent = memagent.UUIDAgent
	tmp_resp.Balance = memacct.Balance
	tmp_resp.Ispwd = memacct.Ispwd
	tmp_resp.Token ="Bearer " + newtoken(basicinfo.UUID, req.Mobile, memagent.InvcodeSelf, memagent.InvcodeAgent)
	err = usersrv.DsUserBasicinfoDao.SetLastLogin(tx, req.Mobile, req.ClientIp, time.Now().Local())
	if err != nil {
		tx.Rollback()
		fmt.Println(400, 400, "设置登陆记录失败!", err.Error(), tmp_resp_err, rsq)
		return err
	}
	tx.Commit()
	tmp_resp_nor, _ := json.Marshal(tmp_resp)
	fmt.Println(200, 200, "登陆成功!", "登陆成功!", tmp_resp_nor, rsq)
	return nil
}

func (usersrv *UserService) SetPaypwd(ctx context.Context, req *pb.SetPaypwdIn, rsq *pb.CommonOut) error {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	vfcode,_ := redisex.RedisInstanceg().Get(constex.REDIS_USER_VFCODE+req.Mobile).Result()
	if req.Vfcode != vfcode{
		return errr.New("短信验证码无效")
	}
	e := usersrv.DsUserMemberAccountDao.SetPaypwd(req.Uuid,req.Paypwd,req.Vfcode,req.Mobile)
	if e != nil {
		fmt.Println(400, 400, e.Error(), e.Error(), []byte(""), rsq)
		return e
	}
	fmt.Println(200, 200, "支付密码设置成功!", "支付密码设置成功!", []byte(""), rsq)
	return nil
}
func (usersrv *UserService) ModifyBasicPwd(ctx context.Context, req *pb.ModifyBasicPwdIn, rsq *pb.CommonOut) error {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	e := usersrv.DsUserBasicinfoDao.ModifyUserInfo(req.Tag, req.Mobile,req.Vfcode, req.Content)
	if e != nil {
		fmt.Println(400, 400, e.Error(), e.Error(), []byte(""), rsq)
		return e
	}
	fmt.Println(200, 200, "修改成功!", "修改成功!", []byte(""), rsq)
	return nil
}
func (usersrv *UserService) ModifyPayPwd(ctx context.Context, req *pb.ModifyPayPwdIn, rsq *pb.CommonOut) error {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	vfcode,_ := redisex.RedisInstanceg().Get(constex.REDIS_USER_VFCODE+req.Mobile).Result()
	if req.Vfcode != vfcode{
		return errr.New("短信验证码无效")
	}
	e := usersrv.DsUserMemberAccountDao.ModifyPayPwd(req.Uuid, req.Newpwd)
	if e != nil {
		fmt.Println(400, 400, e.Error(), e.Error(), []byte(""), rsq)
		return e
	}
	fmt.Println(200, 200, "修改成功!", "修改成功!", []byte(""), rsq)
	return nil
}
func (usersrv *UserService) ModifyLoginPwd(ctx context.Context, req *pb.ModifyLoginPwdIn, rsq *pb.CommonOut) error {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	basicinfo, errb := usersrv.DsUserBasicinfoDao.GetUserBasicInfoEx(req.Mobile)
	if errb != nil {
		fmt.Println(400, 400, "获取会员旧密码失败!", "获取会员旧密码失败", []byte("false"), rsq)
		return errb
	}
	tmpoldhash := md5.HashForPwd(basicinfo.Salt, req.Oldpwd)
	if tmpoldhash != basicinfo.Hash{
		fmt.Println(400, 400, "旧密码错误!", "旧密码错误", []byte("false"), rsq)
		return errb
	}
	newhash := md5.HashForPwd(basicinfo.Salt, req.Newpwd)
	errx := usersrv.DsUserBasicinfoDao.UpdateLoginPwd(req.Mobile,newhash)
	if errx!=nil{
		fmt.Println(400, 400, "登陆密码修改失败!", "登陆密码修改失败!", []byte("false"), rsq)
		return errx
	}
	fmt.Println(200, 200, "登陆密码修改成功!", "登陆密码修改成功", []byte("true"), rsq)
	return nil
}
func (usersrv *UserService) GetMemberUserAgent(ctx context.Context, req *pb.GetMemberUserAgentIn, rsq *pb.CommonOut) error {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	tmp_resp := rspmdl.MemAgent_rsp{}
	tmp_resp_err, _ := json.Marshal(tmp_resp)
	tx := usersrv.DsUserBasicinfoDao.Begin()
	basicinfo, errb := usersrv.DsUserBasicinfoDao.GetUserBasicInfo(tx, req.Uuid)
	if errb != nil {
		fmt.Println(400, 400, "获取会员信息失败!", "获取会员信息失败", tmp_resp_err, rsq)
		tx.Rollback()
		return errb
	}
	memagent, errm := usersrv.DsUserMemberAgentDao.GetUserMemAgent(tx, basicinfo.UUID)
	if errm != nil {
		fmt.Println(400, 400, "获取会员信息失败!", "获取会员信息失败", tmp_resp_err, rsq)
		tx.Rollback()
		return errm
	}
	memacct, errat := usersrv.DsUserMemberAccountDao.GetSelfMemberAccount(tx, basicinfo.UUID)
	if errat != nil {
		fmt.Println(400, 400, "获取会员信息失败!", "获取会员信息失败", tmp_resp_err, rsq)
		tx.Rollback()
		return errat
	}
	tmp_resp.UUIDSelf = basicinfo.UUID
	tmp_resp.MobileSelf = basicinfo.Mobile
	tmp_resp.InvcodeSelf = memagent.InvcodeSelf
	tmp_resp.AgentName = memagent.AgentName
	tmp_resp.AgentTag = memagent.AgentTag
	tmp_resp.AgentClass = memagent.AgentClass
	tmp_resp.MemberName = memagent.MemberName
	tmp_resp.MemberTag = memagent.MemberTag
	tmp_resp.MemclassSelf = memagent.MemclassSelf
	tmp_resp.InvcodeAgent = memagent.InvcodeAgent
	tmp_resp.MobileAgent = memagent.MobileAgent
	tmp_resp.UUIDAgent = memagent.UUIDAgent
	tmp_resp.Balance = memacct.Balance
	tmp_resp.Ispwd = memacct.Ispwd
	tmp_resp_nor, _ := json.Marshal(tmp_resp)
	fmt.Println(200, 200, "获取会员信息成功!", "获取会员信息成功!", tmp_resp_nor, rsq)
	return nil
}
