package idgenerator

import "ds_server/support/utils/logger"

var Log *logger.Logger

const ID_GENERATOR = "IDGenerator"

func init() {
	Log = logger.InitLog()

	Log.Infoln("IDGenerator init")
	RegisterService(ID_GENERATOR, newSQLIdGenerator)
}
