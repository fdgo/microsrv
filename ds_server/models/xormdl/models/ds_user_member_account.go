package models

import (
	"time"
)

type DsUserMemberAccount struct {
	AddressIn  string    `xorm:"not null default '' comment('收款地址') VARCHAR(255)"`
	AddressOut string    `xorm:"not null default '' comment('付款地址') VARCHAR(255)"`
	Balance    string    `xorm:"not null default 0.00 comment('总金额') DECIMAL(18,2)"`
	CreateTime time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' comment('创建时间') DATETIME"`
	Hash       string    `xorm:"not null default '' comment('密码hash') VARCHAR(64)"`
	Ispwd      int       `xorm:"not null comment('密码是否为空') TINYINT(3)"`
	Mobile     string    `xorm:"not null default '' comment('用户手机号') unique VARCHAR(16)"`
	PrivateKey string    `xorm:"not null default '' VARCHAR(255)"`
	Salt       string    `xorm:"not null default '' comment('支付盐') VARCHAR(16)"`
	Status     int       `xorm:"not null default 000 comment('账户状态0:正常 1:禁用  2:注销') TINYINT(3)"`
	UpdateTime time.Time `xorm:"not null comment('更新时间') DATETIME"`
	Uuid       string    `xorm:"not null pk default '' comment('用户ID') VARCHAR(16)"`
}
