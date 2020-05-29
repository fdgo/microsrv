package dao

import (
	mygormdl "ds_server/models/user/gorm_mysql"
	"ds_server/support/lib/redisex"
	"ds_server/support/utils/constex"
	"ds_server/support/utils/sign/md5"
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
	"time"
)

type DsUserBasicinfoMgr struct {
	*_BaseMgr
}

// DsUserBasicinfoMgr open func
func NewDsUserBasicinfoMgr(db *gorm.DB) *DsUserBasicinfoMgr {
	if db == nil {
		panic(fmt.Errorf("DsUserBasicinfoMgr need init by db"))
	}
	return &DsUserBasicinfoMgr{_BaseMgr: &_BaseMgr{DB: db, isRelated: globalIsRelated}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *DsUserBasicinfoMgr) GetTableName() string {
	return "ds_user_basicinfo"
}

// Get 获取
func (obj *DsUserBasicinfoMgr) Get() (result mygormdl.DsUserBasicinfo, err error) {
	err = obj.DB.Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *DsUserBasicinfoMgr) Gets() (results []*mygormdl.DsUserBasicinfo, err error) {
	err = obj.DB.Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithUUID uuid获取 用户ID号
func (obj *DsUserBasicinfoMgr) WithUUID(uuid string) Option {
	return optionFunc(func(o *options) { o.query["uuid"] = uuid })
}

// WithMobile mobile获取 手机号
func (obj *DsUserBasicinfoMgr) WithMobile(mobile string) Option {
	return optionFunc(func(o *options) { o.query["mobile"] = mobile })
}

// WithLoginPwd login_pwd获取 登录密码
func (obj *DsUserBasicinfoMgr) WithLoginPwd(loginPwd string) Option {
	return optionFunc(func(o *options) { o.query["login_pwd"] = loginPwd })
}

// WithSalt salt获取 密码盐
func (obj *DsUserBasicinfoMgr) WithSalt(salt string) Option {
	return optionFunc(func(o *options) { o.query["salt"] = salt })
}

// WithLastLoginTime last_login_time获取 最后一次登录时间
func (obj *DsUserBasicinfoMgr) WithLastLoginTime(lastLoginTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["last_login_time"] = lastLoginTime })
}

// WithLastLoginIP last_login_ip获取 最后一次登录ip
func (obj *DsUserBasicinfoMgr) WithLastLoginIP(lastLoginIP string) Option {
	return optionFunc(func(o *options) { o.query["last_login_ip"] = lastLoginIP })
}

// WithCreateTime create_time获取 创建时间
func (obj *DsUserBasicinfoMgr) WithCreateTime(createTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["create_time"] = createTime })
}

// WithUpdateTime update_time获取 更新时间
func (obj *DsUserBasicinfoMgr) WithUpdateTime(updateTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["update_time"] = updateTime })
}

// WithDeleted deleted获取 是否删除0:未删除 1:删除
func (obj *DsUserBasicinfoMgr) WithDeleted(deleted int8) Option {
	return optionFunc(func(o *options) { o.query["deleted"] = deleted })
}

// WithRealName real_name获取 真实姓名
func (obj *DsUserBasicinfoMgr) WithRealName(realName string) Option {
	return optionFunc(func(o *options) { o.query["real_name"] = realName })
}

// WithAge age获取 年龄
func (obj *DsUserBasicinfoMgr) WithAge(age int8) Option {
	return optionFunc(func(o *options) { o.query["age"] = age })
}

// WithGender gender获取 性别：0:男，1:女
func (obj *DsUserBasicinfoMgr) WithGender(gender int8) Option {
	return optionFunc(func(o *options) { o.query["gender"] = gender })
}

// WithAvatar avatar获取 头像地址
func (obj *DsUserBasicinfoMgr) WithAvatar(avatar string) Option {
	return optionFunc(func(o *options) { o.query["avatar"] = avatar })
}

// WithNickName nick_name获取 昵称
func (obj *DsUserBasicinfoMgr) WithNickName(nickName string) Option {
	return optionFunc(func(o *options) { o.query["nick_name"] = nickName })
}

// WithBirthday birthday获取 生日
func (obj *DsUserBasicinfoMgr) WithBirthday(birthday time.Time) Option {
	return optionFunc(func(o *options) { o.query["birthday"] = birthday })
}

// WithStatus status获取 账户状态0:正常 1:禁用  2:注销
func (obj *DsUserBasicinfoMgr) WithStatus(status int8) Option {
	return optionFunc(func(o *options) { o.query["status"] = status })
}

