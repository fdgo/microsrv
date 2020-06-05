package user

import (
	"ds_server/client"
	reqmdl "ds_server/models/user/gin_req"
	rspmdl "ds_server/models/user/gin_rsp"
	mygormdl "ds_server/models/user/gorm_mysql"
	useproto "ds_server/proto/user"
	"ds_server/support/utils/constex"
	"strconv"

	"ds_server/support/utils/param"
	"ds_server/support/utils/rsp"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
)

func ExchangeRate(c *gin.Context) {
	var ret_resp rspmdl.DsUserWalletExchangeRate_rsp
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	var rtin useproto.ExchangeRateIn
	ret, err := client.UserClient.ExchangeRate(c, &rtin)
	if err != nil {
		rsp.RespGin(400, 400, err.Error(), err.Error(), ret_resp, c)
		return
	}
	json.Unmarshal((*ret).Data, &ret_resp)
	rsp.RespGin((*ret).Httpcode, (*ret).Innercode, (*ret).Clientmsg, (*ret).Innermsg, ret_resp, c)
}
func GetMemberUserAgent(c *gin.Context) {
	var ret_resp rspmdl.MemAgent_rsp
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	var rtin useproto.GetMemberUserAgentIn
	rtin.Uuid = c.Request.Header.Get("X-Head-Uuid")
	ret, err := client.UserClient.GetMemberUserAgent(c, &rtin)
	if err != nil {
		rsp.RespGin(400, 400, err.Error(), err.Error(), ret_resp, c)
		return
	}
	json.Unmarshal((*ret).Data, &ret_resp)
	rsp.RespGin((*ret).Httpcode, (*ret).Innercode, (*ret).Clientmsg, (*ret).Innermsg, ret_resp, c)
}
func MemberDepositLog(c *gin.Context) {
	var ret_resp rspmdl.MemberDepositLog_rsp
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	ty := c.Query("type")
	tmppagesize := c.Query("pagesize")
	pagesize, _ := strconv.Atoi(tmppagesize)
	tmpindex := c.Query("index")
	index, _ := strconv.Atoi(tmpindex)
	if ty == "all" || ty == "USDT" || ty == "ali_pay" || ty == "wechat_pay" {
		var rtin useproto.MemberDepositLogIn
		rtin.Index = int32(index)
		rtin.PageSize = int32(pagesize)
		rtin.TypeName = ty
		rtin.Mobile = c.Request.Header.Get("X-Head-Mobile")
		rtin.Uuid = c.Request.Header.Get("X-Head-Uuid")
		ret, err := client.UserClient.MemberDepositLog(c, &rtin)
		if err != nil {
			rsp.RespGin(400, 400, "服务调用失败!", err.Error(), ret_resp, c)
			return
		}
		json.Unmarshal((*ret).Data, &ret_resp)
		rsp.RespGin((*ret).Httpcode, (*ret).Innercode, (*ret).Clientmsg, (*ret).Innermsg, ret_resp, c)
		return
	}
	rsp.RespGin(400, 400, "充值类型错误!", "充值类型错误!", ret_resp, c)
	return
}
func MemberDeposit(c *gin.Context) {
	var ret_resp rspmdl.MemAgent_rsp
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	var in reqmdl.MemberDeposit_req
	if err := c.ShouldBindJSON(&in); err != nil {
		rsp.RespGin(400, 400, err.Error(), err.Error(), ret_resp, c)
		return
	}
	isok, _ := param.IsParam(in)
	if !isok {
		rsp.RespGin(400, 400, "输入有误,请重写输入!", "参数有误,请重写输入!", ret_resp, c)
		return
	}
	var rtin useproto.MemberDepositIn
	rtin.DepositNum = in.DepositNum
	rtin.DepositType = in.DepositType
	rtin.DepositName = in.DepositName
	rtin.Uuid = c.Request.Header.Get("X-Head-Uuid")
	rtin.Mobile = c.Request.Header.Get("X-Head-Mobile")
	rtin.InvcodeSelf = c.Request.Header.Get("X-Head-InvCodeSelf")
	rtin.InvcodeAgent = c.Request.Header.Get("X-Head-InvCodeAgent")
	ret, err := client.UserClient.MemberDeposit(c, &rtin)
	if err != nil {
		rsp.RespGin(400, 400, err.Error(), err.Error(), ret_resp, c)
		return
	}
	json.Unmarshal((*ret).Data, &ret_resp)
	rsp.RespGin((*ret).Httpcode, (*ret).Innercode, (*ret).Clientmsg, (*ret).Innermsg, ret_resp, c)
}

