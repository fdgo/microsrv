package dao

import (
	"ds_server/services/message/model"
)

func (dao *DsBannerUrlDao) SeachBanner(bype model.BannerType) (result []model.DsBannerUrl, err error) {
	db := dao.db
	if bype > 0 {
		db = db.Where("banner_type = ?", bype)
	}
	err = db.Model(&model.DsBannerUrl{}).Order("ID desc").Find(&result).Error

	return
}
