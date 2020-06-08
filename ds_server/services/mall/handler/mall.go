package handler

import (
	"context"
	pb "ds_server/proto/mall"
	"ds_server/services/mall/model"
	"ds_server/services/mall/service"
	"ds_server/support/utils/logex"
	"encoding/json"
	"fmt"
	"time"
)

type MallHandler struct {
	Mall *service.MallService
}

func (branch *MallHandler) GetProductByID(c context.Context, req *pb.ProductByID, rsp *pb.Response) error {
	mService := service.NewMallService()
	result, err := mService.GetProductByID(c, req.ID, req.Lat1, req.Lng1)
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

func (branch *MallHandler) GetProductDetailByID(c context.Context, req *pb.ProductByID, rsp *pb.Response) error {
	mService := service.NewMallService()
	result, err := mService.GetProductDetailByID(c, req.ID)
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

func (branch *MallHandler) SelectProducts(c context.Context, req *pb.SelectProductRequest, rsp *pb.Response) error {
	mService := service.NewMallService()
	condition := new(model.ProductSearchCondition)
	condition.BranchID = req.BranchID
	condition.ProductName = req.ProductName
	condition.BranchName = req.BranchName
	condition.ProductType = model.ProductType(req.ProductType)
	condition.Remark = req.Remark
	condition.IsOnSale = model.YesOrNo(req.IsOnSale)
	condition.IsRecommend = model.YesOrNo(req.IsRecommend)
	result, count, err := mService.SearchProductPaging(condition, int(req.PageNum), int(req.PageSize))
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

func (branch *MallHandler) CreateAddress(c context.Context, req *pb.CreateAddressRequest, rsp *pb.Response) error {
	mService := service.NewMallService()
	address := &model.DsReceivingAddress{
		UserID:      req.UserID,
		UserAccount: req.UserAccount,
		UserAlias:   req.UserAlias,
		Province:    req.Province,
		City:        req.City,
		Aera:        req.Aera,
		Address:     req.Address,
		Tel:         req.Tel,
		Sex:         model.Sex(req.Sex),
		Label:       model.Label(req.Label),
		TelName:     req.TelName,
	}
	err := mService.CreateAddress(address)
	if err != nil {
		Log.Errorln(err)
		return err
	}
	return nil
}

func (branch *MallHandler) SearchAddress(c context.Context, req *pb.RequestAddressID, rsp *pb.Response) error {
	mService := service.NewMallService()
	result, err := mService.SearchAddress(req.UserID)
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

func (branch *MallHandler) CreateOrder(c context.Context, req *pb.CreateOrderRequest, rsp *pb.Response) error {
	oService := service.NewOrderService()
	order := model.DsOrder{}
	order.UserID = req.UserID
	order.ProductNum = int(req.ProductNum)
	order.ProductID = req.ProductID
	order.AddressID = req.AddressID
	err := oService.CreateOrder(c, &order)
	if err != nil {
		Log.Errorln(err)
		return err
	}
	b, err := json.Marshal(order)
	if err != nil {
		Log.Errorln(err)
		return err
	}
	rsp.Data = b
	return nil
}

func (branch *MallHandler) SearchOrder(c context.Context, req *pb.SelectOrderRequest, rsp *pb.Response) error {
	mService := service.NewOrderService()
	condition := new(model.OrderSearchCondition)
	if req.CreateStartTime != "" {
		condition.CreateStartTime, _ = time.ParseInLocation("2006-01-02 15:04:05", req.CreateStartTime, time.Local)
	}
	if req.CreateEndTime != "" {
		condition.CreateEndTime, _ = time.ParseInLocation("2006-01-02 15:04:05", req.CreateEndTime, time.Local)
	}
	condition.UserID = req.UserID
	condition.OrderState = model.OrderState(req.OrderState)
	result, count, _, err := mService.SearchOrderPaging(condition, int(req.PageNum), int(req.PageSize))
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

func (branch *MallHandler) PayOrder(c context.Context, req *pb.PayOrderRequest, rsp *pb.Response) error {
	mService := service.NewOrderService()
	order := new(model.DsOrder)
	order.UserID = req.UserID
	order.ID = req.OrderID
	order.OrderState = model.OrderStateBeDelivered

	//transactionPassword := req.TransactionPassword
	err := mService.UpdateOrderState(order, c)
	if err != nil {
		return err
	}
	return nil
}

func (branch *MallHandler) GetOrderByID(c context.Context, req *pb.ProductByID, rsp *pb.Response) error {
	mService := service.NewOrderService()
	order, err := mService.GetOrderByID(req.ID)
	if err != nil {
		return err
	}
	b, err := json.Marshal(order)
	if err != nil {
		return err
	}
	rsp.Data = b
	return nil
}
