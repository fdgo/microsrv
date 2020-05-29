package models

import (
	"time"
)

type DsUserMemberDepositHistory struct {
	AddressIn    string    `xorm:"not null default '' comment('收款地址') VARCHAR(255)"`
	AddressOut   string    `xorm:"not null default '' comment('支付地址') VARCHAR(255)"`
	Balance      string    `xorm:"not null default 0.00 comment('金额') DECIMAL(18,2)"`
	BalanceSrc   string    `xorm:"default 0.00 comment('原始金额') DECIMAL(18,2)"`
	CreateTime   time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' comment('创建时间') DATETIME"`
	DepositName  string    `xorm:"not null default '' comment('充值名字：购买商品，商品退款') VARCHAR(16)"`
	DepositType  int       `xorm:"not null default 000 comment('充值类型 0:扣款,1:充值') TINYINT(3)"`
	InvcodeAgent string    `xorm:"not null default '' comment('代理邀请码') index VARCHAR(16)"`
	InvcodeSelf  string    `xorm:"not null default '' comment('自身邀请码') VARCHAR(16)"`
	Mobile       string    `xorm:"not null default '' comment('用户手机号') VARCHAR(16)"`
	Rate         string    `xorm:"default 0.000000 comment('汇率') DECIMAL(18,6)"`
	SourceId     string    `xorm:"not null default '' comment('业务订单Id') index VARCHAR(64)"`
	Status       int       `xorm:"not null default 0000 comment('会员账户状态 0:正常，1:禁止 2:销户') TINYINT(4)"`
	UpdateTime   time.Time `xorm:"not null comment('更新时间') DATETIME"`
	Uuid         string    `xorm:"not null default '' comment('用户ID') VARCHAR(16)"`
}
