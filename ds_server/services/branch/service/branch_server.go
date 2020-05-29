package service

import (
	"ds_server/services/branch/dao"
	"ds_server/services/branch/model"
	db "ds_server/support/lib/mysqlex"
	"ds_server/support/utils/idgenerator"
	"ds_server/support/utils/logex"
	"errors"
)

type BranchService struct {
}

var branchService *BranchService

func NewBranchService() *BranchService {
	if branchService == nil {
		l.Lock()
		if branchService == nil {
			branchService = &BranchService{}
		}
		l.Unlock()
	}
	return branchService
}

func (service BranchService) CreateBranchUrl(m model.DsBranchUrl) error {
	db := db.MysqlInstanceg()
	BranchDao := dao.NewDsBranchUrlDao(db)
	return BranchDao.Create(&m)
}

func (service BranchService) GetBranchByID(id int64) (request model.DsBranch, err error) {
	db := db.MysqlInstanceg()
	branchDao := dao.NewDsBranchDao(db)
	if id == 0 {
		return model.DsBranch{}, errors.New("id is null")
	}
	data := model.DsBranch{ID: id}
	err = branchDao.Get(&data)
	if err != nil {
		Log.Error(err)
		return model.DsBranch{}, err
	}
	return data, nil
}

func (service BranchService) GetBranchDynamicByID(id int64) (request model.DsBranchDynamic, err error) {
	db := db.MysqlInstanceg()
	branchDao := dao.NewDsBranchDynamicDao(db)
	if id == 0 {
		return model.DsBranchDynamic{}, errors.New("id is null")
	}
	data := model.DsBranchDynamic{ID: id}
	err = branchDao.Get(&data)
	if err != nil {
		Log.Error(err)
		return model.DsBranchDynamic{}, err
	}
	return data, nil
}

func (service BranchService) CreateBranch(m model.DsBranch) error {
	db := db.MysqlInstanceg()
	branchDao := dao.NewDsBranchDao(db)
	return branchDao.Create(&m)
}

func (service BranchService) CreateBranchDynamic(m model.DsBranchDynamic) error {
	db := db.MysqlInstanceg()
	branchdDao := dao.NewDsBranchDynamicDao(db)
	err := branchdDao.Create(&m)
	if err != nil {
		logex.Error(err)
		return err
	}
	return err
}

func (service BranchService) SearchBranchrPaging(condition *model.BranchSearchCondition, pageNum int, pageSize int) (request []model.DsBranch, count int, err error) {
	rowbound := model.NewRowBound(pageNum, pageSize)
	return service.searchBranch(condition, &rowbound)
}

func (service BranchService) SearchBranchrOutPaging(condition *model.BranchSearchCondition) (request []model.DsBranch, count int, err error) {
	return service.searchBranch(condition, nil)
}

func (service BranchService) searchBranch(condition *model.BranchSearchCondition, rowbound *model.RowBound) (request []model.DsBranch, count int, err error) {
	db := db.MysqlInstanceg()
	branchDao := dao.NewDsBranchDao(db)
	request, count, err = branchDao.SeachBranch(condition, rowbound)
	if err != nil {
		Log.Error(err)
		return nil, 0, err
	}
	return request, count, nil
}

func (service BranchService) GetBranchID() (count int64, err error) {
	gen := idgenerator.Instance()
	id := gen.GenerateLongID("user")
	return id, nil
}

func (service BranchService) SelectBranchPaging(condition *model.BranchSearchCondition, pageNum int, pageSize int) (request []model.DsBranch, count int, err error) {
	rowbound := model.NewRowBound(pageNum, pageSize)
	return service.selectBranchBranch(condition, &rowbound)
}

func (service BranchService) SelectBranchOutPaging(condition *model.BranchSearchCondition) (request []model.DsBranch, count int, err error) {
	return service.selectBranchBranch(condition, nil)
}

func (service BranchService) selectBranchBranch(condition *model.BranchSearchCondition, rowbound *model.RowBound) (request []model.DsBranch, count int, err error) {
	db := db.MysqlInstanceg()
	branchDao := dao.NewDsBranchDao(db)
	request, count, err = branchDao.SeachBranch(condition, rowbound)
	if err != nil {
		Log.Error(err)
		return nil, 0, err
	}
	return request, count, nil
}

func (service BranchService) SelectBranchDynamicsPaging(condition *model.BranchDynamicSearchCondition, pageNum int, pageSize int) (request []model.DsBranchDynamic, count int, err error) {
	rowbound := model.NewRowBound(pageNum, pageSize)
	return service.selectBranchDynamics(condition, &rowbound)
}

func (service BranchService) SelectBranchDynamicsOutPaging(condition *model.BranchDynamicSearchCondition) (request []model.DsBranchDynamic, count int, err error) {
	return service.selectBranchDynamics(condition, nil)
}

func (service BranchService) selectBranchDynamics(condition *model.BranchDynamicSearchCondition, rowbound *model.RowBound) (request []model.DsBranchDynamic, count int, err error) {
	db := db.MysqlInstanceg()
	branchDao := dao.NewDsBranchDao(db)
	result, count, err := branchDao.SelectBranchDynamics(condition, rowbound)
	if err != nil {
		Log.Error(err)
		return nil, 0, err
	}
	branchuDao := dao.NewDsBranchUrlDao(db)
	for i, v := range result {
		urs, err := branchuDao.Find(&model.DsBranchUrl{DsBranchDynamicID: v.ID})
		if err != nil {
			Log.Error(err)
			return nil, 0, err
		}
		result[i].Urls = urs
	}

	return result, count, nil
}
