package handler

import (
	"context"
	pb "ds_server/proto/message"
	"ds_server/services/message/model"
	"ds_server/services/message/service"
	"ds_server/support/utils/logex"
	"ds_server/support/utils/logex"
	"encoding/json"
	"fmt"
)

type MessageHandler struct {
	Message *service.MessageService
}

func (message *MessageHandler) GetMessageByID(c context.Context, req *pb.IdRequest, rsp *pb.Response) error {
	mService := service.NewMessageService()
	result, err := mService.GetMessageByID(req.ID)
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

func (message *MessageHandler) GetMessageDetailByID(c context.Context, req *pb.IdRequest, rsp *pb.Response) error {
	mService := service.NewMessageService()
	result, err := mService.GetMessageDetailByID(req.ID)
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

func (message *MessageHandler) SelectMessage(c context.Context, req *pb.SelectMessageRequest, rsp *pb.Response) error {
	fmt.Println("============", logger.FormatStruct(req))
	mService := service.NewMessageService()
	condition := new(model.MessageSearchCondition)
	condition.MessageType = req.MessageType
	condition.UserID = req.UserID
	condition.UserType = int(req.UserType)
	//rowBound := dao.NewRowBound(req.PageNum, req.PageSize)
	condition.Swith = 1
	result, count, err := mService.SearchMessagerPaging(condition, int(req.PageNum), int(req.PageSize))
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

func (message *MessageHandler) CreateMessage(c context.Context, req *pb.CreateRequest, rsp *pb.Response) error {
	ress := model.DsMessage{}
	ress.Swith = 2
	err := json.Unmarshal(req.Data, &ress)
	if err != nil {
		logex.Error(err)
		return err
	}
	mService := service.NewMessageService()
	return mService.CreateMessage(ress)
}

func (message *MessageHandler) CreateMessageDetail(c context.Context, req *pb.CreateRequest, rsp *pb.Response) error {
	ress := model.DsMessageDetail{}
	err := json.Unmarshal(req.Data, &ress)
	if err != nil {
		logex.Error(err)
		return err
	}
	mService := service.NewMessageService()
	err = mService.CreateMessageDetail(ress)
	if err != nil {
		logex.Error(err)
		return err
	}
	return nil
}

func (message *MessageHandler) GetMessageID(c context.Context, req *pb.IdRequest, rsp *pb.IdResponse) error {
	mService := service.NewMessageService()
	id, err := mService.GetMessageID()
	if err != nil {
		logex.Error(err)
		return err
	}
	rsp.ID = id
	return nil
}

func (message *MessageHandler) CreateMessageUrl(c context.Context, req *pb.UrlRequest, rsp *pb.Response) error {
	mService := service.NewMessageService()
	ress := model.DsMessageUrl{}
	ress.MessageID = req.MessageID
	ress.Url = req.Url
	err := mService.CreateMessageUrl(ress)
	if err != nil {
		logex.Error(err)
		return err
	}
	return nil
}

func (message *MessageHandler) SearchBanner(c context.Context, req *pb.SelectBannerRequest, rsp *pb.Response) error {
	fmt.Println("============", logger.FormatStruct(req))
	mService := service.NewMessageService()
	result, err := mService.SearchBanner(model.BannerType(req.BannerType))
	if err != nil {
		return err
	}
	b, err := json.Marshal(result)
	if err != nil {
		return err
	}
	rsp.Data = b
	fmt.Println("============", logger.FormatStruct(rsp))
	return nil
}
