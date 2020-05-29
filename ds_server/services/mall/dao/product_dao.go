package dao

import (
	"ds_server/services/mall/model"
	"errors"

	"github.com/jinzhu/gorm"
)

type DsProductDao struct {
	db *gorm.DB
}

func NewDsProductDao(db *gorm.DB) *DsProductDao {
	return &DsProductDao{
		db: db,
	}
}

func (dao *DsProductDao) Create(m *model.DsProduct) error {
	return dao.db.Create(m).Error
}

func (dao *DsProductDao) Find(m *model.DsProduct) (result []model.DsProduct, err error) {
	err = dao.db.Find(&result, m).Error
	return
}

func (dao *DsProductDao) FindOne(m *model.DsProduct) error {
	return dao.db.First(m, m).Error
}

func (dao *DsProductDao) FindLast(m *model.DsProduct) error {
	return dao.db.Last(m, m).Error
}

func (dao *DsProductDao) FindPage(m *model.DsProduct, rowbound model.RowBound) (result []model.DsProduct, count int, err error) {
	err = dao.db.Model(&model.DsProduct{}).Count(&count).Limit(rowbound.Limit).Offset(rowbound.Offset).Find(&result, m).Error
	return
}

func (dao *DsProductDao) Get(m *model.DsProduct) error {
	if dao.db.NewRecord(m) {
		return errors.New("id is nil")
	}
	return dao.db.Find(m).Error
}

func (dao *DsProductDao) BatchGet(idbatch []int64) (result []model.DsProduct, err error) {
	err = dao.db.Model(&model.DsProduct{}).Where("ID in (?)", idbatch).Find(&result).Error
	return
}

func (dao *DsProductDao) GetForUpdate(m *model.DsProduct) error {
	if dao.db.NewRecord(m) {
		return errors.New("id is nil")
	}
	return dao.db.Set("gorm:query_option", "FOR UPDATE").Find(m).Error
}

func (dao *DsProductDao) Save(m *model.DsProduct) error {
	return dao.db.Save(m).Error
}

func (dao *DsProductDao) Delete(m *model.DsProduct) error {
	if dao.db.NewRecord(m) {
		return errors.New("id is nil")
	}
	return dao.db.Delete(m).Error
}

func (dao *DsProductDao) BatchDelete(idbatch []int64) error {
	return dao.db.Where("ID in (?)", idbatch).Delete(&model.DsProduct{}).Error
}

func (dao *DsProductDao) Updates(id int64, attrs map[string]interface{}) error {
	return dao.db.Model(&model.DsProduct{}).Where("ID = ?", id).Updates(attrs).Error
}

func (dao *DsProductDao) Update(id int64, attr string, value interface{}) error {
	return dao.db.Model(&model.DsProduct{}).Where("ID = ?", id).Update(attr, value).Error
}

func (dao *DsProductDao) BatchUpdaterAttrs(idbatch []int64, attrs map[string]interface{}) error {
	return dao.db.Model(&model.DsProduct{}).Where("ID in (?)", idbatch).Updates(attrs).Error
}

func (dao *DsProductDao) Found(m *model.DsProduct) bool {
	find := dao.db.First(m, m).RecordNotFound()
	return !find
}
