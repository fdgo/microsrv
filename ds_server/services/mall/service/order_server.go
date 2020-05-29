package service

import (
	"context"
	rspmdl "ds_server/models/user/gin_rsp"
	userpro "ds_server/proto/user"
	"ds_server/services/mall/dao"
	"ds_server/services/mall/model"
	db "ds_server/support/lib/mysqlex"
	"ds_server/support/utils/constex"
	"ds_server/support/utils/idgenerator"
	"ds_server/support/utils/logger"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/micro/go-micro/client"
	"github.com/shopspring/decimal"
)

type OrderService struct {
}

var orderService *OrderService

func NewOrderService() *OrderService {
	if orderService == nil {
		l.Lock()
		if orderService == nil {
			orderService = &OrderService{}
		}
		l.Unlock()
	}
	return orderService
}

func (service OrderService) SearchOrderPaging(condition *model.OrderSearchCondition, pageNum int, pageSize int) (request []model.DsOrder, count int, total string, err error) {
	rowbound := model.NewRowBound(pageNum, pageSize)
	return service.searchOrder(condition, &rowbound)
}

func (service OrderService) SearchOrderWithOutPaging(condition *model.OrderSearchCondition) (request []model.DsOrder, count int, total string, err error) {
	return service.searchOrder(condition, nil)
}

func (service OrderService) searchOrder(condition *model.OrderSearchCondition, rowbound *model.RowBound) (request []model.DsOrder, count int, total string, err error) {
	db := db.MysqlInstanceg()
	orderDao := dao.NewDsOrderDao(db)
	result, count, total, err := orderDao.QueryOrders(condition, rowbound)
	if err != nil {
		Log.Error(err)
		return nil, 0, total, err
	}

	if err != nil {
		Log.Error(err)
		return nil, 0, total, err
	}
	return result, count, total, err
}

func (service OrderService) CreateOrder(c context.Context, order *model.DsOrder) (err error) {
	db := db.MysqlInstanceg()
	tx := db.Begin()
	defer closeTx(tx, &err)
	orderDao := dao.NewDsOrderDao(tx)
	//判断商品库存够不够
	product, err := NewMallService().GetProductByID(c, order.ProductID, "", "")
	if err != nil {
		Log.Error(err)
		return errors.New("没找到对应商品")
	}
	ret_resp := rspmdl.MemAgent_rsp{}
	client := userpro.NewUserService(constex.SRV_USER, client.DefaultClient)
	var rtin userpro.GetMemberUserAgentIn
	rtin.Uuid = order.UserID
	ret, err := client.GetMemberUserAgent(c, &rtin)
	if err != nil {
		Log.Errorln(err)
		return errors.New("未找到人员信息1")
	}

	err = json.Unmarshal((*ret).Data, &ret_resp)
	if err != nil {
		Log.Errorln(err)
		return errors.New("未找到人员信息2")
	}
	order.TotalAmount = order.OrderPrice.Mul(decimal.RequireFromString(fmt.Sprint(order.ProductNum)))

	if decimal.NewFromFloat(ret_resp.Balance).LessThan(order.TotalAmount) {
		if err != nil {
			return errors.New("金额不足")
		}
	}
	order.UserName = ret_resp.MemberName
	order.Account = ret_resp.MobileSelf

	order.ProductName = product.ProductName
	order.ProductType = model.ProductType(product.ProductType)
	order.BranchID = product.BranchID
	order.BranchName = product.BranchName
	order.ProductUrl = product.Url
	order.ProductPrice = product.DsProductStore.Price
	order.OrderPrice = product.DsProductStore.RealPrice
	if product.DsProductStore.StoreNum < order.ProductNum {
		return errors.New("商品库存不足")
	}
	addrsDao := dao.NewDsReceivingAddressDao(tx)
	err = addrsDao.Get(&model.DsReceivingAddress{ID: order.AddressID})
	if err != nil {
		return errors.New("收货地址不存在，请新建收货地址")
	}
	//判断用户是否存在
	// 生成全局ID
	// 待付款
	order.OrderState = model.OrderStateBePaid
	gen := idgenerator.Instance()
	id := gen.GenerateLongID("order")
	order.ID = id
	// 将ID转换为Code
	code := gen.ChaosID(id, gen.Suffix("order"))
	order.OrderNo = code
	Log.Infoln("order ", logger.FormatStruct(order))

	err = NewProductStoreService().PlayerDecreaseNum(tx, order.UserID, order.ProductID, order.ProductNum, order.ID, func(wallet *model.DsProductStore, number int) {
		wallet.StoreNumTotal = wallet.StoreNumTotal + number
	})
	if err != nil {
		return errors.New("商品库存不足2")
	}
	err = orderDao.Create(order)
	if err != nil {
		return errors.New("下单失败，请联系管理员")
	}

	return nil
}

