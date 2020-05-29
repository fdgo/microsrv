package dao

import (
	mygormdl "ds_server/models/user/gorm_mysql"
	pb "ds_server/proto/user"
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
	"strconv"
	"time"
)

type DsUserMemberDepositHistoryMgr struct {
	*_BaseMgr
}

// DsUserMemberDepositHistoryMgr open func
func NewDsUserMemberDepositHistoryMgr(db *gorm.DB) *DsUserMemberDepositHistoryMgr {
	if db == nil {
		panic(fmt.Errorf("DsUserMemberDepositHistoryMgr need init by db"))
	}
	return &DsUserMemberDepositHistoryMgr{_BaseMgr: &_BaseMgr{DB: db, isRelated: globalIsRelated}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *DsUserMemberDepositHistoryMgr) GetTableName() string {
	return "ds_user_member_deposit_history"
}

// Get 获取
func (obj *DsUserMemberDepositHistoryMgr) Get() (result mygormdl.DsUserMemberDepositHistory, err error) {
	err = obj.DB.Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *DsUserMemberDepositHistoryMgr) Gets() (results []*mygormdl.DsUserMemberDepositHistory, err error) {
	err = obj.DB.Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithUUID uuid获取 用户ID
func (obj *DsUserMemberDepositHistoryMgr) WithUUID(uuid string) Option {
	return optionFunc(func(o *options) { o.query["uuid"] = uuid })
}

// WithMobile mobile获取 用户手机号
func (obj *DsUserMemberDepositHistoryMgr) WithMobile(mobile string) Option {
	return optionFunc(func(o *options) { o.query["mobile"] = mobile })
}

// WithSourceID source_id获取 业务订单Id
func (obj *DsUserMemberDepositHistoryMgr) WithSourceID(sourceID string) Option {
	return optionFunc(func(o *options) { o.query["source_id"] = sourceID })
}

// WithBalance balance获取 金额
func (obj *DsUserMemberDepositHistoryMgr) WithBalance(balance float64) Option {
	return optionFunc(func(o *options) { o.query["balance"] = balance })
}

// WithDepositType deposit_type获取 充值类型 0:扣款,1:充值
func (obj *DsUserMemberDepositHistoryMgr) WithDepositType(depositType int8) Option {
	return optionFunc(func(o *options) { o.query["deposit_type"] = depositType })
}

// WithDepositName deposit_name获取 充值名字：购买商品，商品退款
func (obj *DsUserMemberDepositHistoryMgr) WithDepositName(depositName string) Option {
	return optionFunc(func(o *options) { o.query["deposit_name"] = depositName })
}

// WithAddressIn address_in获取 收款地址
func (obj *DsUserMemberDepositHistoryMgr) WithAddressIn(addressIn string) Option {
	return optionFunc(func(o *options) { o.query["address_in"] = addressIn })
}

// WithAddressOut address_out获取 支付地址
func (obj *DsUserMemberDepositHistoryMgr) WithAddressOut(addressOut string) Option {
	return optionFunc(func(o *options) { o.query["address_out"] = addressOut })
}

// WithCreateTime create_time获取 创建时间
func (obj *DsUserMemberDepositHistoryMgr) WithCreateTime(createTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["create_time"] = createTime })
}

// WithUpdateTime update_time获取 更新时间
func (obj *DsUserMemberDepositHistoryMgr) WithUpdateTime(updateTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["update_time"] = updateTime })
}

// WithStatus status获取 会员账户状态 0:正常，1:禁止 2:销户
func (obj *DsUserMemberDepositHistoryMgr) WithStatus(status int8) Option {
	return optionFunc(func(o *options) { o.query["status"] = status })
}

