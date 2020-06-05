package user

import (
	"context"
	"ds_server/client"
	reqmdl "ds_server/models/user/gin_req"
	rspmdl "ds_server/models/user/gin_rsp"
	useproto "ds_server/proto/user"
	"ds_server/support/utils/param"
	rsp "ds_server/support/utils/rsp"
	"ds_server/support/utils/trace"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"io"
	"log"
	"net/http"
)

func Regist(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	var in reqmdl.Regist_req
	if err := c.ShouldBindJSON(&in); err != nil {
		rsp.RespGin(400, 400, "输入有误,请重写输入!", "参数有误", err.Error(), c)
		return
	}
	isok, _ := param.IsParam(in)
	if !isok {
		rsp.RespGin(400, 400, "输入有误,请重写输入!", "参数有误", "参数有误", c)
		return
	}
	ctx, ok := trace.ContextWithSpan(c)
	if !ok {
		log.Warn("不存在context")
		c.JSON(200, gin.H{
			"code": -1,
			"msg":  "不存在context",
		})
		return
	}
	var rtin useproto.RegistIn
	rtin.Invcodeagent = in.InvCodeAgent
	rtin.Mobile = in.Mobile
	rtin.Pwd = in.Pwd
	rtin.Verifycode = in.VerifyCode
	rtin.ClientIp = c.Request.RemoteAddr
	ret, err := client.UserClient.Regist(ctx, &rtin)
	if err != nil {
		rsp.RespGin(400, 400, "注册失败!", err.Error(), "注册失败!", c)
		return
	}
	rsp.RespGin((*ret).Httpcode, (*ret).Innercode, (*ret).Clientmsg, (*ret).Innermsg, string((*ret).Data), c)
}
func Login(c *gin.Context) {
	var ret_resp rspmdl.Login_rsp
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	var in reqmdl.Login_req
	if err := c.ShouldBindJSON(&in); err != nil {
		rsp.RespGin(400, 400, "输入有误,请重写输入!", "参数有误", ret_resp, c)
		return
	}
	isok, _ := param.IsParam(in)
	if !isok {
		rsp.RespGin(400, 400, "输入有误,请重写输入!", "参数有误", ret_resp, c)
		return
	}
	var logIn useproto.LoginIn
	logIn.Mobidecode = in.VfCode
	logIn.ClientIp = c.Request.RemoteAddr
	logIn.Mobile = in.Mobile
	logIn.Pwd = in.Pwd
	logIn.Type = in.Type
	ret, err := client.UserClient.Login(c, &logIn)
	if err != nil {
		rsp.RespGin(400, 400, err.Error(), err.Error(), ret_resp, c)
		return
	}
	json.Unmarshal((*ret).Data, &ret_resp)
	rsp.RespGin((*ret).Httpcode, (*ret).Innercode, (*ret).Clientmsg, (*ret).Innermsg, ret_resp, c)
}
func ModifyLoginPwd(c *gin.Context) {
	//wmh
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	var in reqmdl.ModifyLoginPwd_req
	if err := c.ShouldBindJSON(&in); err != nil {
		rsp.RespGin(400, 400, err.Error(), err.Error(), "编辑失败！", c)
		return
	}
	isok, _ := param.IsParam(in)
	if !isok {
		rsp.RespGin(400, 400, "输入有误,请重写输入!", "参数有误！", "编辑失败！", c)
		return
	}
	var etin useproto.ModifyLoginPwdIn
	etin.Mobile = in.Mobile
	etin.Oldpwd = in.OldPwd
	etin.Newpwd = in.NewPwd
	ret, err := client.UserClient.ModifyLoginPwd(c, &etin)
	if err != nil {
		rsp.RespGin(400, 400, err.Error(), err.Error(), "", c)
		return
	}
	rsp.RespGin((*ret).Httpcode, (*ret).Innercode, (*ret).Clientmsg, (*ret).Innermsg, string((*ret).Data), c)
}
func ForgetPwd(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	var in reqmdl.ModifyBasicPwd_req
	if err := c.ShouldBindJSON(&in); err != nil {
		rsp.RespGin(400, 400, err.Error(), err.Error(), "编辑失败！", c)
		return
	}
	isok, _ := param.IsParam(in)
	if !isok {
		rsp.RespGin(400, 400, "输入有误,请重写输入!", "参数有误！", "编辑失败！", c)
		return
	}
	var etin useproto.ModifyBasicPwdIn
	etin.Content = in.Content
	etin.Tag = in.Tag
	etin.Vfcode = in.VfCode
	etin.Mobile = in.Mobile
	ret, err := client.UserClient.ModifyBasicPwd(c, &etin)
	if err != nil {
		rsp.RespGin(400, 400, err.Error(), err.Error(), "", c)
		return
	}
	rsp.RespGin((*ret).Httpcode, (*ret).Innercode, (*ret).Clientmsg, (*ret).Innermsg, string((*ret).Data), c)
}
func SetPaypwd(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	var in reqmdl.SetPaypwd
	if err := c.ShouldBindJSON(&in); err != nil {
		rsp.RespGin(400, 400, err.Error(), err.Error(), "！", c)
		return
	}
	isok, _ := param.IsParam(in)
	if !isok {
		rsp.RespGin(400, 400, "输入有误,请重写输入!", "参数有误！", "！", c)
		return
	}
	var etin useproto.SetPaypwdIn
	etin.Paypwd = in.PayPwd
	etin.Vfcode = in.VerifyCode
	etin.Mobile = in.Mobile
	etin.Uuid = c.Request.Header.Get("X-Head-Uuid")
	ret, err := client.UserClient.SetPaypwd(c, &etin)
	if err != nil {
		rsp.RespGin(400, 400, err.Error(), err.Error(), "", c)
		return
	}
	rsp.RespGin((*ret).Httpcode, (*ret).Innercode, (*ret).Clientmsg, (*ret).Innermsg, string((*ret).Data), c)

}
func ModifyPayPwd(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	var in reqmdl.ModifyPayPwd_req
	if err := c.ShouldBindJSON(&in); err != nil {
		rsp.RespGin(400, 400, err.Error(), err.Error(), "参数有误！", c)
		return
	}
	isok, _ := param.IsParam(in)
	if !isok {
		rsp.RespGin(400, 400, "输入有误,请重写输入!", "参数有误！", "！", c)
		return
	}
	var etin useproto.ModifyPayPwdIn
	etin.Newpwd = in.NewPwd
	etin.Vfcode = in.VerifyCode
	etin.Mobile = in.Mobile
	etin.Uuid = c.Request.Header.Get("X-Head-Uuid")
	ret, err := client.UserClient.ModifyPayPwd(c, &etin)
	if err != nil {
		rsp.RespGin(400, 400, err.Error(), err.Error(), "", c)
		return
	}
	rsp.RespGin((*ret).Httpcode, (*ret).Innercode, (*ret).Clientmsg, (*ret).Innermsg, string((*ret).Data), c)
}

