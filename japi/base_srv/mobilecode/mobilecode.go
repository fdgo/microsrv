package mobilecode

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/dysmsapi"
	"github.com/go-redis/redis"
	"io/ioutil"
	tx "microservice/jzapi/basic/time_ex"
	"net"
	"net/http"
	"time"
)

var (
	host                = "https://cdcxdxjk.market.alicloudapi.com/chuangxin/dxjk?"
	appcode             = "39b7bca9b42342fa8ff7d4298a4f1a8e"
	CodeMobileKeyPrefix = "code:mobile:"
)

type MobileCodeTime struct {
	Times    int64  `json:"times"`
	Mobile   string `json:"mobile"`
	Code     string `json:"code"`
	Time     int64  `json:"time"`
}

func SendMobileCode(rcli *redis.Client, mct *MobileCodeTime, times, timediff int64) error {
	result, err := rcli.BLPop(time.Millisecond*500, CodeMobileKeyPrefix+mct.Mobile).Result()
	if err == redis.Nil {
		mct.Times = mct.Times + 1
		err := MobileCodeNew(mct.Mobile, mct.Code)
		if err != nil {
			return errors.New("阿里发送接口调用失败!")
		}
		return SaveCodeToRedis(rcli, mct)
	} else {
		var tmp MobileCodeTime
		json.Unmarshal([]byte(result[1]), &tmp)
		if (tx.GetCurrentTimeStamp()-tmp.Time) > timediff {
			mct.Times = mct.Times + 1
			err := MobileCodeNew(mct.Mobile, mct.Code)
			if err != nil {
				return errors.New("阿里发送接口调用失败!")
			}
			return SaveCodeToRedis(rcli, mct)
		} else {
			mct.Times = tmp.Times + 1
			if mct.Times > times {
				SaveCodeToRedis(rcli, mct) //必须留一个记录
				return errors.New(fmt.Sprintf("%d秒内发送验证码超过%d次！", timediff, times))
			}
			err := MobileCodeNew(mct.Mobile, mct.Code)
			if err != nil {
				return errors.New("阿里发送接口调用失败!")
			}
			return SaveCodeToRedis(rcli, mct)
		}
	}
}
func SaveCodeToRedis(rcli *redis.Client, mct *MobileCodeTime) error {
	tmp_mct, _ := json.Marshal(*mct)
	return rcli.RPush(CodeMobileKeyPrefix+mct.Mobile, tmp_mct).Err()
}
func GetCodeFromRedis(rcli *redis.Client, strmobile string) (error, string) {
	mct, err := rcli.LIndex(strmobile, 0).Result()
	if err != nil {
		return err, "NULL"
	}
	return err, mct
}

func MobileCodeNew(mobile string, str6num string) error {
	client, err := dysmsapi.NewClientWithAccessKey("cn-hangzhou", "LTAI4Fn43K5Xpep7XedX6qKj", "QF3OC5csVY0UltKpcJL2D7SNP7hscz")
	request := dysmsapi.CreateSendSmsRequest()
	request.Scheme = "https"
	request.PhoneNumbers = mobile
	request.SignName = "饺子"
	request.TemplateCode = "SMS_175539027"
	request.TemplateParam = `{"code":"` + str6num + `"}`
	request.Content = []byte("验证码" + str6num + "，有效期10分钟，请勿告知他人")
	//request.Content = []byte("验证码"+"123456，"+"有效期10分钟，请勿告知他人")

	_, err = client.SendSms(request)
	return err
}
func MobileCodeOld(code, mobile string) {
	client := &http.Client{
		Timeout: 30 * time.Second,
		Transport: &http.Transport{
			IdleConnTimeout: 3 * time.Minute,
			//MaxConnsPerHost: 10000,
			TLSHandshakeTimeout: 10 * time.Second,
			DialContext: (&net.Dialer{
				Timeout:   30 * time.Second,
				KeepAlive: 10 * time.Minute,
				DualStack: true,
			}).DialContext,
		},
	}
	fmt.Println("【嚼指云游戏】.........................")
	querys := "content=【嚼指云游戏】您的验证码是：" + code + "，10分钟内有效！&mobile=" + mobile
	req := bytes.NewBuffer([]byte(""))
	reqest, _ := http.NewRequest("POST", host+querys, req)
	reqest.Header.Set("Authorization", "APPCODE "+appcode)
	reqest.Header.Set("Content-Type", "application/json")
	fmt.Println("client.Do(reqest).........................")
	response, _ := client.Do(reqest)
	defer response.Body.Close()
	ioutil.ReadAll(response.Body)
}
