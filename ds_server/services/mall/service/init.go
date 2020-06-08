package service

import (
	"ds_server/support/utils/logex"
	"sync"

	"github.com/jinzhu/gorm"
)

var (
	Log *logger.Logger
	l   sync.Mutex
)

func init() {
	Log = logger.InitLog()
}

const (
	TIME_FORMAT_WITH_MS         = "2006-01-02 15:04:05.000"
	TIME_FORMAT                 = "2006-01-02 15:04:05"
	TIME_FORMAT_WO_SEC_COMPACT  = "200601021504"
	TIME_FORMAT_COMPACT         = "20060102150405"
	TIME_FORMAT_WITH_MS_COMPACT = "20060102150405000"
	DATE_FORMAT                 = "2006-01-02"
	DATE_FORMAT_COMPACT         = "20060102"
	MONTH_FORMAT                = "2006-01"
)

func closeTx(tx *gorm.DB, err *error) {
	r := recover()
	if r != nil {
		tx.Rollback()
		Log.Error(r)
		return
	}

	if *err != nil {
		tx.Rollback()
		Log.Errorf("%+v", *err)
		return
	}
	tx.Commit()
}
