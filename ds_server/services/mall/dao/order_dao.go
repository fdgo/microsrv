package dao

import (
	"ds_server/services/mall/model"
	"errors"

	"github.com/jinzhu/gorm"
)

type DsOrderDao struct {
	db *gorm.DB
}

func NewDsOrderDao(db *gorm.DB) *DsOrderDao {
	return &DsOrderDao{
		db: db,
	}
}

func (dao *DsOrderDao) Create(m *model.DsOrder) error {
	return dao.db.Create(m).Error
}

func (dao *DsOrderDao) Find(m *model.DsOrder) (result []model.DsOrder, err error) {
	err = dao.db.Find(&result, m).Error
	return
}

func (dao *DsOrderDao) FindOne(m *model.DsOrder) error {
	return dao.db.First(m, m).Error
}

func (dao *DsOrderDao) FindLast(m *model.DsOrder) error {
	return dao.db.Last(m, m).Error
}

func (dao *DsOrderDao) FindPage(m *model.DsOrder, rowbound model.RowBound) (result []model.DsOrder, count int, err error) {
	err = dao.db.Model(&model.DsOrder{}).Count(&count).Limit(rowbound.Limit).Offset(rowbound.Offset).Find(&result, m).Error
	return
}

func (dao *DsOrderDao) Get(m *model.DsOrder) error {
	if dao.db.NewRecord(m) {
		return errors.New("id is nil")
	}
	return dao.db.Find(m).Error
}

func (dao *DsOrderDao) BatchGet(idbatch []int64) (result []model.DsOrder, err error) {
	err = dao.db.Model(&model.DsOrder{}).Where("ID in (?)", idbatch).Find(&result).Error
	return
}

func (dao *DsOrderDao) GetForUpdate(m *model.DsOrder) error {
	if dao.db.NewRecord(m) {
		return errors.New("id is nil")
	}
	return dao.db.Set("gorm:query_option", "FOR UPDATE").Find(m).Error
}

func (dao *DsOrderDao) Save(m *model.DsOrder) error {
	return dao.db.Save(m).Error
}

func (dao *DsOrderDao) Delete(m *model.DsOrder) error {
	if dao.db.NewRecord(m) {
		return errors.New("id is nil")
	}
	return dao.db.Delete(m).Error
}

func (dao *DsOrderDao) BatchDelete(idbatch []int64) error {
	return dao.db.Where("ID in (?)", idbatch).Delete(&model.DsOrder{}).Error
}

func (dao *DsOrderDao) Updates(id int64, attrs map[string]interface{}) error {
	return dao.db.Model(&model.DsOrder{}).Where("ID = ?", id).Updates(attrs).Error
}

func (dao *DsOrderDao) Update(id int64, attr string, value interface{}) error {
	return dao.db.Model(&model.DsOrder{}).Where("ID = ?", id).Update(attr, value).Error
}

func (dao *DsOrderDao) BatchUpdaterAttrs(idbatch []int64, attrs map[string]interface{}) error {
	return dao.db.Model(&model.DsOrder{}).Where("ID in (?)", idbatch).Updates(attrs).Error
}

func (dao *DsOrderDao) Found(m *model.DsOrder) bool {
	find := dao.db.First(m, m).RecordNotFound()
	return !find
}
