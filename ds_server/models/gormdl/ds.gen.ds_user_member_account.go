package gormdl

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"time"
)

type _DsUserMemberAccountMgr struct {
	*_BaseMgr
}

// DsUserMemberAccountMgr open func
func DsUserMemberAccountMgr(db *gorm.DB) *_DsUserMemberAccountMgr {
	if db == nil {
		panic(fmt.Errorf("DsUserMemberAccountMgr need init by db"))
	}
	return &_DsUserMemberAccountMgr{_BaseMgr: &_BaseMgr{DB: db, isRelated: globalIsRelated}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_DsUserMemberAccountMgr) GetTableName() string {
	return "ds_user_member_account"
}

// Get 获取
func (obj *_DsUserMemberAccountMgr) Get() (result DsUserMemberAccount, err error) {
	err = obj.DB.Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_DsUserMemberAccountMgr) Gets() (results []*DsUserMemberAccount, err error) {
	err = obj.DB.Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithUUID uuid获取 用户ID
func (obj *_DsUserMemberAccountMgr) WithUUID(uuid string) Option {
	return optionFunc(func(o *options) { o.query["uuid"] = uuid })
}

// WithMobile mobile获取 用户手机号
func (obj *_DsUserMemberAccountMgr) WithMobile(mobile string) Option {
	return optionFunc(func(o *options) { o.query["mobile"] = mobile })
}

// WithBalance balance获取 总金额
func (obj *_DsUserMemberAccountMgr) WithBalance(balance float64) Option {
	return optionFunc(func(o *options) { o.query["balance"] = balance })
}

// WithPrivateKey private_key获取
func (obj *_DsUserMemberAccountMgr) WithPrivateKey(privateKey string) Option {
	return optionFunc(func(o *options) { o.query["private_key"] = privateKey })
}

// WithSalt salt获取 支付盐
func (obj *_DsUserMemberAccountMgr) WithSalt(salt string) Option {
	return optionFunc(func(o *options) { o.query["salt"] = salt })
}

// WithHash hash获取 密码hash
func (obj *_DsUserMemberAccountMgr) WithHash(hash string) Option {
	return optionFunc(func(o *options) { o.query["hash"] = hash })
}

// WithAddressIn address_in获取 收款地址
func (obj *_DsUserMemberAccountMgr) WithAddressIn(addressIn string) Option {
	return optionFunc(func(o *options) { o.query["address_in"] = addressIn })
}

// WithAddressOut address_out获取 付款地址
func (obj *_DsUserMemberAccountMgr) WithAddressOut(addressOut string) Option {
	return optionFunc(func(o *options) { o.query["address_out"] = addressOut })
}

// WithStatus status获取 账户状态0:正常 1:禁用  2:注销
func (obj *_DsUserMemberAccountMgr) WithStatus(status int8) Option {
	return optionFunc(func(o *options) { o.query["status"] = status })
}

// WithCreateTime create_time获取 创建时间
func (obj *_DsUserMemberAccountMgr) WithCreateTime(createTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["create_time"] = createTime })
}

// WithUpdateTime update_time获取 更新时间
func (obj *_DsUserMemberAccountMgr) WithUpdateTime(updateTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["update_time"] = updateTime })
}

// WithIspwd ispwd获取 密码是否为空
func (obj *_DsUserMemberAccountMgr) WithIspwd(ispwd uint8) Option {
	return optionFunc(func(o *options) { o.query["ispwd"] = ispwd })
}

