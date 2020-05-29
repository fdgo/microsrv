package model

import "time"

type DsMessage struct {
	ID                int64 `gorm:"primary_key"`
	CreatedAt         *time.Time
	UpdatedAt         *time.Time
	DeletedAt         *time.Time
	Swith             int    `gorm:" comment('开关: 1:开, 2:关') TINYINT(4)"`
	Name              string `gorm:" default '' comment('名称') VARCHAR(128)"`
	UserType          int    `gorm:" comment('类型: 1:医院, 2:管理员') TINYINT(4)"`
	UserID            string `gorm:" comment('用户类型是医院则填写 branch id') TINYINT(4)"`
	MessageType       int    `gorm:" comment('类型: 1:广告, 2:资讯(动态) ') TINYINT(4)"`
	Title             string `gorm:" comment('标题') VARCHAR(128)"`
	Picture           string `gorm:" comment('图片') VARCHAR(128)"`
	Content           string `gorm:" comment('内容') VARCHAR(128)"`
	Url               string
	DsMessageDetailId int64          `gorm:" comment('详情id') TINYINT(4)"`
	DsMessageUrl      []DsMessageUrl `gorm:"-"`
}

func (model DsMessage) TableName() string {
	return "Ds_Message"
}

type DsMessageDetail struct {
	ID        int64 `gorm:"primary_key"`
	CreatedAt *time.Time
	UpdatedAt *time.Time
	DeletedAt *time.Time
	MessageID int64  `gorm:"-"`
	Title     string `gorm:"-"`
	Name      string `gorm:"-"`
	Content   string `sql:"type:longtext;"`
}

func (model DsMessageDetail) TableName() string {
	return "ds_message_detail"
}

type DsMessageUrl struct {
	ID        int64 `gorm:"primary_key"`
	CreatedAt *time.Time
	UpdatedAt *time.Time
	DeletedAt *time.Time
	Url       string `gorm:" comment('内容') VARCHAR(128)"`
	Content   string `gorm:" comment('保留字段') VARCHAR(128)"`
	MessageID int64  `gorm:" comment('消息id') TINYINT(4)"`
}

func (model DsMessageUrl) TableName() string {
	return "ds_message_url"
}

type MessageSearchCondition struct {
	Swith       int
	Name        string
	MessageType string
	UserType    int
	UserID      string
}

type DsBannerUrl struct {
	ID         int64 `gorm:"primary_key"`
	CreatedAt  *time.Time
	UpdatedAt  *time.Time
	DeletedAt  *time.Time
	BannerType BannerType
	//图片链接
	Url string

	Content string `gorm:" comment('保留字段') VARCHAR(128)"`
	//当作h5链接 暂时不做跳转
	Detail string
}

func (model DsBannerUrl) TableName() string {
	return "ds_banner_url"
}
