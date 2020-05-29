package model

import "time"

// TGameImage 游戏图片
type TGameImage struct {
	ID           int `gorm:"primary_key"`
	GameID       int `gorm:"index"` // 游戏ID
	ImgType      int                // 1-图片；2-gif；3-视频
	PositionType int                // 1-封面；2-loading；3-icon
	URL          string             // URL
	Desc         string             // 描述
	CreatedAt    time.Time          // 创建时间
	UpdatedAt    time.Time          // 更新时间
	DeletedAt    time.Time
	IsDeleted    bool // 是否删除，1-删除; 0-未删除;
}
