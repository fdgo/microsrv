package model

import "time"

// TClientVersion 客户端版本
type TClientVersion struct {
	ID              int `gorm:"primary_key"`
	ClientChannelID int `gorm:"index"` // 客户端渠道id
	Version         string             // 版本字符串
	VersionName     string             // 版本名称
	VersionCode     int                // 版本code，命名规则 年后两位+月+日+第几次更新，如19110100
	Desc            string             // 更新简介
	URL             string             // 下载链接
	IsValid         bool               // 是否启用，1-启用; 0-未启用;
	PublishTime     time.Time          // 启用时间
	UpdateType      int                // 更新方式，0-正常更新; 1-静默更新
	ForceUpdate     bool               // 是否强制更新，1-强制；0-非强制;
	CreatedAt       time.Time          // 创建时间
	UpdatedAt       time.Time          // 更新时间
	DeletedAt       time.Time
	IsDeleted       bool // 是否删除，1-删除; 0-未删除;
}
