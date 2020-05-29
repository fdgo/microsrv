package dao

import (
	"ds_server/services/mall/model"
	"errors"

	"github.com/jinzhu/gorm"
)

type DsProductStoreChangeDao struct {
	db *gorm.DB
}

func NewDsProductStoreChangeDao(db *gorm.DB) *DsProductStoreChangeDao {
	return &DsProductStoreChangeDao{
		db: db,
	}
}

func (dao *DsProductStoreChangeDao) Create(m *model.DsProductStoreChange) error {
	return dao.db.Create(m).Error
}

func (dao *DsProductStoreChangeDao) Find(m *model.DsProductStoreChange) (result []model.DsProductStoreChange, err error) {
	err = dao.db.Find(&result, m).Error
	return
}

func (dao *DsProductStoreChangeDao) FindOne(m *model.DsProductStoreChange) error {
	return dao.db.First(m, m).Error
}

func (dao *DsProductStoreChangeDao) FindLast(m *model.DsProductStoreChange) error {
	return dao.db.Last(m, m).Error
}

func (dao *DsProductStoreChangeDao) FindPage(m *model.DsProductStoreChange, rowbound model.RowBound) (result []model.DsProductStoreChange, count int, err error) {
	err = dao.db.Model(&model.DsProductStoreChange{}).Count(&count).Limit(rowbound.Limit).Offset(rowbound.Offset).Find(&result, m).Error
	return
}

func (dao *DsProductStoreChangeDao) Get(m *model.DsProductStoreChange) error {
	if dao.db.NewRecord(m) {
		return errors.New("id is nil")
	}
	return dao.db.Find(m).Error
}

func (dao *DsProductStoreChangeDao) BatchGet(idbatch []int64) (result []model.DsProductStoreChange, err error) {
	err = dao.db.Model(&model.DsProductStoreChange{}).Where("ID in (?)", idbatch).Find(&result).Error
	return
}

func (dao *DsProductStoreChangeDao) GetForUpdate(m *model.DsProductStoreChange) error {
	if dao.db.NewRecord(m) {
		return errors.New("id is nil")
	}
	return dao.db.Set("gorm:query_option", "FOR UPDATE").Find(m).Error
}

func (dao *DsProductStoreChangeDao) Save(m *model.DsProductStoreChange) error {
	return dao.db.Save(m).Error
}

func (dao *DsProductStoreChangeDao) Delete(m *model.DsProductStoreChange) error {
	if dao.db.NewRecord(m) {
		return errors.New("id is nil")
	}
	return dao.db.Delete(m).Error
}

func (dao *DsProductStoreChangeDao) BatchDelete(idbatch []int64) error {
	return dao.db.Where("ID in (?)", idbatch).Delete(&model.DsProductStoreChange{}).Error
}

func (dao *DsProductStoreChangeDao) Updates(id int64, attrs map[string]interface{}) error {
	return dao.db.Model(&model.DsProductStoreChange{}).Where("ID = ?", id).Updates(attrs).Error
}

func (dao *DsProductStoreChangeDao) Update(id int64, attr string, value interface{}) error {
	return dao.db.Model(&model.DsProductStoreChange{}).Where("ID = ?", id).Update(attr, value).Error
}

func (dao *DsProductStoreChangeDao) BatchUpdaterAttrs(idbatch []int64, attrs map[string]interface{}) error {
	return dao.db.Model(&model.DsProductStoreChange{}).Where("ID in (?)", idbatch).Updates(attrs).Error
}

func (dao *DsProductStoreChangeDao) Found(m *model.DsProductStoreChange) bool {
	find := dao.db.First(m, m).RecordNotFound()
	return !find
}
