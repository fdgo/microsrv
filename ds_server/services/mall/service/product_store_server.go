package service

import (
	"ds_server/services/mall/dao"
	"ds_server/services/mall/model"
	"errors"

	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
)

type ProductStoreService struct{}

var productStoreService *ProductStoreService

func NewProductStoreService() *ProductStoreService {
	if productStoreService == nil {
		l.Lock()
		if productStoreService == nil {
			productStoreService = new(ProductStoreService)
		}
		l.Unlock()
	}
	return productStoreService
}

type WalletModifyCallbackFunc func(*model.DsProductStore, int)

func (service *ProductStoreService) PlayerIncreaceNum(tx *gorm.DB, userID string, productID int64, num int, fundSourceID int64, callback WalletModifyCallbackFunc) error {
	fundchange := model.DsProductStoreChange{
		UserID:          userID,
		ProductID:       productID,
		UserType:        1,
		Number:          num,
		StoreSourceType: model.StoreSourceTypeOrderIncrease,
		FundSourceID:    fundSourceID,
	}
	fundchange.StoreChangeType = model.StoreIncrease
	if err := service.createFundchange(tx, &fundchange, false, callback); err != nil {
		Log.Error(err)
		return err
	}
	return nil
}

func (service *ProductStoreService) PlayerDecreaseNum(tx *gorm.DB, userID string, productID int64, num int, fundSourceID int64, callback WalletModifyCallbackFunc) error {
	fundchange := model.DsProductStoreChange{
		UserID:          userID,
		ProductID:       productID,
		UserType:        1,
		Number:          num,
		StoreSourceType: model.StoreSourceTypeOrderDecrease,
		FundSourceID:    fundSourceID,
	}

	fundchange.StoreChangeType = model.StoreDecrease
	if err := service.createFundchange(tx, &fundchange, false, callback); err != nil {
		Log.Error(err)
		return err
	}

	return nil
}

func (service *ProductStoreService) AdminIncreaceNum(tx *gorm.DB, userID string, productID int64, num int, fundSourceType model.StoreSourceType, fundSourceID int64, callback WalletModifyCallbackFunc) error {
	fundchange := model.DsProductStoreChange{
		UserID:          userID,
		ProductID:       productID,
		UserType:        2,
		Number:          num,
		StoreSourceType: fundSourceType,
		FundSourceID:    fundSourceID,
	}
	fundchange.StoreChangeType = model.StoreIncrease
	if err := service.createFundchange(tx, &fundchange, false, callback); err != nil {
		Log.Error(err)
		return err
	}
	return nil
}

func (service *ProductStoreService) AdminDecreaseNum(tx *gorm.DB, userID string, userType int, productID int64, num int, fundSourceType model.StoreSourceType, fundSourceID int64, callback WalletModifyCallbackFunc) error {
	fundchange := model.DsProductStoreChange{
		UserID:          userID,
		ProductID:       productID,
		UserType:        2,
		Number:          num,
		StoreSourceType: fundSourceType,
		FundSourceID:    fundSourceID,
	}

	fundchange.StoreChangeType = model.StoreDecrease
	if err := service.createFundchange(tx, &fundchange, false, callback); err != nil {
		Log.Error(err)
		return err
	}

	return nil
}

func (service *ProductStoreService) createFundchange(tx *gorm.DB, store *model.DsProductStoreChange, isFreezeEvent bool, callback WalletModifyCallbackFunc) (err error) {
	Log.WithField("store", *store).Debug("FundchangeService.createFundchange")
	//用户ID不能为空
	if store.ProductID == 0 {
		err = errors.New("ProductID is nil")
		return err
	}
	if store.UserID == "" {
		err = errors.New("UserID is nil")
		return err
	}
	//金额为0不添加账变
	if store.Number == 0 {
		return errors.New("StoreNum is 0")
	}

	fundchagneDao := dao.NewDsProductStoreChangeDao(tx)
	walletDao := dao.NewDsProductStoreDao(tx)

	wallet := model.DsProductStore{
		ProductID: store.ProductID,
	}

	err = walletDao.FindForUpdate(&wallet)
	if err != nil {
		Log.Error(err)
		return err
	}

	store.BeforeNumber = wallet.StoreNum

	//计算账变后金额
	if store.StoreChangeType == model.StoreIncrease {
		store.AfterNumber = store.BeforeNumber + store.Number
	}

	//计算账变后金额
	if store.StoreChangeType == model.StoreDecrease {
		store.AfterNumber = store.BeforeNumber - store.Number
	}

	if store.AfterNumber < 0 {
		err = errors.New("More than the balance")
		Log.WithFields(logrus.Fields{
			"AfterNumber": store.AfterNumber,
		}).Error(err)
		return err
	}

	if err := fundchagneDao.Create(store); err != nil {
		Log.Error(err)
		return err
	}
	wallet.StoreNum = store.AfterNumber

	//更新钱包统计数据
	if callback != nil {
		callback(&wallet, store.Number)
	}

	err = walletDao.Save(&wallet)
	if err != nil {
		Log.Error(err)
		return err
	}

	return nil
}
