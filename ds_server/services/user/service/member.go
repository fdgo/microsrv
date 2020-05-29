package service

import (
	"context"
	rspmdl "ds_server/models/user/gin_rsp"
	mygormdl "ds_server/models/user/gorm_mysql"
	pb "ds_server/proto/user"
	"ds_server/support/lib/redisex"
	"ds_server/support/utils/constex"
	"ds_server/support/utils/httpex"
	_ "ds_server/support/utils/httpex"
	"ds_server/support/utils/rsp"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/shopspring/decimal"
	"strconv"
)

func (usersrv *UserService) ExchangeRate(ctx context.Context, req *pb.ExchangeRateIn, rsq *pb.CommonOut) error {
	StatusCode, body, err := httpex.Get(constex.ExchangeRateCfg.Address)
	if err != nil {
		rsp.RespSrv(StatusCode, StatusCode, "货币实时汇率获取失败!", "货币实时汇率获取失败!", []byte(""), rsq)
		return err
	} else {
		if StatusCode == 200 {
			rsp.RespSrv(StatusCode, StatusCode, "货币实时汇率获取成功!", "货币实时汇率获取成功!", []byte(body), rsq)
			return err
		} else {
			rsp.RespSrv(StatusCode, StatusCode, "货币实时汇率获取失败!", "货币实时汇率获取失败!", []byte(""), rsq)
			return err
		}
	}
}
func (usersrv *UserService) MemberDepositLog(ctx context.Context, req *pb.MemberDepositLogIn, rsq *pb.CommonOut) error {
	var tmp_resp rspmdl.MemberDepositLog_rsp
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	tx := usersrv.DsUserMemberDepositHistoryDao.Begin()
	num, tmpret, _ := usersrv.DsUserMemberDepositHistoryDao.QueryMemberDepositHistoryLog(tx,req.TypeName, req.Uuid, req.PageSize, req.Index)
	for _, v := range tmpret {
		tmp_resp.Memdephistory = append(tmp_resp.Memdephistory, v)
	}
	tx.Commit()
	tmp_resp.Num = num
	tmp_resp_nor, _ := json.Marshal(tmp_resp)
	rsp.RespSrv(200, 200, "获取充值记录成功!", "获取充值记录成功!", tmp_resp_nor, rsq)
	return nil
}