// GetByOption 功能选项模式获取
func (obj *_DsUserMemberAccountMgr) GetByOption(opts ...Option) (result DsUserMemberAccount, err error) {
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
func (obj *_DsUserMemberAccountMgr) GetByOptions(opts ...Option) (results []*DsUserMemberAccount, err error) {
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
func (obj *_DsUserMemberAccountMgr) GetFromUUID(uuid string) (result DsUserMemberAccount, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("uuid = ?", UUID).Find(&result).Error

	return
}

// GetBatchFromUUID 批量唯一主键查找 用户ID
func (obj *_DsUserMemberAccountMgr) GetBatchFromUUID(uuids []string) (results []*DsUserMemberAccount, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("uuid IN (?)", uuids).Find(&results).Error

	return
}

// GetFromMobile 通过mobile获取内容 用户手机号
func (obj *_DsUserMemberAccountMgr) GetFromMobile(mobile string) (result DsUserMemberAccount, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("mobile = ?", Mobile).Find(&result).Error

	return
}

// GetBatchFromMobile 批量唯一主键查找 用户手机号
func (obj *_DsUserMemberAccountMgr) GetBatchFromMobile(mobiles []string) (results []*DsUserMemberAccount, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("mobile IN (?)", mobiles).Find(&results).Error

	return
}

// GetFromBalance 通过balance获取内容 总金额
func (obj *_DsUserMemberAccountMgr) GetFromBalance(balance float64) (results []*DsUserMemberAccount, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("balance = ?", balance).Find(&results).Error

	return
}

// GetBatchFromBalance 批量唯一主键查找 总金额
func (obj *_DsUserMemberAccountMgr) GetBatchFromBalance(balances []float64) (results []*DsUserMemberAccount, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("balance IN (?)", balances).Find(&results).Error

	return
}

// GetFromPrivateKey 通过private_key获取内容
func (obj *_DsUserMemberAccountMgr) GetFromPrivateKey(privateKey string) (results []*DsUserMemberAccount, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("private_key = ?", privateKey).Find(&results).Error

	return
}

// GetBatchFromPrivateKey 批量唯一主键查找
func (obj *_DsUserMemberAccountMgr) GetBatchFromPrivateKey(privateKeys []string) (results []*DsUserMemberAccount, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("private_key IN (?)", privateKeys).Find(&results).Error

	return
}

// GetFromSalt 通过salt获取内容 支付盐
func (obj *_DsUserMemberAccountMgr) GetFromSalt(salt string) (results []*DsUserMemberAccount, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("salt = ?", salt).Find(&results).Error

	return
}

// GetBatchFromSalt 批量唯一主键查找 支付盐
func (obj *_DsUserMemberAccountMgr) GetBatchFromSalt(salts []string) (results []*DsUserMemberAccount, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("salt IN (?)", salts).Find(&results).Error

	return
}

// GetFromHash 通过hash获取内容 密码hash
func (obj *_DsUserMemberAccountMgr) GetFromHash(hash string) (results []*DsUserMemberAccount, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("hash = ?", hash).Find(&results).Error

	return
}

// GetBatchFromHash 批量唯一主键查找 密码hash
func (obj *_DsUserMemberAccountMgr) GetBatchFromHash(hashs []string) (results []*DsUserMemberAccount, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("hash IN (?)", hashs).Find(&results).Error

	return
}

// GetFromAddressIn 通过address_in获取内容 收款地址
func (obj *_DsUserMemberAccountMgr) GetFromAddressIn(addressIn string) (results []*DsUserMemberAccount, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("address_in = ?", addressIn).Find(&results).Error

	return
}

// GetBatchFromAddressIn 批量唯一主键查找 收款地址
func (obj *_DsUserMemberAccountMgr) GetBatchFromAddressIn(addressIns []string) (results []*DsUserMemberAccount, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("address_in IN (?)", addressIns).Find(&results).Error

	return
}

// GetFromAddressOut 通过address_out获取内容 付款地址
func (obj *_DsUserMemberAccountMgr) GetFromAddressOut(addressOut string) (results []*DsUserMemberAccount, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("address_out = ?", addressOut).Find(&results).Error

	return
}

// GetBatchFromAddressOut 批量唯一主键查找 付款地址
func (obj *_DsUserMemberAccountMgr) GetBatchFromAddressOut(addressOuts []string) (results []*DsUserMemberAccount, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("address_out IN (?)", addressOuts).Find(&results).Error

	return
}

// GetFromStatus 通过status获取内容 账户状态0:正常 1:禁用  2:注销
func (obj *_DsUserMemberAccountMgr) GetFromStatus(status int8) (results []*DsUserMemberAccount, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("status = ?", status).Find(&results).Error

	return
}

// GetBatchFromStatus 批量唯一主键查找 账户状态0:正常 1:禁用  2:注销
func (obj *_DsUserMemberAccountMgr) GetBatchFromStatus(statuss []int8) (results []*DsUserMemberAccount, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("status IN (?)", statuss).Find(&results).Error

	return
}

// GetFromCreateTime 通过create_time获取内容 创建时间
func (obj *_DsUserMemberAccountMgr) GetFromCreateTime(createTime time.Time) (results []*DsUserMemberAccount, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("create_time = ?", createTime).Find(&results).Error

	return
}

// GetBatchFromCreateTime 批量唯一主键查找 创建时间
func (obj *_DsUserMemberAccountMgr) GetBatchFromCreateTime(createTimes []time.Time) (results []*DsUserMemberAccount, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("create_time IN (?)", createTimes).Find(&results).Error

	return
}

// GetFromUpdateTime 通过update_time获取内容 更新时间
func (obj *_DsUserMemberAccountMgr) GetFromUpdateTime(updateTime time.Time) (results []*DsUserMemberAccount, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("update_time = ?", updateTime).Find(&results).Error

	return
}

// GetBatchFromUpdateTime 批量唯一主键查找 更新时间
func (obj *_DsUserMemberAccountMgr) GetBatchFromUpdateTime(updateTimes []time.Time) (results []*DsUserMemberAccount, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("update_time IN (?)", updateTimes).Find(&results).Error

	return
}

// GetFromIspwd 通过ispwd获取内容 密码是否为空
func (obj *_DsUserMemberAccountMgr) GetFromIspwd(ispwd uint8) (results []*DsUserMemberAccount, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("ispwd = ?", ispwd).Find(&results).Error

	return
}

// GetBatchFromIspwd 批量唯一主键查找 密码是否为空
func (obj *_DsUserMemberAccountMgr) GetBatchFromIspwd(ispwds []uint8) (results []*DsUserMemberAccount, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("ispwd IN (?)", ispwds).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_DsUserMemberAccountMgr) FetchByPrimaryKey(uuid string) (result DsUserMemberAccount, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("uuid = ?", uuid).Find(&result).Error

	return
}

// FetchByUnique primay or index 获取唯一内容
func (obj *_DsUserMemberAccountMgr) FetchByUnique(mobile string) (result DsUserMemberAccount, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("mobile = ?", mobile).Find(&result).Error

	return
}
