package services



import (
	//"microservice/jzapi/user_srv/dao"
	//r6 "microservice/jzapi/basic/randstring"
	//mdl "microservice/jzapi/user_srv/model"
	//"microservice/jzapi/lib/db"
	"time"
	"microservice/jzapi/basic/cfg/config"
	"github.com/dgrijalva/jwt-go"
	"microservice/jzapi/lib/token"
)

func nickName() string {
	//_, first := dao.NewDaoNickName(db.GetDB()).NickNameFirstPart(r6.GetRandSize(1, 1000))
	//_, second := dao.NewDaoNickName(db.GetDB()).NickNameSecondPart(r6.GetRandSize(1, 1000))
	return ""//first + second
}
func addLoginRecord(acctid, mobile, deviceid string,  channel int32, clientip string) string {
	//err := dao.NewDaoLoginRecord(db.GetDB()).AddLoginRecord(&mdl.User_Login_Record{
	//	UUid:  acctid,
	//	Mobile:     mobile,
	//	DeviceId:   deviceid,
	//	Channel:    uint(channel),
	//	Ip:         clientip,
	//	Operate:    "",
	//	Remark:     "",
	//	ClientType: "",
	//	ClientVer:  "",
	//})
	//if err != nil {
	//	return "更新登陆记录失败!"
	//}
	return "更新登陆记录成功!"
}

func CreateToken(mobile string,  account_id string, device_id string , issue string )string  {
	config.C().App(jwtkey, jwtcfg)
	type Token struct {
		*token.Subject
		siginKey string
	}
	tmptoken := Token{
		&token.Subject{mobile,account_id, device_id,
			jwt.StandardClaims{ExpiresAt: time.Now().Add(2 * time.Minute).Unix(), Issuer: issue}},
		jwtcfg.SecretKey,
	}
	jwttk := &token.JwtToken{SigningKey: []byte(jwtcfg.SecretKey)}
	tptoken, _ := jwttk.CreateToken(*tmptoken.Subject)
	return tptoken
}