package dao

import (
	"ds_server/services/branch/model"
	"errors"

	"github.com/jinzhu/gorm"
)

type DsBranchDao struct {
	db *gorm.DB
}

func NewDsBranchDao(db *gorm.DB) *DsBranchDao {
	return &DsBranchDao{
		db: db,
	}
}

func (dao *DsBranchDao) Create(m *model.DsBranch) error {
	return dao.db.Create(m).Error
}

func (dao *DsBranchDao) Find(m *model.DsBranch) (result []model.DsBranch, err error) {
	err = dao.db.Find(&result, m).Error
	return
}

func (dao *DsBranchDao) FindOne(m *model.DsBranch) error {
	return dao.db.First(m, m).Error
}

func (dao *DsBranchDao) FindLast(m *model.DsBranch) error {
	return dao.db.Last(m, m).Error
}

func (dao *DsBranchDao) FindPage(m *model.DsBranch, rowbound model.RowBound) (result []model.DsBranch, count int, err error) {
	err = dao.db.Model(&model.DsBranch{}).Count(&count).Limit(rowbound.Limit).Offset(rowbound.Offset).Find(&result, m).Error
	return
}

func (dao *DsBranchDao) Get(m *model.DsBranch) error {
	if dao.db.NewRecord(m) {
		return errors.New("id is nil")
	}
	return dao.db.Find(m).Error
}

func (dao *DsBranchDao) BatchGet(idbatch []int64) (result []model.DsBranch, err error) {
	err = dao.db.Model(&model.DsBranch{}).Where("ID in (?)", idbatch).Find(&result).Error
	return
}

func (dao *DsBranchDao) GetForUpdate(m *model.DsBranch) error {
	if dao.db.NewRecord(m) {
		return errors.New("id is nil")
	}
	return dao.db.Set("gorm:query_option", "FOR UPDATE").Find(m).Error
}

func (dao *DsBranchDao) Save(m *model.DsBranch) error {
	return dao.db.Save(m).Error
}

func (dao *DsBranchDao) Delete(m *model.DsBranch) error {
	if dao.db.NewRecord(m) {
		return errors.New("id is nil")
	}
	return dao.db.Delete(m).Error
}

func (dao *DsBranchDao) BatchDelete(idbatch []int64) error {
	return dao.db.Where("ID in (?)", idbatch).Delete(&model.DsBranch{}).Error
}

func (dao *DsBranchDao) Updates(id int64, attrs map[string]interface{}) error {
	return dao.db.Model(&model.DsBranch{}).Where("ID = ?", id).Updates(attrs).Error
}

func (dao *DsBranchDao) Update(id int64, attr string, value interface{}) error {
	return dao.db.Model(&model.DsBranch{}).Where("ID = ?", id).Update(attr, value).Error
}

func (dao *DsBranchDao) BatchUpdaterAttrs(idbatch []int64, attrs map[string]interface{}) error {
	return dao.db.Model(&model.DsBranch{}).Where("ID in (?)", idbatch).Updates(attrs).Error
}

func (dao *DsBranchDao) Found(m *model.DsBranch) bool {
	find := dao.db.First(m, m).RecordNotFound()
	return !find
}
