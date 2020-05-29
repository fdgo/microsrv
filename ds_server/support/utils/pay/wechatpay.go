package pay

import (
	"fmt"
	"net/http"
	"strconv"
	"time"
	"encoding/json"

	"ds_server/support/utils/constex"
	"github.com/iGoogle-ink/gopay"
	"github.com/iGoogle-ink/gopay/wechat"
	"github.com/shopspring/decimal"
)

// UnifiedOrder统一下单接口：title 标题,ordernum 内部订单号,price 付款金额,clientip 客户IP
func GetWechatpayAppParam(title, ordernum string, price decimal.Decimal, clientip string) (string, error) {
	client := wechat.NewClient(constex.WechatpayCfg.AppID, constex.WechatpayCfg.MchID, constex.WechatpayCfg.ApiKey, false)

	//设置国家
	client.SetCountry(wechat.China)

	//初始化参数Map
	bm := make(gopay.BodyMap)
	bm.Set("nonce_str", gopay.GetRandomString(32))
	bm.Set("body", title)
	bm.Set("out_trade_no", ordernum)
	bm.Set("total_fee", price.Mul(decimal.New(100, 2)).IntPart())
	bm.Set("spbill_create_ip", clientip)
	bm.Set("notify_url", constex.WechatpayCfg.CallBackURL)
	bm.Set("trade_type", wechat.TradeType_App)
	bm.Set("device_info", "WEB")
	bm.Set("sign_type", wechat.SignType_MD5)
	sign := wechat.GetParamSign(constex.WechatpayCfg.AppID, constex.WechatpayCfg.MchID, constex.WechatpayCfg.ApiKey, bm)
	bm.Set("sign", sign)

	//请求支付下单，成功后得到结果
	wxRsp, err := client.UnifiedOrder(bm)
	if err != nil {
		fmt.Println("Error:", err)
		return "", err
	}
	fmt.Println("wxRsp:", *wxRsp)
	timeStamp := strconv.FormatInt(time.Now().Unix(), 10)

	paySign := wechat.GetAppPaySign(constex.WechatpayCfg.AppID, constex.WechatpayCfg.PartnerID, wxRsp.NonceStr, wxRsp.PrepayId, wechat.SignType_MD5, timeStamp, constex.WechatpayCfg.ApiKey)
	appparam, _ := json.Marshal(map[string]interface{}{
		"appid":     constex.WechatpayCfg.AppID,
		"partnerid": constex.WechatpayCfg.PartnerID,
		"prepayid":  wxRsp.PrepayId,
		"package":   "Sign=WXPay",
		"noncestr":  wxRsp.NonceStr,
		"timestamp": timeStamp,
		"sign":      paySign,
	})
	return string(appparam), nil
}

// 返回三个参数 是否成功
func WechatPayCallBack(req *http.Request) (bool, decimal.Decimal, error) {
	client := wechat.NewClient(constex.WechatpayCfg.AppID, constex.WechatpayCfg.MchID, constex.WechatpayCfg.ApiKey, false)
	notifyReq, err := wechat.ParseRefundNotifyResult(req)
	// ==解密退款异步通知的加密参数 req_info ==
	refundNotify, err := wechat.DecryptRefundNotifyReqInfo(notifyReq.ReqInfo, constex.WechatpayCfg.ApiKey)
}

func WechatPayResponse(ok bool)string{
	rsp := new(wechat.NotifyResponse) // 回复微信的数据
	rsp.ReturnCode = gopay.SUCCESS
	if ok {
		rsp.ReturnMsg = gopay.OK
	}else{
		rsp.ReturnMsg = gopay.FAIL
	}
	return rsp.ToXmlString()

}