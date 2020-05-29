package gormdl

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"time"
)

type _DsUserMemberDepositHistoryMgr struct {
	*_BaseMgr
}

// DsUserMemberDepositHistoryMgr open func
func DsUserMemberDepositHistoryMgr(db *gorm.DB) *_DsUserMemberDepositHistoryMgr {
	if db == nil {
		panic(fmt.Errorf("DsUserMemberDepositHistoryMgr need init by db"))
	}
	return &_DsUserMemberDepositHistoryMgr{_BaseMgr: &_BaseMgr{DB: db, isRelated: globalIsRelated}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_DsUserMemberDepositHistoryMgr) GetTableName() string {
	return "ds_user_member_deposit_history"
}

// Get 获取
func (obj *_DsUserMemberDepositHistoryMgr) Get() (result DsUserMemberDepositHistory, err error) {
	err = obj.DB.Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_DsUserMemberDepositHistoryMgr) Gets() (results []*DsUserMemberDepositHistory, err error) {
	err = obj.DB.Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithUUID uuid获取 用户ID
func (obj *_DsUserMemberDepositHistoryMgr) WithUUID(uuid string) Option {
	return optionFunc(func(o *options) { o.query["uuid"] = uuid })
}

// WithMobile mobile获取 用户手机号
func (obj *_DsUserMemberDepositHistoryMgr) WithMobile(mobile string) Option {
	return optionFunc(func(o *options) { o.query["mobile"] = mobile })
}

// WithSourceID source_id获取 业务订单Id
func (obj *_DsUserMemberDepositHistoryMgr) WithSourceID(sourceID string) Option {
	return optionFunc(func(o *options) { o.query["source_id"] = sourceID })
}

// WithBalance balance获取 金额
func (obj *_DsUserMemberDepositHistoryMgr) WithBalance(balance float64) Option {
	return optionFunc(func(o *options) { o.query["balance"] = balance })
}

// WithRate rate获取 汇率
func (obj *_DsUserMemberDepositHistoryMgr) WithRate(rate float64) Option {
	return optionFunc(func(o *options) { o.query["rate"] = rate })
}

// WithBalanceSrc balance_src获取 原始金额
func (obj *_DsUserMemberDepositHistoryMgr) WithBalanceSrc(balanceSrc float64) Option {
	return optionFunc(func(o *options) { o.query["balance_src"] = balanceSrc })
}

// WithDepositType deposit_type获取 充值类型 0:扣款,1:充值
func (obj *_DsUserMemberDepositHistoryMgr) WithDepositType(depositType uint8) Option {
	return optionFunc(func(o *options) { o.query["deposit_type"] = depositType })
}

// WithDepositName deposit_name获取 充值名字：购买商品，商品退款
func (obj *_DsUserMemberDepositHistoryMgr) WithDepositName(depositName string) Option {
	return optionFunc(func(o *options) { o.query["deposit_name"] = depositName })
}

// WithAddressIn address_in获取 收款地址
func (obj *_DsUserMemberDepositHistoryMgr) WithAddressIn(addressIn string) Option {
	return optionFunc(func(o *options) { o.query["address_in"] = addressIn })
}

// WithAddressOut address_out获取 支付地址
func (obj *_DsUserMemberDepositHistoryMgr) WithAddressOut(addressOut string) Option {
	return optionFunc(func(o *options) { o.query["address_out"] = addressOut })
}

// WithCreateTime create_time获取 创建时间
func (obj *_DsUserMemberDepositHistoryMgr) WithCreateTime(createTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["create_time"] = createTime })
}

// WithUpdateTime update_time获取 更新时间
func (obj *_DsUserMemberDepositHistoryMgr) WithUpdateTime(updateTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["update_time"] = updateTime })
}

// WithStatus status获取 会员账户状态 0:正常，1:禁止 2:销户
func (obj *_DsUserMemberDepositHistoryMgr) WithStatus(status int8) Option {
	return optionFunc(func(o *options) { o.query["status"] = status })
}

// WithInvcodeSelf invcode_self获取 自身邀请码
func (obj *_DsUserMemberDepositHistoryMgr) WithInvcodeSelf(invcodeSelf string) Option {
	return optionFunc(func(o *options) { o.query["invcode_self"] = invcodeSelf })
}

// WithInvcodeAgent invcode_agent获取 代理邀请码
func (obj *_DsUserMemberDepositHistoryMgr) WithInvcodeAgent(invcodeAgent string) Option {
	return optionFunc(func(o *options) { o.query["invcode_agent"] = invcodeAgent })
}

