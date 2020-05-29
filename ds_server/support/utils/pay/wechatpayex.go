package pay

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"
)

//----------------------------------------------------------------------------

type WeixinClient struct {
	AppID       string
	AppSecret   string
	AccessToken string
	ExpireTime  int64
}

func NewWeixinClient(id string, secret string) *WeixinClient {
	r := new(WeixinClient)

	r.AppID = id
	r.AppSecret = secret
	r.AccessToken = ""
	r.ExpireTime = 0

	return r
}

//----------------------------------------------------------------------------

func (wx *WeixinClient) httpGet(url string) (map[string]interface{}, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	buf, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	obj := make(map[string]interface{})

	err = json.Unmarshal(buf, &obj)
	if err != nil {
		return nil, err
	}

	errMsg, okay := obj["errmsg"].(string)
	if okay {
		return obj, errors.New(errMsg)
	}

	return obj, nil
}

func (wx *WeixinClient) httpPost(url string, body string) (map[string]interface{}, error) {
	resp, err := http.Post(url, "application/x-www-form-urlencoded", strings.NewReader(body))
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	buf, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	fmt.Println(string(buf))

	obj := make(map[string]interface{})

	err = json.Unmarshal(buf, &obj)
	if err != nil {
		return nil, err
	}

	errMsg, okay := obj["errmsg"].(string)
	if !okay {
		return obj, ERR_INVALID_RESPONSE
	}
	if errMsg != "ok" {
		return obj, errors.New(errMsg)
	}

	return obj, nil
}

//----------------------------------------------------------------------------

func (wx *WeixinClient) CheckAccessToken() error {
	now := time.Now().Unix()
	if now < wx.ExpireTime-5 {
		return nil
	}

	return wx.UpdateAccessToken()
}

//----------------------------------------------------------------------------

func (wx *WeixinClient) UpdateAccessToken() error {
	url := "https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=" + wx.AppID +
		"&secret=" + wx.AppSecret

	obj, err := wx.httpGet(url)
	if err != nil {
		return err
	}

	accessToken, okay := obj["access_token"].(string)
	if !okay {
		return ERR_INVALID_RESPONSE
	}

	expiresIn, okay := obj["expires_in"].(float64)
	if !okay {
		return ERR_INVALID_RESPONSE
	}

	wx.AccessToken = accessToken
	wx.ExpireTime = time.Now().Unix() + int64(expiresIn)

	return nil
}

//----------------------------------------------------------------------------

func (wx *WeixinClient) Login(code string) (string, error) {
	url := "https://api.weixin.qq.com/sns/jscode2session?appid=" + wx.AppID +
		"&secret=" + wx.AppSecret +
		"&js_code=" + code +
		"&grant_type=authorization_code"

	obj, err := wx.httpGet(url)
	if err != nil {
		return "", err
	}

	openID, okay := obj["openid"].(string)
	if !okay {
		return "", ERR_INVALID_RESPONSE
	}

	return openID, nil
}

//----------------------------------------------------------------------------

func (wx *WeixinClient) GetTemplates(offset int, count int) (map[string]string, error) {
	err := wx.CheckAccessToken()
	if err != nil {
		return nil, err
	}

	url := "https://api.weixin.qq.com/cgi-bin/wxopen/template/list?access_token=" + wx.AccessToken

	obj, err := wx.httpPost(url, `{"offset":`+strconv.Itoa(offset)+`,"count":`+strconv.Itoa(count)+`}`)
	if err != nil {
		return nil, err
	}

	ls, okay := obj["list"].([]interface{})
	if !okay {
		return nil, ERR_INVALID_RESPONSE
	}

	m := make(map[string]string)
	for i := 0; i < len(ls); i++ {
		item, okay := ls[i].(map[string]interface{})
		if !okay {
			continue
		}

		id, okay := item["template_id"].(string)
		if !okay {
			continue
		}
		title, okay := item["title"].(string)
		if !okay {
			continue
		}

		m[id] = title
	}
	return m, nil
}

//----------------------------------------------------------------------------

func (wx *WeixinClient) SendMessage(openID, templateID, page, formID, data string) error {
	err := wx.CheckAccessToken()
	if err != nil {
		return err
	}

	url := "https://api.weixin.qq.com/cgi-bin/message/wxopen/template/send?access_token=" + wx.AccessToken
	body := `{` +
		`"touser":"` + openID + `",` +
		`"template_id":"` + templateID + `",` +
		`"page":"` + page + `",` +
		`"form_id":"` + formID + `",` +
		`"data":` + data +
		`}`

	_, err = wx.httpPost(url, body)
	if err != nil {
		return err
	}

	return nil
}

//----------------------------------------------------------------------------
