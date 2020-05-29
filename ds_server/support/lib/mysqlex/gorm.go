package mysqlex

import (
	"ds_server/support/utils/constex"
	"log"
	"sync"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var gdbLock sync.Mutex
var mygsqlInstarnce *gorm.DB

// 得到唯一的主库实例
func MysqlInstanceg() *gorm.DB {
	if mygsqlInstarnce != nil {
		return mygsqlInstarnce
	}
	gdbLock.Lock()
	defer gdbLock.Unlock()

	if mygsqlInstarnce != nil {
		return mygsqlInstarnce
	}
	return mysqlDb()
}
func mysqlDb() *gorm.DB {
	db, err := gorm.Open("mysql", constex.MysqlCfg.Url)
	if err != nil {
		log.Fatal("dbhelper.InstanceDbMaster DB error ", err)
		return nil
	}
	db.DB().SetMaxIdleConns(100)
	db.DB().SetMaxOpenConns(100)
	db.DB().SetConnMaxLifetime(time.Duration(300) * time.Second)
	db = db.Set("gorm:table_options", "ENGINE=InnoDB CHARSET=utf8")

	db.SingularTable(true)

	mygsqlInstarnce = db

	return mygsqlInstarnce
}
