package client

import (
	userpro "ds_server/proto/user"
	basepro "ds_server/proto/base"
	"ds_server/support/utils/constex"
	"github.com/micro/go-micro/client"
)

var (
	UserClient = userpro.NewUserService(constex.SRV_USER, client.DefaultClient)
	BaseClient = basepro.NewBaseService(constex.SRV_BASE, client.DefaultClient)
)