package dao

import (
	"fmt"
	"github.com/jinzhu/gorm"
	mygormdl "ds_server/models/user/gorm_mysql"
)

type DsSysInfoMgr struct {
	*_BaseMgr
}

// DsSysInfoMgr open func
func NewDsSysInfoMgr(db *gorm.DB) *DsSysInfoMgr {
	if db == nil {
		panic(fmt.Errorf("DsSysInfoMgr need init by db"))
	}
	return &DsSysInfoMgr{_BaseMgr: &_BaseMgr{DB: db, isRelated: globalIsRelated}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *DsSysInfoMgr) GetTableName() string {
	return "ds_sys_info"
}

// Get 获取
func (obj *DsSysInfoMgr) Get() (result mygormdl.DsSysInfo, err error) {
	err = obj.DB.Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *DsSysInfoMgr) Gets() (results []*mygormdl.DsSysInfo, err error) {
	err = obj.DB.Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithConnectUs connect_us获取 联系我们
func (obj *DsSysInfoMgr) WithConnectUs(connectUs string) Option {
	return optionFunc(func(o *options) { o.query["connect_us"] = connectUs })
}

// GetByOption 功能选项模式获取
func (obj *DsSysInfoMgr) GetByOption(opts ...Option) (result mygormdl.DsSysInfo, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.Table(obj.GetTableName()).Where(options.query).Find(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *DsSysInfoMgr) GetByOptions(opts ...Option) (results []*mygormdl.DsSysInfo, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.Table(obj.GetTableName()).Where(options.query).Find(&results).Error

	return
}

//////////////////////////enume case ////////////////////////////////////////////

// GetFromConnectUs 通过connect_us获取内容 联系我们
func (obj *DsSysInfoMgr) GetFromConnectUs(connectUs string) (results []*mygormdl.DsSysInfo, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("connect_us = ?", connectUs).Find(&results).Error

	return
}

// GetBatchFromConnectUs 批量唯一主键查找 联系我们
func (obj *DsSysInfoMgr) GetBatchFromConnectUs(connectUss []string) (results []*mygormdl.DsSysInfo, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("connect_us IN (?)", connectUss).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////
func (obj *DsSysInfoMgr) ConnectUs() (result mygormdl.DsSysInfo, err error) {
	err = obj.DB.Table(obj.GetTableName()).Find(&result).Error
	return
}
