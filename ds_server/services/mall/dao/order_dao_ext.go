package dao

import (
	"ds_server/services/mall/model"
)

func (dao *DsOrderDao) FindForUpdate(id int64) (*model.DsOrder, error) {
	order := &model.DsOrder{}
	err := dao.db.Set("gorm:query_option", "FOR UPDATE").First(order, id).Error
	return order, err
}

func (dao *DsOrderDao) FindOrderByID(order *model.DsOrder) error {
	return dao.db.Model(&model.DsOrder{}).Where("id = ?", order.ID).Preload("DsOrderParameter").Find(&order).Error
}

func (dao *DsOrderDao) FindOrderByOrderNo(order *model.DsOrder) error {
	return dao.db.Model(&model.DsOrder{}).Where("Order_no = ?", order.OrderNo).Preload("DsOrderParameter").Find(&order).Error
}

func (dao *DsOrderDao) QueryOrders(condition *model.OrderSearchCondition, rowBound *model.RowBound) (result []model.DsOrder, count int, total string, err error) {
	db := dao.db
	if condition.UserID != "" {
		db = db.Where("user_id in (?) ", condition.UserID)
	}
	// 根据订单号查询时 去掉创建时间的条件
	if !condition.CreateEndTime.IsZero() && condition.OrderNo == "" {
		db = db.Where("created_at BETWEEN ? AND ?", condition.CreateStartTime, condition.CreateEndTime)
	}
	// 完成时间
	if !condition.CompleteStartTime.IsZero() && !condition.CompleteEndTime.IsZero() {
		db = db.Where("complete_time BETWEEN ? AND ?", condition.CompleteStartTime, condition.CompleteEndTime)
	} else if !condition.CompleteStartTime.IsZero() {
		db = db.Where("complete_time >= ?", condition.CompleteStartTime)
	} else if !condition.CompleteEndTime.IsZero() {
		db = db.Where("complete_time <= ?", condition.CompleteEndTime)
	}

	if condition.OrderNo != "" {
		db = db.Where("order_no like   ?", condition.OrderNo+"%")
	}

	if condition.IP != "" {
		db = db.Where("IP like   ?", condition.IP+"%")
	}

	if condition.Account != "" {
		db = db.Where("account like  ? ", condition.Account+"%")
	}

	if condition.Remark != "" {
		db = db.Where("remark like   ? ", condition.Remark+"%")
	}

	if condition.OrderState > 0 {
		db = db.Where("order_state   = ?  ", condition.OrderState)
	}

	rows, err := db.Model(&model.DsOrder{}).Select(" IFNULL(sum(order_price),0) AS total").Rows()
	defer rows.Close()
	if err != nil {
		Log.Error(err)
		return nil, 0, "", nil
	}

	if rows.Next() {
		err := rows.Scan(&total)
		if err != nil {
			Log.Error(err)
			return nil, 0, "", nil
		}
	}

	if rowBound == nil {
		err = db.Model(&model.DsOrder{}).Order("ID desc").Preload("DsOrderParameter").Count(&count).Find(&result).Error
	} else {
		err = db.Model(&model.DsOrder{}).Order("ID desc").Preload("DsOrderParameter").Count(&count).Offset(rowBound.Offset).Limit(rowBound.Limit).Find(&result).Error
	}

	return result, count, total, err
}