func (usersrv *UserService) MemberDeposit(ctx context.Context, req *pb.MemberDepositIn, rsq *pb.CommonOut) error {
	var tmp_resp rspmdl.MemAgent_rsp
	tmp_resp_err, _ := json.Marshal(tmp_resp)
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	StatusCode, body, err := httpex.Get(constex.ExchangeRateCfg.Address)
	if err != nil {
		rsp.RespSrv(StatusCode, StatusCode, "货币实时汇率获取失败!", err.Error(), tmp_resp_err, rsq)
		return err
	} else {
		if StatusCode != 200 {
			rsp.RespSrv(StatusCode, StatusCode, "货币实时汇率获取失败!", err.Error(), tmp_resp_err, rsq)
			return err
		}
	}
	var fBuyPri, mBuyPri float64
	var rate rspmdl.DsUserWalletExchangeRate_rsp
	err = json.Unmarshal([]byte(body), &rate)
	if err != nil {
		rsp.RespSrv(StatusCode, StatusCode, "货币实时汇率获取失败!", err.Error(), tmp_resp_err, rsq)
		return err
	}
	for _, v := range rate.Result {
		fmt.Println(v.Data1)
		if v.Data1.Name == "美元" {
			FBuyPri, _ := strconv.ParseFloat(v.Data1.FBuyPri, 32)
			fBuyPri = FBuyPri
			MBuyPri, _ := strconv.ParseFloat(v.Data1.MBuyPri, 32)
			mBuyPri = MBuyPri
			break
		}
	}
	var rate_money float64
	if req.DepositName == "ali_pay" || req.DepositName == "wechat_pay" {
		rate_money = float64(req.DepositNum) / ((fBuyPri) / 100)
	} else {
		rate_money = float64(req.DepositNum)
	}
	tx := usersrv.DsUserMemberDepositHistoryDao.Begin()
	//创建充值记录
	err = usersrv.DsUserMemberDepositHistoryDao.CreateMemberDepositHistory(tx, req, rate_money, float64(req.DepositNum), (fBuyPri+mBuyPri)/200)
	if err != nil {
		tx.Rollback()
		rsp.RespSrv(400, 400, "充值失败!", err.Error(), tmp_resp_err, rsq)
		return err
	}
	histyslice, errz := usersrv.DsUserMemberDepositHistoryDao.QueryMemberAllDepositHistory(tx, req.Uuid)
	if errz != nil {
		tx.Rollback()
		rsp.RespSrv(400, 400, "充值失败!", errz.Error(), tmp_resp_err, rsq)
		return errz
	}
	var all_mem_money float64
	var posi_money float64
	for _, v := range histyslice {
		all_mem_money += v.Balance
		if v.Balance > 0 {
			posi_money += v.Balance
		}
	}
	fmt.Println(req.Uuid,"total acct,",all_mem_money)
	//总账户
	err = usersrv.DsUserMemberAccountDao.UpdateMemberAccount(tx, req.Uuid, all_mem_money)
	if err != nil {
		tx.Rollback()
		rsp.RespSrv(400, 400, "充值失败!", err.Error(), tmp_resp_err, rsq)
		return err
	}
	var exx error
	var member_tagex, member_class, member_name string
	member_class, member_tagex, member_name, exx = usersrv.DsUserMemberClassDao.MemberClassGetRedis(redisex.RedisInstanceg(), posi_money)
	mem_classex, _ := strconv.Atoi(member_class)
	var tmpmem_class int8
	if exx!=nil {
		member_tagex, tmpmem_class, member_name ,exx = usersrv.DsUserMemberClassDao.MemberClassGetMysql(tx, posi_money)
		if exx!=nil {
			tx.Rollback()
			rsp.RespSrv(400, 400, "会员等级配置获取失败!", "会员等级配置获取失败!", tmp_resp_err, rsq)
			return errors.New("会员等级配置获取失败！")
		} else {
			//设置会员等级
			err = usersrv.DsUserMemberAgentDao.SetUserMemClass(tx, member_tagex, tmpmem_class, member_name, req.Uuid)
			if err != nil {
				tx.Rollback()
				rsp.RespSrv(400, 400, "设置会员等级失败!", err.Error(), tmp_resp_err, rsq)
				return err
			}
		}
	} else {
		//设置会员等级
		err = usersrv.DsUserMemberAgentDao.SetUserMemClass(tx, member_tagex, int8(mem_classex), member_name, req.Uuid)
		if err != nil {
			tx.Rollback()
			rsp.RespSrv(400, 400, "设置会员等级失败!", err.Error(), tmp_resp_err, rsq)
			return err
		}
	}
	//统计充值用户代理商下线的所有充值费用
	var allmoney_agent float64
	allmem, errey := usersrv.DsUserMemberDepositHistoryDao.GetAllMemDeposiMonForAgent(tx, req.InvcodeAgent)
	if errey != nil {
		tx.Rollback()
		rsp.RespSrv(400, 400, "充值失败!", errey.Error(), tmp_resp_err, rsq)
		return errey
	}
	for _, v := range allmem {
		allmoney_agent += v.Balance
	}
	var exxy error
	var agent_tag, agent_class, agent_name string
	agent_class, agent_tag, agent_name, exxy = usersrv.DsUserAgentClassDao.AgentClassGetRedis(redisex.RedisInstanceg(), allmoney_agent)
	agentclassex, _ := strconv.Atoi(agent_class)
	var tmpagent_class int8
	if exxy != nil {
		agent_tag, tmpagent_class, agent_name,exxy = usersrv.DsUserAgentClassDao.AgentClassGetMysql(tx, allmoney_agent)
		if exxy!=nil {
			tx.Rollback()
			rsp.RespSrv(400, 400, "代理等级配置获取失败!", "代理等级配置获取失败!", tmp_resp_err, rsq)
			return errors.New("代理等级配置获取失败！")
		} else {
			//设置会员等级
			exxy = usersrv.DsUserMemberAgentDao.SetAgentClass(tx, tmpagent_class, agent_tag, agent_name, req.InvcodeAgent)
			if exxy != nil {
				tx.Rollback()
				rsp.RespSrv(400, 400, "设置代理等级失败!", exxy.Error(), tmp_resp_err, rsq)
				return exxy
			}
		}
	} else {
		//设置会员等级
		exxy = usersrv.DsUserMemberAgentDao.SetAgentClass(tx, int8(agentclassex), agent_tag, agent_name, req.InvcodeAgent)
		if exxy != nil {
			tx.Rollback()
			rsp.RespSrv(400, 400, "设置代理等级失败!", exxy.Error(), tmp_resp_err, rsq)
			return exxy
		}
	}
	fmt.Println(exxy)
	tx.Commit()
	retMoney, _ := decimal.NewFromFloat(all_mem_money).Round(6).Float64()
	tmp_resp.Balance = retMoney
	tmp_resp_nor, _ := json.Marshal(tmp_resp)
	rsp.RespSrv(200, 200, "充值成功!", "充值成功!", tmp_resp_nor, rsq)
	return nil
}

