package models

type DsUserAgentClass struct {
	AgentMoney string `xorm:"not null comment('合伙人等级金额界限') index DECIMAL(18,2)"`
	AgentName  string `xorm:"not null comment('合伙人等级名称  代理合伙人，高级合伙人，城市合伙人，区域合伙人，全球合伙人') VARCHAR(16)"`
	AgentTag   int    `xorm:"not null comment('合伙人等级1,2,3,4,5') index TINYINT(4)"`
	AgentTagex string `xorm:"not null comment('合伙人等级标记  D1，D2，D3，D4，D5') VARCHAR(16)"`
}
