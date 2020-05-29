package model

type MobileCode_req struct {
	Mobile string `json:"mobile" validate:"required,min=11,max=11"`
}

type Regist_req struct {
	InvCodeAgent string `json:"invcodeagent" validate:"min=6,max=6"`
	VerifyCode   string `json:"verifycode" validate:"required,min=6,max=6"`
	Mobile       string `json:"mobile" validate:"required,min=11,max=11"`
	Pwd          string `json:"pwd" validate:"required,min=6,max=16"`
}
type Login_req struct {
	Type   string `json:"type" validate:"required,min=1,max=1"`
	Mobile string `json:"mobile" validate:"required,min=11,max=11"`
	Pwd    string `json:"pwd"`
	VfCode string `json:"vfcode"`
}
type ModifyLoginPwd_req struct {
	OldPwd string `json:"oldpwd" validate:"required,min=6,max=6"`
	NewPwd string `json:"newpwd" validate:"required,min=6,max=6"`
	Mobile string `json:"mobile" validate:"required,min=11,max=11"`
}
type ModifyBasicPwd_req struct {
	Tag     string `json:"tag" validate:"required"`
	Mobile  string `json:"mobile" validate:"required,min=11,max=11"`
	VfCode  string `json:"vfcode" validate:"required,min=6,max=6"`
	Content string `json:"content" validate:"required"`
}
type ModifyPayPwd_req struct {
	Mobile     string `json:"mobile" validate:"required,min=11,max=11"`
	NewPwd     string `json:"newpwd" validate:"required,min=6,max=6"`
	VerifyCode string `json:"vfcode" validate:"required,min=6,max=6"`
}
type SetPaypwd struct {
	Mobile     string `json:"mobile" validate:"required,min=11,max=11"`
	VerifyCode string `json:"vfcode" validate:"required,min=6,max=6"`
	PayPwd     string `json:"paypwd" validate:"required,min=6,max=6"`
}

type Publish_req struct {
	Msg string `json:"msg" validate:"required,min=6,max=100"`
}
type Webst_req struct {
	Count int64 `json:"count"`
}
type MemberDeposit_req struct {
	Mobile      string  `json:"mobile" validate:"required,min=11,max=11"`
	DepositType int32   `json:"deposittype" validate:"required,min=1,max=16"`
	DepositNum  float32 `json:"depositnum" validate:"required,min=0.009,max=1000000000000"`
	DepositName string  `json:"depositname" validate:"required,min=2,max=16"`
	AddressIn   string  `json:"addressin"`
	AddressOut  string  `json:"addressout"`
}
type OnlinePay_req struct {
	SrcId       string  `json:"srcid" validate:"required,min=6,max=128"`
	DepositType int32   `json:"deposittype" validate:"required,min=1,max=16"`
	DepositNum  float32 `json:"depositnum" validate:"required,min=0.009,max=1000000000000"`
	DepositName string  `json:"depositname" validate:"required,min=2,max=16"`
	AddressIn   string  `json:"addressin"`
	AddressOut  string  `json:"addressout"`
}
type MemberDepositLog_req struct {
	Mobile string `json:"mobile" validate:"required,min=11,max=11"`
	Type   string `json:"type" validate:"required,min=1,max=16"`
}

type MemClassSet_req struct {
	Money1 float32 `json:"money1" validate:"required,min=100,max=10000000"`
	Money2 float32 `json:"money2" validate:"required,min=100,max=10000000"`
	Money3 float32 `json:"money3" validate:"required,min=100,max=10000000"`
	Money4 float32 `json:"money4" validate:"required,min=100,max=10000000"`
	Money5 float32 `json:"money5" validate:"required,min=100,max=10000000"`

	Tag1 int32 `json:"tag1"`
	Tag2 int32 `json:"tag2"`
	Tag3 int32 `json:"tag3"`
	Tag4 int32 `json:"tag4"`
	Tag5 int32 `json:"tag5"`

	Tagex1 string `json:"tagex1"`
	Tagex2 string `json:"tagex2"`
	Tagex3 string `json:"tagex3"`
	Tagex4 string `json:"tagex4"`
	Tagex5 string `json:"tagex5"`

	MemName1 string `json:"memname1"`
	MemName2 string `json:"memname2"`
	MemName3 string `json:"memname3"`
	MemName4 string `json:"memname4"`
	MemName5 string `json:"memname5"`
}

type AgentClassSet_req struct {
	Money1 float32 `json:"money1" validate:"required,min=100,max=100000000"`
	Money2 float32 `json:"money2" validate:"required,min=100,max=100000000"`
	Money3 float32 `json:"money3" validate:"required,min=100,max=100000000"`
	Money4 float32 `json:"money4" validate:"required,min=100,max=100000000"`
	Money5 float32 `json:"money5" validate:"required,min=100,max=100000000"`

	Tag1 int32 `json:"tag1"`
	Tag2 int32 `json:"tag2"`
	Tag3 int32 `json:"tag3"`
	Tag4 int32 `json:"tag4"`
	Tag5 int32 `json:"tag5"`

	Tagex1 string `json:"tagex1"`
	Tagex2 string `json:"tagex2"`
	Tagex3 string `json:"tagex3"`
	Tagex4 string `json:"tagex4"`
	Tagex5 string `json:"tagex5"`

	AgentName1 string `json:"agentname1"`
	AgentName2 string `json:"agentname2"`
	AgentName3 string `json:"agentname3"`
	AgentName4 string `json:"agentname4"`
	AgentName5 string `json:"agentname5"`
}