func UserInfo(c *gin.Context) {

}

//****************************************************************************
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

func isExpectedClose(err error) bool {
	if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway) {
		log.Println("Unexpected websocket close: ", err)
		return false
	}
	return true
}
func Do(cli useproto.UserService, ws *websocket.Conn) error {
	var req useproto.WsIn
	err := ws.ReadJSON(&req)
	if err != nil {
		return err
	}
	go func() {
		for {
			if _, _, err := ws.NextReader(); err != nil {
				break
			}
		}
	}()
	log.Printf("Received Request: %v", req)
	stream, err := cli.ServerStream(context.Background(), &req)
	if err != nil {
		return err
	}
	defer stream.Close()
	for {
		rsp, err := stream.Recv()
		if err != nil {
			if err != io.EOF {
				return err
			}
			break
		}
		fmt.Println("888:", rsp.Data)
		err = ws.WriteJSON(string(rsp.Data))
		if err != nil {
			if isExpectedClose(err) {
				log.Println("Expected Close on socket", err)
				break
			} else {
				return err
			}
		}
	}
	return nil
}

func WebsocketMsg(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Fatal("Upgrade: ", err)
		return
	}
	defer conn.Close()
	if err := Do(client.UserClient, conn); err != nil {
		log.Fatal("Echo: ", err)
		return
	}
}
