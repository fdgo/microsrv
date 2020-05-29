package dao

import "ds_server/services/branch/model"

func (dao *DsBranchDynamicDao) GetBranchID() (count int64, err error) {
	db := dao.db
	rows, err := db.Model(&model.DsBranchDynamic{}).Select(" IFNULL(max(id),0)+20 count ").Rows()
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
