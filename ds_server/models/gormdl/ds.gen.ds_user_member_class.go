package gormdl

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

type _DsUserMemberClassMgr struct {
	*_BaseMgr
}

// DsUserMemberClassMgr open func
func DsUserMemberClassMgr(db *gorm.DB) *_DsUserMemberClassMgr {
	if db == nil {
		panic(fmt.Errorf("DsUserMemberClassMgr need init by db"))
	}
	return &_DsUserMemberClassMgr{_BaseMgr: &_BaseMgr{DB: db, isRelated: globalIsRelated}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_DsUserMemberClassMgr) GetTableName() string {
	return "ds_user_member_class"
}

// Get 获取
func (obj *_DsUserMemberClassMgr) Get() (result DsUserMemberClass, err error) {
	err = obj.DB.Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_DsUserMemberClassMgr) Gets() (results []*DsUserMemberClass, err error) {
	err = obj.DB.Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithMemMoney mem_money获取 会员等级金额界限
func (obj *_DsUserMemberClassMgr) WithMemMoney(memMoney float64) Option {
	return optionFunc(func(o *options) { o.query["mem_money"] = memMoney })
}

// WithMemTag mem_tag获取 会员等级1,2,3,4,5
func (obj *_DsUserMemberClassMgr) WithMemTag(memTag int8) Option {
	return optionFunc(func(o *options) { o.query["mem_tag"] = memTag })
}

// WithMemTagex mem_tagex获取 会员等级标记  M1，M2，M3，M4，M5
func (obj *_DsUserMemberClassMgr) WithMemTagex(memTagex string) Option {
	return optionFunc(func(o *options) { o.query["mem_tagex"] = memTagex })
}

// WithMemName mem_name获取 会员等级名称  普通卡，铜卡，银卡，金卡，钻石卡
func (obj *_DsUserMemberClassMgr) WithMemName(memName string) Option {
	return optionFunc(func(o *options) { o.query["mem_name"] = memName })
}

// GetByOption 功能选项模式获取
func (obj *_DsUserMemberClassMgr) GetByOption(opts ...Option) (result DsUserMemberClass, err error) {
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
func (obj *_DsUserMemberClassMgr) GetByOptions(opts ...Option) (results []*DsUserMemberClass, err error) {
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

// GetFromMemMoney 通过mem_money获取内容 会员等级金额界限
func (obj *_DsUserMemberClassMgr) GetFromMemMoney(memMoney float64) (results []*DsUserMemberClass, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("mem_money = ?", memMoney).Find(&results).Error

	return
}

// GetBatchFromMemMoney 批量唯一主键查找 会员等级金额界限
func (obj *_DsUserMemberClassMgr) GetBatchFromMemMoney(memMoneys []float64) (results []*DsUserMemberClass, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("mem_money IN (?)", memMoneys).Find(&results).Error

	return
}

// GetFromMemTag 通过mem_tag获取内容 会员等级1,2,3,4,5
func (obj *_DsUserMemberClassMgr) GetFromMemTag(memTag int8) (results []*DsUserMemberClass, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("mem_tag = ?", memTag).Find(&results).Error

	return
}

// GetBatchFromMemTag 批量唯一主键查找 会员等级1,2,3,4,5
func (obj *_DsUserMemberClassMgr) GetBatchFromMemTag(memTags []int8) (results []*DsUserMemberClass, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("mem_tag IN (?)", memTags).Find(&results).Error

	return
}

// GetFromMemTagex 通过mem_tagex获取内容 会员等级标记  M1，M2，M3，M4，M5
func (obj *_DsUserMemberClassMgr) GetFromMemTagex(memTagex string) (results []*DsUserMemberClass, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("mem_tagex = ?", memTagex).Find(&results).Error

	return
}

// GetBatchFromMemTagex 批量唯一主键查找 会员等级标记  M1，M2，M3，M4，M5
func (obj *_DsUserMemberClassMgr) GetBatchFromMemTagex(memTagexs []string) (results []*DsUserMemberClass, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("mem_tagex IN (?)", memTagexs).Find(&results).Error

	return
}

// GetFromMemName 通过mem_name获取内容 会员等级名称  普通卡，铜卡，银卡，金卡，钻石卡
func (obj *_DsUserMemberClassMgr) GetFromMemName(memName string) (results []*DsUserMemberClass, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("mem_name = ?", memName).Find(&results).Error

	return
}

// GetBatchFromMemName 批量唯一主键查找 会员等级名称  普通卡，铜卡，银卡，金卡，钻石卡
func (obj *_DsUserMemberClassMgr) GetBatchFromMemName(memNames []string) (results []*DsUserMemberClass, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("mem_name IN (?)", memNames).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByIndex  获取多个内容
func (obj *_DsUserMemberClassMgr) FetchByIndex(memMoney float64, memTag int8) (results []*DsUserMemberClass, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("mem_money = ? AND mem_tag = ?", memMoney, memTag).Find(&results).Error

	return
}
