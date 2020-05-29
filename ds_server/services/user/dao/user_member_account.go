package dao

import (
	mygormdl "ds_server/models/user/gorm_mysql"
	"ds_server/support/lib/redisex"
	"ds_server/support/utils/constex"
	"ds_server/support/utils/sign/md5"
	string_ex "ds_server/support/utils/stringex"
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
	"time"
)

type DsUserMemberAccountMgr struct {
	*_BaseMgr
}

// DsUserMemberAccountMgr open func
func NewDsUserMemberAccountMgr(db *gorm.DB) *DsUserMemberAccountMgr {
	if db == nil {
		panic(fmt.Errorf("DsUserMemberAccountMgr need init by db"))
	}
	return &DsUserMemberAccountMgr{_BaseMgr: &_BaseMgr{DB: db, isRelated: globalIsRelated}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *DsUserMemberAccountMgr) GetTableName() string {
	return "ds_user_member_account"
}

// Get 获取
func (obj *DsUserMemberAccountMgr) Get() (result mygormdl.DsUserMemberAccount, err error) {
	err = obj.DB.Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *DsUserMemberAccountMgr) Gets() (results []*mygormdl.DsUserMemberAccount, err error) {
	err = obj.DB.Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithUUID uuid获取 用户ID
func (obj *DsUserMemberAccountMgr) WithUUID(uuid string) Option {
	return optionFunc(func(o *options) { o.query["uuid"] = uuid })
}

// WithMobile mobile获取 用户手机号
func (obj *DsUserMemberAccountMgr) WithMobile(mobile string) Option {
	return optionFunc(func(o *options) { o.query["mobile"] = mobile })
}

// WithBalance balance获取 总金额
func (obj *DsUserMemberAccountMgr) WithBalance(balance float64) Option {
	return optionFunc(func(o *options) { o.query["balance"] = balance })
}

// WithPrivateKey private_key获取
func (obj *DsUserMemberAccountMgr) WithPrivateKey(privateKey string) Option {
	return optionFunc(func(o *options) { o.query["private_key"] = privateKey })
}

// WithPayPwd pay_pwd获取 支付密码
func (obj *DsUserMemberAccountMgr) WithPayPwd(payPwd string) Option {
	return optionFunc(func(o *options) { o.query["pay_pwd"] = payPwd })
}

// WithSalt salt获取 支付盐
func (obj *DsUserMemberAccountMgr) WithSalt(salt string) Option {
	return optionFunc(func(o *options) { o.query["salt"] = salt })
}

// WithAddressIn address_in获取 收款地址
func (obj *DsUserMemberAccountMgr) WithAddressIn(addressIn string) Option {
	return optionFunc(func(o *options) { o.query["address_in"] = addressIn })
}

// WithAddressOut address_out获取 付款地址
func (obj *DsUserMemberAccountMgr) WithAddressOut(addressOut string) Option {
	return optionFunc(func(o *options) { o.query["address_out"] = addressOut })
}

// WithStatus status获取 账户状态0:正常 1:禁用  2:注销
func (obj *DsUserMemberAccountMgr) WithStatus(status int8) Option {
	return optionFunc(func(o *options) { o.query["status"] = status })
}

// WithCreateTime create_time获取 创建时间
func (obj *DsUserMemberAccountMgr) WithCreateTime(createTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["create_time"] = createTime })
}

// WithUpdateTime update_time获取 更新时间
func (obj *DsUserMemberAccountMgr) WithUpdateTime(updateTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["update_time"] = updateTime })
}

