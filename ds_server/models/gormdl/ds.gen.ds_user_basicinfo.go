package gormdl

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"time"
)

type _DsUserBasicinfoMgr struct {
	*_BaseMgr
}

// DsUserBasicinfoMgr open func
func DsUserBasicinfoMgr(db *gorm.DB) *_DsUserBasicinfoMgr {
	if db == nil {
		panic(fmt.Errorf("DsUserBasicinfoMgr need init by db"))
	}
	return &_DsUserBasicinfoMgr{_BaseMgr: &_BaseMgr{DB: db, isRelated: globalIsRelated}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_DsUserBasicinfoMgr) GetTableName() string {
	return "ds_user_basicinfo"
}

// Get 获取
func (obj *_DsUserBasicinfoMgr) Get() (result DsUserBasicinfo, err error) {
	err = obj.DB.Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_DsUserBasicinfoMgr) Gets() (results []*DsUserBasicinfo, err error) {
	err = obj.DB.Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithUUID uuid获取 用户ID号
func (obj *_DsUserBasicinfoMgr) WithUUID(uuid string) Option {
	return optionFunc(func(o *options) { o.query["uuid"] = uuid })
}

// WithMobile mobile获取 手机号
func (obj *_DsUserBasicinfoMgr) WithMobile(mobile string) Option {
	return optionFunc(func(o *options) { o.query["mobile"] = mobile })
}

// WithSalt salt获取 密码盐
func (obj *_DsUserBasicinfoMgr) WithSalt(salt string) Option {
	return optionFunc(func(o *options) { o.query["salt"] = salt })
}

// WithHash hash获取 密码hash
func (obj *_DsUserBasicinfoMgr) WithHash(hash string) Option {
	return optionFunc(func(o *options) { o.query["hash"] = hash })
}

// WithLastLoginTime last_login_time获取 最后一次登录时间
func (obj *_DsUserBasicinfoMgr) WithLastLoginTime(lastLoginTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["last_login_time"] = lastLoginTime })
}

// WithLastLoginIP last_login_ip获取 最后一次登录ip
func (obj *_DsUserBasicinfoMgr) WithLastLoginIP(lastLoginIP string) Option {
	return optionFunc(func(o *options) { o.query["last_login_ip"] = lastLoginIP })
}

// WithCreateTime create_time获取 创建时间
func (obj *_DsUserBasicinfoMgr) WithCreateTime(createTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["create_time"] = createTime })
}

// WithUpdateTime update_time获取 更新时间
func (obj *_DsUserBasicinfoMgr) WithUpdateTime(updateTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["update_time"] = updateTime })
}

// WithDeleted deleted获取 是否删除0:未删除 1:删除
func (obj *_DsUserBasicinfoMgr) WithDeleted(deleted int8) Option {
	return optionFunc(func(o *options) { o.query["deleted"] = deleted })
}

// WithRealName real_name获取 真实姓名
func (obj *_DsUserBasicinfoMgr) WithRealName(realName string) Option {
	return optionFunc(func(o *options) { o.query["real_name"] = realName })
}

// WithAge age获取 年龄
func (obj *_DsUserBasicinfoMgr) WithAge(age int8) Option {
	return optionFunc(func(o *options) { o.query["age"] = age })
}

// WithGender gender获取 性别：0:男，1:女
func (obj *_DsUserBasicinfoMgr) WithGender(gender int8) Option {
	return optionFunc(func(o *options) { o.query["gender"] = gender })
}

// WithAvatar avatar获取 头像地址
func (obj *_DsUserBasicinfoMgr) WithAvatar(avatar string) Option {
	return optionFunc(func(o *options) { o.query["avatar"] = avatar })
}

// WithNickName nick_name获取 昵称
func (obj *_DsUserBasicinfoMgr) WithNickName(nickName string) Option {
	return optionFunc(func(o *options) { o.query["nick_name"] = nickName })
}

// WithBirthday birthday获取 生日
func (obj *_DsUserBasicinfoMgr) WithBirthday(birthday time.Time) Option {
	return optionFunc(func(o *options) { o.query["birthday"] = birthday })
}

// WithStatus status获取 账户状态0:正常 1:禁用  2:注销
func (obj *_DsUserBasicinfoMgr) WithStatus(status int8) Option {
	return optionFunc(func(o *options) { o.query["status"] = status })
}

