package gormdl

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"time"
)

type _DsUserMemberAgentMgr struct {
	*_BaseMgr
}

// DsUserMemberAgentMgr open func
func DsUserMemberAgentMgr(db *gorm.DB) *_DsUserMemberAgentMgr {
	if db == nil {
		panic(fmt.Errorf("DsUserMemberAgentMgr need init by db"))
	}
	return &_DsUserMemberAgentMgr{_BaseMgr: &_BaseMgr{DB: db, isRelated: globalIsRelated}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_DsUserMemberAgentMgr) GetTableName() string {
	return "ds_user_member_agent"
}

// Get 获取
func (obj *_DsUserMemberAgentMgr) Get() (result DsUserMemberAgent, err error) {
	err = obj.DB.Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_DsUserMemberAgentMgr) Gets() (results []*DsUserMemberAgent, err error) {
	err = obj.DB.Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithUUIDSelf uuid_self获取 新注册的用户ID号
func (obj *_DsUserMemberAgentMgr) WithUUIDSelf(uuidSelf string) Option {
	return optionFunc(func(o *options) { o.query["uuid_self"] = uuidSelf })
}

// WithMobileSelf mobile_self获取 新注册的用户手机号
func (obj *_DsUserMemberAgentMgr) WithMobileSelf(mobileSelf string) Option {
	return optionFunc(func(o *options) { o.query["mobile_self"] = mobileSelf })
}

// WithInvcodeSelf invcode_self获取 新注册的用户自身邀请码
func (obj *_DsUserMemberAgentMgr) WithInvcodeSelf(invcodeSelf string) Option {
	return optionFunc(func(o *options) { o.query["invcode_self"] = invcodeSelf })
}

// WithUUIDAgent uuid_agent获取 代理的ID号
func (obj *_DsUserMemberAgentMgr) WithUUIDAgent(uuidAgent string) Option {
	return optionFunc(func(o *options) { o.query["uuid_agent"] = uuidAgent })
}

// WithMobileAgent mobile_agent获取 代理的手机号
func (obj *_DsUserMemberAgentMgr) WithMobileAgent(mobileAgent string) Option {
	return optionFunc(func(o *options) { o.query["mobile_agent"] = mobileAgent })
}

// WithInvcodeAgent invcode_agent获取 代理的邀请码
func (obj *_DsUserMemberAgentMgr) WithInvcodeAgent(invcodeAgent string) Option {
	return optionFunc(func(o *options) { o.query["invcode_agent"] = invcodeAgent })
}

// WithMemclassSelf memclass_self获取 会员等级
func (obj *_DsUserMemberAgentMgr) WithMemclassSelf(memclassSelf int8) Option {
	return optionFunc(func(o *options) { o.query["memclass_self"] = memclassSelf })
}

// WithMemberTag member_tag获取 会员标识
func (obj *_DsUserMemberAgentMgr) WithMemberTag(memberTag string) Option {
	return optionFunc(func(o *options) { o.query["member_tag"] = memberTag })
}

// WithMemberName member_name获取 会员名称
func (obj *_DsUserMemberAgentMgr) WithMemberName(memberName string) Option {
	return optionFunc(func(o *options) { o.query["member_name"] = memberName })
}

// WithAgentClass agent_class获取 代理等级
func (obj *_DsUserMemberAgentMgr) WithAgentClass(agentClass uint8) Option {
	return optionFunc(func(o *options) { o.query["agent_class"] = agentClass })
}

// WithAgentTag agent_tag获取 代理标识
func (obj *_DsUserMemberAgentMgr) WithAgentTag(agentTag string) Option {
	return optionFunc(func(o *options) { o.query["agent_tag"] = agentTag })
}

// WithAgentName agent_name获取 代理名称
func (obj *_DsUserMemberAgentMgr) WithAgentName(agentName string) Option {
	return optionFunc(func(o *options) { o.query["agent_name"] = agentName })
}

// WithCreateTime create_time获取 创建时间
func (obj *_DsUserMemberAgentMgr) WithCreateTime(createTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["create_time"] = createTime })
}

// WithUpdateTime update_time获取 更新时间
func (obj *_DsUserMemberAgentMgr) WithUpdateTime(updateTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["update_time"] = updateTime })
}

