package dao

import (
	mygormdl "ds_server/models/user/gorm_mysql"
	"fmt"
	"github.com/jinzhu/gorm"
	"time"
)

type DsUserMemberAgentMgr struct {
	*_BaseMgr
}

// DsUserMemberAgentMgr open func
func NewDsUserMemberAgentMgr(db *gorm.DB) *DsUserMemberAgentMgr {
	if db == nil {
		panic(fmt.Errorf("DsUserMemberAgentMgr need init by db"))
	}
	return &DsUserMemberAgentMgr{_BaseMgr: &_BaseMgr{DB: db, isRelated: globalIsRelated}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *DsUserMemberAgentMgr) GetTableName() string {
	return "ds_user_member_agent"
}

// Get 获取
func (obj *DsUserMemberAgentMgr) Get() (result mygormdl.DsUserMemberAgent, err error) {
	err = obj.DB.Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *DsUserMemberAgentMgr) Gets() (results []*mygormdl.DsUserMemberAgent, err error) {
	err = obj.DB.Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithUUIDSelf uuid_self获取 新注册的用户ID号
func (obj *DsUserMemberAgentMgr) WithUUIDSelf(uuidSelf string) Option {
	return optionFunc(func(o *options) { o.query["uuid_self"] = uuidSelf })
}

// WithMobileSelf mobile_self获取 新注册的用户手机号
func (obj *DsUserMemberAgentMgr) WithMobileSelf(mobileSelf string) Option {
	return optionFunc(func(o *options) { o.query["mobile_self"] = mobileSelf })
}

// WithInvcodeSelf invcode_self获取 新注册的用户自身邀请码
func (obj *DsUserMemberAgentMgr) WithInvcodeSelf(invcodeSelf string) Option {
	return optionFunc(func(o *options) { o.query["invcode_self"] = invcodeSelf })
}

// WithUUIDAgent uuid_agent获取 代理的ID号
func (obj *DsUserMemberAgentMgr) WithUUIDAgent(uuidAgent string) Option {
	return optionFunc(func(o *options) { o.query["uuid_agent"] = uuidAgent })
}

// WithMobileAgent mobile_agent获取 代理的手机号
func (obj *DsUserMemberAgentMgr) WithMobileAgent(mobileAgent string) Option {
	return optionFunc(func(o *options) { o.query["mobile_agent"] = mobileAgent })
}

// WithInvcodeAgent invcode_agent获取 代理的邀请码
func (obj *DsUserMemberAgentMgr) WithInvcodeAgent(invcodeAgent string) Option {
	return optionFunc(func(o *options) { o.query["invcode_agent"] = invcodeAgent })
}

// WithMemclassSelf memclass_self获取 会员等级
func (obj *DsUserMemberAgentMgr) WithMemclassSelf(memclassSelf int8) Option {
	return optionFunc(func(o *options) { o.query["memclass_self"] = memclassSelf })
}

// WithMemberTag member_tag获取 会员标识
func (obj *DsUserMemberAgentMgr) WithMemberTag(memberTag string) Option {
	return optionFunc(func(o *options) { o.query["member_tag"] = memberTag })
}

// WithMemberName member_name获取 会员名称
func (obj *DsUserMemberAgentMgr) WithMemberName(memberName string) Option {
	return optionFunc(func(o *options) { o.query["member_name"] = memberName })
}

// WithAgentClass agent_class获取 代理等级
func (obj *DsUserMemberAgentMgr) WithAgentClass(agentClass uint8) Option {
	return optionFunc(func(o *options) { o.query["agent_class"] = agentClass })
}

// WithAgentTag agent_tag获取 代理标识
func (obj *DsUserMemberAgentMgr) WithAgentTag(agentTag string) Option {
	return optionFunc(func(o *options) { o.query["agent_tag"] = agentTag })
}

// WithAgentName agent_name获取 代理名称
func (obj *DsUserMemberAgentMgr) WithAgentName(agentName string) Option {
	return optionFunc(func(o *options) { o.query["agent_name"] = agentName })
}

// GetByOption 功能选项模式获取
func (obj *DsUserMemberAgentMgr) GetByOption(opts ...Option) (result mygormdl.DsUserMemberAgent, err error) {
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
func (obj *DsUserMemberAgentMgr) GetByOptions(opts ...Option) (results []*mygormdl.DsUserMemberAgent, err error) {
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
func (obj *DsUserMemberAgentMgr) GetFromUUIDSelf(uuidSelf string) (result mygormdl.DsUserMemberAgent, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("uuid_self = ?", uuidSelf).Find(&result).Error

	return
}

// GetBatchFromUUIDSelf 批量唯一主键查找 新注册的用户ID号
func (obj *DsUserMemberAgentMgr) GetBatchFromUUIDSelf(uuidSelfs []string) (results []*mygormdl.DsUserMemberAgent, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("uuid_self IN (?)", uuidSelfs).Find(&results).Error

	return
}

// GetFromMobileSelf 通过mobile_self获取内容 新注册的用户手机号
func (obj *DsUserMemberAgentMgr) GetFromMobileSelf(mobileSelf string) (result mygormdl.DsUserMemberAgent, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("mobile_self = ?", mobileSelf).Find(&result).Error

	return
}

// GetBatchFromMobileSelf 批量唯一主键查找 新注册的用户手机号
func (obj *DsUserMemberAgentMgr) GetBatchFromMobileSelf(mobileSelfs []string) (results []*mygormdl.DsUserMemberAgent, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("mobile_self IN (?)", mobileSelfs).Find(&results).Error

	return
}

// GetFromInvcodeSelf 通过invcode_self获取内容 新注册的用户自身邀请码
func (obj *DsUserMemberAgentMgr) GetFromInvcodeSelf(invcodeSelf string) (result mygormdl.DsUserMemberAgent, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("invcode_self = ?", invcodeSelf).Find(&result).Error

	return
}

// GetBatchFromInvcodeSelf 批量唯一主键查找 新注册的用户自身邀请码
func (obj *DsUserMemberAgentMgr) GetBatchFromInvcodeSelf(invcodeSelfs []string) (results []*mygormdl.DsUserMemberAgent, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("invcode_self IN (?)", invcodeSelfs).Find(&results).Error

	return
}

// GetFromUUIDAgent 通过uuid_agent获取内容 代理的ID号
func (obj *DsUserMemberAgentMgr) GetFromUUIDAgent(uuidAgent string) (results []*mygormdl.DsUserMemberAgent, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("uuid_agent = ?", uuidAgent).Find(&results).Error

	return
}

// GetBatchFromUUIDAgent 批量唯一主键查找 代理的ID号
func (obj *DsUserMemberAgentMgr) GetBatchFromUUIDAgent(uuidAgents []string) (results []*mygormdl.DsUserMemberAgent, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("uuid_agent IN (?)", uuidAgents).Find(&results).Error

	return
}

// GetFromMobileAgent 通过mobile_agent获取内容 代理的手机号
func (obj *DsUserMemberAgentMgr) GetFromMobileAgent(mobileAgent string) (results []*mygormdl.DsUserMemberAgent, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("mobile_agent = ?", mobileAgent).Find(&results).Error

	return
}

// GetBatchFromMobileAgent 批量唯一主键查找 代理的手机号
func (obj *DsUserMemberAgentMgr) GetBatchFromMobileAgent(mobileAgents []string) (results []*mygormdl.DsUserMemberAgent, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("mobile_agent IN (?)", mobileAgents).Find(&results).Error

	return
}

// GetFromInvcodeAgent 通过invcode_agent获取内容 代理的邀请码
func (obj *DsUserMemberAgentMgr) GetFromInvcodeAgent(invcodeAgent string) (results []*mygormdl.DsUserMemberAgent, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("invcode_agent = ?", invcodeAgent).Find(&results).Error

	return
}

// GetBatchFromInvcodeAgent 批量唯一主键查找 代理的邀请码
func (obj *DsUserMemberAgentMgr) GetBatchFromInvcodeAgent(invcodeAgents []string) (results []*mygormdl.DsUserMemberAgent, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("invcode_agent IN (?)", invcodeAgents).Find(&results).Error

	return
}

// GetFromMemclassSelf 通过memclass_self获取内容 会员等级
func (obj *DsUserMemberAgentMgr) GetFromMemclassSelf(memclassSelf int8) (results []*mygormdl.DsUserMemberAgent, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("memclass_self = ?", memclassSelf).Find(&results).Error

	return
}

// GetBatchFromMemclassSelf 批量唯一主键查找 会员等级
func (obj *DsUserMemberAgentMgr) GetBatchFromMemclassSelf(memclassSelfs []int8) (results []*mygormdl.DsUserMemberAgent, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("memclass_self IN (?)", memclassSelfs).Find(&results).Error

	return
}

// GetFromMemberTag 通过member_tag获取内容 会员标识
func (obj *DsUserMemberAgentMgr) GetFromMemberTag(memberTag string) (results []*mygormdl.DsUserMemberAgent, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("member_tag = ?", memberTag).Find(&results).Error

	return
}

// GetBatchFromMemberTag 批量唯一主键查找 会员标识
func (obj *DsUserMemberAgentMgr) GetBatchFromMemberTag(memberTags []string) (results []*mygormdl.DsUserMemberAgent, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("member_tag IN (?)", memberTags).Find(&results).Error

	return
}

// GetFromMemberName 通过member_name获取内容 会员名称
func (obj *DsUserMemberAgentMgr) GetFromMemberName(memberName string) (results []*mygormdl.DsUserMemberAgent, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("member_name = ?", memberName).Find(&results).Error

	return
}

// GetBatchFromMemberName 批量唯一主键查找 会员名称
func (obj *DsUserMemberAgentMgr) GetBatchFromMemberName(memberNames []string) (results []*mygormdl.DsUserMemberAgent, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("member_name IN (?)", memberNames).Find(&results).Error

	return
}

// GetFromAgentClass 通过agent_class获取内容 代理等级
func (obj *DsUserMemberAgentMgr) GetFromAgentClass(agentClass uint8) (results []*mygormdl.DsUserMemberAgent, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("agent_class = ?", agentClass).Find(&results).Error

	return
}

// GetBatchFromAgentClass 批量唯一主键查找 代理等级
func (obj *DsUserMemberAgentMgr) GetBatchFromAgentClass(agentClasss []uint8) (results []*mygormdl.DsUserMemberAgent, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("agent_class IN (?)", agentClasss).Find(&results).Error

	return
}

// GetFromAgentTag 通过agent_tag获取内容 代理标识
func (obj *DsUserMemberAgentMgr) GetFromAgentTag(agentTag string) (results []*mygormdl.DsUserMemberAgent, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("agent_tag = ?", agentTag).Find(&results).Error

	return
}

// GetBatchFromAgentTag 批量唯一主键查找 代理标识
func (obj *DsUserMemberAgentMgr) GetBatchFromAgentTag(agentTags []string) (results []*mygormdl.DsUserMemberAgent, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("agent_tag IN (?)", agentTags).Find(&results).Error

	return
}

// GetFromAgentName 通过agent_name获取内容 代理名称
func (obj *DsUserMemberAgentMgr) GetFromAgentName(agentName string) (results []*mygormdl.DsUserMemberAgent, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("agent_name = ?", agentName).Find(&results).Error

	return
}

// GetBatchFromAgentName 批量唯一主键查找 代理名称
func (obj *DsUserMemberAgentMgr) GetBatchFromAgentName(agentNames []string) (results []*mygormdl.DsUserMemberAgent, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("agent_name IN (?)", agentNames).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *DsUserMemberAgentMgr) FetchByPrimaryKey(uuidSelf string) (result mygormdl.DsUserMemberAgent, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("uuid_self = ?", uuidSelf).Find(&result).Error

	return
}

// FetchByUnique primay or index 获取唯一内容
func (obj *DsUserMemberAgentMgr) FetchByUnique(mobileSelf string, invcodeSelf string) (result mygormdl.DsUserMemberAgent, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("mobile_self = ? AND invcode_self = ?", mobileSelf, invcodeSelf).Find(&result).Error

	return
}
func (obj *DsUserMemberAgentMgr) SetUserMemClass(tx *gorm.DB,  member_tagex string ,member_class int8,member_name string, uuid string ) error {

	return tx.Table(obj.GetTableName()).Where("uuid_self= ?",uuid ).Updates(
		mygormdl.DsUserMemberAgent{
			MemclassSelf:member_class,
			MemberTag:member_tagex,
			MemberName:member_name,
			CreateTime:time.Now(),
			UpdateTime:time.Now(),
		}).Error
}
func (obj *DsUserMemberAgentMgr) SetAgentClass(tx *gorm.DB, agent_tag int8 ,agent_class ,member_name string, InvcodeSelf string ) error {

	return tx.Table(obj.GetTableName()).Where("invcode_self= ?",InvcodeSelf ).Updates(
		mygormdl.DsUserMemberAgent{
			AgentClass:agent_tag,
			AgentTag: agent_class,
			AgentName:member_name,
			CreateTime:time.Now(),
			UpdateTime:time.Now(),
		}).Error
}

func (obj *DsUserMemberAgentMgr) IsInvCodeExist(invcode_agent string) (result mygormdl.DsUserMemberAgent, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("invcode_agent = ?", invcode_agent).Find(&result).Error
	return
}

func (obj *DsUserMemberAgentMgr) GetUserMemAgent(tx *gorm.DB,uuid string) (results mygormdl.DsUserMemberAgent, err error) {
	err = tx.Table(obj.GetTableName()).Where("uuid_self = ?", uuid).Find(&results).Error
	return
}
func (obj *DsUserMemberAgentMgr) GetUserMemAgentEx(tx *gorm.DB, invcodeRev string) (results mygormdl.DsUserMemberAgent, err error) {
	err = tx.Table(obj.GetTableName()).Where("invcode_self = ?", invcodeRev).Find(&results).Error
	return
}
