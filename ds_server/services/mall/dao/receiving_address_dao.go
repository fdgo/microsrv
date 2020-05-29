package dao

import (
	"ds_server/services/mall/model"
	"errors"

	"github.com/jinzhu/gorm"
)

type DsReceivingAddressDao struct {
	db *gorm.DB
}

func NewDsReceivingAddressDao(db *gorm.DB) *DsReceivingAddressDao {
	return &DsReceivingAddressDao{
		db: db,
	}
}

func (dao *DsReceivingAddressDao) Create(m *model.DsReceivingAddress) error {
	return dao.db.Create(m).Error
}

func (dao *DsReceivingAddressDao) Find(m *model.DsReceivingAddress) (result []model.DsReceivingAddress, err error) {
	err = dao.db.Find(&result, m).Error
	return
}

func (dao *DsReceivingAddressDao) FindOne(m *model.DsReceivingAddress) error {
	return dao.db.First(m, m).Error
}

func (dao *DsReceivingAddressDao) FindLast(m *model.DsReceivingAddress) error {
	return dao.db.Last(m, m).Error
}

func (dao *DsReceivingAddressDao) FindPage(m *model.DsReceivingAddress, rowbound model.RowBound) (result []model.DsReceivingAddress, count int, err error) {
	err = dao.db.Model(&model.DsReceivingAddress{}).Count(&count).Limit(rowbound.Limit).Offset(rowbound.Offset).Find(&result, m).Error
	return
}

func (dao *DsReceivingAddressDao) Get(m *model.DsReceivingAddress) error {
	if dao.db.NewRecord(m) {
		return errors.New("id is nil")
	}
	return dao.db.Find(m).Error
}

func (dao *DsReceivingAddressDao) BatchGet(idbatch []int64) (result []model.DsReceivingAddress, err error) {
	err = dao.db.Model(&model.DsReceivingAddress{}).Where("ID in (?)", idbatch).Find(&result).Error
	return
}

func (dao *DsReceivingAddressDao) GetForUpdate(m *model.DsReceivingAddress) error {
	if dao.db.NewRecord(m) {
		return errors.New("id is nil")
	}
	return dao.db.Set("gorm:query_option", "FOR UPDATE").Find(m).Error
}

func (dao *DsReceivingAddressDao) Save(m *model.DsReceivingAddress) error {
	return dao.db.Save(m).Error
}

func (dao *DsReceivingAddressDao) Delete(m *model.DsReceivingAddress) error {
	if dao.db.NewRecord(m) {
		return errors.New("id is nil")
	}
	return dao.db.Delete(m).Error
}

func (dao *DsReceivingAddressDao) BatchDelete(idbatch []int64) error {
	return dao.db.Where("ID in (?)", idbatch).Delete(&model.DsReceivingAddress{}).Error
}

func (dao *DsReceivingAddressDao) Updates(id int64, attrs map[string]interface{}) error {
	return dao.db.Model(&model.DsReceivingAddress{}).Where("ID = ?", id).Updates(attrs).Error
}

func (dao *DsReceivingAddressDao) Update(id int64, attr string, value interface{}) error {
	return dao.db.Model(&model.DsReceivingAddress{}).Where("ID = ?", id).Update(attr, value).Error
}

func (dao *DsReceivingAddressDao) BatchUpdaterAttrs(idbatch []int64, attrs map[string]interface{}) error {
	return dao.db.Model(&model.DsReceivingAddress{}).Where("ID in (?)", idbatch).Updates(attrs).Error
}

func (dao *DsReceivingAddressDao) Found(m *model.DsReceivingAddress) bool {
	find := dao.db.First(m, m).RecordNotFound()
	return !find
}
