package dao

import (
	"ds_server/services/mall/model"
)

func (dao *DsProductDao) FindProductByID(product *model.DsProduct) error {
	return dao.db.Model(&model.DsProduct{}).Where("id = ?", product.ID).Preload("DsProductStore").Find(&product).Error
}

func (dao *DsProductDao) QueryProducts(condition *model.ProductSearchCondition, rowBound *model.RowBound) (result []model.DsProduct, count int, err error) {
	db := dao.db

	if condition.BranchID > 0 {
		db = db.Where("branch_id =   ?", condition.BranchID)
	}
	if condition.ProductType > 0 {
		db = db.Where("product_type =   ?", condition.ProductType)
	}

	if condition.BranchName != "" {
		db = db.Where("branch_name like   ?", condition.BranchName+"%")
	}

	if condition.Remark != "" {
		db = db.Where("remark like   ?", condition.Remark+"%")
	}

	if condition.ProductName != "" {
		db = db.Where("product_name like   ?", condition.ProductName+"%")
	}

	if condition.IsOnSale > 0 {
		db = db.Where("Is_On_Sale = ?", condition.IsOnSale)
	}

	if condition.IsRecommend > 0 {
		db = db.Where("is_recommend = ?", condition.IsRecommend)
	}

	if rowBound == nil {
		err = db.Model(&model.DsProduct{}).Order("ID desc").Preload("DsProductStore").Count(&count).Find(&result).Error
	} else {
		err = db.Model(&model.DsProduct{}).Order("ID desc").Preload("DsProductStore").Count(&count).Offset(rowBound.Offset).Limit(rowBound.Limit).Find(&result).Error
	}

	return result, count, err
}
