package dao

import (
	"ds_server/services/branch/model"
	"fmt"
)

func (dao *DsBranchDao) SeachCount(data *model.BranchSearchCondition) (count int, err error) {
	db := dao.db
	if data.Name != "" {
		db = db.Where("name like ?", "%"+data.Name+"%")
	}

	err = db.Table("Ds_Branch").Count(&count).Error
	if err != nil {
		Log.Errorln(err)
		return 0, err
	}
	return count, nil
}

func (dao *DsBranchDao) SeachBranch(data *model.BranchSearchCondition, rowbound *model.RowBound) (request []model.DsBranch, count int, err error) {
	db := dao.db
	count, err = dao.SeachCount(data)
	if err != nil {
		Log.Errorln(err)
		return []model.DsBranch{}, 0, err
	}

	l := data.Latitude
	d := data.Longitude
	sql := "select *," +
		"ROUND(" +
		"   6378.138 * 2 * ASIN(" +
		" SQRT(" +
		" POW(" +
		" SIN(" +
		" (" +
		" " + l + " * PI() / 180 - Latitude * PI() / 180" +
		"          ) / 2" +
		"         )," +
		"        2" +
		"     ) + COS(" + l + " * PI() / 180) * COS(Latitude * PI() / 180) * POW(" +
		"         SIN(" +
		"           (" +
		"               " + d + " * PI() / 180 - Longitude * PI() / 180" +
		"            ) / 2" +
		"         )," +
		"          2 " +
		"       )" +
		"     )" +
		"   )  " +
		" ) AS Juli" +
		" FROM Ds_Branch " +
		" WHERE 1 =1 "
	if data.Name != "" {
		sql = sql + " and name like '%" + data.Name + "%'"
	}
	sql = sql + " order by juli asc limit " + fmt.Sprint(rowbound.Limit) + " OFFSET " + fmt.Sprint(rowbound.Offset)

	err = db.Raw(sql).Scan(&request).Error
	if err != nil {
		Log.Errorln(err)
		return []model.DsBranch{}, 0, err
	}

	return request, count, nil

}

func (dao *DsBranchDao) SelectBranchDynamics(data *model.BranchDynamicSearchCondition, rowbound *model.RowBound) (result []model.DsBranchDynamic, count int, err error) {
	db := dao.db
	if data.BranchID > 0 {
		db = db.Where("Branch_ID = ?", data.BranchID)
	}

	if rowbound == nil {
		err = db.Model(&model.DsBranchDynamic{}).Order("ID desc").Count(&count).Find(&result).Error
	} else {
		err = db.Model(&model.DsBranchDynamic{}).Order("ID desc").Count(&count).Offset(rowbound.Offset).Limit(rowbound.Limit).Find(&result).Error
	}
	return result, count, nil
}
