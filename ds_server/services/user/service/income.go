package service

import (
	"context"
	pb "ds_server/proto/user"
	_ "ds_server/support/utils/httpex"
	"fmt"
)

//直推和间推的比率
type KJRate struct {
	K float32
	J float32
}

func (usersrv *UserService) MemberUsdtRecharge(ctx context.Context, req *pb.MemberUsdtRechargeIn, rsq *pb.CommonOut) error {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	uuid := req.Uuid
	amount := req.Amount

	//0、需要先记录充值前的会员状态
	userMemberAgentBefore, err := usersrv.DsUserMemberAgentDao.FetchByPrimaryKey(uuid)
	if err != nil {
		return nil
	}

	//1、根据充值金额更新会员卡信息
	_ = usersrv.updateMember(uuid, amount);

	//2、根据充值金额计算上级收益(需要先记录充值钱的会员状态)

	_ = usersrv.countProfit(uuid, amount, userMemberAgentBefore.MemclassSelf);

	//3、根据充值金额更新上级代理等级
	//usersrv.updateAgent(uuid, amount);

	fmt.Println(200, 200, "处理成功!", "处理成功!", []byte{}, rsq)
	return nil
}

//根据充值金额更新会员卡信息
func (usersrv *UserService) updateMember(uuid string, amount float32) error {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	//先获取用户
	result, err := usersrv.DsUserMemberAccountDao.GetFromUUID(uuid)
	//用户不存在什么都不做
	if err != nil {
		return nil
	}
	//更新余额
	all_mem_money := result.Balance + float64(amount)
	tx := usersrv.DsUserMemberAccountDao.DB
	_ = usersrv.DsUserMemberAccountDao.UpdateMemberAccount(tx, uuid, all_mem_money)
	//
	userMemberAgent, err := usersrv.DsUserMemberAgentDao.FetchByPrimaryKey(uuid)
	//用户不存在什么都不做
	if err != nil {
		return nil
	}
	//历史记录日志
	req := &pb.MemberDepositIn{}
	req.Uuid = uuid
	req.Mobile = result.Mobile
	req.InvcodeSelf = userMemberAgent.InvcodeSelf
	req.InvcodeAgent = userMemberAgent.InvcodeAgent
	req.DepositType = 1
	req.DepositName = "USDT"
	req.DepositNum = 0
	req.Addressin = result.AddressIn
	req.Addressout = result.AddressOut
	userMemberDepositHistoryTx := usersrv.DsUserMemberDepositHistoryDao.DB
	rate_money := float64(amount)
	src_money := float64(amount)
	rate := 1.0
	_ = usersrv.DsUserMemberDepositHistoryDao.CreateMemberDepositHistory(userMemberDepositHistoryTx, req, rate_money, src_money, rate)
	//获取累计充值金额
	commonTotal, _ := usersrv.DsUserMemberDepositHistoryDao.GetAllCharge(uuid)
	fmt.Println("累计充值:", commonTotal.Total)
	var memTag int8
	if commonTotal.Total<1000.0 {
		return nil
	} else if commonTotal.Total>=1000.0 && commonTotal.Total<5000.0 {
		memTag = 1
	} else if commonTotal.Total>=5000.0 && commonTotal.Total<10000.0 {
		memTag = 2
	} else if commonTotal.Total>=10000.0 && commonTotal.Total<30000.0 {
		memTag = 3
	} else if commonTotal.Total>=30000.0 && commonTotal.Total<50000.0 {
		memTag = 4
	} else if commonTotal.Total>=50000.0 {
		memTag = 5
	}
	userMemberClass, err := usersrv.DsUserMemberClassDao.GetOneFromMemTag(memTag)
	//vip等级信息未找到不处理
	if err != nil {
		return nil
	}
	//更新自己的会员信息
	tx = usersrv.DsUserMemberAgentDao.DB
	_ = usersrv.DsUserMemberAgentDao.SetUserMemClass(tx, userMemberClass.MemTagex, userMemberClass.MemTag, userMemberClass.MemName, uuid)
	return nil
}

