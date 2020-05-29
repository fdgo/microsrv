package dao

import (
	"ds_server/services/branch/model"
	"errors"

	"github.com/jinzhu/gorm"
)

type DsBranchUrlDao struct {
	db *gorm.DB
}

func NewDsBranchUrlDao(db *gorm.DB) *DsBranchUrlDao {
	return &DsBranchUrlDao{
		db: db,
	}
}

func (dao *DsBranchUrlDao) Create(m *model.DsBranchUrl) error {
	return dao.db.Create(m).Error
}

func (dao *DsBranchUrlDao) Find(m *model.DsBranchUrl) (result []model.DsBranchUrl, err error) {
	err = dao.db.Find(&result, m).Error
	return
}

func (dao *DsBranchUrlDao) FindOne(m *model.DsBranchUrl) error {
	return dao.db.First(m, m).Error
}

func (dao *DsBranchUrlDao) FindLast(m *model.DsBranchUrl) error {
	return dao.db.Last(m, m).Error
}

func (dao *DsBranchUrlDao) FindPage(m *model.DsBranchUrl, rowbound model.RowBound) (result []model.DsBranchUrl, count int, err error) {
	err = dao.db.Model(&model.DsBranchUrl{}).Count(&count).Limit(rowbound.Limit).Offset(rowbound.Offset).Find(&result, m).Error
	return
}

func (dao *DsBranchUrlDao) Get(m *model.DsBranchUrl) error {
	if dao.db.NewRecord(m) {
		return errors.New("id is nil")
	}
	return dao.db.Find(m).Error
}

func (dao *DsBranchUrlDao) BatchGet(idbatch []int64) (result []model.DsBranchUrl, err error) {
	err = dao.db.Model(&model.DsBranchUrl{}).Where("ID in (?)", idbatch).Find(&result).Error
	return
}

func (dao *DsBranchUrlDao) GetForUpdate(m *model.DsBranchUrl) error {
	if dao.db.NewRecord(m) {
		return errors.New("id is nil")
	}
	return dao.db.Set("gorm:query_option", "FOR UPDATE").Find(m).Error
}

func (dao *DsBranchUrlDao) Save(m *model.DsBranchUrl) error {
	return dao.db.Save(m).Error
}

func (dao *DsBranchUrlDao) Delete(m *model.DsBranchUrl) error {
	if dao.db.NewRecord(m) {
		return errors.New("id is nil")
	}
	return dao.db.Delete(m).Error
}

func (dao *DsBranchUrlDao) BatchDelete(idbatch []int64) error {
	return dao.db.Where("ID in (?)", idbatch).Delete(&model.DsBranchUrl{}).Error
}

func (dao *DsBranchUrlDao) Updates(id int64, attrs map[string]interface{}) error {
	return dao.db.Model(&model.DsBranchUrl{}).Where("ID = ?", id).Updates(attrs).Error
}

func (dao *DsBranchUrlDao) Update(id int64, attr string, value interface{}) error {
	return dao.db.Model(&model.DsBranchUrl{}).Where("ID = ?", id).Update(attr, value).Error
}

func (dao *DsBranchUrlDao) BatchUpdaterAttrs(idbatch []int64, attrs map[string]interface{}) error {
	return dao.db.Model(&model.DsBranchUrl{}).Where("ID in (?)", idbatch).Updates(attrs).Error
}

func (dao *DsBranchUrlDao) Found(m *model.DsBranchUrl) bool {
	find := dao.db.First(m, m).RecordNotFound()
	return !find
}
