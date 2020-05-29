package dao

import (
	"ds_server/services/mall/model"
	"errors"

	"github.com/jinzhu/gorm"
)

type DsProductStoreDao struct {
	db *gorm.DB
}

func NewDsProductStoreDao(db *gorm.DB) *DsProductStoreDao {
	return &DsProductStoreDao{
		db: db,
	}
}

func (dao *DsProductStoreDao) Create(m *model.DsProductStore) error {
	return dao.db.Create(m).Error
}

func (dao *DsProductStoreDao) Find(m *model.DsProductStore) (result []model.DsProductStore, err error) {
	err = dao.db.Find(&result, m).Error
	return
}

func (dao *DsProductStoreDao) FindOne(m *model.DsProductStore) error {
	return dao.db.First(m, m).Error
}

func (dao *DsProductStoreDao) FindLast(m *model.DsProductStore) error {
	return dao.db.Last(m, m).Error
}

func (dao *DsProductStoreDao) FindPage(m *model.DsProductStore, rowbound model.RowBound) (result []model.DsProductStore, count int, err error) {
	err = dao.db.Model(&model.DsProductStore{}).Count(&count).Limit(rowbound.Limit).Offset(rowbound.Offset).Find(&result, m).Error
	return
}

func (dao *DsProductStoreDao) Get(m *model.DsProductStore) error {
	if dao.db.NewRecord(m) {
		return errors.New("id is nil")
	}
	return dao.db.Find(m).Error
}

func (dao *DsProductStoreDao) BatchGet(idbatch []int64) (result []model.DsProductStore, err error) {
	err = dao.db.Model(&model.DsProductStore{}).Where("ID in (?)", idbatch).Find(&result).Error
	return
}

func (dao *DsProductStoreDao) GetForUpdate(m *model.DsProductStore) error {
	if dao.db.NewRecord(m) {
		return errors.New("id is nil")
	}
	return dao.db.Set("gorm:query_option", "FOR UPDATE").Find(m).Error
}

func (dao *DsProductStoreDao) Save(m *model.DsProductStore) error {
	return dao.db.Save(m).Error
}

func (dao *DsProductStoreDao) Delete(m *model.DsProductStore) error {
	if dao.db.NewRecord(m) {
		return errors.New("id is nil")
	}
	return dao.db.Delete(m).Error
}

func (dao *DsProductStoreDao) BatchDelete(idbatch []int64) error {
	return dao.db.Where("ID in (?)", idbatch).Delete(&model.DsProductStore{}).Error
}

func (dao *DsProductStoreDao) Updates(id int64, attrs map[string]interface{}) error {
	return dao.db.Model(&model.DsProductStore{}).Where("ID = ?", id).Updates(attrs).Error
}

func (dao *DsProductStoreDao) Update(id int64, attr string, value interface{}) error {
	return dao.db.Model(&model.DsProductStore{}).Where("ID = ?", id).Update(attr, value).Error
}

func (dao *DsProductStoreDao) BatchUpdaterAttrs(idbatch []int64, attrs map[string]interface{}) error {
	return dao.db.Model(&model.DsProductStore{}).Where("ID in (?)", idbatch).Updates(attrs).Error
}

func (dao *DsProductStoreDao) Found(m *model.DsProductStore) bool {
	find := dao.db.First(m, m).RecordNotFound()
	return !find
}
