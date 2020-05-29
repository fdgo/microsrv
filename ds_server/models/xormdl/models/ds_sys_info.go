package models

type DsSysInfo struct {
	ConnectUs string `xorm:"not null default '' comment('联系我们') VARCHAR(255)"`
}
