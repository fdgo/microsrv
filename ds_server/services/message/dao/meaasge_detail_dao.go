package dao

import (
	"ds_server/services/message/model"
	"errors"

	"github.com/jinzhu/gorm"
)

type DsMessageDetailDao struct {
	db *gorm.DB
}

func NewDsMessageDetailDao(db *gorm.DB) *DsMessageDetailDao {
	return &DsMessageDetailDao{
		db: db,
	}
}

func (dao *DsMessageDetailDao) Create(m *model.DsMessageDetail) error {
	return dao.db.Create(m).Error
}

func (dao *DsMessageDetailDao) Find(m *model.DsMessageDetail) (result []model.DsMessageDetail, err error) {
	err = dao.db.Find(&result, m).Error
	return
}

func (dao *DsMessageDetailDao) FindOne(m *model.DsMessageDetail) error {
	return dao.db.First(m, m).Error
}

func (dao *DsMessageDetailDao) FindLast(m *model.DsMessageDetail) error {
	return dao.db.Last(m, m).Error
}

func (dao *DsMessageDetailDao) FindPage(m *model.DsMessageDetail, rowbound model.RowBound) (result []model.DsMessageDetail, count int, err error) {
	err = dao.db.Model(&model.DsMessageDetail{}).Count(&count).Limit(rowbound.Limit).Offset(rowbound.Offset).Find(&result, m).Error
	return
}

func (dao *DsMessageDetailDao) Get(m *model.DsMessageDetail) error {
	if dao.db.NewRecord(m) {
		return errors.New("id is nil")
	}
	return dao.db.Find(m).Error
}

func (dao *DsMessageDetailDao) BatchGet(idbatch []int64) (result []model.DsMessageDetail, err error) {
	err = dao.db.Model(&model.DsMessageDetail{}).Where("ID in (?)", idbatch).Find(&result).Error
	return
}

func (dao *DsMessageDetailDao) GetForUpdate(m *model.DsMessageDetail) error {
	if dao.db.NewRecord(m) {
		return errors.New("id is nil")
	}
	return dao.db.Set("gorm:query_option", "FOR UPDATE").Find(m).Error
}

func (dao *DsMessageDetailDao) Save(m *model.DsMessageDetail) error {
	return dao.db.Save(m).Error
}

func (dao *DsMessageDetailDao) Delete(m *model.DsMessageDetail) error {
	if dao.db.NewRecord(m) {
		return errors.New("id is nil")
	}
	return dao.db.Delete(m).Error
}

func (dao *DsMessageDetailDao) BatchDelete(idbatch []int64) error {
	return dao.db.Where("ID in (?)", idbatch).Delete(&model.DsMessageDetail{}).Error
}

func (dao *DsMessageDetailDao) Updates(id int64, attrs map[string]interface{}) error {
	return dao.db.Model(&model.DsMessageDetail{}).Where("ID = ?", id).Updates(attrs).Error
}

func (dao *DsMessageDetailDao) Update(id int64, attr string, value interface{}) error {
	return dao.db.Model(&model.DsMessageDetail{}).Where("ID = ?", id).Update(attr, value).Error
}

func (dao *DsMessageDetailDao) BatchUpdaterAttrs(idbatch []int64, attrs map[string]interface{}) error {
	return dao.db.Model(&model.DsMessageDetail{}).Where("ID in (?)", idbatch).Updates(attrs).Error
}

func (dao *DsMessageDetailDao) Found(m *model.DsMessageDetail) bool {
	find := dao.db.First(m, m).RecordNotFound()
	return !find
}
