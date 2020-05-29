package models

import (
	"time"
)

type DsUserMemberAgent struct {
	AgentClass   int       `xorm:"not null default 000 comment('代理等级') TINYINT(3)"`
	AgentName    string    `xorm:"not null default '' comment('代理名称') VARCHAR(16)"`
	AgentTag     string    `xorm:"not null default '' comment('代理标识') VARCHAR(16)"`
	CreateTime   time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' comment('创建时间') DATETIME"`
	InvcodeAgent string    `xorm:"not null default '' comment('代理的邀请码') VARCHAR(16)"`
	InvcodeSelf  string    `xorm:"not null default '' comment('新注册的用户自身邀请码') unique VARCHAR(16)"`
	MemberName   string    `xorm:"not null default '' comment('会员名称') VARCHAR(16)"`
	MemberTag    string    `xorm:"not null default '' comment('会员标识') VARCHAR(16)"`
	MemclassSelf int       `xorm:"not null default 000 comment('会员等级') TINYINT(3)"`
	MobileAgent  string    `xorm:"not null default '' comment('代理的手机号') VARCHAR(16)"`
	MobileSelf   string    `xorm:"not null default '' comment('新注册的用户手机号') unique VARCHAR(16)"`
	UpdateTime   time.Time `xorm:"not null comment('更新时间') DATETIME"`
	UuidAgent    string    `xorm:"not null default '' comment('代理的ID号') VARCHAR(16)"`
	UuidSelf     string    `xorm:"not null pk default '' comment('新注册的用户ID号') VARCHAR(16)"`
}
