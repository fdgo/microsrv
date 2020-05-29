package dao

import (
	"ds_server/support/utils/logger"
)

var Log *logger.Logger

func init() {
	Log = logger.InitLog()
}
