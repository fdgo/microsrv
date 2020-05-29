package model

import "time"

// TClientChannelInfo 客户端渠道信息
type TClientChannelInfo struct {
	ID          int    `gorm:"primary_key"`
	ClientType  string `gorm:"unique_index:idx_client_channel"` // 客户端，iOS；Android ; web; h5; wechat;
	Channel     string `gorm:"unique_index:idx_client_channel"` // 渠道
	ProductName string `gorm:"unique_index:idx_client_channel"` // 产品唯一标识，以包名识别
	ProductType int                                             // 是否是合集：0-单包；1-合集
	CreatedAt   time.Time                                       // 创建时间
	UpdatedAt   time.Time                                       // 更新时间
	DeletedAt   time.Time
	IsDeleted   bool // 是否删除，1-删除; 0-未删除;
}
