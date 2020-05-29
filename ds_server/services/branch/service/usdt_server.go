package service

import (
	"crypto/hmac"
	"crypto/sha256"
	"ds_server/services/branch/service/dto"
	"ds_server/support/utils/encoder"
	"ds_server/support/utils/logger"
	"ds_server/support/utils/query"
	"encoding/hex"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/prometheus/common/log"
)

type UsdtService struct {
}

var usdtService *UsdtService

func NewUsdtService() *UsdtService {
	if usdtService == nil {
		l.Lock()
		if usdtService == nil {
			usdtService = &UsdtService{}
		}
		l.Unlock()
	}
	return usdtService
}

func (service UsdtService) GetAddress() (res2 dto.GetAddressResponse, err error) {
	res := dto.GetAddressResponse{}
	url := "https://hoo.com/api/open/vip/v1/address"
	var req dto.GetAddressRequest
	req.ClientID = "Lu2jyts9qdYT3FbFTK1s6CJKUZ2zfX"
	req.CoinName = "ETH"
	req.Num = 1
	ss, _, _ := query.Values(req)
	enStr := ss.Encode()
	log.Infoln("===========enStr======s=======", enStr)

	sign := service.ComputeHmacSha256(enStr, "6fpvAq4X2aKCJ1inpLeY1GsU5CchTVNr9tQvnZMkBtjENubWVS")
	log.Infoln("===========sign=============", sign)

	req.Sign = sign

	log.Infoln("=========================", logger.FormatStruct(req))
	// reqJSON, err := json.Marshal(req)
	// if err != nil {
	// 	Log.Error(err)
	// 	return res, err
	// }

	err = service.UsPost(url, req, &res)
	if err != nil {
		Log.Error(err)
		return res, err
	}
	log.Infoln("=============res============", logger.FormatStruct(res))
	return res, nil
}

func (service UsdtService) GetBalance() (err error) {
	url := "https://hoo.com/api/open/vip/v1/accounts"
	var req dto.GetBalanceRequest
	req.ClientID = "Lu2jyts9qdYT3FbFTK1s6CJKUZ2zfX"
	req.CoinName = "ETH"
	ss, _, _ := query.Values(req)
	enStr := ss.Encode()
	log.Infoln("===========GetBalanceenStr=============", enStr)

	sign := service.ComputeHmacSha256(enStr, "6fpvAq4X2aKCJ1inpLeY1GsU5CchTVNr9tQvnZMkBtjENubWVS")
	req.Sign = sign
	log.Infoln("=============GetBalance============", logger.FormatStruct(req))
	// reqJSON, err := json.Marshal(req)
	// if err != nil {
	// 	Log.Error(err)
	// 	return
	// }
	res := &dto.GetBalanceResponse{}
	service.UsPost(url, req, res)
	log.Infoln("=============res============", logger.FormatStruct(res))
	return nil
}

func (service UsdtService) UsPost(url string, reqJSON interface{}, resp interface{}) (err error) {

	// var queryrequest JetRequest
	// queryrequest.TranType = "queryOrder"
	// queryrequest.Param = string([]byte(xml.Header)) + string(output)

	resvalues, _, _ := query.Values(reqJSON)

	//Log.Debug("paramres: ", paramres)
	//--记录request logex------
	//------------------
	body, err := http.PostForm(url, resvalues)

	//	body, err := httpex.Post(url, "application/json", strings.NewReader(string(encoder.Base64Encode(reqJSON))))
	if err != nil {
		Log.Error(err)
		return
	}
	body2, _ := ioutil.ReadAll(body.Body)
	Log.Debug("body: %v", string(body2))
	emsg := ""
	if err != nil {
		Log.Errorf("errs = %v", err)
		emsg = fmt.Sprintf("errs = %v", err)
		err = errors.New("请求三方服务器失败")
	} else if len(body2) == 0 {
		emsg = "三方服务器无效返回"
		err = errors.New(emsg)
	} else if err = encoder.UnserializeFromJson(string(body2), &resp); err != nil {
		emsg = "解析三方服务器返回消息失败"
		err = errors.New(emsg)
	} else {
		emsg = ""
	}
	return err
}

func (service UsdtService) ComputeHmacSha256(instr string, secret string) string {
	key := []byte(secret)
	h := hmac.New(sha256.New, key)
	h.Write([]byte(instr))
	sha := hex.EncodeToString(h.Sum(nil))
	//	hex.EncodeToString(h.Sum(nil))
	return sha
}