// GetByOption 功能选项模式获取
func (obj *_DsUserMemberAgentMgr) GetByOption(opts ...Option) (result DsUserMemberAgent, err error) {
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
func (obj *_DsUserMemberAgentMgr) GetByOptions(opts ...Option) (results []*DsUserMemberAgent, err error) {
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

// GetFromUUIDSelf 通过uuid_self获取内容 新注册的用户ID号
func (obj *_DsUserMemberAgentMgr) GetFromUUIDSelf(uuidSelf string) (result DsUserMemberAgent, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("uuid_self = ?", UUIDSelf).Find(&result).Error

	return
}

// GetBatchFromUUIDSelf 批量唯一主键查找 新注册的用户ID号
func (obj *_DsUserMemberAgentMgr) GetBatchFromUUIDSelf(uuidSelfs []string) (results []*DsUserMemberAgent, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("uuid_self IN (?)", uuidSelfs).Find(&results).Error

	return
}

// GetFromMobileSelf 通过mobile_self获取内容 新注册的用户手机号
func (obj *_DsUserMemberAgentMgr) GetFromMobileSelf(mobileSelf string) (result DsUserMemberAgent, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("mobile_self = ?", MobileSelf).Find(&result).Error

	return
}

// GetBatchFromMobileSelf 批量唯一主键查找 新注册的用户手机号
func (obj *_DsUserMemberAgentMgr) GetBatchFromMobileSelf(mobileSelfs []string) (results []*DsUserMemberAgent, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("mobile_self IN (?)", mobileSelfs).Find(&results).Error

	return
}

// GetFromInvcodeSelf 通过invcode_self获取内容 新注册的用户自身邀请码
func (obj *_DsUserMemberAgentMgr) GetFromInvcodeSelf(invcodeSelf string) (result DsUserMemberAgent, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("invcode_self = ?", InvcodeSelf).Find(&result).Error

	return
}

// GetBatchFromInvcodeSelf 批量唯一主键查找 新注册的用户自身邀请码
func (obj *_DsUserMemberAgentMgr) GetBatchFromInvcodeSelf(invcodeSelfs []string) (results []*DsUserMemberAgent, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("invcode_self IN (?)", invcodeSelfs).Find(&results).Error

	return
}

// GetFromUUIDAgent 通过uuid_agent获取内容 代理的ID号
func (obj *_DsUserMemberAgentMgr) GetFromUUIDAgent(uuidAgent string) (results []*DsUserMemberAgent, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("uuid_agent = ?", uuidAgent).Find(&results).Error

	return
}

// GetBatchFromUUIDAgent 批量唯一主键查找 代理的ID号
func (obj *_DsUserMemberAgentMgr) GetBatchFromUUIDAgent(uuidAgents []string) (results []*DsUserMemberAgent, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("uuid_agent IN (?)", uuidAgents).Find(&results).Error

	return
}

// GetFromMobileAgent 通过mobile_agent获取内容 代理的手机号
func (obj *_DsUserMemberAgentMgr) GetFromMobileAgent(mobileAgent string) (results []*DsUserMemberAgent, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("mobile_agent = ?", mobileAgent).Find(&results).Error

	return
}

// GetBatchFromMobileAgent 批量唯一主键查找 代理的手机号
func (obj *_DsUserMemberAgentMgr) GetBatchFromMobileAgent(mobileAgents []string) (results []*DsUserMemberAgent, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("mobile_agent IN (?)", mobileAgents).Find(&results).Error

	return
}

// GetFromInvcodeAgent 通过invcode_agent获取内容 代理的邀请码
func (obj *_DsUserMemberAgentMgr) GetFromInvcodeAgent(invcodeAgent string) (results []*DsUserMemberAgent, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("invcode_agent = ?", invcodeAgent).Find(&results).Error

	return
}

// GetBatchFromInvcodeAgent 批量唯一主键查找 代理的邀请码
func (obj *_DsUserMemberAgentMgr) GetBatchFromInvcodeAgent(invcodeAgents []string) (results []*DsUserMemberAgent, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("invcode_agent IN (?)", invcodeAgents).Find(&results).Error

	return
}

// GetFromMemclassSelf 通过memclass_self获取内容 会员等级
func (obj *_DsUserMemberAgentMgr) GetFromMemclassSelf(memclassSelf int8) (results []*DsUserMemberAgent, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("memclass_self = ?", memclassSelf).Find(&results).Error

	return
}

// GetBatchFromMemclassSelf 批量唯一主键查找 会员等级
func (obj *_DsUserMemberAgentMgr) GetBatchFromMemclassSelf(memclassSelfs []int8) (results []*DsUserMemberAgent, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("memclass_self IN (?)", memclassSelfs).Find(&results).Error

	return
}

// GetFromMemberTag 通过member_tag获取内容 会员标识
func (obj *_DsUserMemberAgentMgr) GetFromMemberTag(memberTag string) (results []*DsUserMemberAgent, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("member_tag = ?", memberTag).Find(&results).Error

	return
}

// GetBatchFromMemberTag 批量唯一主键查找 会员标识
func (obj *_DsUserMemberAgentMgr) GetBatchFromMemberTag(memberTags []string) (results []*DsUserMemberAgent, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("member_tag IN (?)", memberTags).Find(&results).Error

	return
}

// GetFromMemberName 通过member_name获取内容 会员名称
func (obj *_DsUserMemberAgentMgr) GetFromMemberName(memberName string) (results []*DsUserMemberAgent, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("member_name = ?", memberName).Find(&results).Error

	return
}

// GetBatchFromMemberName 批量唯一主键查找 会员名称
func (obj *_DsUserMemberAgentMgr) GetBatchFromMemberName(memberNames []string) (results []*DsUserMemberAgent, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("member_name IN (?)", memberNames).Find(&results).Error

	return
}

// GetFromAgentClass 通过agent_class获取内容 代理等级
func (obj *_DsUserMemberAgentMgr) GetFromAgentClass(agentClass uint8) (results []*DsUserMemberAgent, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("agent_class = ?", agentClass).Find(&results).Error

	return
}

// GetBatchFromAgentClass 批量唯一主键查找 代理等级
func (obj *_DsUserMemberAgentMgr) GetBatchFromAgentClass(agentClasss []uint8) (results []*DsUserMemberAgent, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("agent_class IN (?)", agentClasss).Find(&results).Error

	return
}

// GetFromAgentTag 通过agent_tag获取内容 代理标识
func (obj *_DsUserMemberAgentMgr) GetFromAgentTag(agentTag string) (results []*DsUserMemberAgent, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("agent_tag = ?", agentTag).Find(&results).Error

	return
}

// GetBatchFromAgentTag 批量唯一主键查找 代理标识
func (obj *_DsUserMemberAgentMgr) GetBatchFromAgentTag(agentTags []string) (results []*DsUserMemberAgent, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("agent_tag IN (?)", agentTags).Find(&results).Error

	return
}

// GetFromAgentName 通过agent_name获取内容 代理名称
func (obj *_DsUserMemberAgentMgr) GetFromAgentName(agentName string) (results []*DsUserMemberAgent, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("agent_name = ?", agentName).Find(&results).Error

	return
}

// GetBatchFromAgentName 批量唯一主键查找 代理名称
func (obj *_DsUserMemberAgentMgr) GetBatchFromAgentName(agentNames []string) (results []*DsUserMemberAgent, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("agent_name IN (?)", agentNames).Find(&results).Error

	return
}

// GetFromCreateTime 通过create_time获取内容 创建时间
func (obj *_DsUserMemberAgentMgr) GetFromCreateTime(createTime time.Time) (results []*DsUserMemberAgent, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("create_time = ?", createTime).Find(&results).Error

	return
}

// GetBatchFromCreateTime 批量唯一主键查找 创建时间
func (obj *_DsUserMemberAgentMgr) GetBatchFromCreateTime(createTimes []time.Time) (results []*DsUserMemberAgent, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("create_time IN (?)", createTimes).Find(&results).Error

	return
}

// GetFromUpdateTime 通过update_time获取内容 更新时间
func (obj *_DsUserMemberAgentMgr) GetFromUpdateTime(updateTime time.Time) (results []*DsUserMemberAgent, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("update_time = ?", updateTime).Find(&results).Error

	return
}

// GetBatchFromUpdateTime 批量唯一主键查找 更新时间
func (obj *_DsUserMemberAgentMgr) GetBatchFromUpdateTime(updateTimes []time.Time) (results []*DsUserMemberAgent, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("update_time IN (?)", updateTimes).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_DsUserMemberAgentMgr) FetchByPrimaryKey(uuidSelf string) (result DsUserMemberAgent, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("uuid_self = ?", uuidSelf).Find(&result).Error

	return
}

// FetchByUnique primay or index 获取唯一内容
func (obj *_DsUserMemberAgentMgr) FetchByUnique(mobileSelf string, invcodeSelf string) (result DsUserMemberAgent, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("mobile_self = ? AND invcode_self = ?", mobileSelf, invcodeSelf).Find(&result).Error

	return
}
