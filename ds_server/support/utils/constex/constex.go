package constex

import (
	"ds_server/support/utils/cfg/common"
	en2 "github.com/go-playground/locales/en"
	zh2 "github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
)

const (
	SRV_GATAWAY    = "ds.srv.gateway"
	SRV_USER       = "ds.srv.user"
	SRV_BASE       = "ds.srv.base"
	SRV_MESSAGE    = "ds.srv.message"
	SRV_BRANCH     = "ds.srv.branch"
	SRV_MALL       = "ds.srv.mall"
	CONFIG_ADDRESS = ":8208"

	REDIS_USER_TOKEN  = "user:login:token:"
	REDIS_USER_VFCODE = "user:login:vfcode:"
	REDIS_USER_MEMBER = "user:member:V:"
	REDIS_USER_AGENT  = "user:agent:D:"

	Agent1money = 3000
	Agent2money = 9000
	Agent3money = 27000
	Agent4money = 81000
	Agent5money = 243000

	Agent1Tag = 1
	Agent2Tag = 2
	Agent3Tag = 3
	Agent4Tag = 4
	Agent5Tag = 5

	Agent1Tagex = "D1"
	Agent2Tagex = "D2"
	Agent3Tagex = "D3"
	Agent4Tagex = "D4"
	Agent5Tagex = "D5"

	Agent1Name = "分销商"
	Agent2Name = "经销商"
	Agent3Name = "区代"
	Agent4Name = "县代"
	Agent5Name = "市代"

	//--------------------------------------

	Member1money = 1000
	Member2money = 5000
	Member3money = 10000
	Member4money = 30000
	Member5money = 50000

	Mem1Tag = 1
	Mem2Tag = 2
	Mem3Tag = 3
	Mem4Tag = 4
	Mem5Tag = 5

	Mem1Tagex = "V1"
	Mem2Tagex = "V2"
	Mem3Tagex = "V3"
	Mem4Tagex = "V4"
	Mem5Tagex = "V5"

	Mem1Name = "普通卡"
	Mem2Name = "铜卡"
	Mem3Name = "银卡"
	Mem4Name = "金卡"
	Mem5Name = "钻石卡"
)

var (
	ConsulCfg       = &common.Consul{}
	MysqlCfg        = &common.Mysql{}
	RedisCfg        = &common.Redis{}
	RedisCluterCfg  = &common.RedisCluster{}
	WechatpayCfg    = &common.Wechatpay{}
	AlipayCfg       = &common.Alipay{}
	JwtCfg          = &common.Jwt{}
	AiossCfg        = &common.Aioss{}
	MemberClassCfg  = &common.MemberClass{}
	AgentClassCfg   = &common.AgentClass{}
	ExchangeRateCfg = &common.ExchangeRate{}
	GlobalCfg       = string("")
	GlobalMsg       = string("")
)

var (
	zh       = zh2.New()
	en       = en2.New()
	Uni      = ut.New(en, zh)
	Validate = validator.New()
	Tans, _  = Uni.GetTranslator("en")
)

