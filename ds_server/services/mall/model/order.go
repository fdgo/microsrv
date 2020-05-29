package model

import (
	"time"

	"github.com/shopspring/decimal"
)

type DsOrder struct {
	ID        int64 `gorm:"primary_key"`
	CreatedAt *time.Time
	UpdatedAt *time.Time
	//订单号（由ID生成的混淆字符串，可被还原为ID）
	OrderNo string
	//所属人ID
	UserID string `gorm:"index"`
	//所属人姓名
	UserName string
	//账号
	Account string
	//医院id
	BranchID int64
	//医院名称
	BranchName string
	//商品id
	ProductID int64
	//商品名称
	ProductName string
	//商品Url
	ProductUrl string
	//商品类型
	ProductType ProductType
	//商品下单数量
	ProductNum int
	//商品单价
	ProductPrice decimal.Decimal `sql:"type:decimal(20,4);"`
	//用户下单金额
	OrderPrice decimal.Decimal `sql:"type:decimal(20,4);"`
	//总金额
	TotalAmount decimal.Decimal `sql:"type:decimal(20,4);"`
	//订单状态（待处理/成功/失败）
	OrderState OrderState
	//订单IP
	IP string
	//订单IP地址
	IPAddr string
	//备注
	Remark string
	//操作人
	Operator string
	//下单地址
	AddressID int64
	//业务完成时间
	CompleteTime *time.Time
	//支付时间
	BeDeliveredTime *time.Time
	//收获时间
	ReceivedTime *time.Time
	//退款时间
	RefundTime *time.Time
	//退款完成时间
	ReComolete         *time.Time
	DsReceivingAddress DsReceivingAddress `gorm:"-"`
	DsOrderParameterID int64              `gorm:"index"`
	DsOrderParameter   DsOrderParameter   `gorm:"ForeignKey:DsOrderParameterID" gorm:"association_foreignkey:OrderID"`
}

func (model DsOrder) TableName() string {
	return "ds_order"
}

type DsOrderParameter struct {
	ID               int64 `gorm:"primary_key"`
	CreatedAt        *time.Time
	UpdatedAt        *time.Time
	OrderID          int64
	DailyWithdrawCnt int
}

func (model DsOrderParameter) TableName() string {
	return "ds_order_parameter"
}

type OrderSearchCondition struct {
	//用户id
	UserID string
	//玩家账号
	Account string
	//创建开始时间
	CreateStartTime time.Time
	//创建截止时间
	CreateEndTime time.Time
	//完成开始时间
	CompleteStartTime time.Time
	//完成截止时间
	CompleteEndTime time.Time
	//订单号
	OrderNo string
	//订单状态
	OrderState OrderState
	//取款Ip
	IP string
	//备注
	Remark string
}
