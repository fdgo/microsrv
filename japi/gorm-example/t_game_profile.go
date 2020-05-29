package model

import "time"

// TGameProfile 游戏存档
type TGameProfile struct {
	ID         int `gorm:"primary_key"`
	UserID     int `gorm:"unique_index:idx_game_user"` // 用户id
	GameID     int `gorm:"unique_index:idx_game_user"` // 游戏id
	DeviceID   string                                  // 设备id
	GsIP       string                                  // gs_ip
	GpID       string                                  // gs_id
	Status     int                                     // 存档状态：1-上传成功，2-上传中，3-上传失败
	URL        string                                  // 存档链接
	UploadTime time.Time                               // 存档上传时间
	CreatedAt  time.Time                               // 创建时间
	UpdatedAt  time.Time                               // 更新时间
	DeletedAt  time.Time
	IsDeleted  bool // 是否删除，1-删除; 0-未删除;
}
