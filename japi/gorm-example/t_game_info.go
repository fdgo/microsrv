package model

import "time"

// TGameInfo 游戏信息表
type TGameInfo struct {
	ID                  int    `gorm:"primary_key"`
	Name                string `gorm:"unique"` // 游戏名称，全局唯一
	NameEn              string                 // 游戏英文名，比如SSFIV.
	NamePy              string                 // 游戏拼音，便于搜索
	Status              int                    // 1-正常；0-下线;2-待上线；
	Title               string                 // 游戏一句话介绍
	Summary             string                 // 游戏简介
	Desc                string                 // 游戏详情
	Company             string                 // 厂商
	Website             string                 // 官网
	PublishDate         time.Time              // 游戏上线时间
	Contact             string                 // 联系方式，格式，QQ:231111;email:xx@xx.com
	VideoWidth          int                    // 默认宽
	VideoHeight         int                    // 默认高
	VideoBitrate        int                    // 默认码率
	GameType            int                    // 游戏类型：单机、同屏联机、跨屏联机
	MaxPlayer           int                    // 支持玩家数
	CPULoad             int                    // cpu负载
	GpuLoad             int                    // gpu负载
	MemoryLoad          int                    // 内存负载
	ProfileType         int                    // 存档类型，默认0，不存档；1，普通存档。
	ControlType         int                    // 1，手柄；2，全键盘；4，鼠标；8, 自定义
	UseClientResolution bool                   // 是否使用客户端分辨率，默认0，不使用
	CreatedAt           time.Time              // 创建时间
	UpdatedAt           time.Time              // 更新时间
	DeletedAt           time.Time
	IsDeleted           bool // 是否删除，1-删除; 0-未删除;
}
