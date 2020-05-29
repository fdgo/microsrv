package dao

import (
	"ds_server/services/mall/model"
)

func (dao *DsProductStoreDao) FindForUpdate(m *model.DsProductStore) error {
	return dao.db.Set("gorm:query_option", "FOR UPDATE").Find(m, m).Error
}