func OnlinePay(c *gin.Context) {
	var ret_resp mygormdl.DsUserMemberDepositHistory
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	var in reqmdl.OnlinePay_req
	if err := c.ShouldBindJSON(&in); err != nil {
		rsp.RespGin(400, 400, err.Error(), err.Error(), ret_resp, c)
		return
	}
	isok, _ := param.IsParam(in)
	if !isok {
		rsp.RespGin(400, 400, "输入有误,请重写输入!", "参数有误", ret_resp, c)
		return
	}
	var rtin useproto.OnlinePayIn
	rtin.SrcId = in.SrcId
	rtin.DepositNum = in.DepositNum
	rtin.DepositType = in.DepositType
	rtin.Uuid = c.Request.Header.Get("X-Head-Uuid")
	//rtin.Mobile = c.Request.Header.Get("X-Head-Mobile")
	//rtin.InvcodeSelf = c.Request.Header.Get("X-Head-InvCodeSelf")
	//rtin.InvcodeAgent = c.Request.Header.Get("X-Head-InvCodeAgent")
	ret, err := client.UserClient.OnlinePay(c, &rtin)
	if err != nil {
		rsp.RespGin(400, 400, err.Error(), err.Error(), ret_resp, c)
		return
	}
	json.Unmarshal((*ret).Data, &ret_resp)
	rsp.RespGin((*ret).Httpcode, (*ret).Innercode, (*ret).Clientmsg, (*ret).Innermsg, ret_resp, c)
}
func AgentClassSet(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	var in reqmdl.AgentClassSet_req
	if err := c.ShouldBindJSON(&in); err != nil {
		rsp.RespGin(400, 400, "输入有误,请重写输入!", "参数有误", "", c)
		return
	}
	isok, _ := param.IsParam(in)
	if !isok {
		rsp.RespGin(400, 400, "输入有误,请重写输入!", "参数有误", "", c)
		return
	}
	if in.Money1 > in.Money2 || in.Money2 > in.Money3 || in.Money3 > in.Money4 || in.Money4 > in.Money5 {
		rsp.RespGin(400, 400, "金额从money1到money5依次递增", "金额从money1到money5依次递增", "金额从money1到money5依次递增", c)
		return
	}
	var etin useproto.AgentClassSetIn
	if in.Money1 == 0 {
		in.Money1 = float32(constex.AgentClassCfg.Agent1Money)
	}
	if in.Money2 == 0 {
		in.Money2 = float32(constex.AgentClassCfg.Agent2Money)
	}
	if in.Money3 == 0 {
		in.Money3 = float32(constex.AgentClassCfg.Agent3Money)
	}
	if in.Money4 == 0 {
		in.Money4 = float32(constex.AgentClassCfg.Agent4Money)
	}
	if in.Money5 == 0 {
		in.Money5 = float32(constex.AgentClassCfg.Agent5Money)
	}
	if in.Tag1 == 0 {
		in.Tag1 = constex.AgentClassCfg.Agent1Tag
	}
	if in.Tag2 == 0 {
		in.Tag2 = constex.AgentClassCfg.Agent2Tag
	}
	if in.Tag3 == 0 {
		in.Tag3 = constex.AgentClassCfg.Agent3Tag
	}
	if in.Tag4 == 0 {
		in.Tag4 = constex.AgentClassCfg.Agent4Tag
	}
	if in.Tag5 == 0 {
		in.Tag5 = constex.AgentClassCfg.Agent5Tag
	}
	if in.Tagex1 == "" {
		in.Tagex1 = constex.AgentClassCfg.Agent1Tagex
	}
	if in.Tagex2 == "" {
		in.Tagex2 = constex.AgentClassCfg.Agent2Tagex
	}
	if in.Tagex3 == "" {
		in.Tagex3 = constex.AgentClassCfg.Agent3Tagex
	}
	if in.Tagex4 == "" {
		in.Tagex4 = constex.AgentClassCfg.Agent4Tagex
	}
	if in.Tagex5 == "" {
		in.Tagex5 = constex.AgentClassCfg.Agent5Tagex
	}
	if in.AgentName1 == "" {
		in.AgentName1 = constex.AgentClassCfg.Agent1Name
	}
	if in.AgentName2 == "" {
		in.AgentName2 = constex.AgentClassCfg.Agent2Name
	}
	if in.AgentName3 == "" {
		in.AgentName3 = constex.AgentClassCfg.Agent3Name
	}
	if in.AgentName4 == "" {
		in.AgentName4 = constex.AgentClassCfg.Agent4Name
	}
	if in.AgentName5 == "" {
		in.AgentName5 = constex.AgentClassCfg.Agent5Name
	}

	etin.Money1 = in.Money1
	etin.Money2 = in.Money2
	etin.Money3 = in.Money3
	etin.Money4 = in.Money4
	etin.Money5 = in.Money5
	etin.Tag1 = int32(in.Tag1)
	etin.Tag2 = int32(in.Tag2)
	etin.Tag3 = int32(in.Tag3)
	etin.Tag4 = int32(in.Tag4)
	etin.Tag5 = int32(in.Tag5)
	etin.Tagex1 = in.Tagex1
	etin.Tagex2 = in.Tagex2
	etin.Tagex3 = in.Tagex3
	etin.Tagex4 = in.Tagex4
	etin.Tagex5 = in.Tagex5
	etin.Agentname1 = in.AgentName1
	etin.Agentname2 = in.AgentName2
	etin.Agentname3 = in.AgentName3
	etin.Agentname4 = in.AgentName4
	etin.Agentname5 = in.AgentName5
	ret, err := client.UserClient.AgentClassSet(c, &etin)
	if err != nil {
		rsp.RespGin(400, 400, err.Error(), err.Error(), "", c)
		return
	}
	rsp.RespGin((*ret).Httpcode, (*ret).Innercode, (*ret).Clientmsg, (*ret).Innermsg, string((*ret).Data), c)
}
func MemerClassSet(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	var in reqmdl.MemClassSet_req
	if err := c.ShouldBindJSON(&in); err != nil {
		rsp.RespGin(400, 400, "输入有误,请重写输入!", "参数有误", "", c)
		return
	}
	isok, _ := param.IsParam(in)
	if !isok {
		rsp.RespGin(400, 400, "输入有误,请重写输入!", "参数有误", "", c)
		return
	}
	if in.Money1 > in.Money2 || in.Money2 > in.Money3 || in.Money3 > in.Money4 || in.Money4 > in.Money5 {
		rsp.RespGin(400, 400, "金额从money1到money5依次递增", "金额从money1到money5依次递增", "金额从money1到money5依次递增", c)
		return
	}
	var etin useproto.MemerClassSetIn
	if in.Money1 == 0 {
		in.Money1 = float32(constex.MemberClassCfg.Member1Money)
	}
	if in.Money2 == 0 {
		in.Money2 = float32(constex.MemberClassCfg.Member2Money)
	}
	if in.Money3 == 0 {
		in.Money3 = float32(constex.MemberClassCfg.Member3Money)
	}
	if in.Money4 == 0 {
		in.Money4 = float32(constex.MemberClassCfg.Member4Money)
	}
	if in.Money5 == 0 {
		in.Money5 = float32(constex.MemberClassCfg.Member5Money)
	}
	if in.Tag1 == 0 {
		in.Tag1 = constex.MemberClassCfg.Mem1Tag
	}
	if in.Tag2 == 0 {
		in.Tag2 = constex.MemberClassCfg.Mem2Tag
	}
	if in.Tag3 == 0 {
		in.Tag3 = constex.MemberClassCfg.Mem3Tag
	}
	if in.Tag4 == 0 {
		in.Tag4 = constex.MemberClassCfg.Mem4Tag
	}
	if in.Tag5 == 0 {
		in.Tag5 = constex.MemberClassCfg.Mem5Tag
	}
	if in.Tagex1 == "" {
		in.Tagex1 = constex.MemberClassCfg.Mem1Tagex
	}
	if in.Tagex2 == "" {
		in.Tagex2 = constex.MemberClassCfg.Mem2Tagex
	}
	if in.Tagex3 == "" {
		in.Tagex3 = constex.MemberClassCfg.Mem3Tagex
	}
	if in.Tagex4 == "" {
		in.Tagex4 = constex.MemberClassCfg.Mem4Tagex
	}
	if in.Tagex5 == "" {
		in.Tagex5 = constex.MemberClassCfg.Mem5Tagex
	}
	if in.MemName1 == "" {
		in.MemName1 = constex.MemberClassCfg.Mem1Name
	}
	if in.MemName2 == "" {
		in.MemName2 = constex.MemberClassCfg.Mem2Name
	}
	if in.MemName3 == "" {
		in.MemName3 = constex.MemberClassCfg.Mem3Name
	}
	if in.MemName4 == "" {
		in.MemName4 = constex.MemberClassCfg.Mem4Name
	}
	if in.MemName5 == "" {
		in.MemName5 = constex.MemberClassCfg.Mem5Name
	}
	etin.Money1 = in.Money1
	etin.Money2 = in.Money2
	etin.Money3 = in.Money3
	etin.Money4 = in.Money4
	etin.Money5 = in.Money5
	etin.Tag1 = int32(in.Tag1)
	etin.Tag2 = int32(in.Tag2)
	etin.Tag3 = int32(in.Tag3)
	etin.Tag4 = int32(in.Tag4)
	etin.Tag5 = int32(in.Tag5)
	etin.Tagex1 = in.Tagex1
	etin.Tagex2 = in.Tagex2
	etin.Tagex3 = in.Tagex3
	etin.Tagex4 = in.Tagex4
	etin.Tagex5 = in.Tagex5
	etin.Memname1 = in.MemName1
	etin.Memname2 = in.MemName2
	etin.Memname3 = in.MemName3
	etin.Memname4 = in.MemName4
	etin.Memname5 = in.MemName5
	ret, err := client.UserClient.MemerClassSet(c, &etin)
	if err != nil {
		rsp.RespGin(400, 400, err.Error(), err.Error(), "", c)
		return
	}
	rsp.RespGin((*ret).Httpcode, (*ret).Innercode, (*ret).Clientmsg, (*ret).Innermsg, string((*ret).Data), c)

}

