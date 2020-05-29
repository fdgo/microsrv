package common

type Wechatpay struct {
	AppID       string `json:"appId"`
	MchID       string `json:"mchId"`
	ApiKey      string `json:"apiKey"`
	PartnerID   string `json:"partnerId"`
	CallBackURL string `json:"callBackURL"`
}

type Alipay struct {
	AppID       string `json:"appId"`
	PriKey      string `json:"priKey"`
	CallBackURL string `json:"callBackURL"`
}
