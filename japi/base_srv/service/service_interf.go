package service

import (
	"fmt"
	r "github.com/go-redis/redis"
	"microservice/jzapi/lib/redis"
	baseproto "microservice/jzapi/proto/base"
	"sync"
	"context"
)

var (
	s  *service
	Ca *r.Client
	m  sync.RWMutex
)

// service 服务
type service struct {
}

// Service 用户服务类
type Service interface {
	SendMobileCode(ctx context.Context, in *baseproto.SendMobileCodeIn) *baseproto.CommonOutput
	IsMobileCodeOk(ctx context.Context, in *baseproto.IsMobileCodeOkIn) *baseproto.CommonOutput
}

// GetService 获取服务类
func GetService() (Service, error) {
	if s == nil {
		return nil, fmt.Errorf("[GetService] GetService 未初始化")
	}
	return s, nil
}

// Init 初始化用户服务层
func Init() {
	m.Lock()
	defer m.Unlock()

	if s != nil {
		return
	}
	Ca = redis.Redis()
	s = &service{}
}
