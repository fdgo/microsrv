package db

import (
	"github.com/jinzhu/gorm"
	"github.com/micro/go-micro/util/log"
	"microservice/jzapi/basic/cfg/config"
	"time"
	"fmt"
)

type db struct {
	Mysql Mysql `json："mysql"`
}

// Mysql mySQL 配置
type Mysql struct {
	URL               string `json:"url"`
	Enable            bool   `json:"enabled"`
	MaxIdleConnection int    `json:"maxIdleConnection"`
	MaxOpenConnection int    `json:"maxOpenConnection"`
}

func initMysql() {
	log.Logf("[initMysql] 初始化Mysql")

	c := config.C()
	cfg := &db{}

	err := c.App("db", cfg)
	if err != nil {
		log.Logf("[initMysql] %s", err)
	}

	if !cfg.Mysql.Enable {
		log.Logf("[initMysql] 未启用Mysql")
		return
	}
	mysqlDB, err = gorm.Open("mysql", cfg.Mysql.URL)
	if err != nil {
		log.Fatal("failed to connect database：", err)
		panic(err)
	}
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return "t_" + defaultTableName
	}
	mysqlDB.SingularTable(true)

	mysqlDB.LogMode(true)

	mysqlDB.Callback().Create().Replace("gorm:update_time_stamp", updateTimeStampForCreateCallback)
	mysqlDB.Callback().Update().Replace("gorm:update_time_stamp", updateTimeStampForUpdateCallback)
	mysqlDB.Callback().Delete().Replace("gorm:delete", deleteCallback)
	// 最大连接数
	mysqlDB.DB().SetMaxOpenConns(cfg.Mysql.MaxOpenConnection)
	// 最大闲置数
	mysqlDB.DB().SetMaxIdleConns(cfg.Mysql.MaxIdleConnection)
	// 空闲连接最大15妙关闭
	mysqlDB.DB().SetConnMaxLifetime(time.Second*15)

	// 激活链接
	if err = mysqlDB.DB().Ping(); err != nil {
		log.Fatal(err)
	}
	log.Logf("[initMysql] Mysql 连接成功")
}
// updateTimeStampForCreateCallback will set `CreatedOn`, `ModifiedOn` when creating
func updateTimeStampForCreateCallback(scope *gorm.Scope) {
	if !scope.HasError() {
		nowTime := time.Now().Unix()
		if createTimeField, ok := scope.FieldByName("CreatedAt"); ok {
			if createTimeField.IsBlank {
				_ = createTimeField.Set(nowTime)
			}
		}

		if modifyTimeField, ok := scope.FieldByName("ModifiedAt"); ok {
			if modifyTimeField.IsBlank {
				_ = modifyTimeField.Set(nowTime)
			}
		}
	}
}

// updateTimeStampForUpdateCallback will set `ModifiedOn` when updating
func updateTimeStampForUpdateCallback(scope *gorm.Scope) {
	if _, ok := scope.Get("gorm:update_column"); !ok {
		_ = scope.SetColumn("ModifiedAt", time.Now().Unix())
	}
}

func deleteCallback(scope *gorm.Scope) {
	if !scope.HasError() {
		var extraOption string
		if str, ok := scope.Get("gorm:delete_option"); ok {
			extraOption = fmt.Sprint(str)
		}

		deletedOnField, hasDeletedOnField := scope.FieldByName("DeletedAt")

		if !scope.Search.Unscoped && hasDeletedOnField {
			scope.Raw(fmt.Sprintf(
				"UPDATE %v SET %v=%v%v%v",
				scope.QuotedTableName(),
				scope.Quote(deletedOnField.DBName),
				scope.AddToVars(time.Now().Unix()),
				addExtraSpaceIfExist(scope.CombinedConditionSql()),
				addExtraSpaceIfExist(extraOption),
			)).Exec()
		} else {
			scope.Raw(fmt.Sprintf(
				"DELETE FROM %v%v%v",
				scope.QuotedTableName(),
				addExtraSpaceIfExist(scope.CombinedConditionSql()),
				addExtraSpaceIfExist(extraOption),
			)).Exec()
		}
	}
}

func addExtraSpaceIfExist(str string) string {
	if str != "" {
		return " " + str
	}
	return ""
}
