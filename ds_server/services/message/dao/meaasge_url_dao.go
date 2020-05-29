package dao

import (
	"ds_server/services/message/model"
	"errors"

	"github.com/jinzhu/gorm"
)

type DsMessageUrlDao struct {
	db *gorm.DB
}

func NewDsMessageUrlDao(db *gorm.DB) *DsMessageUrlDao {
	return &DsMessageUrlDao{
		db: db,
	}
}

func (dao *DsMessageUrlDao) Create(m *model.DsMessageUrl) error {
	return dao.db.Create(m).Error
}

func (dao *DsMessageUrlDao) Find(m *model.DsMessageUrl) (result []model.DsMessageUrl, err error) {
	err = dao.db.Find(&result, m).Error
	return
}

func (dao *DsMessageUrlDao) FindOne(m *model.DsMessageUrl) error {
	return dao.db.First(m, m).Error
}

func (dao *DsMessageUrlDao) FindLast(m *model.DsMessageUrl) error {
	return dao.db.Last(m, m).Error
}

func (dao *DsMessageUrlDao) FindPage(m *model.DsMessageUrl, rowbound model.RowBound) (result []model.DsMessageUrl, count int, err error) {
	err = dao.db.Model(&model.DsMessageUrl{}).Count(&count).Limit(rowbound.Limit).Offset(rowbound.Offset).Find(&result, m).Error
	return
}

func (dao *DsMessageUrlDao) Get(m *model.DsMessageUrl) error {
	if dao.db.NewRecord(m) {
		return errors.New("id is nil")
	}
	return dao.db.Find(m).Error
}

func (dao *DsMessageUrlDao) BatchGet(idbatch []int64) (result []model.DsMessageUrl, err error) {
	err = dao.db.Model(&model.DsMessageUrl{}).Where("ID in (?)", idbatch).Find(&result).Error
	return
}

func (dao *DsMessageUrlDao) GetForUpdate(m *model.DsMessageUrl) error {
	if dao.db.NewRecord(m) {
		return errors.New("id is nil")
	}
	return dao.db.Set("gorm:query_option", "FOR UPDATE").Find(m).Error
}

func (dao *DsMessageUrlDao) Save(m *model.DsMessageUrl) error {
	return dao.db.Save(m).Error
}

func (dao *DsMessageUrlDao) Delete(m *model.DsMessageUrl) error {
	if dao.db.NewRecord(m) {
		return errors.New("id is nil")
	}
	return dao.db.Delete(m).Error
}

func (dao *DsMessageUrlDao) BatchDelete(idbatch []int64) error {
	return dao.db.Where("ID in (?)", idbatch).Delete(&model.DsMessageUrl{}).Error
}

func (dao *DsMessageUrlDao) Updates(id int64, attrs map[string]interface{}) error {
	return dao.db.Model(&model.DsMessageUrl{}).Where("ID = ?", id).Updates(attrs).Error
}

func (dao *DsMessageUrlDao) Update(id int64, attr string, value interface{}) error {
	return dao.db.Model(&model.DsMessageUrl{}).Where("ID = ?", id).Update(attr, value).Error
}

func (dao *DsMessageUrlDao) BatchUpdaterAttrs(idbatch []int64, attrs map[string]interface{}) error {
	return dao.db.Model(&model.DsMessageUrl{}).Where("ID in (?)", idbatch).Updates(attrs).Error
}

func (dao *DsMessageUrlDao) Found(m *model.DsMessageUrl) bool {
	find := dao.db.First(m, m).RecordNotFound()
	return !find
}
