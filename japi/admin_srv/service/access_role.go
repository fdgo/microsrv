package service

import (
	"context"
	"encoding/json"
	adminproto "microservice/jzapi/proto/admin"
	"net/http"
)

func (s *service) GetAccessList(ctx context.Context, in *adminproto.GetAccessListInput) *adminproto.CommonOutput {
	output := new(adminproto.CommonOutput)
	type res struct {
		Uuid      string `json:"uuid"`
		UserName  string `json:"user_name"`
		RealName  string `json:"real_name"`
		RoleId    uint   `json:"role_id"`
		RoleName  string `json:"role_name"`
		Mobile    string `json:"mobile"`
		Address   string `json:"address"`
		LastLogin string `json:"last_login"`
		ChannelId uint	 `json:"channel_id"`
	}
	var tmpres res
	tmpslice := make([]res,0)
	rows, err := db.GetDB().Raw("select t_user_basic.uuid,user_name,real_name,role_id,role_name, mobile, address, last_login,channel_id from t_user_basic  inner join  t_user_back  on t_user_basic.uuid = t_user_back.uuid").Rows() // (*sql.Rows, error)
	if err != nil {
		output.HttpCode = http.StatusInternalServerError
		output.Code = http.StatusInternalServerError
		output.Detail = "服务正芒，请稍后再试！"
		output.Msg = "数据库连接失败！"
		output.Data = []byte{}
		return output
	}
	defer rows.Close()
	for rows.Next() {
		rows.Scan(&tmpres.Uuid, &tmpres.UserName, &tmpres.RealName, &tmpres.RoleId, &tmpres.RoleName, &tmpres.Mobile, &tmpres.Address, &tmpres.LastLogin, &tmpres.ChannelId)
		tmpslice = append(tmpslice,tmpres)
	}
	if len(tmpslice) == 0 {
		output.HttpCode = http.StatusOK
		output.Code = http.StatusOK
		output.Detail = "没有数据！"
		output.Msg = "数据库没有数据！"
		tmpdata, _ := json.Marshal(tmpslice)
		output.Data = tmpdata
		return output
	}
	output.HttpCode = http.StatusOK
	output.Code = http.StatusOK
	output.Msg = "查询成功！"
	output.Detail = "查询成功！"
	tmpdata, _ := json.Marshal(tmpslice)
	output.Data = tmpdata
	return output
}
func (s *service) GetRoleList(ctx context.Context, in *adminproto.GetRoleListInput) *adminproto.CommonOutput {

}
func (s *service) AddRole(ctx context.Context, in *adminproto.AddRoleInput) *adminproto.CommonOutput {

}
func (s *service) EditRole(ctx context.Context, in *adminproto.EditRoleInput) *adminproto.CommonOutput {

}
func (s *service) DeleteRole(ctx context.Context, in *adminproto.DeleteRoleInput) *adminproto.CommonOutput {

}
