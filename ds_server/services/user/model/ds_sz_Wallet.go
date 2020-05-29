package model

import (
	"time"

	"github.com/shopspring/decimal"
)

type DsSzWallet struct {
	ID        int64
	CreatedAt *time.Time
	UpdatedAt *time.Time
	DeletedAt *time.Time
	//用户ID
	UserID string `gorm:"index"`
	//数字类型
	CoinName string `gorm:"index"`
	//数字类型
	CoinAddress string `gorm:"index"`
	//钱包余额
	Balance decimal.Decimal `sql:"type:decimal(20,4);"`
	//冻结余额
	Frozen decimal.Decimal `sql:"type:decimal(20,4);"`
	//充值总额
	DepositTotal decimal.Decimal `sql:"type:decimal(20,4);"`
	//取款总额
	WithdrawTotal decimal.Decimal `sql:"type:decimal(20,4);"`
}

func (model DsSzWallet) TableName() string {
	return "ds_sz_wallet"
}
