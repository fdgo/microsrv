package model

import (
	"time"

	"github.com/shopspring/decimal"
)

type DsProduct struct {
	ID          int64 `gorm:"primary_key"`
	CreatedAt   *time.Time
	UpdatedAt   *time.Time
	DeletedAt   *time.Time
	ProductName string
	//医院id
	BranchID int64
	//医院名字
	BranchName string
	//商品类型
	ProductType int
	//商品是否在售
	IsOnSale YesOrNo
	//是否推荐商品
	IsRecommend YesOrNo
	//库存id
	DsProductStoreID int64          `gorm:"index"`
	DsProductStore   DsProductStore `gorm:"ForeignKey:DsProductStoreID" gorm:"association_foreignkey:ProductID"`

	//商品详情id
	DsProductDetailID int64 `gorm:"index"`
	//商品详情
	DsProductDetail DsProductDetail `gorm:"-"`
	//库存
	//显示url
	Url string `gorm:"-"`
	//详情页面
	DsProductUrls []DsProductUrl `gorm:"-"`
	//距离
	Distance string `gorm:"-"`
}

func (model DsProduct) TableName() string {
	return "ds_product"
}

type DsProductStore struct {
	ID            int64 `gorm:"primary_key"`
	CreatedAt     *time.Time
	UpdatedAt     *time.Time
	DeletedAt     *time.Time
	StoreNum      int
	Price         decimal.Decimal `sql:"type:decimal(20,4);"`
	RealPrice     decimal.Decimal `sql:"type:decimal(20,4);"`
	StoreNumTotal int
	ProductID     int64
}

func (model DsProductStore) TableName() string {
	return "ds_product_store"
}

type DsProductStoreChange struct {
	ID              int64 `gorm:"primary_key"`
	CreatedAt       *time.Time
	UpdatedAt       *time.Time
	DeletedAt       *time.Time
	UserID          string `gorm:"index"`
	UserType        int
	ProductID       int64 `gorm:"index"`
	StoreChangeType StoreChangeType
	StoreSourceType StoreSourceType
	Number          int
	BeforeNumber    int
	AfterNumber     int
	FundSourceID    int64
}

func (model DsProductStoreChange) TableName() string {
	return "ds_product_store_change"
}

type DsProductUrl struct {
	ID        int64 `gorm:"primary_key"`
	CreatedAt *time.Time
	UpdatedAt *time.Time
	DeletedAt *time.Time
	UrlType   UrlType //1、商品展示页面（静态） 2、购买页面（显示多张）
	ProductID int64
	Url       string
}

func (model DsProductUrl) TableName() string {
	return "Ds_Product_Url"
}

type DsProductDetail struct {
	ID        int64 `gorm:"primary_key"`
	CreatedAt *time.Time
	UpdatedAt *time.Time
	DeletedAt *time.Time
	//商品id
	DsProductID int64
	Content     string `sql:"type:longtext;"`
}

func (model DsProductDetail) TableName() string {
	return "ds_product_detail"
}

type ProductSearchCondition struct {
	ProductName string
	BranchID    int64
	BranchName  string
	ProductType ProductType
	Remark      string
	IsOnSale    YesOrNo
	IsRecommend YesOrNo
}

type DsReceivingAddress struct {
	ID          int64 `gorm:"primary_key"`
	CreatedAt   *time.Time
	UpdatedAt   *time.Time
	DeletedAt   *time.Time
	UserID      string
	UserAccount string
	UserAlias   string
	Province    string
	City        string
	Aera        string
	Address     string
	Tel         string
	Sex         Sex
	Label       Label
	TelName     string
	IsDefault   YesOrNo
}

func (model DsReceivingAddress) TableName() string {
	return "ds_receiving_address"
}
