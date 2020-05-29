package service

import (
	"context"
	pbm "ds_server/proto/branch"
	"ds_server/services/mall/dao"
	"ds_server/services/mall/model"
	"ds_server/services/mall/service/dto"
	"errors"
	"strconv"

	db "ds_server/support/lib/mysqlex"
	"ds_server/support/utils/constex"
	"encoding/json"
	"fmt"
	"math"

	"github.com/jinzhu/gorm"
	"github.com/micro/go-micro/client"
)

type MallService struct {
}

var mallService *MallService

func NewMallService() *MallService {
	if mallService == nil {
		l.Lock()
		if mallService == nil {
			mallService = &MallService{}
		}
		l.Unlock()
	}
	return mallService
}

func EarthDistance(lat1, lat2, lng1, lng2 float64) float64 {
	radius := 6371000.0 //6378137.0
	rad := math.Pi / 180.0
	lat1 = lat1 * rad
	lng1 = lng1 * rad
	lat2 = lat2 * rad
	lng2 = lng2 * rad
	theta := lng2 - lng1
	dist := math.Acos(math.Sin(lat1)*math.Sin(lat2) + math.Cos(lat1)*math.Cos(lat2)*math.Cos(theta))
	return dist * radius / 1000
}

func (service MallService) GetProductByID(c context.Context, id int64, lat1 string, lng1 string) (request model.DsProduct, err error) {
	db := db.MysqlInstanceg()
	orderDao := dao.NewDsProductDao(db)
	productUrlDao := dao.NewDsProductUrlDao(db)
	productDetailDao := dao.NewDsProductDetailDao(db)
	result := model.DsProduct{ID: id}
	err = orderDao.FindProductByID(&result)
	if err != nil {
		Log.Error(err)
		return model.DsProduct{}, err
	}
	urls, err := productUrlDao.Find(&model.DsProductUrl{ProductID: id, UrlType: model.BuyUrlType})
	if err != nil {
		Log.Error(err)
		return model.DsProduct{}, err
	}
	result.DsProductUrls = urls

	ss := model.DsProductUrl{ProductID: id, UrlType: model.BuyUrlType}
	err = productUrlDao.FindOne(&ss)
	if err != nil {
		Log.Error(err)
		return model.DsProduct{}, err
	}
	result.Url = ss.Url
	//查询距离
	branchService := pbm.NewBranchService(constex.SRV_BRANCH, client.DefaultClient)
	response, err := branchService.GetBranchByID(c, &pbm.IdRequest{ID: result.BranchID})
	if err != nil {
		Log.Error(err)
		return model.DsProduct{}, err
	}

	ress := dto.DsBranch{}
	err = json.Unmarshal(response.Data, &ress)
	if err != nil {
		Log.Error(err)
		return model.DsProduct{}, err
	}

	if lat1 != "" && lng1 != "" {
		lat1f, err := strconv.ParseFloat(lat1, 64)
		if err != nil {
			Log.Error(err)
			return model.DsProduct{}, err
		}

		lng1f, err := strconv.ParseFloat(lng1, 64)
		if err != nil {
			Log.Error(err)
			return model.DsProduct{}, err
		}
		latitude, err := strconv.ParseFloat(ress.Latitude, 64)
		if err != nil {
			Log.Error(err)
			return model.DsProduct{}, err
		}
		longitude, err := strconv.ParseFloat(ress.Longitude, 64)
		if err != nil {
			Log.Error(err)
			return model.DsProduct{}, err
		}
		distance := EarthDistance(lat1f, latitude, lng1f, longitude)
		sa := math.Trunc(distance*1e2+0.5) * 1e-2
		result.Distance = fmt.Sprintf("%.1fkm", sa)
	}

	detail := model.DsProductDetail{}
	err = productDetailDao.FindOne(&detail)
	if err != nil {
		Log.Error(err)
		return model.DsProduct{}, err
	}
	result.DsProductDetail = detail
	return result, nil

}

func (service MallService) SearchProductPaging(condition *model.ProductSearchCondition, pageNum int, pageSize int) (request []model.DsProduct, count int, err error) {
	rowbound := model.NewRowBound(pageNum, pageSize)
	return service.searchProduct(condition, &rowbound)
}

func (service MallService) SearchProductWithOutPaging(condition *model.ProductSearchCondition) (request []model.DsProduct, count int, err error) {
	return service.searchProduct(condition, nil)
}

func (service MallService) searchProduct(condition *model.ProductSearchCondition, rowbound *model.RowBound) (request []model.DsProduct, count int, err error) {
	db := db.MysqlInstanceg()
	orderDao := dao.NewDsProductDao(db)
	productUrlDao := dao.NewDsProductUrlDao(db)
	condition.IsOnSale = model.Yes //商品在售
	result, count, err := orderDao.QueryProducts(condition, rowbound)
	if err != nil {
		Log.Error(err)
		return nil, 0, err
	}
	for i, v := range result {
		productUrl := &model.DsProductUrl{UrlType: model.StaticUrlType, ProductID: v.ID}
		err = productUrlDao.FindOne(productUrl)
		if err != nil {
			Log.Error(err)
			return nil, 0, err
		}
		result[i].Url = productUrl.Url
	}
	return result, count, err
}

func (service MallService) SearchAddress(id string) (request []model.DsReceivingAddress, err error) {
	db := db.MysqlInstanceg()
	orderDao := dao.NewDsReceivingAddressDao(db)
	return orderDao.Find(&model.DsReceivingAddress{UserID: id})
}

func (service MallService) CreateAddress(address *model.DsReceivingAddress) (err error) {
	db := db.MysqlInstanceg()
	orderDao := dao.NewDsReceivingAddressDao(db)
	result, err := orderDao.Find(&model.DsReceivingAddress{UserID: address.UserID})
	if !gorm.IsRecordNotFoundError(err) {
		if len(result) > 4 {
			return errors.New("收货地址不能超过5条！")
		}
	}
	return orderDao.Create(address)
}

func (service MallService) GetProductDetailByID(c context.Context, id int64) (request model.DsProductDetail, err error) {
	db := db.MysqlInstanceg()
	productDetailDao := dao.NewDsProductDetailDao(db)
	detail := model.DsProductDetail{DsProductID: id}
	err = productDetailDao.FindOne(&detail)
	if err != nil {
		Log.Error(err)
		return model.DsProductDetail{}, err
	}
	return detail, nil

}
