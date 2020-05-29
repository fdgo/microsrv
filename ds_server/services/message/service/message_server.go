package service

import (
	"ds_server/services/message/dao"
	"ds_server/services/message/model"
	"ds_server/services/message/service/dto"
	db "ds_server/support/lib/mysqlex"
	"ds_server/support/utils/idgenerator"
	"ds_server/support/utils/logex"
	"errors"
)

type MessageService struct {
}

var messageService *MessageService

func NewMessageService() *MessageService {
	if messageService == nil {
		l.Lock()
		if messageService == nil {
			messageService = &MessageService{}
		}
		l.Unlock()
	}
	return messageService
}

func (service MessageService) GetMessageByID(id int64) (request model.DsMessage, err error) {
	db := db.MysqlInstanceg()
	messageDao := dao.NewDsMessageDao(db)
	if id == 0 {
		return model.DsMessage{}, errors.New("id is null")
	}
	data := model.DsMessage{ID: id}
	err = messageDao.Get(&data)
	if err != nil {
		Log.Error(err)
		return model.DsMessage{}, nil
	}
	return data, nil
}
func (service MessageService) GetMessageDetailByID(id int64) (request model.DsMessageDetail, err error) {
	db := db.MysqlInstanceg()
	messageDao := dao.NewDsMessageDetailDao(db)
	if id == 0 {
		return model.DsMessageDetail{}, errors.New("id is null")
	}
	data := model.DsMessageDetail{ID: id}
	err = messageDao.Get(&data)
	if err != nil {
		Log.Error(err)
		return model.DsMessageDetail{}, err
	}
	return data, nil
}

func (service MessageService) CreateMessage(m model.DsMessage) (err error) {
	db := db.MysqlInstanceg()
	tx := db.Begin()
	defer closeTx(tx, &err)
	messageDao := dao.NewDsMessageDao(tx)
	err = messageDao.Create(&m)
	if err != nil {
		logex.Error(err)
		return err
	}
	messageurlDao := dao.NewDsMessageUrlDao(tx)

	if m.Url != "" {
		err = messageurlDao.Create(&model.DsMessageUrl{MessageID: m.ID, Url: m.Url})
		if err != nil {
			logex.Error(err)
			return err
		}
	}
	return nil
}

func (service MessageService) CreateMessageDetail(m model.DsMessageDetail) (err error) {
	db := db.MysqlInstanceg()
	tx := db.Begin()
	defer closeTx(tx, &err)
	messagedDao := dao.NewDsMessageDetailDao(tx)
	err = messagedDao.Create(&m)
	if err != nil {
		logex.Error(err)
		return err
	}
	messagedfDao := dao.NewDsMessageDao(tx)
	if m.MessageID > 0 {
		err = messagedfDao.Create(&model.DsMessage{ID: m.MessageID, Swith: 2, DsMessageDetailId: m.ID, MessageType: 2, Title: m.Title, Name: m.Name})
		if err != nil {
			logex.Error(err)
			return err
		}
	}
	return err
}

func (service MessageService) SearchMessagerPaging(condition *model.MessageSearchCondition, pageNum int, pageSize int) (request []model.DsMessage, count int, err error) {
	rowbound := model.NewRowBound(pageNum, pageSize)
	return service.searchMessage(condition, &rowbound)
}

func (service MessageService) SearchMessagerOutPaging(condition *model.MessageSearchCondition) (request []model.DsMessage, count int, err error) {
	return service.searchMessage(condition, nil)
}

func (service MessageService) searchMessage(condition *model.MessageSearchCondition, rowbound *model.RowBound) (request []model.DsMessage, count int, err error) {
	db := db.MysqlInstanceg()
	messageDao := dao.NewDsMessageDao(db)
	request, count, err = messageDao.SeachMessage(condition, rowbound)
	if err != nil {
		Log.Error(err)
		return nil, 0, err
	}
	messageurlDao := dao.NewDsMessageUrlDao(db)

	for i, v := range request {
		re, err := messageurlDao.Find(&model.DsMessageUrl{MessageID: v.ID})
		if err != nil {
			Log.Error(err)
			return nil, 0, err
		}
		request[i].DsMessageUrl = re
	}
	return request, count, nil
}

func (service MessageService) GetMessageID() (count int64, err error) {
	gen := idgenerator.Instance()
	id := gen.GenerateLongID("user")
	return id, nil
}

func (service MessageService) CreateMessageUrl(m model.DsMessageUrl) error {
	db := db.MysqlInstanceg()
	messageDao := dao.NewDsMessageUrlDao(db)
	return messageDao.Create(&m)
}

func (service MessageService) SearchBanner(bannerType model.BannerType) (request []dto.Banners, err error) {
	db := db.MysqlInstanceg()
	messageDao := dao.NewDsBannerUrlDao(db)

	result, err := messageDao.SeachBanner(bannerType)
	if err != nil {
		Log.Error(err)
		return []dto.Banners{}, err
	}

	modelMap := make(map[model.BannerType][]model.DsBannerUrl)
	for _, v := range result {
		if k, ok := modelMap[v.BannerType]; ok {
			ss := append(k, v)
			modelMap[v.BannerType] = ss
		} else {
			modelMap[v.BannerType] = []model.DsBannerUrl{v}
		}
	}

	ss2 := []dto.Banners{}
	for k, v := range modelMap {
		banner := dto.Banners{BannerType: k, BannerUrl: v}
		ss2 = append(ss2, banner)
	}
	return ss2, err
}
