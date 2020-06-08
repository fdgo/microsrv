package handler

import (
	"ds_server/support/utils/logex"
)

var Log *logger.Logger

func init() {
	Log = logger.InitLog()
}