// GetByOption 功能选项模式获取
func (obj *_DsUserBasicinfoMgr) GetByOption(opts ...Option) (result DsUserBasicinfo, err error) {
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
func (obj *_DsUserBasicinfoMgr) GetByOptions(opts ...Option) (results []*DsUserBasicinfo, err error) {
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

// GetFromUUID 通过uuid获取内容 用户ID号
func (obj *_DsUserBasicinfoMgr) GetFromUUID(uuid string) (result DsUserBasicinfo, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("uuid = ?", UUID).Find(&result).Error

	return
}

// GetBatchFromUUID 批量唯一主键查找 用户ID号
func (obj *_DsUserBasicinfoMgr) GetBatchFromUUID(uuids []string) (results []*DsUserBasicinfo, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("uuid IN (?)", uuids).Find(&results).Error

	return
}

// GetFromMobile 通过mobile获取内容 手机号
func (obj *_DsUserBasicinfoMgr) GetFromMobile(mobile string) (result DsUserBasicinfo, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("mobile = ?", Mobile).Find(&result).Error

	return
}

// GetBatchFromMobile 批量唯一主键查找 手机号
func (obj *_DsUserBasicinfoMgr) GetBatchFromMobile(mobiles []string) (results []*DsUserBasicinfo, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("mobile IN (?)", mobiles).Find(&results).Error

	return
}

// GetFromSalt 通过salt获取内容 密码盐
func (obj *_DsUserBasicinfoMgr) GetFromSalt(salt string) (results []*DsUserBasicinfo, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("salt = ?", salt).Find(&results).Error

	return
}

// GetBatchFromSalt 批量唯一主键查找 密码盐
func (obj *_DsUserBasicinfoMgr) GetBatchFromSalt(salts []string) (results []*DsUserBasicinfo, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("salt IN (?)", salts).Find(&results).Error

	return
}

// GetFromHash 通过hash获取内容 密码hash
func (obj *_DsUserBasicinfoMgr) GetFromHash(hash string) (results []*DsUserBasicinfo, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("hash = ?", hash).Find(&results).Error

	return
}

// GetBatchFromHash 批量唯一主键查找 密码hash
func (obj *_DsUserBasicinfoMgr) GetBatchFromHash(hashs []string) (results []*DsUserBasicinfo, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("hash IN (?)", hashs).Find(&results).Error

	return
}

// GetFromLastLoginTime 通过last_login_time获取内容 最后一次登录时间
func (obj *_DsUserBasicinfoMgr) GetFromLastLoginTime(lastLoginTime time.Time) (results []*DsUserBasicinfo, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("last_login_time = ?", lastLoginTime).Find(&results).Error

	return
}

// GetBatchFromLastLoginTime 批量唯一主键查找 最后一次登录时间
func (obj *_DsUserBasicinfoMgr) GetBatchFromLastLoginTime(lastLoginTimes []time.Time) (results []*DsUserBasicinfo, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("last_login_time IN (?)", lastLoginTimes).Find(&results).Error

	return
}

// GetFromLastLoginIP 通过last_login_ip获取内容 最后一次登录ip
func (obj *_DsUserBasicinfoMgr) GetFromLastLoginIP(lastLoginIP string) (results []*DsUserBasicinfo, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("last_login_ip = ?", lastLoginIP).Find(&results).Error

	return
}

// GetBatchFromLastLoginIP 批量唯一主键查找 最后一次登录ip
func (obj *_DsUserBasicinfoMgr) GetBatchFromLastLoginIP(lastLoginIPs []string) (results []*DsUserBasicinfo, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("last_login_ip IN (?)", lastLoginIPs).Find(&results).Error

	return
}

// GetFromCreateTime 通过create_time获取内容 创建时间
func (obj *_DsUserBasicinfoMgr) GetFromCreateTime(createTime time.Time) (results []*DsUserBasicinfo, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("create_time = ?", createTime).Find(&results).Error

	return
}

// GetBatchFromCreateTime 批量唯一主键查找 创建时间
func (obj *_DsUserBasicinfoMgr) GetBatchFromCreateTime(createTimes []time.Time) (results []*DsUserBasicinfo, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("create_time IN (?)", createTimes).Find(&results).Error

	return
}

// GetFromUpdateTime 通过update_time获取内容 更新时间
func (obj *_DsUserBasicinfoMgr) GetFromUpdateTime(updateTime time.Time) (results []*DsUserBasicinfo, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("update_time = ?", updateTime).Find(&results).Error

	return
}

// GetBatchFromUpdateTime 批量唯一主键查找 更新时间
func (obj *_DsUserBasicinfoMgr) GetBatchFromUpdateTime(updateTimes []time.Time) (results []*DsUserBasicinfo, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("update_time IN (?)", updateTimes).Find(&results).Error

	return
}

// GetFromDeleted 通过deleted获取内容 是否删除0:未删除 1:删除
func (obj *_DsUserBasicinfoMgr) GetFromDeleted(deleted int8) (results []*DsUserBasicinfo, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("deleted = ?", deleted).Find(&results).Error

	return
}

// GetBatchFromDeleted 批量唯一主键查找 是否删除0:未删除 1:删除
func (obj *_DsUserBasicinfoMgr) GetBatchFromDeleted(deleteds []int8) (results []*DsUserBasicinfo, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("deleted IN (?)", deleteds).Find(&results).Error

	return
}

// GetFromRealName 通过real_name获取内容 真实姓名
func (obj *_DsUserBasicinfoMgr) GetFromRealName(realName string) (results []*DsUserBasicinfo, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("real_name = ?", realName).Find(&results).Error

	return
}

// GetBatchFromRealName 批量唯一主键查找 真实姓名
func (obj *_DsUserBasicinfoMgr) GetBatchFromRealName(realNames []string) (results []*DsUserBasicinfo, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("real_name IN (?)", realNames).Find(&results).Error

	return
}

// GetFromAge 通过age获取内容 年龄
func (obj *_DsUserBasicinfoMgr) GetFromAge(age int8) (results []*DsUserBasicinfo, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("age = ?", age).Find(&results).Error

	return
}

// GetBatchFromAge 批量唯一主键查找 年龄
func (obj *_DsUserBasicinfoMgr) GetBatchFromAge(ages []int8) (results []*DsUserBasicinfo, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("age IN (?)", ages).Find(&results).Error

	return
}

// GetFromGender 通过gender获取内容 性别：0:男，1:女
func (obj *_DsUserBasicinfoMgr) GetFromGender(gender int8) (results []*DsUserBasicinfo, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("gender = ?", gender).Find(&results).Error

	return
}

// GetBatchFromGender 批量唯一主键查找 性别：0:男，1:女
func (obj *_DsUserBasicinfoMgr) GetBatchFromGender(genders []int8) (results []*DsUserBasicinfo, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("gender IN (?)", genders).Find(&results).Error

	return
}

// GetFromAvatar 通过avatar获取内容 头像地址
func (obj *_DsUserBasicinfoMgr) GetFromAvatar(avatar string) (results []*DsUserBasicinfo, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("avatar = ?", avatar).Find(&results).Error

	return
}

// GetBatchFromAvatar 批量唯一主键查找 头像地址
func (obj *_DsUserBasicinfoMgr) GetBatchFromAvatar(avatars []string) (results []*DsUserBasicinfo, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("avatar IN (?)", avatars).Find(&results).Error

	return
}

// GetFromNickName 通过nick_name获取内容 昵称
func (obj *_DsUserBasicinfoMgr) GetFromNickName(nickName string) (results []*DsUserBasicinfo, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("nick_name = ?", nickName).Find(&results).Error

	return
}

// GetBatchFromNickName 批量唯一主键查找 昵称
func (obj *_DsUserBasicinfoMgr) GetBatchFromNickName(nickNames []string) (results []*DsUserBasicinfo, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("nick_name IN (?)", nickNames).Find(&results).Error

	return
}

// GetFromBirthday 通过birthday获取内容 生日
func (obj *_DsUserBasicinfoMgr) GetFromBirthday(birthday time.Time) (results []*DsUserBasicinfo, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("birthday = ?", birthday).Find(&results).Error

	return
}

// GetBatchFromBirthday 批量唯一主键查找 生日
func (obj *_DsUserBasicinfoMgr) GetBatchFromBirthday(birthdays []time.Time) (results []*DsUserBasicinfo, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("birthday IN (?)", birthdays).Find(&results).Error

	return
}

// GetFromStatus 通过status获取内容 账户状态0:正常 1:禁用  2:注销
func (obj *_DsUserBasicinfoMgr) GetFromStatus(status int8) (results []*DsUserBasicinfo, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("status = ?", status).Find(&results).Error

	return
}

// GetBatchFromStatus 批量唯一主键查找 账户状态0:正常 1:禁用  2:注销
func (obj *_DsUserBasicinfoMgr) GetBatchFromStatus(statuss []int8) (results []*DsUserBasicinfo, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("status IN (?)", statuss).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_DsUserBasicinfoMgr) FetchByPrimaryKey(uuid string) (result DsUserBasicinfo, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("uuid = ?", uuid).Find(&result).Error

	return
}

// FetchByUnique primay or index 获取唯一内容
func (obj *_DsUserBasicinfoMgr) FetchByUnique(mobile string) (result DsUserBasicinfo, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("mobile = ?", mobile).Find(&result).Error

	return
}
