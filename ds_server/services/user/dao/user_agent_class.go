package dao

import (
	mygormdl "ds_server/models/user/gorm_mysql"
	rs "ds_server/support/lib/redisex"
	"ds_server/support/utils/constex"
	"errors"
	"fmt"
	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
)

type DsUserAgentClassMgr struct {
	*_BaseMgr
}

// DsUserAgentClassMgr open func
func NewDsUserAgentClassMgr(db *gorm.DB) *DsUserAgentClassMgr {
	if db == nil {
		panic(fmt.Errorf("DsUserAgentClassMgr need init by db"))
	}
	return &DsUserAgentClassMgr{_BaseMgr: &_BaseMgr{DB: db, isRelated: globalIsRelated}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *DsUserAgentClassMgr) GetTableName() string {
	return "ds_user_agent_class"
}

// Get 获取
func (obj *DsUserAgentClassMgr) Get() (result mygormdl.DsUserAgentClass, err error) {
	err = obj.DB.Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *DsUserAgentClassMgr) Gets() (results []*mygormdl.DsUserAgentClass, err error) {
	err = obj.DB.Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithAgentMoney agent_money获取 合伙人等级金额界限
func (obj *DsUserAgentClassMgr) WithAgentMoney(agentMoney float64) Option {
	return optionFunc(func(o *options) { o.query["agent_money"] = agentMoney })
}

// WithAgentTag agent_tag获取 合伙人等级1,2,3,4,5
func (obj *DsUserAgentClassMgr) WithAgentTag(agentTag int8) Option {
	return optionFunc(func(o *options) { o.query["agent_tag"] = agentTag })
}

// WithAgentTagex agent_tagex获取 合伙人等级标记  D1，D2，D3，D4，D5
func (obj *DsUserAgentClassMgr) WithAgentTagex(agentTagex string) Option {
	return optionFunc(func(o *options) { o.query["agent_tagex"] = agentTagex })
}

// WithAgentName agent_name获取 合伙人等级名称  代理合伙人，高级合伙人，城市合伙人，区域合伙人，全球合伙人
func (obj *DsUserAgentClassMgr) WithAgentName(agentName string) Option {
	return optionFunc(func(o *options) { o.query["agent_name"] = agentName })
}

// GetByOption 功能选项模式获取
func (obj *DsUserAgentClassMgr) GetByOption(opts ...Option) (result mygormdl.DsUserAgentClass, err error) {
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
func (obj *DsUserAgentClassMgr) GetByOptions(opts ...Option) (results []*mygormdl.DsUserAgentClass, err error) {
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
func (obj *DsUserAgentClassMgr) GetFromAgentMoney(agentMoney float64) (results []*mygormdl.DsUserAgentClass, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("agent_money = ?", agentMoney).Find(&results).Error

	return
}

// GetBatchFromAgentMoney 批量唯一主键查找 合伙人等级金额界限
func (obj *DsUserAgentClassMgr) GetBatchFromAgentMoney(agentMoneys []float64) (results []*mygormdl.DsUserAgentClass, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("agent_money IN (?)", agentMoneys).Find(&results).Error

	return
}

// GetFromAgentTag 通过agent_tag获取内容 合伙人等级1,2,3,4,5
func (obj *DsUserAgentClassMgr) GetFromAgentTag(agentTag int8) (results []*mygormdl.DsUserAgentClass, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("agent_tag = ?", agentTag).Find(&results).Error

	return
}

// GetBatchFromAgentTag 批量唯一主键查找 合伙人等级1,2,3,4,5
func (obj *DsUserAgentClassMgr) GetBatchFromAgentTag(agentTags []int8) (results []*mygormdl.DsUserAgentClass, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("agent_tag IN (?)", agentTags).Find(&results).Error

	return
}

// GetFromAgentTagex 通过agent_tagex获取内容 合伙人等级标记  D1，D2，D3，D4，D5
func (obj *DsUserAgentClassMgr) GetFromAgentTagex(agentTagex string) (results []*mygormdl.DsUserAgentClass, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("agent_tagex = ?", agentTagex).Find(&results).Error

	return
}

// GetBatchFromAgentTagex 批量唯一主键查找 合伙人等级标记  D1，D2，D3，D4，D5
func (obj *DsUserAgentClassMgr) GetBatchFromAgentTagex(agentTagexs []string) (results []*mygormdl.DsUserAgentClass, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("agent_tagex IN (?)", agentTagexs).Find(&results).Error

	return
}

// GetFromAgentName 通过agent_name获取内容 合伙人等级名称  代理合伙人，高级合伙人，城市合伙人，区域合伙人，全球合伙人
func (obj *DsUserAgentClassMgr) GetFromAgentName(agentName string) (results []*mygormdl.DsUserAgentClass, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("agent_name = ?", agentName).Find(&results).Error

	return
}

// GetBatchFromAgentName 批量唯一主键查找 合伙人等级名称  代理合伙人，高级合伙人，城市合伙人，区域合伙人，全球合伙人
func (obj *DsUserAgentClassMgr) GetBatchFromAgentName(agentNames []string) (results []*mygormdl.DsUserAgentClass, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("agent_name IN (?)", agentNames).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByIndex  获取多个内容
func (obj *DsUserAgentClassMgr) FetchByIndex(agentMoney float64, agentTag int8) (results []*mygormdl.DsUserAgentClass, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("agent_money = ? AND agent_tag = ?", agentMoney, agentTag).Find(&results).Error

	return
}


func (obj *DsUserAgentClassMgr) AgentClassSetMysql(uacslice []*mygormdl.DsUserAgentClass) (isok bool, err error) {
	tx := obj.Begin()
	err =tx.Exec("delete from "+ obj.GetTableName() ).Error
	if err != nil {
		tx.Rollback()
		return false, err
	}
	for _, v := range uacslice {
		err =tx.Exec("Insert into "+ obj.GetTableName() + " (agent_money,agent_tag,agent_tagex,agent_name) values(?,?,?,?)",v.AgentMoney,v.AgentTag,v.AgentTagex,v.AgentName).Error
		if err != nil {
			tx.Rollback()
			return false, err
		}
	}
	tx.Commit()
	return true, nil
}
func (obj *DsUserAgentClassMgr)AgentClassGetMysql(tx *gorm.DB, money float64) (agent_tagex string, agent_tag int8 , agent_name string, exx error) {
	var mem_money float64
	if money > 0 && money < constex.AgentClassCfg.Agent1Money {
		return "",0,"",nil
		//mem_money = constex.AgentClassCfg.Agent1Money
	} else if money >= constex.AgentClassCfg.Agent1Money && money < constex.AgentClassCfg.Agent2Money {
		mem_money = constex.AgentClassCfg.Agent1Money
	} else if money >=  constex.AgentClassCfg.Agent2Money && money <  constex.AgentClassCfg.Agent3Money{
		mem_money = constex.AgentClassCfg.Agent2Money
	} else if money >=  constex.AgentClassCfg.Agent3Money && money <  constex.AgentClassCfg.Agent4Money {
		mem_money = constex.AgentClassCfg.Agent3Money
	} else if money >=  constex.AgentClassCfg.Agent4Money&& money <  constex.AgentClassCfg.Agent5Money {
		mem_money = constex.AgentClassCfg.Agent4Money
	}else {
		mem_money = constex.AgentClassCfg.Agent5Money
	}
	var agentclass mygormdl.DsUserAgentClass
	err := tx.Table(obj.GetTableName()).Where("agent_money = ?", mem_money).Find(&agentclass).Error
	if err != nil{
		return "",0,"", errors.New("mysqlex get error!")
	}
	return agentclass.AgentTagex,agentclass.AgentTag,agentclass.AgentName,nil
}

func (obj *DsUserAgentClassMgr) AgentClassSetRedis(agentxmoney ...map[string]interface{}) (isok bool, err error) {
	for k, v := range agentxmoney {
		if err := rs.RedisInstanceg().HMSet(fmt.Sprintf(constex.REDIS_USER_AGENT+"%d", k+1), v).Err(); err != nil {
			return false, err
		}
	}
	return true, nil
}
func (obj *DsUserAgentClassMgr)AgentClassGetRedis(redisdb *redis.Client, money float64) (agent_tag , agent_class, agent_name string, ex error) {
	var result []interface{}
	var Suffix,_,AgentxTag, AgentxTagex, AgentxName string
fmt.Println("money:",money,"contex.AgentClassCfg.Agent1Money:",constex.AgentClassCfg.Agent1Money)
	if money > 0 && money < constex.AgentClassCfg.Agent1Money {
		return "","","",nil
		//Suffix="1"
		//AgentxTag = "Agent1Tag"
		//AgentxTagex = "Agent1Tagex"
		//AgentxName = "Agent1Name"
	} else if money >= constex.AgentClassCfg.Agent1Money && money < constex.AgentClassCfg.Agent2Money {
		Suffix="1"
		AgentxTag = "Agent1Tag"
		AgentxTagex = "Agent1Tagex"
		AgentxName = "Agent1Name"
	} else if money >= constex.AgentClassCfg.Agent2Money && money < constex.AgentClassCfg.Agent3Money {
		Suffix="2"
		AgentxTag = "Agent2Tag"
		AgentxTagex = "Agent2Tagex"
		AgentxName = "Agent2Name"
	} else if money >= constex.AgentClassCfg.Agent3Money && money < constex.AgentClassCfg.Agent4Money {
		Suffix="3"
		AgentxTag = "Agent3Tag"
		AgentxTagex = "Agent3Tagex"
		AgentxName = "Agent3Name"
	} else if money >= constex.AgentClassCfg.Agent4Money && money < constex.AgentClassCfg.Agent5Money {
		Suffix="4"
		AgentxTag = "Agent4Tag"
		AgentxTagex = "Agent4Tagex"
		AgentxName = "Agent4Name"
	}else {
		Suffix="5"
		AgentxTag = "Agent5Tag"
		AgentxTagex = "Agent5Tagex"
		AgentxName = "Agent5Name"
	}
	result, ex = redisdb.HMGet(constex.REDIS_USER_AGENT+Suffix, AgentxTag, AgentxTagex, AgentxName).Result()
	if ex != nil { //服务器断开 "EOF"
		return "","","",errors.New("redisex hmget error!")
	}
	var isMemxTag, isMemxTagex, isMemxName bool
	isMemxTag, _ = redisdb.HExists(constex.REDIS_USER_AGENT+Suffix, AgentxTag).Result()
	isMemxTagex, _ = redisdb.HExists(constex.REDIS_USER_AGENT+Suffix, AgentxTagex).Result()
	isMemxName, _ = redisdb.HExists(constex.REDIS_USER_AGENT+Suffix, AgentxName).Result()
	if !isMemxTag || !isMemxTagex || !isMemxName {
		 return "","","",errors.New("redisex hmget error!")
	}
	agent_tag = result[0].(string)
	agent_class = result[1].(string)
	agent_name = result[2].(string)
	return agent_tag, agent_class,agent_name,nil
}
