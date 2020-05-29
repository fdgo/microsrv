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

type DsUserMemberClassMgr struct {
	*_BaseMgr
}

// DsUserMemberClassMgr open func
func NewDsUserMemberClassMgr(db *gorm.DB) *DsUserMemberClassMgr {
	if db == nil {
		panic(fmt.Errorf("DsUserMemberClassMgr need init by db"))
	}
	return &DsUserMemberClassMgr{_BaseMgr: &_BaseMgr{DB: db, isRelated: globalIsRelated}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *DsUserMemberClassMgr) GetTableName() string {
	return "ds_user_member_class"
}

// Get 获取
func (obj *DsUserMemberClassMgr) Get() (result mygormdl.DsUserMemberClass, err error) {
	err = obj.DB.Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *DsUserMemberClassMgr) Gets() (results []*mygormdl.DsUserMemberClass, err error) {
	err = obj.DB.Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithMemMoney mem_money获取 会员等级金额界限
func (obj *DsUserMemberClassMgr) WithMemMoney(memMoney float64) Option {
	return optionFunc(func(o *options) { o.query["mem_money"] = memMoney })
}

// WithMemTag mem_tag获取 会员等级1,2,3,4,5
func (obj *DsUserMemberClassMgr) WithMemTag(memTag int8) Option {
	return optionFunc(func(o *options) { o.query["mem_tag"] = memTag })
}

// WithMemTagex mem_tagex获取 会员等级标记  M1，M2，M3，M4，M5
func (obj *DsUserMemberClassMgr) WithMemTagex(memTagex string) Option {
	return optionFunc(func(o *options) { o.query["mem_tagex"] = memTagex })
}

// WithMemName mem_name获取 会员等级名称  普通卡，铜卡，银卡，金卡，钻石卡
func (obj *DsUserMemberClassMgr) WithMemName(memName string) Option {
	return optionFunc(func(o *options) { o.query["mem_name"] = memName })
}

// GetByOption 功能选项模式获取
func (obj *DsUserMemberClassMgr) GetByOption(opts ...Option) (result mygormdl.DsUserMemberClass, err error) {
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
func (obj *DsUserMemberClassMgr) GetByOptions(opts ...Option) (results []*mygormdl.DsUserMemberClass, err error) {
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
func (obj *DsUserMemberClassMgr) GetFromMemMoney(memMoney float64) (results []*mygormdl.DsUserMemberClass, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("mem_money = ?", memMoney).Find(&results).Error

	return
}

// GetBatchFromMemMoney 批量唯一主键查找 会员等级金额界限
func (obj *DsUserMemberClassMgr) GetBatchFromMemMoney(memMoneys []float64) (results []*mygormdl.DsUserMemberClass, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("mem_money IN (?)", memMoneys).Find(&results).Error

	return
}

// GetFromMemTag 通过mem_tag获取内容 会员等级1,2,3,4,5
func (obj *DsUserMemberClassMgr) GetFromMemTag(memTag int8) (results []*mygormdl.DsUserMemberClass, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("mem_tag = ?", memTag).Find(&results).Error

	return
}

// GetOneFromMemTag 通过memTag获取内容 会员等级1,2,3,4,5
func (obj *DsUserMemberClassMgr) GetOneFromMemTag(memTag int8) (results mygormdl.DsUserMemberClass, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("mem_tag = ?", memTag).First(&results).Error

	return
}

// GetBatchFromMemTag 批量唯一主键查找 会员等级1,2,3,4,5
func (obj *DsUserMemberClassMgr) GetBatchFromMemTag(memTags []int8) (results []*mygormdl.DsUserMemberClass, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("mem_tag IN (?)", memTags).Find(&results).Error

	return
}

// GetFromMemTagex 通过mem_tagex获取内容 会员等级标记  M1，M2，M3，M4，M5
func (obj *DsUserMemberClassMgr) GetFromMemTagex(memTagex string) (results []*mygormdl.DsUserMemberClass, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("mem_tagex = ?", memTagex).Find(&results).Error

	return
}

// GetBatchFromMemTagex 批量唯一主键查找 会员等级标记  M1，M2，M3，M4，M5
func (obj *DsUserMemberClassMgr) GetBatchFromMemTagex(memTagexs []string) (results []*mygormdl.DsUserMemberClass, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("mem_tagex IN (?)", memTagexs).Find(&results).Error

	return
}

// GetFromMemName 通过mem_name获取内容 会员等级名称  普通卡，铜卡，银卡，金卡，钻石卡
func (obj *DsUserMemberClassMgr) GetFromMemName(memName string) (results []*mygormdl.DsUserMemberClass, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("mem_name = ?", memName).Find(&results).Error

	return
}

// GetBatchFromMemName 批量唯一主键查找 会员等级名称  普通卡，铜卡，银卡，金卡，钻石卡
func (obj *DsUserMemberClassMgr) GetBatchFromMemName(memNames []string) (results []*mygormdl.DsUserMemberClass, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("mem_name IN (?)", memNames).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByIndex  获取多个内容
func (obj *DsUserMemberClassMgr) FetchByIndex(memMoney float64, memTag int8) (results []*mygormdl.DsUserMemberClass, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("mem_money = ? AND mem_tag = ?", memMoney, memTag).Find(&results).Error

	return
}

func (obj *DsUserMemberClassMgr) MemerClassSetMysql(umcslice []*mygormdl.DsUserMemberClass) (isok bool, err error) {

	tx := obj.Begin()
	err = tx.Exec("delete from " + obj.GetTableName()).Error
	if err != nil {
		tx.Rollback()
		return false, err
	}
	for _, v := range umcslice {
		err = tx.Exec("Insert into "+obj.GetTableName()+" (mem_money,mem_tag,mem_tagex,mem_name) values(?,?,?,?)", v.MemMoney, v.MemTag, v.MemTagex, v.MemName).Error
		if err != nil {
			tx.Rollback()
			return false, err
		}
	}
	tx.Commit()
	return true, nil
}
func (obj *DsUserMemberClassMgr) MemberClassGetMysql(tx *gorm.DB, money float64) (member_tagex string, member_tag int8, member_name string, exx error) {
	var mem_money float64
	if money > 0 && money < constex.MemberClassCfg.Member1Money {
		return "",0,"",nil
		//mem_money = constex.MemberClassCfg.Member1Money
	} else if money >= constex.MemberClassCfg.Member1Money && money < constex.MemberClassCfg.Member2Money {
		mem_money = constex.MemberClassCfg.Member1Money
	} else if money >= constex.MemberClassCfg.Member2Money && money < constex.MemberClassCfg.Member3Money {
		mem_money = constex.MemberClassCfg.Member2Money
	} else if money >= constex.MemberClassCfg.Member3Money && money < constex.MemberClassCfg.Member4Money {
		mem_money = constex.MemberClassCfg.Member3Money
	} else if money >= constex.MemberClassCfg.Member4Money && money < constex.MemberClassCfg.Member5Money {
		mem_money = constex.MemberClassCfg.Member4Money
	} else {
		mem_money = constex.MemberClassCfg.Member5Money
	}
	var memberclass mygormdl.DsUserMemberClass
	err := tx.Table(obj.GetTableName()).Where("mem_money = ?", mem_money).Find(&memberclass).Error
	if err != nil {
		return "", 0, "", errors.New("mysqlex get error!")
	}
	return memberclass.MemTagex, memberclass.MemTag, memberclass.MemName,nil

}
func (obj *DsUserMemberClassMgr) MemerClassSetRedis(memberxmoney ...map[string]interface{}) (isok bool, err error) {
	for k, v := range memberxmoney {
		if err := rs.RedisInstanceg().HMSet(fmt.Sprintf(constex.REDIS_USER_MEMBER+"%d", k+1), v).Err(); err != nil {
			return false, err
		}
	}
	return true, nil
}

func (obj *DsUserMemberClassMgr) MemberClassGetRedis(redisdb *redis.Client, money float64) (member_tag, member_class, member_name string, exx error) {
	var err error
	var result []interface{}
	var Suffix, _, MemxTag, MemxTagex, MemxName string

	if money > 0 && money < constex.MemberClassCfg.Member1Money {
		//Suffix="1"
		//MemxTag = "Mem1Tag"
		//MemxTagex = "Mem1Tagex"
		//MemxName = "Mem1Name"
		return "","","",nil
	} else if money >= constex.MemberClassCfg.Member1Money && money < constex.MemberClassCfg.Member2Money {
		Suffix = "1"
		MemxTag = "Mem1Tag"
		MemxTagex = "Mem1Tagex"
		MemxName = "Mem1Name"
	} else if money >= constex.MemberClassCfg.Member2Money && money < constex.MemberClassCfg.Member3Money {
		Suffix = "2"
		MemxTag = "Mem2Tag"
		MemxTagex = "Mem2Tagex"
		MemxName = "Mem2Name"
	} else if money >= constex.MemberClassCfg.Member3Money && money < constex.MemberClassCfg.Member4Money {
		Suffix = "3"
		MemxTag = "Mem3Tag"
		MemxTagex = "Mem3Tagex"
		MemxName = "Mem3Name"
	} else if money >= constex.MemberClassCfg.Member4Money && money < constex.MemberClassCfg.Member5Money {
		Suffix = "4"
		MemxTag = "Mem4Tag"
		MemxTagex = "Mem4Tagex"
		MemxName = "Mem4Name"
	} else {
		Suffix = "5"
		MemxTag = "Mem5Tag"
		MemxTagex = "Mem5Tagex"
		MemxName = "Mem5Name"
	}
	result, err = redisdb.HMGet(constex.REDIS_USER_MEMBER+Suffix, MemxTag, MemxTagex, MemxName).Result()
	if err != nil { //服务器断开 "EOF"
		return "","","",errors.New("redisex get error!")
	}
	var isMemxTag, isMemxTagex, isMemxName bool
	isMemxTag, err = redisdb.HExists(constex.REDIS_USER_MEMBER+Suffix, MemxTag).Result()
	isMemxTagex, err = redisdb.HExists(constex.REDIS_USER_MEMBER+Suffix, MemxTagex).Result()
	isMemxName, err = redisdb.HExists(constex.REDIS_USER_MEMBER+Suffix, MemxName).Result()
	if !isMemxTag || !isMemxTagex || !isMemxName {
		return "","","",errors.New("redisex get error!")
	}
	member_tag = result[0].(string)
	member_class = result[1].(string)
	member_name = result[2].(string)
	return member_tag,member_class,member_name,nil
}
