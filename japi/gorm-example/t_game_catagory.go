package model

import "time"

// TGameCatagory 游戏分类
type TGameCatagory struct {
	ID        int `gorm:"primary_key"`
	ParentID  int                    // 父类ID
	Name      string `gorm:"unique"` // 分类名称
	Desc      string                 // 分类描述
	Weight    int                    // 权重，越大排名越前
	CreatedAt time.Time              // 创建时间
	UpdatedAt time.Time              // 更新时间
	DeletedAt time.Time
	IsDeleted bool // 是否删除，1-删除; 0-未删除;
}
