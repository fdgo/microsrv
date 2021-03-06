package user

import (
	"ds_server/client"
	mygormdl "ds_server/models/user/gorm_mysql"
	useproto "ds_server/proto/user"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
)

func ConnectUs(c *gin.Context) {
	ret_resp := mygormdl.DsSysInfo{}
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	var rtin useproto.ConnectUsIn
	ret, err := client.UserClient.ConnectUs(c, &rtin)
	if err != nil {
		fmt.Println(400, 400, err.Error(), err.Error(), ret_resp, c)
		return
	}
	json.Unmarshal((*ret).Data, &ret_resp)
	fmt.Println((*ret).Httpcode, (*ret).Innercode, (*ret).Clientmsg, (*ret).Innermsg, ret_resp, c)
}
