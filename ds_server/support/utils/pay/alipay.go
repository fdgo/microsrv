package pay

import (
	"ds_server/support/utils/constex"
	"github.com/iGoogle-ink/gopay"
	"github.com/iGoogle-ink/gopay/alipay"
	"github.com/shopspring/decimal"
)

func GetAlipayAppParam(title, ordernum string, price decimal.Decimal) (string, error) {
	client := alipay.NewClient(constex.AlipayCfg.AppID, constex.AlipayCfg.PriKey, false)
	client.SetCharset("utf-8").
		SetSignType(alipay.RSA2).
		SetPrivateKeyType(alipay.PKCS1).
		SetNotifyUrl(constex.AlipayCfg.CallBackURL)

	//请求参数
	body := make(gopay.BodyMap)
	body.Set("subject", title)
	body.Set("out_trade_no", ordernum)
	body.Set("total_amount", price.String())
	//手机APP支付参数请求
	payParam, err := client.TradeAppPay(body)
	if err != nil {
		//fmt.Println("err:", err)
		return "",err
	}
	//fmt.Println("payParam:", payParam)
	return payParam,err
}