//数据库余额和收益日志的处理
func (usersrv *UserService) doBalanceAndIncomeLog(uuid string, amount float32) error {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	//先获取用户
	result, err := usersrv.DsUserMemberAccountDao.GetFromUUID(uuid)
	//用户不存在什么都不做
	if err != nil {
		return nil
	}
	//更新余额
	all_mem_money := result.Balance + float64(amount)
	tx := usersrv.DsUserMemberAccountDao.DB
	_ = usersrv.DsUserMemberAccountDao.UpdateMemberAccount(tx, uuid, all_mem_money)
	//
	userMemberAgent, err := usersrv.DsUserMemberAgentDao.FetchByPrimaryKey(uuid)
	//用户不存在什么都不做
	if err != nil {
		return nil
	}
	//历史记录日志
	req := &pb.MemberDepositIn{}
	req.Uuid = uuid
	req.Mobile = result.Mobile
	req.InvcodeSelf = userMemberAgent.InvcodeSelf
	req.InvcodeAgent = userMemberAgent.InvcodeAgent
	req.DepositType = 1
	req.DepositName = "USDT"
	req.DepositNum = 0
	req.Addressin = result.AddressIn
	req.Addressout = result.AddressOut
	userMemberDepositHistoryTx := usersrv.DsUserMemberDepositHistoryDao.DB
	rate_money := float64(amount)
	src_money := float64(amount)
	rate := 1.0
	_ = usersrv.DsUserMemberDepositHistoryDao.CreateMemberDepositHistory(userMemberDepositHistoryTx, req, rate_money, src_money, rate)

	return nil
}

//根据充值金额计算上级收益
func (usersrv *UserService) countProfit(uuid string, amount float32, userMemberAgentBeforeGrade int8) error {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	if userMemberAgentBeforeGrade < 1 {
		return nil
	}
	//查找父级
	userMemberAgentParent, err := usersrv.DsUserMemberAgentDao.FetchByPrimaryKey(uuid)
	if err != nil {
		return nil
	}
	_ = usersrv.sendProfit(userMemberAgentParent.UUIDAgent, amount, "k", userMemberAgentBeforeGrade, userMemberAgentParent.MemclassSelf)
	//查找祖级
	userMemberAgentParent2, err := usersrv.DsUserMemberAgentDao.FetchByPrimaryKey(userMemberAgentParent.UUIDAgent)
	if err != nil {
		return nil
	}
	_ = usersrv.sendProfit(userMemberAgentParent2.UUIDAgent, amount, "j", userMemberAgentBeforeGrade, userMemberAgentParent2.MemclassSelf)

	return nil
}

/*
	实际发放上级的收益
	uuid 用户id
	amount 用户充值金额
	flag 直推还是间推(k j)
	userMemberAgentBeforeGrade 充值用户充值前的等级
	memClassSelf 获得收益的用户的会员等级
*/
func (usersrv *UserService) sendProfit(uuid string, amount float32, flag string, userMemberAgentBeforeGrade, memClassSelf int8) error {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	//不满足V1
	if memClassSelf < 1 {
		return nil
	}
	m := make(map[int8]KJRate)
	m[1] = KJRate{K:0.05, J:0.02}
	m[2] = KJRate{K:0.10, J:0.04}
	m[3] = KJRate{K:0.15, J:0.06}
	m[4] = KJRate{K:0.20, J:0.08}
	m[5] = KJRate{K:0.25, J:0.10}
	if item, ok := m[userMemberAgentBeforeGrade]; ok {
		rate := float32(0)
		if flag == "k" {
			rate = item.K
		} else if flag == "j" {
			rate = item.J
		}
		if rate > 0 {
			_ = usersrv.doBalanceAndIncomeLog(uuid, amount * rate)
		}
	}

	return nil
}