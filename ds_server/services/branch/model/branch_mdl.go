package model

import "time"

type DsBranch struct {
	ID            int64 `gorm:"primary_key"`
	CreatedAt     *time.Time
	UpdatedAt     *time.Time
	DeletedAt     *time.Time
	Name          string `gorm:" default '' comment('名称') VARCHAR(128)"`
	Address       string `gorm:" default '' comment('名称') VARCHAR(128)"`
	Tel           string
	BusinessHours string
	Title         string `gorm:" comment('标题') VARCHAR(128)"`
	Content       string `gorm:" comment('内容') VARCHAR(128)"`
	Url           string `gorm:" comment('图片保存路径') VARCHAR(128)"`
	Latitude      string `gorm:" comment('经度') VARCHAR(128)"`
	Longitude     string `gorm:" comment('纬度') VARCHAR(128)"`
	GeoHashCode   string `gorm:" comment('GeoHashCode') VARCHAR(128)"`
	Juli          string `gorm:" comment('GeoHashCode') VARCHAR(128)"`
}

func (model DsBranch) TableName() string {
	return "Ds_Branch"
}

type DsBranchDynamic struct {
	ID        int64 `gorm:"primary_key"`
	CreatedAt *time.Time
	UpdatedAt *time.Time
	DeletedAt *time.Time
	BranchID  int64  `gorm:" comment('详情id') TINYINT(4)"`
	Title     string `gorm:" comment('标题') VARCHAR(128)"`
	Content   string `sql:"type:longtext;"`
	Urls      []DsBranchUrl
	Name      string
}

func (model DsBranchDynamic) TableName() string {
	return "Ds_Branch_Dynamic"
}

type DsBranchUrl struct {
	ID                int64 `gorm:"primary_key"`
	CreatedAt         *time.Time
	UpdatedAt         *time.Time
	DeletedAt         *time.Time
	Url               string `gorm:" comment('内容') VARCHAR(128)"`
	Content           string `gorm:" comment('保留字段') VARCHAR(128)"`
	DsBranchDynamicID int64  `gorm:" comment('消息id') TINYINT(4)"`
}

func (model DsBranchUrl) TableName() string {
	return "Ds_Branch_Url"
}

type BranchSearchCondition struct {
	Name        string
	Latitude    string //经度
	Longitude   string
	GeoHashCode string
}

type BranchDynamicSearchCondition struct {
	BranchID int64
}
