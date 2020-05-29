package model

// swagger:dao
type (
	MessageType int
)

const (
	SiteBillName   = "费用账单"
	GlobalBillName = "结算账单"
)

const (
	DateFormatAll         = "2006-01-02"
	DateFormatYearMonth   = "2006年01月"
	DateFormatYearMonthMC = "2006-01"
)

type IEnum interface {
	Val() int
}

const (
	_             MessageType = iota
	Advertisement             //广告	1
	Information               //资讯	2
)
