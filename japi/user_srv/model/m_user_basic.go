package model

type User_Basic struct {
	Model
	RealName string `json:"realname"`
	Gender   string `json:"gender"`
	Age      int    `json:"age"`
	IdCard   string `json:"idcard"`
	Mobile   string `json:"mobile"`
	Eamil    string `json:"email"`
	Address  string `json:"address"`
}
