package model

// swagger:dao
type (
	ProductType     int
	UrlType         int
	YesOrNo         int
	Sex             int
	Label           int
	StoreChangeType int
	StoreSourceType int
	OrderState      int
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
	_         ProductType = iota
	YxProduct             //严选商品	1
	MyProduct             //母婴商品	2
	KqProduct             //口腔商品 3
	ZyProduct             //中医商品
	GjProduct             //国际商品
)

const (
	_             UrlType = iota
	StaticUrlType         //静态页面列表展示	1
	BuyUrlType            //购买页面显示多个图片	2
)

const (
	_   YesOrNo = iota
	Yes         //	1
	No          //	2
)

const (
	_    Sex = iota
	Man      //	1
	Male     //	2
)

const (
	_       Label = iota
	Home          //	1 家
	Company       //	2 公司
	School        //3学校
)

const (
	_             StoreChangeType = iota
	StoreIncrease                 // 1
	StoreDecrease                 //账户余额减少 2
)

const (
	_                            StoreSourceType = iota
	StoreSourceTypeOrderDecrease                 //下单成功减小 存库
	StoreSourceTypeOrderIncrease                 //下单失败增加 存库
)

const (
	_                     OrderState = iota
	OrderStateBePaid                 //待付款
	OrderStateBeDelivered            //待发货
	OrderStateBeReceived             //待收货
	OrderStateRefund                 //退款
	OrderStateReComolete             //退款完成
	OrderStateComplete               //完成
	OrderStateColse                  //关闭
)
