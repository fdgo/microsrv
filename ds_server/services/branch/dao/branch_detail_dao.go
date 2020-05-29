package dao

import (
	"ds_server/services/branch/model"
	"errors"

	"github.com/jinzhu/gorm"
)

type DsBranchDynamicDao struct {
	db *gorm.DB
}

func NewDsBranchDynamicDao(db *gorm.DB) *DsBranchDynamicDao {
	return &DsBranchDynamicDao{
		db: db,
	}
}


func (dao *DsBranchDynamicDao) Create(m *model.DsBranchDynamic) error {
	return dao.db.Create(m).Error
}

func (dao *DsBranchDynamicDao) Find(m *model.DsBranchDynamic) (result []model.DsBranchDynamic, err error) {
	err = dao.db.Find(&result, m).Error
	return
}

func (dao *DsBranchDynamicDao) FindOne(m *model.DsBranchDynamic) error {
	return dao.db.First(m, m).Error
}

func (dao *DsBranchDynamicDao) FindLast(m *model.DsBranchDynamic) error {
	return dao.db.Last(m, m).Error
}

func (dao *DsBranchDynamicDao) FindPage(m *model.DsBranchDynamic, rowbound model.RowBound) (result []model.DsBranchDynamic, count int, err error) {
	err = dao.db.Model(&model.DsBranchDynamic{}).Count(&count).Limit(rowbound.Limit).Offset(rowbound.Offset).Find(&result, m).Error
	return
}

func (dao *DsBranchDynamicDao) Get(m *model.DsBranchDynamic) error {
	if dao.db.NewRecord(m) {
		return errors.New("id is nil")
	}
	return dao.db.Find(m).Error
}

func (dao *DsBranchDynamicDao) BatchGet(idbatch []int64) (result []model.DsBranchDynamic, err error) {
	err = dao.db.Model(&model.DsBranchDynamic{}).Where("ID in (?)", idbatch).Find(&result).Error
	return
}

func (dao *DsBranchDynamicDao) GetForUpdate(m *model.DsBranchDynamic) error {
	if dao.db.NewRecord(m) {
		return errors.New("id is nil")
	}
	return dao.db.Set("gorm:query_option", "FOR UPDATE").Find(m).Error
}

func (dao *DsBranchDynamicDao) Save(m *model.DsBranchDynamic) error {
	return dao.db.Save(m).Error
}

func (dao *DsBranchDynamicDao) Delete(m *model.DsBranchDynamic) error {
	if dao.db.NewRecord(m) {
		return errors.New("id is nil")
	}
	return dao.db.Delete(m).Error
}

func (dao *DsBranchDynamicDao) BatchDelete(idbatch []int64) error {
	return dao.db.Where("ID in (?)", idbatch).Delete(&model.DsBranchDynamic{}).Error
}

func (dao *DsBranchDynamicDao) Updates(id int64, attrs map[string]interface{}) error {
	return dao.db.Model(&model.DsBranchDynamic{}).Where("ID = ?", id).Updates(attrs).Error
}

func (dao *DsBranchDynamicDao) Update(id int64, attr string, value interface{}) error {
	return dao.db.Model(&model.DsBranchDynamic{}).Where("ID = ?", id).Update(attr, value).Error
}

func (dao *DsBranchDynamicDao) BatchUpdaterAttrs(idbatch []int64, attrs map[string]interface{}) error {
	return dao.db.Model(&model.DsBranchDynamic{}).Where("ID in (?)", idbatch).Updates(attrs).Error
}

func (dao *DsBranchDynamicDao) Found(m *model.DsBranchDynamic) bool {
	find := dao.db.First(m, m).RecordNotFound()
	return !find
}
