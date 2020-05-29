package db

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/micro/go-micro/util/log"
	glocfg "microservice/jzapi/basic/cfg"
	"sync"
)

var (
	inited  bool
	mysqlDB *gorm.DB
	m       sync.RWMutex
)

func init() {
	glocfg.Register(initDB)
}

// initDB 初始化数据库
func initDB() {
	m.Lock()
	defer m.Unlock()

	var err error

	if inited {
		err = fmt.Errorf("[initDB] db 已经初始化过")
		log.Logf(err.Error())
		return
	}

	initMysql()

	inited = true
}

// GetDB 获取db
func GetDB() *gorm.DB {
	return mysqlDB
}
func CloseDB() {
	defer mysqlDB.Close()
}