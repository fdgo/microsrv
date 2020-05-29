package model

import "time"

// TGameCategoryRelation 游戏分类关联表
type TGameCategoryRelation struct {
	ID           int `gorm:"primary_key"`
	GameID       int `gorm:"index"` // 游戏id
	CategoryID   int `gorm:"index"` // 分类Id
	CatagoryName string             // 分类名称
	Weight       int                // 权重，越大排名越前
	CreatedAt    time.Time          // 创建时间
	UpdatedAt    time.Time          // 更新时间
	DeletedAt    time.Time
	IsDeleted    bool // 是否删除，1-删除; 0-未删除;
}