// GetByOption 功能选项模式获取
func (obj *DsUserMemberDepositHistoryMgr) GetByOption(opts ...Option) (result mygormdl.DsUserMemberDepositHistory, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.Table(obj.GetTableName()).Where(options.query).Find(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *DsUserMemberDepositHistoryMgr) GetByOptions(opts ...Option) (results []*mygormdl.DsUserMemberDepositHistory, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.Table(obj.GetTableName()).Where(options.query).Find(&results).Error

	return
}

//////////////////////////enume case ////////////////////////////////////////////

// GetFromUUID 通过uuid获取内容 用户ID
func (obj *DsUserMemberDepositHistoryMgr) GetFromUUID(uuid string) (result mygormdl.DsUserMemberDepositHistory, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("uuid = ?", uuid).Find(&result).Error

	return
}

// GetBatchFromUUID 批量唯一主键查找 用户ID
func (obj *DsUserMemberDepositHistoryMgr) GetBatchFromUUID(uuids []string) (results []*mygormdl.DsUserMemberDepositHistory, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("uuid IN (?)", uuids).Find(&results).Error

	return
}

// GetFromMobile 通过mobile获取内容 用户手机号
func (obj *DsUserMemberDepositHistoryMgr) GetFromMobile(mobile string) (results []*mygormdl.DsUserMemberDepositHistory, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("mobile = ?", mobile).Find(&results).Error

	return
}

// GetBatchFromMobile 批量唯一主键查找 用户手机号
func (obj *DsUserMemberDepositHistoryMgr) GetBatchFromMobile(mobiles []string) (results []*mygormdl.DsUserMemberDepositHistory, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("mobile IN (?)", mobiles).Find(&results).Error

	return
}

// GetFromSourceID 通过source_id获取内容 业务订单Id
func (obj *DsUserMemberDepositHistoryMgr) GetFromSourceID(sourceID string) (results []*mygormdl.DsUserMemberDepositHistory, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("source_id = ?", sourceID).Find(&results).Error

	return
}

// GetBatchFromSourceID 批量唯一主键查找 业务订单Id
func (obj *DsUserMemberDepositHistoryMgr) GetBatchFromSourceID(sourceIDs []string) (results []*mygormdl.DsUserMemberDepositHistory, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("source_id IN (?)", sourceIDs).Find(&results).Error

	return
}

// GetFromBalance 通过balance获取内容 金额
func (obj *DsUserMemberDepositHistoryMgr) GetFromBalance(balance float64) (results []*mygormdl.DsUserMemberDepositHistory, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("balance = ?", balance).Find(&results).Error

	return
}

// GetBatchFromBalance 批量唯一主键查找 金额
func (obj *DsUserMemberDepositHistoryMgr) GetBatchFromBalance(balances []float64) (results []*mygormdl.DsUserMemberDepositHistory, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("balance IN (?)", balances).Find(&results).Error

	return
}

// GetFromDepositType 通过deposit_type获取内容 充值类型 0:扣款,1:充值
func (obj *DsUserMemberDepositHistoryMgr) GetFromDepositType(depositType int8) (results []*mygormdl.DsUserMemberDepositHistory, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("deposit_type = ?", depositType).Find(&results).Error

	return
}

// GetBatchFromDepositType 批量唯一主键查找 充值类型 0:扣款,1:充值
func (obj *DsUserMemberDepositHistoryMgr) GetBatchFromDepositType(depositTypes []int8) (results []*mygormdl.DsUserMemberDepositHistory, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("deposit_type IN (?)", depositTypes).Find(&results).Error

	return
}

// GetFromDepositName 通过deposit_name获取内容 充值名字：购买商品，商品退款
func (obj *DsUserMemberDepositHistoryMgr) GetFromDepositName(depositName string) (results []*mygormdl.DsUserMemberDepositHistory, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("deposit_name = ?", depositName).Find(&results).Error

	return
}

// GetBatchFromDepositName 批量唯一主键查找 充值名字：购买商品，商品退款
func (obj *DsUserMemberDepositHistoryMgr) GetBatchFromDepositName(depositNames []string) (results []*mygormdl.DsUserMemberDepositHistory, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("deposit_name IN (?)", depositNames).Find(&results).Error

	return
}

// GetFromAddressIn 通过address_in获取内容 收款地址
func (obj *DsUserMemberDepositHistoryMgr) GetFromAddressIn(addressIn string) (results []*mygormdl.DsUserMemberDepositHistory, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("address_in = ?", addressIn).Find(&results).Error

	return
}

// GetBatchFromAddressIn 批量唯一主键查找 收款地址
func (obj *DsUserMemberDepositHistoryMgr) GetBatchFromAddressIn(addressIns []string) (results []*mygormdl.DsUserMemberDepositHistory, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("address_in IN (?)", addressIns).Find(&results).Error

	return
}

// GetFromAddressOut 通过address_out获取内容 支付地址
func (obj *DsUserMemberDepositHistoryMgr) GetFromAddressOut(addressOut string) (results []*mygormdl.DsUserMemberDepositHistory, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("address_out = ?", addressOut).Find(&results).Error

	return
}

// GetBatchFromAddressOut 批量唯一主键查找 支付地址
func (obj *DsUserMemberDepositHistoryMgr) GetBatchFromAddressOut(addressOuts []string) (results []*mygormdl.DsUserMemberDepositHistory, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("address_out IN (?)", addressOuts).Find(&results).Error

	return
}

// GetFromCreateTime 通过create_time获取内容 创建时间
func (obj *DsUserMemberDepositHistoryMgr) GetFromCreateTime(createTime time.Time) (results []*mygormdl.DsUserMemberDepositHistory, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("create_time = ?", createTime).Find(&results).Error

	return
}

// GetBatchFromCreateTime 批量唯一主键查找 创建时间
func (obj *DsUserMemberDepositHistoryMgr) GetBatchFromCreateTime(createTimes []time.Time) (results []*mygormdl.DsUserMemberDepositHistory, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("create_time IN (?)", createTimes).Find(&results).Error

	return
}

// GetFromUpdateTime 通过update_time获取内容 更新时间
func (obj *DsUserMemberDepositHistoryMgr) GetFromUpdateTime(updateTime time.Time) (results []*mygormdl.DsUserMemberDepositHistory, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("update_time = ?", updateTime).Find(&results).Error

	return
}

// GetBatchFromUpdateTime 批量唯一主键查找 更新时间
func (obj *DsUserMemberDepositHistoryMgr) GetBatchFromUpdateTime(updateTimes []time.Time) (results []*mygormdl.DsUserMemberDepositHistory, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("update_time IN (?)", updateTimes).Find(&results).Error

	return
}

// GetFromStatus 通过status获取内容 会员账户状态 0:正常，1:禁止 2:销户
func (obj *DsUserMemberDepositHistoryMgr) GetFromStatus(status int8) (results []*mygormdl.DsUserMemberDepositHistory, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("status = ?", status).Find(&results).Error

	return
}

// GetBatchFromStatus 批量唯一主键查找 会员账户状态 0:正常，1:禁止 2:销户
func (obj *DsUserMemberDepositHistoryMgr) GetBatchFromStatus(statuss []int8) (results []*mygormdl.DsUserMemberDepositHistory, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("status IN (?)", statuss).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *DsUserMemberDepositHistoryMgr) FetchByPrimaryKey(uuid string) (result mygormdl.DsUserMemberDepositHistory, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("uuid = ?", uuid).Find(&result).Error

	return
}

// FetchByIndex  获取多个内容
func (obj *DsUserMemberDepositHistoryMgr) FetchByIndex(sourceID string) (results []*mygormdl.DsUserMemberDepositHistory, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("source_id = ?", sourceID).Find(&results).Error

	return
}

//change
func (obj *DsUserMemberDepositHistoryMgr) CreateMemberDepositHistory(tx *gorm.DB, req *pb.MemberDepositIn, rate_money, src_money, rate float64) (err error) {
	dephistry := mygormdl.DsUserMemberDepositHistory{
		UUID:         req.Uuid,
		Mobile:       req.Mobile,
		SourceID:     req.DepositName,
		Balance:      rate_money,
		BalanceSrc:   src_money,
		Rate:         rate,
		DepositType:  int8(req.DepositType),
		DepositName:  req.DepositName,
		Status:       0,
		AddressIn:    req.Addressin,
		AddressOut:   req.Addressout,
		InvcodeAgent: req.InvcodeAgent,
		InvcodeSelf:  req.InvcodeSelf,
		CreateTime:   time.Now(),
		UpdateTime:   time.Now(),
	}
	return tx.Create(&dephistry).Error
}

//change
func (obj *DsUserMemberDepositHistoryMgr) GetAllMemDeposiMonForAgent(tx *gorm.DB, invcodeagent string) (results []*mygormdl.DsUserMemberDepositHistory, err error) {
	err = tx.Table(obj.GetTableName()).Where("invcode_agent=? and balance > ?", invcodeagent, 0).Find(&results).Error
	return
}
//change
func (obj *DsUserMemberDepositHistoryMgr) CreateMemberWithDrawHistory(tx *gorm.DB, paymoney float64, req *pb.OnlinePayIn,ret mygormdl.DsUserMemberAccount) (dephistry mygormdl.DsUserMemberDepositHistory, err error) {
	leftmoney, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", paymoney-2*paymoney), 64)
	dephistry = mygormdl.DsUserMemberDepositHistory{
		UUID:         req.Uuid,
		Mobile:       ret.Mobile,
		SourceID:     req.SrcId,
		Balance:      leftmoney,
		Rate:         1,
		BalanceSrc:   0,
		DepositType:  3,
		DepositName:  "USDT",
		Status:       1,
		AddressIn:    req.Addressin,
		AddressOut:   req.Addressout,
		CreateTime:   time.Now(),
		UpdateTime:   time.Now(),
		InvcodeAgent: req.InvcodeAgent,
		InvcodeSelf:  req.InvcodeSelf,
	}
	err = tx.Create(&dephistry).Error
	return
}
//change
func (obj *DsUserMemberDepositHistoryMgr) QueryMemberDepositHistory(tx *gorm.DB, uuid string) (results []*mygormdl.DsUserMemberDepositHistory, err error) {
	err = tx.Table(obj.GetTableName()).Where("uuid=? and balance > ? ", uuid, 0).Find(&results).Error
	return
}
//change
func (obj *DsUserMemberDepositHistoryMgr) QueryMemberAllDepositHistory(tx *gorm.DB, uuid string) (results []*mygormdl.DsUserMemberDepositHistory, err error) {
	err = tx.Table(obj.GetTableName()).Where("uuid=?", uuid).Find(&results).Error
	return
}
//change
func (obj *DsUserMemberDepositHistoryMgr) QueryAgentDepositHistory(tx *gorm.DB, uuid string) (results []*mygormdl.DsUserMemberDepositHistory, err error) {
	err = tx.Table(obj.GetTableName()).Where("uuid=? and balance > ? ", uuid, 0).Find(&results).Error
	return
}
//change
func (obj *DsUserMemberDepositHistoryMgr) QueryMemberDepositHistoryLog(tx *gorm.DB,ty, uuid string, pagesize int32, index int32) (num int, results []*mygormdl.DsUserMemberDepositHistory, err error) {
	var tmp []*mygormdl.DsUserMemberDepositHistory
	if ty == "USDT" {
		err = tx.Table(obj.GetTableName()).Where("uuid=? and deposit_name=?", uuid, ty).Find(&tmp).Error
		if err!=nil{
			tx.Rollback()
			return 0, nil, err
		}
		err = tx.Table(obj.GetTableName()).Where("uuid=? and deposit_name=?", uuid, ty).Limit(pagesize).Offset(pagesize * index).Order("create_time desc").Find(&results).Error
		if err!=nil{
			tx.Rollback()
			return 0, nil, err
		}
		return len(tmp),results,err
	} else if ty == "ali_pay" {
		err = tx.Table(obj.GetTableName()).Where("uuid=? and deposit_name=?", uuid, ty).Find(&tmp).Error
		if err!=nil{
			tx.Rollback()
			return 0, nil, err
		}
		err = tx.Table(obj.GetTableName()).Where("uuid=? and deposit_name=?", uuid, ty).Limit(pagesize).Offset(pagesize * index).Order("create_time desc").Find(&results).Error
		if err!=nil{
			tx.Rollback()
			return 0, nil, err
		}
		return len(tmp),results,err
	} else if ty == "wechat_pay" {
		err = tx.Table(obj.GetTableName()).Where("uuid=? and deposit_name=?", uuid, ty).Find(&tmp).Error
		if err!=nil{
			tx.Rollback()
			return 0, nil, err
		}
		err = tx.Table(obj.GetTableName()).Where("uuid=? and deposit_name=?", uuid, ty).Limit(pagesize).Offset(pagesize * index).Order("create_time desc").Find(&results).Error
		if err!=nil{
			tx.Rollback()
			return 0, nil, err
		}
		return len(tmp),results,err
	} else if ty == "all" {
		err = tx.Table(obj.GetTableName()).Where("uuid=? ", uuid).Find(&tmp).Error
		if err!=nil{
			tx.Rollback()
			return 0, nil, err
		}
		err = tx.Table(obj.GetTableName()).Where("uuid=? ", uuid).Limit(pagesize).Offset(pagesize * index).Order("create_time desc").Find(&results).Error
		if err!=nil{
			tx.Rollback()
			return 0, nil, err
		}
		return len(tmp),results,err
	}
	return 0, nil, errors.New("充值或者支付类型错误！")
}

// GetAllCharge 获取某用户的累计充值金额
func (obj *DsUserMemberDepositHistoryMgr) GetAllCharge(uuid string) (result mygormdl.CommonTotal, err error) {
	err = obj.DB.Select("ifnull(sum(balance),0) as total").Table(obj.GetTableName()).Where("uuid=? && deposit_type=1", uuid).First(&result).Error

	return
}
