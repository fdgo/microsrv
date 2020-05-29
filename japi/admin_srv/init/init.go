package init

import (
	ser "microservice/jzapi/admin_srv/service"
)

// Init 初始化模型层
func InitSrvs() {
	ser.Init()
}
