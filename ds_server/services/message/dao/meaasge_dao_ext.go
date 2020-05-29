package dao

import (
	"ds_server/services/message/model"
)

func (dao *DsMessageDao) SeachMessage(data *model.MessageSearchCondition, rowbound *model.RowBound) (result []model.DsMessage, count int, err error) {
	db := dao.db
	if data.MessageType != "" {
		db = db.Where("Message_Type = ?", data.MessageType)
	}

	if data.Name != "" {
		db = db.Where("name = ?", data.Name)
	}

	if data.Swith > 0 {
		db = db.Where("Swith = ?", data.Swith)
	}

	if data.UserID != "" {
		db = db.Where("User_Id = ?", data.UserID)
	}
	if data.UserType > 0 {
		db = db.Where("User_Type = ?", data.UserType)
	}

	if rowbound == nil {
		err = db.Model(&model.DsMessage{}).Order("ID desc").Count(&count).Find(&result).Error
	} else {
		err = db.Model(&model.DsMessage{}).Order("ID desc").Count(&count).Offset(rowbound.Offset).Limit(rowbound.Limit).Find(&result).Error
	}
	return

}

func (dao *DsMessageDao) GetMessageID() (count int64, err error) {
	db := dao.db
	rows, err := db.Model(&model.DsMessage{}).Select(" IFNULL(max(id),0)+20 count").Rows()
	defer rows.Close()
	if err != nil {
		Log.Error(err)
		return 0, nil
	}
	if rows.Next() {
		err := rows.Scan(&count)
		if err != nil {
			Log.Error(err)
			return 0, nil
		}
	}
	return count, nil
}