func UsdtRechargeCallback_xx(c *gin.Context) {
	var ret_resp rspmdl.SuccessOrFail_rsp
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	var rtin useproto.MemberUsdtRechargeIn
	rtin.Uuid = c.Request.Header.Get("X-Head-Uuid")
	Amount, _ := strconv.ParseFloat(c.Request.Header.Get("X-Head-Amount"), 64)
	rtin.Amount = float32(Amount)
	ret, err := client.UserClient.MemberUsdtRecharge(c, &rtin)
	if err != nil {
		rsp.RespGin(400, 400, err.Error(), err.Error(), ret_resp, c)
		return
	}
	json.Unmarshal((*ret).Data, &ret_resp)
	rsp.RespGin((*ret).Httpcode, (*ret).Innercode, (*ret).Clientmsg, (*ret).Innermsg, ret_resp, c)
}

func UsdtRechargeCallback(c *gin.Context) {
	fmt.Println("UsdtRechargeCallback")
	fmt.Println(c.Request)
	var ret_resp rspmdl.RechargeCallback_rsp
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	if err := c.ShouldBindJSON(&ret_resp); err != nil {
		rsp.RespGin(400, 400, err.Error(), err.Error(), ret_resp, c)
		return
	}
	fmt.Println(ret_resp)
	rsp.RespGin(200, 0, "", "", ret_resp, c)
}