// GetByOption 功能选项模式获取
func (obj *DsUserBasicinfoMgr) GetByOption(opts ...Option) (result mygormdl.DsUserBasicinfo, err error) {
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
func (obj *DsUserBasicinfoMgr) GetByOptions(opts ...Option) (results []*mygormdl.DsUserBasicinfo, err error) {
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
func (obj *DsUserBasicinfoMgr) GetFromUUID(uuid string) (result mygormdl.DsUserBasicinfo, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("uuid = ?", uuid).Find(&result).Error

	return
}

// GetBatchFromUUID 批量唯一主键查找 用户ID号
func (obj *DsUserBasicinfoMgr) GetBatchFromUUID(uuids []string) (results []*mygormdl.DsUserBasicinfo, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("uuid IN (?)", uuids).Find(&results).Error

	return
}

// GetFromMobile 通过mobile获取内容 手机号
func (obj *DsUserBasicinfoMgr) GetFromMobile(mobile string) (results []*mygormdl.DsUserBasicinfo, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("mobile = ?", mobile).Find(&results).Error

	return
}

// GetBatchFromMobile 批量唯一主键查找 手机号
func (obj *DsUserBasicinfoMgr) GetBatchFromMobile(mobiles []string) (results []*mygormdl.DsUserBasicinfo, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("mobile IN (?)", mobiles).Find(&results).Error

	return
}

// GetFromLoginPwd 通过login_pwd获取内容 登录密码
func (obj *DsUserBasicinfoMgr) GetFromLoginPwd(loginPwd string) (results []*mygormdl.DsUserBasicinfo, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("login_pwd = ?", loginPwd).Find(&results).Error

	return
}

// GetBatchFromLoginPwd 批量唯一主键查找 登录密码
func (obj *DsUserBasicinfoMgr) GetBatchFromLoginPwd(loginPwds []string) (results []*mygormdl.DsUserBasicinfo, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("login_pwd IN (?)", loginPwds).Find(&results).Error

	return
}

// GetFromSalt 通过salt获取内容 密码盐
func (obj *DsUserBasicinfoMgr) GetFromSalt(salt string) (results []*mygormdl.DsUserBasicinfo, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("salt = ?", salt).Find(&results).Error

	return
}

// GetBatchFromSalt 批量唯一主键查找 密码盐
func (obj *DsUserBasicinfoMgr) GetBatchFromSalt(salts []string) (results []*mygormdl.DsUserBasicinfo, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("salt IN (?)", salts).Find(&results).Error

	return
}

// GetFromLastLoginTime 通过last_login_time获取内容 最后一次登录时间
func (obj *DsUserBasicinfoMgr) GetFromLastLoginTime(lastLoginTime time.Time) (results []*mygormdl.DsUserBasicinfo, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("last_login_time = ?", lastLoginTime).Find(&results).Error

	return
}

// GetBatchFromLastLoginTime 批量唯一主键查找 最后一次登录时间
func (obj *DsUserBasicinfoMgr) GetBatchFromLastLoginTime(lastLoginTimes []time.Time) (results []*mygormdl.DsUserBasicinfo, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("last_login_time IN (?)", lastLoginTimes).Find(&results).Error

	return
}

// GetFromLastLoginIP 通过last_login_ip获取内容 最后一次登录ip
func (obj *DsUserBasicinfoMgr) GetFromLastLoginIP(lastLoginIP string) (results []*mygormdl.DsUserBasicinfo, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("last_login_ip = ?", lastLoginIP).Find(&results).Error

	return
}

// GetBatchFromLastLoginIP 批量唯一主键查找 最后一次登录ip
func (obj *DsUserBasicinfoMgr) GetBatchFromLastLoginIP(lastLoginIPs []string) (results []*mygormdl.DsUserBasicinfo, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("last_login_ip IN (?)", lastLoginIPs).Find(&results).Error

	return
}

// GetFromCreateTime 通过create_time获取内容 创建时间
func (obj *DsUserBasicinfoMgr) GetFromCreateTime(createTime time.Time) (results []*mygormdl.DsUserBasicinfo, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("create_time = ?", createTime).Find(&results).Error

	return
}

// GetBatchFromCreateTime 批量唯一主键查找 创建时间
func (obj *DsUserBasicinfoMgr) GetBatchFromCreateTime(createTimes []time.Time) (results []*mygormdl.DsUserBasicinfo, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("create_time IN (?)", createTimes).Find(&results).Error

	return
}

// GetFromUpdateTime 通过update_time获取内容 更新时间
func (obj *DsUserBasicinfoMgr) GetFromUpdateTime(updateTime time.Time) (results []*mygormdl.DsUserBasicinfo, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("update_time = ?", updateTime).Find(&results).Error

	return
}

// GetBatchFromUpdateTime 批量唯一主键查找 更新时间
func (obj *DsUserBasicinfoMgr) GetBatchFromUpdateTime(updateTimes []time.Time) (results []*mygormdl.DsUserBasicinfo, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("update_time IN (?)", updateTimes).Find(&results).Error

	return
}

// GetFromDeleted 通过deleted获取内容 是否删除0:未删除 1:删除
func (obj *DsUserBasicinfoMgr) GetFromDeleted(deleted int8) (results []*mygormdl.DsUserBasicinfo, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("deleted = ?", deleted).Find(&results).Error

	return
}

// GetBatchFromDeleted 批量唯一主键查找 是否删除0:未删除 1:删除
func (obj *DsUserBasicinfoMgr) GetBatchFromDeleted(deleteds []int8) (results []*mygormdl.DsUserBasicinfo, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("deleted IN (?)", deleteds).Find(&results).Error

	return
}

// GetFromRealName 通过real_name获取内容 真实姓名
func (obj *DsUserBasicinfoMgr) GetFromRealName(realName string) (results []*mygormdl.DsUserBasicinfo, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("real_name = ?", realName).Find(&results).Error

	return
}

// GetBatchFromRealName 批量唯一主键查找 真实姓名
func (obj *DsUserBasicinfoMgr) GetBatchFromRealName(realNames []string) (results []*mygormdl.DsUserBasicinfo, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("real_name IN (?)", realNames).Find(&results).Error

	return
}

// GetFromAge 通过age获取内容 年龄
func (obj *DsUserBasicinfoMgr) GetFromAge(age int8) (results []*mygormdl.DsUserBasicinfo, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("age = ?", age).Find(&results).Error

	return
}

// GetBatchFromAge 批量唯一主键查找 年龄
func (obj *DsUserBasicinfoMgr) GetBatchFromAge(ages []int8) (results []*mygormdl.DsUserBasicinfo, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("age IN (?)", ages).Find(&results).Error

	return
}

// GetFromGender 通过gender获取内容 性别：0:男，1:女
func (obj *DsUserBasicinfoMgr) GetFromGender(gender int8) (results []*mygormdl.DsUserBasicinfo, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("gender = ?", gender).Find(&results).Error

	return
}

// GetBatchFromGender 批量唯一主键查找 性别：0:男，1:女
func (obj *DsUserBasicinfoMgr) GetBatchFromGender(genders []int8) (results []*mygormdl.DsUserBasicinfo, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("gender IN (?)", genders).Find(&results).Error

	return
}

// GetFromAvatar 通过avatar获取内容 头像地址
func (obj *DsUserBasicinfoMgr) GetFromAvatar(avatar string) (results []*mygormdl.DsUserBasicinfo, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("avatar = ?", avatar).Find(&results).Error

	return
}

// GetBatchFromAvatar 批量唯一主键查找 头像地址
func (obj *DsUserBasicinfoMgr) GetBatchFromAvatar(avatars []string) (results []*mygormdl.DsUserBasicinfo, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("avatar IN (?)", avatars).Find(&results).Error

	return
}

// GetFromNickName 通过nick_name获取内容 昵称
func (obj *DsUserBasicinfoMgr) GetFromNickName(nickName string) (results []*mygormdl.DsUserBasicinfo, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("nick_name = ?", nickName).Find(&results).Error

	return
}

// GetBatchFromNickName 批量唯一主键查找 昵称
func (obj *DsUserBasicinfoMgr) GetBatchFromNickName(nickNames []string) (results []*mygormdl.DsUserBasicinfo, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("nick_name IN (?)", nickNames).Find(&results).Error

	return
}

// GetFromBirthday 通过birthday获取内容 生日
func (obj *DsUserBasicinfoMgr) GetFromBirthday(birthday time.Time) (results []*mygormdl.DsUserBasicinfo, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("birthday = ?", birthday).Find(&results).Error

	return
}

// GetBatchFromBirthday 批量唯一主键查找 生日
func (obj *DsUserBasicinfoMgr) GetBatchFromBirthday(birthdays []time.Time) (results []*mygormdl.DsUserBasicinfo, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("birthday IN (?)", birthdays).Find(&results).Error

	return
}

// GetFromStatus 通过status获取内容 账户状态0:正常 1:禁用  2:注销
func (obj *DsUserBasicinfoMgr) GetFromStatus(status int8) (results []*mygormdl.DsUserBasicinfo, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("status = ?", status).Find(&results).Error

	return
}

// GetBatchFromStatus 批量唯一主键查找 账户状态0:正常 1:禁用  2:注销
func (obj *DsUserBasicinfoMgr) GetBatchFromStatus(statuss []int8) (results []*mygormdl.DsUserBasicinfo, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("status IN (?)", statuss).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *DsUserBasicinfoMgr) FetchByPrimaryKey(uuid string) (result mygormdl.DsUserBasicinfo, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("uuid = ?", uuid).Find(&result).Error

	return
}

// FetchByIndex  获取多个内容
func (obj *DsUserBasicinfoMgr) FetchByIndex(mobile string) (results []*mygormdl.DsUserBasicinfo, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("mobile = ?", mobile).Find(&results).Error

	return
}
func (obj *DsUserBasicinfoMgr) GetHashSalt( mobile string) (result mygormdl.DsUserBasicinfo,err error) {
	err =  obj.DB.Table(obj.GetTableName()).Where("mobile = ?", mobile).Find(&result).Error
	return
}
func (obj *DsUserBasicinfoMgr) LoginVfcode(mobile string) (result mygormdl.DsUserBasicinfo,err error) {
	err =  obj.DB.Table(obj.GetTableName()).Where("mobile = ?", mobile).Find(&result).Error
	return
}

func (obj *DsUserBasicinfoMgr) GetUserBasicInfo(tx *gorm.DB, uuid string) (results mygormdl.DsUserBasicinfo, err error) {
	err = tx.Table(obj.GetTableName()).Where("uuid = ?", uuid).Find(&results).Error
	return
}
func (obj *DsUserBasicinfoMgr) GetUserBasicInfoEx(mobile string) (results mygormdl.DsUserBasicinfo, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("mobile = ?", mobile).Find(&results).Error
	return
}
func (obj *DsUserBasicinfoMgr) UpdateLoginPwd(mobile, hash string) (err error) {
	fmt.Println("new hash:",hash,mobile)
	return obj.DB.Table(obj.GetTableName()).Where("mobile= ?",mobile).Updates(mygormdl.DsUserBasicinfo{Hash: hash}).Error
}
func (obj *DsUserBasicinfoMgr) GetUserBasicInfoex(tx *gorm.DB, mobile string) (results mygormdl.DsUserBasicinfo, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("mobile = ?", mobile).Find(&results).Error
	return
}
func (obj *DsUserBasicinfoMgr) SetLastLogin(tx *gorm.DB, uuid,ip string, last_login_time time.Time) (err error) {
	return tx.Table(obj.GetTableName()).Where("uuid= ?",uuid).Updates(mygormdl.DsUserBasicinfo{LastLoginIP: ip,LastLoginTime: last_login_time}).Error
}

func (obj *DsUserBasicinfoMgr) ModifyUserInfo(tag,mobile, vfcode, content string) (err error) {
	switch tag{
	case "avatar":
		err =  obj.Table(obj.GetTableName()).Where("mobile= ?",mobile).Updates(mygormdl.DsUserBasicinfo{Avatar: content}).Error
		break
	case "nickname":
		err = obj.Table(obj.GetTableName()).Where("mobile= ?",mobile).Updates(mygormdl.DsUserBasicinfo{NickName: content}).Error
		break
	case "mobile":
		err = obj.Table(obj.GetTableName()).Where("mobile= ?",mobile).Updates(mygormdl.DsUserBasicinfo{Mobile: content}).Error
		break
	case "pwd":
		vf_code,_ := redisex.RedisInstanceg().Get(constex.REDIS_USER_VFCODE+mobile).Result()
		if vf_code != vfcode{
			return errors.New("短信验证码无效")
		}
		tx := obj.DB.Begin()
		var userbasic mygormdl.DsUserBasicinfo
		db := tx.Table(obj.GetTableName()).Where("mobile= ?",mobile).Find(&userbasic)
		if db.RecordNotFound(){
			tx.Rollback()
			err =  errors.New("不存在该手机号！")
			return
		}
		tmphash := md5.HashForPwd(userbasic.Salt,content)
		if tmphash == userbasic.Hash{
			tx.Rollback()
			err =  errors.New("密码相同，更新密码失败！")
			return
		}
		err =tx.Table(obj.GetTableName()).Where("mobile= ?",mobile).Updates(mygormdl.DsUserBasicinfo{Hash: md5.HashForPwd(userbasic.Salt,content)}).Error
		if err!=nil{
			tx.Rollback()
			err =  errors.New("更新密码失败！")
			return
		}
		tx.Commit()
		err = nil
		break
	case "realname":
		err =  obj.Table(obj.GetTableName()).Where("mobile= ?",mobile).Updates(mygormdl.DsUserBasicinfo{RealName: content}).Error
		break
	default:
		err =  errors.New("更改类型错误！")
	}
	return err
}
