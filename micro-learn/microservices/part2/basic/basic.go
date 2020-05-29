package basic

import (
	"github.com/wangmhgo/microservice-project/go-micro-learn/microservices/part2/basic/config"
	"github.com/wangmhgo/microservice-project/go-micro-learn/microservices/part2/basic/db"
	"github.com/wangmhgo/microservice-project/go-micro-learn/microservices/part2/basic/redis"
)

func Init() {
	config.Init()
	db.Init()
	redis.Init()
}
