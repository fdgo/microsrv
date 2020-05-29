package dao

import (
	"ds_server/services/message/model"
	"errors"

	"github.com/jinzhu/gorm"
)

type DsBannerUrlDao struct {
	db *gorm.DB
}

func NewDsBannerUrlDao(db *gorm.DB) *DsBannerUrlDao {
	return &DsBannerUrlDao{
		db: db,
	}
}

func (dao *DsBannerUrlDao) Create(m *model.DsBannerUrl) error {
	return dao.db.Create(m).Error
}

func (dao *DsBannerUrlDao) Find(m *model.DsBannerUrl) (result []model.DsBannerUrl, err error) {
	err = dao.db.Find(&result, m).Error
	return
}

func (dao *DsBannerUrlDao) FindOne(m *model.DsBannerUrl) error {
	return dao.db.First(m, m).Error
}

func (dao *DsBannerUrlDao) FindLast(m *model.DsBannerUrl) error {
	return dao.db.Last(m, m).Error
}

func (dao *DsBannerUrlDao) FindPage(m *model.DsBannerUrl, rowbound model.RowBound) (result []model.DsBannerUrl, count int, err error) {
	err = dao.db.Model(&model.DsBannerUrl{}).Count(&count).Limit(rowbound.Limit).Offset(rowbound.Offset).Find(&result, m).Error
	return
}

func (dao *DsBannerUrlDao) Get(m *model.DsBannerUrl) error {
	if dao.db.NewRecord(m) {
		return errors.New("id is nil")
	}
	return dao.db.Find(m).Error
}

func (dao *DsBannerUrlDao) BatchGet(idbatch []int64) (result []model.DsBannerUrl, err error) {
	err = dao.db.Model(&model.DsBannerUrl{}).Where("ID in (?)", idbatch).Find(&result).Error
	return
}

func (dao *DsBannerUrlDao) GetForUpdate(m *model.DsBannerUrl) error {
	if dao.db.NewRecord(m) {
		return errors.New("id is nil")
	}
	return dao.db.Set("gorm:query_option", "FOR UPDATE").Find(m).Error
}

func (dao *DsBannerUrlDao) Save(m *model.DsBannerUrl) error {
	return dao.db.Save(m).Error
}

func (dao *DsBannerUrlDao) Delete(m *model.DsBannerUrl) error {
	if dao.db.NewRecord(m) {
		return errors.New("id is nil")
	}
	return dao.db.Delete(m).Error
}

func (dao *DsBannerUrlDao) BatchDelete(idbatch []int64) error {
	return dao.db.Where("ID in (?)", idbatch).Delete(&model.DsBannerUrl{}).Error
}

func (dao *DsBannerUrlDao) Updates(id int64, attrs map[string]interface{}) error {
	return dao.db.Model(&model.DsBannerUrl{}).Where("ID = ?", id).Updates(attrs).Error
}

func (dao *DsBannerUrlDao) Update(id int64, attr string, value interface{}) error {
	return dao.db.Model(&model.DsBannerUrl{}).Where("ID = ?", id).Update(attr, value).Error
}

func (dao *DsBannerUrlDao) BatchUpdaterAttrs(idbatch []int64, attrs map[string]interface{}) error {
	return dao.db.Model(&model.DsBannerUrl{}).Where("ID in (?)", idbatch).Updates(attrs).Error
}

func (dao *DsBannerUrlDao) Found(m *model.DsBannerUrl) bool {
	find := dao.db.First(m, m).RecordNotFound()
	return !find
}
