package handler

import (
	"context"
	pb "ds_server/proto/branch"
	"ds_server/services/branch/model"
	"ds_server/services/branch/service"
	"ds_server/support/utils/logex"
	"ds_server/support/utils/logex"
	"encoding/json"
	"errors"
	"fmt"
)

type BranchHandler struct {
	Branch *service.BranchService
}

func (branch *BranchHandler) GetBranchByID(c context.Context, req *pb.IdRequest, rsp *pb.Response) error {
	mService := service.NewBranchService()
	result, err := mService.GetBranchByID(req.ID)
	if err != nil {
		Log.Errorln(err)
		return err
	}
	b, err := json.Marshal(result)
	if err != nil {
		Log.Errorln(err)
		return err
	}
	rsp.Data = b
	return nil
}

func (branch *BranchHandler) GetBranchDynamicByID(c context.Context, req *pb.IdRequest, rsp *pb.Response) error {
	mService := service.NewBranchService()
	result, err := mService.GetBranchDynamicByID(req.ID)
	if err != nil {
		Log.Errorln(err)
		return err
	}
	b, err := json.Marshal(result)
	if err != nil {
		Log.Errorln(err)
		return err
	}
	rsp.Data = b
	return nil
}

func (branch *BranchHandler) SelectBranchDynamics(c context.Context, req *pb.SelectBranchDynamicsRequest, rsp *pb.Response) error {
	fmt.Println("============", logger.FormatStruct(req))
	mService := service.NewBranchService()
	condition := new(model.BranchDynamicSearchCondition)
	condition.BranchID = req.BranchID
	result, count, err := mService.SelectBranchDynamicsPaging(condition, int(req.PageNum), int(req.PageSize))
	if err != nil {
		return err
	}
	b, err := json.Marshal(result)
	if err != nil {
		return err
	}
	rsp.Count = int32(count)
	rsp.Data = b
	fmt.Println("============", logger.FormatStruct(rsp))
	return nil
}

func (branch *BranchHandler) SelectBranch(c context.Context, req *pb.SelectBranchRequest, rsp *pb.Response) error {
	fmt.Println("============", logger.FormatStruct(req))
	mService := service.NewBranchService()
	condition := new(model.BranchSearchCondition)
	condition.Name = req.Name
	condition.Latitude = req.Latitude
	condition.Longitude = req.Longitude
	condition.GeoHashCode = req.GeoHashCode
	result, count, err := mService.SearchBranchrPaging(condition, int(req.PageNum), int(req.PageSize))
	if err != nil {
		return err
	}

	b, err := json.Marshal(result)
	if err != nil {
		return err
	}
	rsp.Count = int32(count)
	rsp.Data = b
	fmt.Println("============", logger.FormatStruct(rsp))

	mwService := service.NewUsdtService()
	mwService.GetBalance()
	return nil
}

func (branch *BranchHandler) CreateBranch(c context.Context, req *pb.CreateRequest, rsp *pb.Response) error {
	ress := model.DsBranch{}
	err := json.Unmarshal(req.Data, &ress)
	if err != nil {
		logex.Error(err)
		return err
	}
	mService := service.NewBranchService()
	return mService.CreateBranch(ress)
}

func (branch *BranchHandler) CreateBranchDynamic(c context.Context, req *pb.CreateBranchDynamicRequest, rsp *pb.Response) error {
	ress := model.DsBranchDynamic{}
	ress.ID = req.BranchDynamicID
	ress.BranchID = req.BranchID
	ress.Title = ress.Title
	ress.Name = ress.Name
	ress.Content = ress.Content
	mService := service.NewBranchService()
	err := mService.CreateBranchDynamic(ress)
	if err != nil {
		logex.Error(err)
		return err
	}
	return nil
}

func (Branch *BranchHandler) GetBranchID(c context.Context, req *pb.IdRequest, rsp *pb.IdResponse) error {
	mService := service.NewBranchService()
	id, err := mService.GetBranchID()
	if err != nil {
		logex.Error(err)
		return err
	}
	rsp.ID = id
	return nil
}

func (Branch *BranchHandler) CreateBranchUrl(c context.Context, req *pb.UrlRequest, rsp *pb.Response) error {
	mService := service.NewBranchService()
	ress := model.DsBranchUrl{}
	ress.DsBranchDynamicID = req.BranchDynamicID
	ress.Url = req.Url
	err := mService.CreateBranchUrl(ress)
	if err != nil {
		logex.Error(err)
		return err
	}
	return nil
}

func (branch *BranchHandler) GetAddress(c context.Context, req *pb.IdRequest, rsp *pb.AddResponse) error {
	mwService := service.NewUsdtService()
	dto, err := mwService.GetAddress()
	//	mwService.GetBalance()
	if err != nil {
		logex.Error(err)
		return err
	}
	if dto.Code != "10000" {
		logex.Error(err)
		return errors.New(dto.Message)
	}
	Log.Infoln("=============================", dto.Data)
	rsp.Address = dto.Data[0]

	return nil
}
