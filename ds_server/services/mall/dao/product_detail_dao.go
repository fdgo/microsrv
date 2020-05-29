package dao

import (
	"ds_server/services/mall/model"
	"errors"

	"github.com/jinzhu/gorm"
)

type DsProductDetailDao struct {
	db *gorm.DB
}

func NewDsProductDetailDao(db *gorm.DB) *DsProductDetailDao {
	return &DsProductDetailDao{
		db: db,
	}
}

func (dao *DsProductDetailDao) Create(m *model.DsProductDetail) error {
	return dao.db.Create(m).Error
}

func (dao *DsProductDetailDao) Find(m *model.DsProductDetail) (result []model.DsProductDetail, err error) {
	err = dao.db.Find(&result, m).Error
	return
}

func (dao *DsProductDetailDao) FindOne(m *model.DsProductDetail) error {
	return dao.db.First(m, m).Error
}

func (dao *DsProductDetailDao) FindLast(m *model.DsProductDetail) error {
	return dao.db.Last(m, m).Error
}

func (dao *DsProductDetailDao) FindPage(m *model.DsProductDetail, rowbound model.RowBound) (result []model.DsProductDetail, count int, err error) {
	err = dao.db.Model(&model.DsProductDetail{}).Count(&count).Limit(rowbound.Limit).Offset(rowbound.Offset).Find(&result, m).Error
	return
}

func (dao *DsProductDetailDao) Get(m *model.DsProductDetail) error {
	if dao.db.NewRecord(m) {
		return errors.New("id is nil")
	}
	return dao.db.Find(m).Error
}

func (dao *DsProductDetailDao) BatchGet(idbatch []int64) (result []model.DsProductDetail, err error) {
	err = dao.db.Model(&model.DsProductDetail{}).Where("ID in (?)", idbatch).Find(&result).Error
	return
}

func (dao *DsProductDetailDao) GetForUpdate(m *model.DsProductDetail) error {
	if dao.db.NewRecord(m) {
		return errors.New("id is nil")
	}
	return dao.db.Set("gorm:query_option", "FOR UPDATE").Find(m).Error
}

func (dao *DsProductDetailDao) Save(m *model.DsProductDetail) error {
	return dao.db.Save(m).Error
}

func (dao *DsProductDetailDao) Delete(m *model.DsProductDetail) error {
	if dao.db.NewRecord(m) {
		return errors.New("id is nil")
	}
	return dao.db.Delete(m).Error
}

func (dao *DsProductDetailDao) BatchDelete(idbatch []int64) error {
	return dao.db.Where("ID in (?)", idbatch).Delete(&model.DsProductDetail{}).Error
}

func (dao *DsProductDetailDao) Updates(id int64, attrs map[string]interface{}) error {
	return dao.db.Model(&model.DsProductDetail{}).Where("ID = ?", id).Updates(attrs).Error
}

func (dao *DsProductDetailDao) Update(id int64, attr string, value interface{}) error {
	return dao.db.Model(&model.DsProductDetail{}).Where("ID = ?", id).Update(attr, value).Error
}

func (dao *DsProductDetailDao) BatchUpdaterAttrs(idbatch []int64, attrs map[string]interface{}) error {
	return dao.db.Model(&model.DsProductDetail{}).Where("ID in (?)", idbatch).Updates(attrs).Error
}

func (dao *DsProductDetailDao) Found(m *model.DsProductDetail) bool {
	find := dao.db.First(m, m).RecordNotFound()
	return !find
}
