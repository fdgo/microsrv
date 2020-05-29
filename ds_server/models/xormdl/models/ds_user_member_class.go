package models

type DsUserMemberClass struct {
	MemMoney string `xorm:"not null comment('会员等级金额界限') index DECIMAL(18,2)"`
	MemName  string `xorm:"not null comment('会员等级名称  普通卡，铜卡，银卡，金卡，钻石卡') VARCHAR(16)"`
	MemTag   int    `xorm:"not null comment('会员等级1,2,3,4,5') index TINYINT(4)"`
	MemTagex string `xorm:"not null comment('会员等级标记  M1，M2，M3，M4，M5') VARCHAR(16)"`
}
