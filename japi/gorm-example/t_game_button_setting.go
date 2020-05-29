package model

import "time"

// TGameButtonSetting 用户游戏按键设置,gameid=0 通用，userid=0 该游戏默认按键
type TGameButtonSetting struct {
	ID         int `gorm:"primary_key"`
	UserID     int `gorm:"index:idx_game_user"` // 用户id
	GameID     int `gorm:"index:idx_game_user"` // 游戏id
	Status     int                              // 状态：0-未启用，1-已启用
	DeviceType int                              // 设备类型：1-Android；2-iOS;3-Web
	Config     string                           // 游戏app按键设置，json结构
	Height     int                              // 分辨率高
	Width      int                              // 分辨率宽
	CreatedAt  time.Time                        // 创建时间
	UpdatedAt  time.Time                        // 更新时间
	DeletedAt  time.Time
	IsDeleted  bool // 是否删除，1-删除; 0-未删除;
}
