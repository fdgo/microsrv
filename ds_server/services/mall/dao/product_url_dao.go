package dao

import (
	"ds_server/services/mall/model"
	"errors"

	"github.com/jinzhu/gorm"
)

type DsProductUrlDao struct {
	db *gorm.DB
}

func NewDsProductUrlDao(db *gorm.DB) *DsProductUrlDao {
	return &DsProductUrlDao{
		db: db,
	}
}

func (dao *DsProductUrlDao) Create(m *model.DsProductUrl) error {
	return dao.db.Create(m).Error
}

func (dao *DsProductUrlDao) Find(m *model.DsProductUrl) (result []model.DsProductUrl, err error) {
	err = dao.db.Find(&result, m).Error
	return
}

func (dao *DsProductUrlDao) FindOne(m *model.DsProductUrl) error {
	return dao.db.First(m, m).Error
}

func (dao *DsProductUrlDao) FindLast(m *model.DsProductUrl) error {
	return dao.db.Last(m, m).Error
}

func (dao *DsProductUrlDao) FindPage(m *model.DsProductUrl, rowbound model.RowBound) (result []model.DsProductUrl, count int, err error) {
	err = dao.db.Model(&model.DsProductUrl{}).Count(&count).Limit(rowbound.Limit).Offset(rowbound.Offset).Find(&result, m).Error
	return
}

func (dao *DsProductUrlDao) Get(m *model.DsProductUrl) error {
	if dao.db.NewRecord(m) {
		return errors.New("id is nil")
	}
	return dao.db.Find(m).Error
}

func (dao *DsProductUrlDao) BatchGet(idbatch []int64) (result []model.DsProductUrl, err error) {
	err = dao.db.Model(&model.DsProductUrl{}).Where("ID in (?)", idbatch).Find(&result).Error
	return
}

func (dao *DsProductUrlDao) GetForUpdate(m *model.DsProductUrl) error {
	if dao.db.NewRecord(m) {
		return errors.New("id is nil")
	}
	return dao.db.Set("gorm:query_option", "FOR UPDATE").Find(m).Error
}

func (dao *DsProductUrlDao) Save(m *model.DsProductUrl) error {
	return dao.db.Save(m).Error
}

func (dao *DsProductUrlDao) Delete(m *model.DsProductUrl) error {
	if dao.db.NewRecord(m) {
		return errors.New("id is nil")
	}
	return dao.db.Delete(m).Error
}

func (dao *DsProductUrlDao) BatchDelete(idbatch []int64) error {
	return dao.db.Where("ID in (?)", idbatch).Delete(&model.DsProductUrl{}).Error
}

func (dao *DsProductUrlDao) Updates(id int64, attrs map[string]interface{}) error {
	return dao.db.Model(&model.DsProductUrl{}).Where("ID = ?", id).Updates(attrs).Error
}

func (dao *DsProductUrlDao) Update(id int64, attr string, value interface{}) error {
	return dao.db.Model(&model.DsProductUrl{}).Where("ID = ?", id).Update(attr, value).Error
}

func (dao *DsProductUrlDao) BatchUpdaterAttrs(idbatch []int64, attrs map[string]interface{}) error {
	return dao.db.Model(&model.DsProductUrl{}).Where("ID in (?)", idbatch).Updates(attrs).Error
}

func (dao *DsProductUrlDao) Found(m *model.DsProductUrl) bool {
	find := dao.db.First(m, m).RecordNotFound()
	return !find
}
