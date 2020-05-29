package model_response

//import "microservice/jzapi/user_srv/model"

type User_RegistQuick struct {
	Uuid        string      `json:"uuid"`
	Password    string      `json:"password "`
	BindAccount BindAccount `json:"bind_account"`
}
type BindAccount struct {
	Uuid            string `json:"uuid"`
	DeviceId        string `json:"device_id"`
	Level           string `json:"level"`
	Exp             string `json:"exp"`
	NickName        string `json:"nickname"`
	Avatar          string `json:"avatar"`
	Vip             string `json:"vip"`
	VipValidityTime string `json:"vip_validity_time"`
	MyCoin          string `json:"mycoin"`
	BindPhone       string `json:"bind_phone"`
	BindEmail       string `json:"bind_email"`
	Stage           int    `json:"stage"`
	Series          int    `json:"series"`
}
type User_MultiLoginMobile struct {
	TbAcct *BindAccount `json:"bind_account"`
	Token  string       `json:"token"`
}
type Admin_Manage_User struct {
	//UserStaff *model.User_Back
}
