package basic

import (
	"github.com/wangmhgo/microservice-project/go-micro-learn/microservices/part1/user-srv/basic/config"
	"github.com/wangmhgo/microservice-project/go-micro-learn/microservices/part1/user-srv/basic/db"
)

func Init() {
	config.Init()
	db.Init()
}
