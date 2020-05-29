package model

import "time"

// TGameTag 游戏标签关联表
type TGameTag struct {
	ID        int `gorm:"primary_key"`
	GameID    int       // 游戏id
	Name      string    // 标签名称
	Desc      string    // 标签描述
	Weight    int       // 权重，越大排名越前
	CreatedAt time.Time // 创建时间
	UpdatedAt time.Time // 更新时间
	DeletedAt time.Time
	IsDeleted bool // 是否删除，1-删除; 0-未删除;
}
