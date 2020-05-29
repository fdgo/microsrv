package gormdl

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

type _DsUserAgentClassMgr struct {
	*_BaseMgr
}

// DsUserAgentClassMgr open func
func DsUserAgentClassMgr(db *gorm.DB) *_DsUserAgentClassMgr {
	if db == nil {
		panic(fmt.Errorf("DsUserAgentClassMgr need init by db"))
	}
	return &_DsUserAgentClassMgr{_BaseMgr: &_BaseMgr{DB: db, isRelated: globalIsRelated}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_DsUserAgentClassMgr) GetTableName() string {
	return "ds_user_agent_class"
}

// Get 获取
func (obj *_DsUserAgentClassMgr) Get() (result DsUserAgentClass, err error) {
	err = obj.DB.Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_DsUserAgentClassMgr) Gets() (results []*DsUserAgentClass, err error) {
	err = obj.DB.Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithAgentMoney agent_money获取 合伙人等级金额界限
func (obj *_DsUserAgentClassMgr) WithAgentMoney(agentMoney float64) Option {
	return optionFunc(func(o *options) { o.query["agent_money"] = agentMoney })
}

// WithAgentTag agent_tag获取 合伙人等级1,2,3,4,5
func (obj *_DsUserAgentClassMgr) WithAgentTag(agentTag int8) Option {
	return optionFunc(func(o *options) { o.query["agent_tag"] = agentTag })
}

// WithAgentTagex agent_tagex获取 合伙人等级标记  D1，D2，D3，D4，D5
func (obj *_DsUserAgentClassMgr) WithAgentTagex(agentTagex string) Option {
	return optionFunc(func(o *options) { o.query["agent_tagex"] = agentTagex })
}

// WithAgentName agent_name获取 合伙人等级名称  代理合伙人，高级合伙人，城市合伙人，区域合伙人，全球合伙人
func (obj *_DsUserAgentClassMgr) WithAgentName(agentName string) Option {
	return optionFunc(func(o *options) { o.query["agent_name"] = agentName })
}

// GetByOption 功能选项模式获取
func (obj *_DsUserAgentClassMgr) GetByOption(opts ...Option) (result DsUserAgentClass, err error) {
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
func (obj *_DsUserAgentClassMgr) GetByOptions(opts ...Option) (results []*DsUserAgentClass, err error) {
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

// GetFromAgentMoney 通过agent_money获取内容 合伙人等级金额界限
func (obj *_DsUserAgentClassMgr) GetFromAgentMoney(agentMoney float64) (results []*DsUserAgentClass, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("agent_money = ?", agentMoney).Find(&results).Error

	return
}

// GetBatchFromAgentMoney 批量唯一主键查找 合伙人等级金额界限
func (obj *_DsUserAgentClassMgr) GetBatchFromAgentMoney(agentMoneys []float64) (results []*DsUserAgentClass, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("agent_money IN (?)", agentMoneys).Find(&results).Error

	return
}

// GetFromAgentTag 通过agent_tag获取内容 合伙人等级1,2,3,4,5
func (obj *_DsUserAgentClassMgr) GetFromAgentTag(agentTag int8) (results []*DsUserAgentClass, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("agent_tag = ?", agentTag).Find(&results).Error

	return
}

// GetBatchFromAgentTag 批量唯一主键查找 合伙人等级1,2,3,4,5
func (obj *_DsUserAgentClassMgr) GetBatchFromAgentTag(agentTags []int8) (results []*DsUserAgentClass, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("agent_tag IN (?)", agentTags).Find(&results).Error

	return
}

// GetFromAgentTagex 通过agent_tagex获取内容 合伙人等级标记  D1，D2，D3，D4，D5
func (obj *_DsUserAgentClassMgr) GetFromAgentTagex(agentTagex string) (results []*DsUserAgentClass, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("agent_tagex = ?", agentTagex).Find(&results).Error

	return
}

// GetBatchFromAgentTagex 批量唯一主键查找 合伙人等级标记  D1，D2，D3，D4，D5
func (obj *_DsUserAgentClassMgr) GetBatchFromAgentTagex(agentTagexs []string) (results []*DsUserAgentClass, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("agent_tagex IN (?)", agentTagexs).Find(&results).Error

	return
}

// GetFromAgentName 通过agent_name获取内容 合伙人等级名称  代理合伙人，高级合伙人，城市合伙人，区域合伙人，全球合伙人
func (obj *_DsUserAgentClassMgr) GetFromAgentName(agentName string) (results []*DsUserAgentClass, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("agent_name = ?", agentName).Find(&results).Error

	return
}

// GetBatchFromAgentName 批量唯一主键查找 合伙人等级名称  代理合伙人，高级合伙人，城市合伙人，区域合伙人，全球合伙人
func (obj *_DsUserAgentClassMgr) GetBatchFromAgentName(agentNames []string) (results []*DsUserAgentClass, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("agent_name IN (?)", agentNames).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByIndex  获取多个内容
func (obj *_DsUserAgentClassMgr) FetchByIndex(agentMoney float64, agentTag int8) (results []*DsUserAgentClass, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("agent_money = ? AND agent_tag = ?", agentMoney, agentTag).Find(&results).Error

	return
}
