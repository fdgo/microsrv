package basic

import (
	"github.com/wangmhgo/microservice-project/go-micro-learn/microservices/part3/basic/config"
	"github.com/wangmhgo/microservice-project/go-micro-learn/microservices/part3/basic/db"
	"github.com/wangmhgo/microservice-project/go-micro-learn/microservices/part3/basic/redis"
)

func Init() {
	config.Init()
	db.Init()
	redis.Init()
}