func (service OrderService) GetOrderForUpdate(tx *gorm.DB, orderNo string) (*model.DsOrder, error) {
	orderDao := dao.NewDsOrderDao(tx)
	orderQuery := &model.DsOrder{
		OrderNo: orderNo,
	}
	if err := orderDao.FindOrderByOrderNo(orderQuery); err != nil {
		Log.Errorln(err)
		return nil, err
	}

	order, err := orderDao.FindForUpdate(orderQuery.ID)

	if err != nil {
		Log.Errorln(err)
		return nil, err
	}
	return order, err
}

func (service OrderService) UpdateOrderState(param *model.DsOrder, c context.Context) (err error) {
	db := db.MysqlInstanceg()
	tx := db.Begin()
	defer closeTx(tx, &err)
	_, err = service.updateOrderState(tx, c, param)
	if err != nil {
		Log.Errorln(err)
		return err
	}
	return nil
}

func (service OrderService) updateOrderState(tx *gorm.DB, c context.Context, param *model.DsOrder) (errstring string, err error) {
	Log.Debug("updateOrderState")
	orderDao := dao.NewDsOrderDao(tx)
	errstring = ""
	//order, err := service.GetOrderForUpdate(tx, param.ID)

	order, err := orderDao.FindForUpdate(param.ID)

	Log.Infoln("========", order)
	if err != nil {
		Log.Errorln(err)
		return errstring, err
	}

	//验证order state是否合法
	// if err := validateOrderStates(order.State, param.State, isWithdrawalAuditRequest); err != nil {
	// 	Log.Error(err)
	// 	errstring = "验证订单错误"
	// 	return "", err
	// }

	attrs := map[string]interface{}{}
	attrs["order_state"] = param.OrderState
	attrs["remark"] = param.Remark

	//订单完成则写上完成时间
	if param.OrderState == model.OrderStateComplete {
		attrs["complete_time"] = time.Now() //完成时间
	}

	if err := orderDao.Updates(param.ID, attrs); err != nil {
		Log.Error(err)
		errstring = "修改订单错误"
		return errstring, err
	}

	// OrderStateBeDelivered            //待发货
	// OrderStateBeReceived             //待收货
	// OrderStateRefund                 //退款
	// OrderStateReComolete             //退款完成
	// OrderStateComplete               //完成
	// OrderStateColse                  //关闭

	switch param.OrderState {

	case model.OrderStateBePaid:

	case model.OrderStateBeDelivered:
		//判断金额够不够
		// return "", errex.New("余额不足请充值")
		// //足够则进行扣钱
		ret_resp := rspmdl.MemAgent_rsp{}
		client := userpro.NewUserService(constex.SRV_USER, client.DefaultClient)
		var rtin userpro.GetMemberUserAgentIn
		rtin.Uuid = order.UserID

		ret, err := client.GetMemberUserAgent(c, &rtin)
		if err != nil {
			Log.Errorln(err)
			return "", err
		}

		err = json.Unmarshal((*ret).Data, &ret_resp)
		if err != nil {
			Log.Errorln(err)
			return "", err
		}

		if decimal.NewFromFloat(ret_resp.Balance).LessThan(order.TotalAmount) {
			if err != nil {
				return "金额不足", errors.New("金额不足")
			}
		}

		var deposit userpro.OnlinePayIn
		deposit.SrcId = fmt.Sprint(order.ID)
		float, err := strconv.ParseFloat(fmt.Sprint(order.TotalAmount), 32)
		if err != nil {
			return "金额格式错误", errors.New("金额格式错误")
		}
		deposit.DepositNum = float32(float)
		deposit.DepositType = 3
		deposit.Uuid = order.UserID
		_, err = client.OnlinePay(c, &deposit)
		if err != nil {
			return "扣钱失败", errors.New("扣钱失败")
		}

	case model.OrderStateBeReceived:

	case model.OrderStateRefund:

	case model.OrderStateReComolete:

	case model.OrderStateComplete:

	case model.OrderStateColse:
		err = NewProductStoreService().PlayerDecreaseNum(tx, order.UserID, order.ProductID, order.ProductNum, order.ID, func(wallet *model.DsProductStore, number int) {
			wallet.StoreNumTotal = wallet.StoreNumTotal - number
		})

		if err != nil {
			return "商品库存不足", errors.New("商品库存不足")
		}
	}

	return errstring, nil
}

func (service OrderService) GetOrderByID(id int64) (order model.DsOrder, err error) {
	db := db.MysqlInstanceg()
	orderDao := dao.NewDsOrderDao(db)
	or := model.DsOrder{}
	or.ID = id
	err = orderDao.Get(&or)
	if err != nil {
		return model.DsOrder{}, err
	}
	addDao := dao.NewDsReceivingAddressDao(db)

	add := model.DsReceivingAddress{}
	add.ID = or.AddressID
	err = addDao.Get(&add)
	if err != nil {
		return model.DsOrder{}, err
	}

	or.DsReceivingAddress = add

	return or, err
}
