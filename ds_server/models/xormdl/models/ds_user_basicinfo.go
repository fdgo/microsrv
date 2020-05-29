package models

import (
	"time"
)

type DsUserBasicinfo struct {
	Age           int       `xorm:"not null default 000 comment('年龄') TINYINT(3)"`
	Avatar        string    `xorm:"not null default '' comment('头像地址') VARCHAR(255)"`
	Birthday      time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' comment('生日') DATETIME"`
	CreateTime    time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' comment('创建时间') DATETIME"`
	Deleted       int       `xorm:"not null comment('是否删除0:未删除 1:删除') TINYINT(4)"`
	Gender        int       `xorm:"not null default 000 comment('性别：0:男，1:女') TINYINT(3)"`
	Hash          string    `xorm:"not null default '' comment('密码hash') VARCHAR(64)"`
	LastLoginIp   string    `xorm:"not null default '' comment('最后一次登录ip') VARCHAR(32)"`
	LastLoginTime time.Time `xorm:"not null comment('最后一次登录时间') DATETIME"`
	Mobile        string    `xorm:"not null default '' comment('手机号') unique VARCHAR(16)"`
	NickName      string    `xorm:"not null default '' comment('昵称') VARCHAR(16)"`
	RealName      string    `xorm:"not null default '' comment('真实姓名') VARCHAR(16)"`
	Salt          string    `xorm:"not null default '' comment('密码盐') VARCHAR(16)"`
	Status        int       `xorm:"not null default 000 comment('账户状态0:正常 1:禁用  2:注销') TINYINT(3)"`
	UpdateTime    time.Time `xorm:"not null comment('更新时间') DATETIME"`
	Uuid          string    `xorm:"not null pk default '' comment('用户ID号') VARCHAR(16)"`
}