// GetByOption 功能选项模式获取
func (obj *DsUserMemberAccountMgr) GetByOption(opts ...Option) (result mygormdl.DsUserMemberAccount, err error) {
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
func (obj *DsUserMemberAccountMgr) GetByOptions(opts ...Option) (results []*mygormdl.DsUserMemberAccount, err error) {
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
func (obj *DsUserMemberAccountMgr) GetFromUUID(uuid string) (result mygormdl.DsUserMemberAccount, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("uuid = ?", uuid).Find(&result).Error

	return
}

// GetBatchFromUUID 批量唯一主键查找 用户ID
func (obj *DsUserMemberAccountMgr) GetBatchFromUUID(uuids []string) (results []*mygormdl.DsUserMemberAccount, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("uuid IN (?)", uuids).Find(&results).Error

	return
}

// GetBatchFromMobile 批量唯一主键查找 用户手机号
func (obj *DsUserMemberAccountMgr) GetBatchFromMobile(mobiles []string) (results []*mygormdl.DsUserMemberAccount, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("mobile IN (?)", mobiles).Find(&results).Error

	return
}

// GetFromBalance 通过balance获取内容 总金额
func (obj *DsUserMemberAccountMgr) GetFromBalance(balance float64) (results []*mygormdl.DsUserMemberAccount, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("balance = ?", balance).Find(&results).Error

	return
}

// GetBatchFromBalance 批量唯一主键查找 总金额
func (obj *DsUserMemberAccountMgr) GetBatchFromBalance(balances []float64) (results []*mygormdl.DsUserMemberAccount, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("balance IN (?)", balances).Find(&results).Error

	return
}

// GetFromPrivateKey 通过private_key获取内容
func (obj *DsUserMemberAccountMgr) GetFromPrivateKey(privateKey string) (results []*mygormdl.DsUserMemberAccount, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("private_key = ?", privateKey).Find(&results).Error

	return
}

// GetBatchFromPrivateKey 批量唯一主键查找
func (obj *DsUserMemberAccountMgr) GetBatchFromPrivateKey(privateKeys []string) (results []*mygormdl.DsUserMemberAccount, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("private_key IN (?)", privateKeys).Find(&results).Error

	return
}

// GetFromPayPwd 通过pay_pwd获取内容 支付密码
func (obj *DsUserMemberAccountMgr) GetFromPayPwd(payPwd string) (results []*mygormdl.DsUserMemberAccount, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("pay_pwd = ?", payPwd).Find(&results).Error

	return
}

// GetBatchFromPayPwd 批量唯一主键查找 支付密码
func (obj *DsUserMemberAccountMgr) GetBatchFromPayPwd(payPwds []string) (results []*mygormdl.DsUserMemberAccount, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("pay_pwd IN (?)", payPwds).Find(&results).Error

	return
}

// GetFromSalt 通过salt获取内容 支付盐
func (obj *DsUserMemberAccountMgr) GetFromSalt(salt string) (results []*mygormdl.DsUserMemberAccount, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("salt = ?", salt).Find(&results).Error

	return
}

// GetBatchFromSalt 批量唯一主键查找 支付盐
func (obj *DsUserMemberAccountMgr) GetBatchFromSalt(salts []string) (results []*mygormdl.DsUserMemberAccount, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("salt IN (?)", salts).Find(&results).Error

	return
}

// GetFromAddressIn 通过address_in获取内容 收款地址
func (obj *DsUserMemberAccountMgr) GetFromAddressIn(addressIn string) (results []*mygormdl.DsUserMemberAccount, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("address_in = ?", addressIn).Find(&results).Error

	return
}

// GetBatchFromAddressIn 批量唯一主键查找 收款地址
func (obj *DsUserMemberAccountMgr) GetBatchFromAddressIn(addressIns []string) (results []*mygormdl.DsUserMemberAccount, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("address_in IN (?)", addressIns).Find(&results).Error

	return
}

// GetFromAddressOut 通过address_out获取内容 付款地址
func (obj *DsUserMemberAccountMgr) GetFromAddressOut(addressOut string) (results []*mygormdl.DsUserMemberAccount, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("address_out = ?", addressOut).Find(&results).Error

	return
}

// GetBatchFromAddressOut 批量唯一主键查找 付款地址
func (obj *DsUserMemberAccountMgr) GetBatchFromAddressOut(addressOuts []string) (results []*mygormdl.DsUserMemberAccount, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("address_out IN (?)", addressOuts).Find(&results).Error

	return
}

// GetFromStatus 通过status获取内容 账户状态0:正常 1:禁用  2:注销
func (obj *DsUserMemberAccountMgr) GetFromStatus(status int8) (results []*mygormdl.DsUserMemberAccount, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("status = ?", status).Find(&results).Error

	return
}

// GetBatchFromStatus 批量唯一主键查找 账户状态0:正常 1:禁用  2:注销
func (obj *DsUserMemberAccountMgr) GetBatchFromStatus(statuss []int8) (results []*mygormdl.DsUserMemberAccount, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("status IN (?)", statuss).Find(&results).Error

	return
}

// GetFromCreateTime 通过create_time获取内容 创建时间
func (obj *DsUserMemberAccountMgr) GetFromCreateTime(createTime time.Time) (results []*mygormdl.DsUserMemberAccount, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("create_time = ?", createTime).Find(&results).Error

	return
}

// GetBatchFromCreateTime 批量唯一主键查找 创建时间
func (obj *DsUserMemberAccountMgr) GetBatchFromCreateTime(createTimes []time.Time) (results []*mygormdl.DsUserMemberAccount, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("create_time IN (?)", createTimes).Find(&results).Error

	return
}

// GetFromUpdateTime 通过update_time获取内容 更新时间
func (obj *DsUserMemberAccountMgr) GetFromUpdateTime(updateTime time.Time) (results []*mygormdl.DsUserMemberAccount, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("update_time = ?", updateTime).Find(&results).Error

	return
}

// GetBatchFromUpdateTime 批量唯一主键查找 更新时间
func (obj *DsUserMemberAccountMgr) GetBatchFromUpdateTime(updateTimes []time.Time) (results []*mygormdl.DsUserMemberAccount, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("update_time IN (?)", updateTimes).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *DsUserMemberAccountMgr) FetchByPrimaryKey(uuid string) (result mygormdl.DsUserMemberAccount, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("uuid = ?", uuid).Find(&result).Error

	return
}

// FetchByUnique primay or index 获取唯一内容
func (obj *DsUserMemberAccountMgr) FetchByUnique(mobile string) (result mygormdl.DsUserMemberAccount, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("mobile = ?", mobile).Find(&result).Error

	return
}

/////////////////////////primary index case ////////////////////////////////////////////
func (obj *DsUserMemberAccountMgr) GetMemberAccountByMobileTy(tx *gorm.DB, mobile string, ty int32) (results mygormdl.DsUserMemberAccount, err error) {
	err = tx.Table(obj.GetTableName()).Where("mobile = ? and coin_type = ?", mobile, ty).Find(&results).Error
	return
}

func (obj *DsUserMemberAccountMgr) GetMemberAccount(tx *gorm.DB, invcode_agent, invcode_self string) (results []*mygormdl.DsUserMemberAccount, err error) {
	err = tx.Table(obj.GetTableName()).Where("invcode_agent = ? and invcode_self <> ?", invcode_agent, invcode_self).Find(&results).Error
	return
}

func (obj *DsUserMemberAccountMgr) UpdateSelfMemberAccount(tx *gorm.DB, uuid string, money float64) (err error) {
	in := mygormdl.DsUserMemberAccount{
		Balance: money,
	}
	err = tx.Table(obj.GetTableName()).Where("uuid= ? ", uuid).Update(&in).Error
	return
}

func (obj *DsUserMemberAccountMgr) UpdateMemberAccount(tx *gorm.DB, uuid string, all_mem_money float64) (err error) { //retwa mygormdl.DsUserMemberAccount,
	err = tx.Table(obj.GetTableName()).Where("uuid=?", uuid).Updates(
		mygormdl.DsUserMemberAccount{
			Balance:    all_mem_money,
			UpdateTime: time.Now(),
		}).Error
	//tx.Table(obj.GetTableName()).Where("uuid=?", uuid).Find(&retwa)
	return
}
func (obj *DsUserMemberAccountMgr) GetMemberAccountOrigin(tx *gorm.DB, invcode_self string) (results mygormdl.DsUserMemberAccount, err error) {
	err = tx.Table(obj.GetTableName()).Where("invcode_self = ?", invcode_self).Find(&results).Error
	return
}
func (obj *DsUserMemberAccountMgr) GetFromMobile(mobile string) (results mygormdl.DsUserMemberAccount, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("mobile = ?", mobile).Find(&results).Error
	return
}

func (obj *DsUserMemberAccountMgr) GetSelfMemberAccount(tx *gorm.DB, uuid string) (results mygormdl.DsUserMemberAccount, err error) {
	err = tx.Table(obj.GetTableName()).Where("uuid = ? ", uuid).Find(&results).Error
	return
}
func (obj *DsUserMemberAccountMgr) InsertMemberAccount(tx *gorm.DB, uuid, mobile string, all_mem_money float64) (wa mygormdl.DsUserMemberAccount, err error) {
	wa = mygormdl.DsUserMemberAccount{
		UUID:       uuid,
		Mobile:     mobile,
		Balance:    0,
		PrivateKey: string_ex.Rand6NumString(),
		AddressIn:  string_ex.Rand6NumString(),
		AddressOut: string_ex.Rand6NumString(),
		Status:     0,
		CreateTime: time.Now(),
		UpdateTime: time.Now(),
	}
	err = tx.Create(&wa).Error
	return
}

func (obj *DsUserMemberAccountMgr) SetPaypwd(uuid, paypwd,  vfcode,mobile string) (err error) {
	rs_vfcode,_ := redisex.RedisInstanceg().Get(constex.REDIS_USER_VFCODE+mobile).Result()
	if rs_vfcode != vfcode{
		return errors.New("短信验证码无效")
	}
	salt := string_ex.GetRandomString(16)
	hash := md5.HashForPwd(salt, paypwd)
	return obj.Table(obj.GetTableName()).Where("uuid= ?",uuid).Updates(mygormdl.DsUserMemberAccount{UpdateTime: time.Now(),Hash: hash,Salt: salt,Ispwd: 1}).Error
}
func (obj *DsUserMemberAccountMgr) ModifyPayPwd(uuid, newpwd string) (err error) {
	tx := obj.DB.Begin()
	var usermemacct mygormdl.DsUserMemberAccount
	db := tx.Table(obj.GetTableName()).Where("uuid= ?", uuid).Find(&usermemacct)
	if db.RecordNotFound() {
		tx.Rollback()
		return errors.New("不存在该手机号！")
	}
	db = tx.Table(obj.GetTableName()).Where("uuid= ?", uuid).Updates(mygormdl.DsUserBasicinfo{Hash: md5.HashForPwd(usermemacct.Salt, newpwd)})
	if db.RowsAffected == 0 {
		tx.Rollback()
		return errors.New("密码相同,更新密码失败！")
	}
	tx.Commit()
	return nil
}