func (usersrv *UserService) OnlinePay(ctx context.Context, req *pb.OnlinePayIn, rsq *pb.CommonOut) error {
	var tmp_resp mygormdl.DsUserMemberDepositHistory
	tmp_resp_err, _ := json.Marshal(tmp_resp)
	tx := usersrv.DsUserMemberDepositHistoryDao.Begin()
	ret, err := usersrv.DsUserMemberAccountDao.GetSelfMemberAccount(tx, req.Uuid)
	if err != nil {
		tx.Rollback()
		rsp.RespSrv(400, 400, err.Error(), err.Error(), tmp_resp_err, rsq)
		return nil
	}
	if ret.Balance > float64(req.DepositNum) {
		TmpMoney, _ := decimal.NewFromFloat(ret.Balance).Sub(decimal.NewFromFloat(float64(req.DepositNum))).Float64()
		retMoney, _ := decimal.NewFromFloat(TmpMoney).Round(6).Float64()
		err = usersrv.DsUserMemberAccountDao.UpdateSelfMemberAccount(tx, req.Uuid, retMoney)
		if err != nil {
			tx.Rollback()
			rsp.RespSrv(400, 400, "在线支付失败!", err.Error(), tmp_resp_err, rsq)
			return err
		}
		tmpret, err := usersrv.DsUserMemberDepositHistoryDao.CreateMemberWithDrawHistory(tx, float64(req.DepositNum), req, ret)
		if err != nil {
			tx.Rollback()
			rsp.RespSrv(400, 400, "支付失败!", err.Error(), tmp_resp_err, rsq)
			return nil
		}
		tmp_resp_nor, _ := json.Marshal(tmpret)
		rsp.RespSrv(200, 200, "在线支付成功!", "在线支付成功!", tmp_resp_nor, rsq)
		tx.Commit()
		return nil
	} else {
		tx.Rollback()
		rsp.RespSrv(400, 400, "在线支付失败!,请充值!", "在线支付失败!,请充值!", tmp_resp_err, rsq)
		return err
	}
}
func (usersrv *UserService) AgentClassSet(ctx context.Context, req *pb.AgentClassSetIn, rsq *pb.CommonOut) error {
	Agent1money := map[string]interface{}{"Agent1money": req.Money1, "Agent1Tag": req.Tag1, "Agent1Tagex": req.Tagex1, "Agent1Name": req.Agentname1}
	Agent2money := map[string]interface{}{"Agent2money": req.Money2, "Agent2Tag": req.Tag2, "Agent2Tagex": req.Tagex2, "Agent2Name": req.Agentname2}
	Agent3money := map[string]interface{}{"Agent3money": req.Money3, "Agent3Tag": req.Tag3, "Agent3Tagex": req.Tagex3, "Agent3Name": req.Agentname3}
	Agent4money := map[string]interface{}{"Agent4money": req.Money4, "Agent4Tag": req.Tag4, "Agent4Tagex": req.Tagex4, "Agent4Name": req.Agentname4}
	Agent5money := map[string]interface{}{"Agent5money": req.Money5, "Agent5Tag": req.Tag5, "Agent5Tagex": req.Tagex5, "Agent5Name": req.Agentname5}
	var tmpUACslice []*mygormdl.DsUserAgentClass
	tmpUACslice = append(tmpUACslice, &mygormdl.DsUserAgentClass{AgentMoney: float64(req.Money1), AgentTag: int8(req.Tag1), AgentTagex: req.Tagex1, AgentName: req.Agentname1})
	tmpUACslice = append(tmpUACslice, &mygormdl.DsUserAgentClass{AgentMoney: float64(req.Money2), AgentTag: int8(req.Tag2), AgentTagex: req.Tagex2, AgentName: req.Agentname2})
	tmpUACslice = append(tmpUACslice, &mygormdl.DsUserAgentClass{AgentMoney: float64(req.Money3), AgentTag: int8(req.Tag3), AgentTagex: req.Tagex3, AgentName: req.Agentname3})
	tmpUACslice = append(tmpUACslice, &mygormdl.DsUserAgentClass{AgentMoney: float64(req.Money4), AgentTag: int8(req.Tag4), AgentTagex: req.Tagex4, AgentName: req.Agentname4})
	tmpUACslice = append(tmpUACslice, &mygormdl.DsUserAgentClass{AgentMoney: float64(req.Money5), AgentTag: int8(req.Tag5), AgentTagex: req.Tagex5, AgentName: req.Agentname5})
	_, err := usersrv.DsUserAgentClassDao.AgentClassSetMysql(tmpUACslice)
	if err != nil {
		rsp.RespSrv(400, 400, "设置合伙人等级失败!", "设置合伙人等级失败!", []byte("false"), rsq)
		return err
	}
	_, err = usersrv.DsUserAgentClassDao.AgentClassSetRedis(Agent1money, Agent2money, Agent3money, Agent4money, Agent5money)
	if err != nil {
		rsp.RespSrv(400, 400, "设置合伙人等级失败!", "设置合伙人等级失败!", []byte("false"), rsq)
		return err
	}
	rsp.RespSrv(200, 200, "设置合伙人等级成功!", "设置合伙人等级成功!", []byte("true"), rsq)
	return nil
}
func (usersrv *UserService) MemerClassSet(ctx context.Context, req *pb.MemerClassSetIn, rsq *pb.CommonOut) error {
	Member1money := map[string]interface{}{"Member1money": req.Money1, "Mem1Tag": req.Tag1, "Mem1Tagex": req.Tagex1, "Mem1Name": req.Memname1}
	Member2money := map[string]interface{}{"Member2money": req.Money2, "Mem2Tag": req.Tag2, "Mem2Tagex": req.Tagex2, "Mem2Name": req.Memname2}
	Member3money := map[string]interface{}{"Member3money": req.Money3, "Mem3Tag": req.Tag3, "Mem3Tagex": req.Tagex3, "Mem3Name": req.Memname3}
	Member4money := map[string]interface{}{"Member4money": req.Money4, "Mem4Tag": req.Tag4, "Mem4Tagex": req.Tagex4, "Mem4Name": req.Memname4}
	Member5money := map[string]interface{}{"Member5money": req.Money5, "Mem5Tag": req.Tag5, "Mem5Tagex": req.Tagex5, "Mem5Name": req.Memname5}
	var tmpUMCslice []*mygormdl.DsUserMemberClass
	tmpUMCslice = append(tmpUMCslice, &mygormdl.DsUserMemberClass{MemMoney: float64(req.Money1), MemTag: int8(req.Tag1), MemTagex: req.Tagex1, MemName: req.Memname1})
	tmpUMCslice = append(tmpUMCslice, &mygormdl.DsUserMemberClass{MemMoney: float64(req.Money2), MemTag: int8(req.Tag2), MemTagex: req.Tagex2, MemName: req.Memname2})
	tmpUMCslice = append(tmpUMCslice, &mygormdl.DsUserMemberClass{MemMoney: float64(req.Money3), MemTag: int8(req.Tag3), MemTagex: req.Tagex3, MemName: req.Memname3})
	tmpUMCslice = append(tmpUMCslice, &mygormdl.DsUserMemberClass{MemMoney: float64(req.Money4), MemTag: int8(req.Tag4), MemTagex: req.Tagex4, MemName: req.Memname4})
	tmpUMCslice = append(tmpUMCslice, &mygormdl.DsUserMemberClass{MemMoney: float64(req.Money5), MemTag: int8(req.Tag5), MemTagex: req.Tagex5, MemName: req.Memname5})
	_, err := usersrv.DsUserMemberClassDao.MemerClassSetMysql(tmpUMCslice)
	if err != nil {
		rsp.RespSrv(400, 400, "设置会员等级失败!", "设置会员等级失败!", []byte("false"), rsq)
		return err
	}
	_, err = usersrv.DsUserMemberClassDao.MemerClassSetRedis(Member1money, Member2money, Member3money, Member4money, Member5money)
	if err != nil {
		rsp.RespSrv(400, 400, "设置会员等级失败!", "设置会员等级失败!", []byte("false"), rsq)
		return err
	}
	rsp.RespSrv(200, 200, "设置会员等级成功!", "设置会员等级成功!", []byte("true"), rsq)
	return nil
}