// GetByOption 功能选项模式获取
func (obj *_DsUserMemberDepositHistoryMgr) GetByOption(opts ...Option) (result DsUserMemberDepositHistory, err error) {
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
func (obj *_DsUserMemberDepositHistoryMgr) GetByOptions(opts ...Option) (results []*DsUserMemberDepositHistory, err error) {
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
func (obj *_DsUserMemberDepositHistoryMgr) GetFromUUID(uuid string) (results []*DsUserMemberDepositHistory, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("uuid = ?", uuid).Find(&results).Error

	return
}

// GetBatchFromUUID 批量唯一主键查找 用户ID
func (obj *_DsUserMemberDepositHistoryMgr) GetBatchFromUUID(uuids []string) (results []*DsUserMemberDepositHistory, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("uuid IN (?)", uuids).Find(&results).Error

	return
}

// GetFromMobile 通过mobile获取内容 用户手机号
func (obj *_DsUserMemberDepositHistoryMgr) GetFromMobile(mobile string) (results []*DsUserMemberDepositHistory, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("mobile = ?", mobile).Find(&results).Error

	return
}

// GetBatchFromMobile 批量唯一主键查找 用户手机号
func (obj *_DsUserMemberDepositHistoryMgr) GetBatchFromMobile(mobiles []string) (results []*DsUserMemberDepositHistory, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("mobile IN (?)", mobiles).Find(&results).Error

	return
}

// GetFromSourceID 通过source_id获取内容 业务订单Id
func (obj *_DsUserMemberDepositHistoryMgr) GetFromSourceID(sourceID string) (results []*DsUserMemberDepositHistory, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("source_id = ?", sourceID).Find(&results).Error

	return
}

// GetBatchFromSourceID 批量唯一主键查找 业务订单Id
func (obj *_DsUserMemberDepositHistoryMgr) GetBatchFromSourceID(sourceIDs []string) (results []*DsUserMemberDepositHistory, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("source_id IN (?)", sourceIDs).Find(&results).Error

	return
}

// GetFromBalance 通过balance获取内容 金额
func (obj *_DsUserMemberDepositHistoryMgr) GetFromBalance(balance float64) (results []*DsUserMemberDepositHistory, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("balance = ?", balance).Find(&results).Error

	return
}

// GetBatchFromBalance 批量唯一主键查找 金额
func (obj *_DsUserMemberDepositHistoryMgr) GetBatchFromBalance(balances []float64) (results []*DsUserMemberDepositHistory, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("balance IN (?)", balances).Find(&results).Error

	return
}

// GetFromRate 通过rate获取内容 汇率
func (obj *_DsUserMemberDepositHistoryMgr) GetFromRate(rate float64) (results []*DsUserMemberDepositHistory, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("rate = ?", rate).Find(&results).Error

	return
}

// GetBatchFromRate 批量唯一主键查找 汇率
func (obj *_DsUserMemberDepositHistoryMgr) GetBatchFromRate(rates []float64) (results []*DsUserMemberDepositHistory, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("rate IN (?)", rates).Find(&results).Error

	return
}

// GetFromBalanceSrc 通过balance_src获取内容 原始金额
func (obj *_DsUserMemberDepositHistoryMgr) GetFromBalanceSrc(balanceSrc float64) (results []*DsUserMemberDepositHistory, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("balance_src = ?", balanceSrc).Find(&results).Error

	return
}

// GetBatchFromBalanceSrc 批量唯一主键查找 原始金额
func (obj *_DsUserMemberDepositHistoryMgr) GetBatchFromBalanceSrc(balanceSrcs []float64) (results []*DsUserMemberDepositHistory, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("balance_src IN (?)", balanceSrcs).Find(&results).Error

	return
}

// GetFromDepositType 通过deposit_type获取内容 充值类型 0:扣款,1:充值
func (obj *_DsUserMemberDepositHistoryMgr) GetFromDepositType(depositType uint8) (results []*DsUserMemberDepositHistory, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("deposit_type = ?", depositType).Find(&results).Error

	return
}

// GetBatchFromDepositType 批量唯一主键查找 充值类型 0:扣款,1:充值
func (obj *_DsUserMemberDepositHistoryMgr) GetBatchFromDepositType(depositTypes []uint8) (results []*DsUserMemberDepositHistory, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("deposit_type IN (?)", depositTypes).Find(&results).Error

	return
}

// GetFromDepositName 通过deposit_name获取内容 充值名字：购买商品，商品退款
func (obj *_DsUserMemberDepositHistoryMgr) GetFromDepositName(depositName string) (results []*DsUserMemberDepositHistory, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("deposit_name = ?", depositName).Find(&results).Error

	return
}

// GetBatchFromDepositName 批量唯一主键查找 充值名字：购买商品，商品退款
func (obj *_DsUserMemberDepositHistoryMgr) GetBatchFromDepositName(depositNames []string) (results []*DsUserMemberDepositHistory, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("deposit_name IN (?)", depositNames).Find(&results).Error

	return
}

// GetFromAddressIn 通过address_in获取内容 收款地址
func (obj *_DsUserMemberDepositHistoryMgr) GetFromAddressIn(addressIn string) (results []*DsUserMemberDepositHistory, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("address_in = ?", addressIn).Find(&results).Error

	return
}

// GetBatchFromAddressIn 批量唯一主键查找 收款地址
func (obj *_DsUserMemberDepositHistoryMgr) GetBatchFromAddressIn(addressIns []string) (results []*DsUserMemberDepositHistory, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("address_in IN (?)", addressIns).Find(&results).Error

	return
}

// GetFromAddressOut 通过address_out获取内容 支付地址
func (obj *_DsUserMemberDepositHistoryMgr) GetFromAddressOut(addressOut string) (results []*DsUserMemberDepositHistory, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("address_out = ?", addressOut).Find(&results).Error

	return
}

// GetBatchFromAddressOut 批量唯一主键查找 支付地址
func (obj *_DsUserMemberDepositHistoryMgr) GetBatchFromAddressOut(addressOuts []string) (results []*DsUserMemberDepositHistory, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("address_out IN (?)", addressOuts).Find(&results).Error

	return
}

// GetFromCreateTime 通过create_time获取内容 创建时间
func (obj *_DsUserMemberDepositHistoryMgr) GetFromCreateTime(createTime time.Time) (results []*DsUserMemberDepositHistory, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("create_time = ?", createTime).Find(&results).Error

	return
}

// GetBatchFromCreateTime 批量唯一主键查找 创建时间
func (obj *_DsUserMemberDepositHistoryMgr) GetBatchFromCreateTime(createTimes []time.Time) (results []*DsUserMemberDepositHistory, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("create_time IN (?)", createTimes).Find(&results).Error

	return
}

// GetFromUpdateTime 通过update_time获取内容 更新时间
func (obj *_DsUserMemberDepositHistoryMgr) GetFromUpdateTime(updateTime time.Time) (results []*DsUserMemberDepositHistory, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("update_time = ?", updateTime).Find(&results).Error

	return
}

// GetBatchFromUpdateTime 批量唯一主键查找 更新时间
func (obj *_DsUserMemberDepositHistoryMgr) GetBatchFromUpdateTime(updateTimes []time.Time) (results []*DsUserMemberDepositHistory, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("update_time IN (?)", updateTimes).Find(&results).Error

	return
}

// GetFromStatus 通过status获取内容 会员账户状态 0:正常，1:禁止 2:销户
func (obj *_DsUserMemberDepositHistoryMgr) GetFromStatus(status int8) (results []*DsUserMemberDepositHistory, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("status = ?", status).Find(&results).Error

	return
}

// GetBatchFromStatus 批量唯一主键查找 会员账户状态 0:正常，1:禁止 2:销户
func (obj *_DsUserMemberDepositHistoryMgr) GetBatchFromStatus(statuss []int8) (results []*DsUserMemberDepositHistory, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("status IN (?)", statuss).Find(&results).Error

	return
}

// GetFromInvcodeSelf 通过invcode_self获取内容 自身邀请码
func (obj *_DsUserMemberDepositHistoryMgr) GetFromInvcodeSelf(invcodeSelf string) (results []*DsUserMemberDepositHistory, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("invcode_self = ?", invcodeSelf).Find(&results).Error

	return
}

// GetBatchFromInvcodeSelf 批量唯一主键查找 自身邀请码
func (obj *_DsUserMemberDepositHistoryMgr) GetBatchFromInvcodeSelf(invcodeSelfs []string) (results []*DsUserMemberDepositHistory, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("invcode_self IN (?)", invcodeSelfs).Find(&results).Error

	return
}

// GetFromInvcodeAgent 通过invcode_agent获取内容 代理邀请码
func (obj *_DsUserMemberDepositHistoryMgr) GetFromInvcodeAgent(invcodeAgent string) (results []*DsUserMemberDepositHistory, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("invcode_agent = ?", invcodeAgent).Find(&results).Error

	return
}

// GetBatchFromInvcodeAgent 批量唯一主键查找 代理邀请码
func (obj *_DsUserMemberDepositHistoryMgr) GetBatchFromInvcodeAgent(invcodeAgents []string) (results []*DsUserMemberDepositHistory, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("invcode_agent IN (?)", invcodeAgents).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByIndex  获取多个内容
func (obj *_DsUserMemberDepositHistoryMgr) FetchByIndex(sourceID string, invcodeAgent string) (results []*DsUserMemberDepositHistory, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("source_id = ? AND invcode_agent = ?", sourceID, invcodeAgent).Find(&results).Error

	return
}
