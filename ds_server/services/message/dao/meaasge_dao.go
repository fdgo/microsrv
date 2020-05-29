package dao

import (
	"ds_server/services/message/model"
	"errors"

	"github.com/jinzhu/gorm"
)

type DsMessageDao struct {
	db *gorm.DB
}

func NewDsMessageDao(db *gorm.DB) *DsMessageDao {
	return &DsMessageDao{
		db: db,
	}
}

func (dao *DsMessageDao) Create(m *model.DsMessage) error {
	return dao.db.Create(m).Error
}

func (dao *DsMessageDao) Find(m *model.DsMessage) (result []model.DsMessage, err error) {
	err = dao.db.Find(&result, m).Error
	return
}

func (dao *DsMessageDao) FindOne(m *model.DsMessage) error {
	return dao.db.First(m, m).Error
}

func (dao *DsMessageDao) FindLast(m *model.DsMessage) error {
	return dao.db.Last(m, m).Error
}

func (dao *DsMessageDao) FindPage(m *model.DsMessage, rowbound model.RowBound) (result []model.DsMessage, count int, err error) {
	err = dao.db.Model(&model.DsMessage{}).Count(&count).Limit(rowbound.Limit).Offset(rowbound.Offset).Find(&result, m).Error
	return
}

func (dao *DsMessageDao) Get(m *model.DsMessage) error {
	if dao.db.NewRecord(m) {
		return errors.New("id is nil")
	}
	return dao.db.Find(m).Error
}

func (dao *DsMessageDao) BatchGet(idbatch []int64) (result []model.DsMessage, err error) {
	err = dao.db.Model(&model.DsMessage{}).Where("ID in (?)", idbatch).Find(&result).Error
	return
}

func (dao *DsMessageDao) GetForUpdate(m *model.DsMessage) error {
	if dao.db.NewRecord(m) {
		return errors.New("id is nil")
	}
	return dao.db.Set("gorm:query_option", "FOR UPDATE").Find(m).Error
}

func (dao *DsMessageDao) Save(m *model.DsMessage) error {
	return dao.db.Save(m).Error
}

func (dao *DsMessageDao) Delete(m *model.DsMessage) error {
	if dao.db.NewRecord(m) {
		return errors.New("id is nil")
	}
	return dao.db.Delete(m).Error
}

func (dao *DsMessageDao) BatchDelete(idbatch []int64) error {
	return dao.db.Where("ID in (?)", idbatch).Delete(&model.DsMessage{}).Error
}

func (dao *DsMessageDao) Updates(id int64, attrs map[string]interface{}) error {
	return dao.db.Model(&model.DsMessage{}).Where("ID = ?", id).Updates(attrs).Error
}

func (dao *DsMessageDao) Update(id int64, attr string, value interface{}) error {
	return dao.db.Model(&model.DsMessage{}).Where("ID = ?", id).Update(attr, value).Error
}

func (dao *DsMessageDao) BatchUpdaterAttrs(idbatch []int64, attrs map[string]interface{}) error {
	return dao.db.Model(&model.DsMessage{}).Where("ID in (?)", idbatch).Updates(attrs).Error
}

func (dao *DsMessageDao) Found(m *model.DsMessage) bool {
	find := dao.db.First(m, m).RecordNotFound()
	return !find
}
