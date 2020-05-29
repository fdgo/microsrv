package model

import "time"

// TGameChannel 客户端渠道和游戏关联
type TGameChannel struct {
	ID              int `gorm:"primary_key"`
	ClientChannelID int `gorm:"unique_index:idx_client_channel"` // 客户端渠道id
	GameID          int `gorm:"unique_index:idx_client_channel"` // 游戏id
	ClientVersion   int                                          // 最低版本code
	Weight          int                                          // 权重，越大排名越前
	VideoWidth      int                                          // 默认宽
	VideoHeight     int                                          // 默认高
	VideoBitrate    int                                          // 默认码率
	CreatedAt       time.Time                                    // 创建时间
	UpdatedAt       time.Time                                    // 更新时间
	DeletedAt       time.Time
	IsDeleted       bool // 是否删除，1-删除; 0-未删除;
}
