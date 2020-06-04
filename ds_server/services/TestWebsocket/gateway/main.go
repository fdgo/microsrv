package main

import (
	webst "ds_server/services/TestWebsocket/websock_srv/proto"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/micro/go-micro/client"
	"net/http"
	"strconv"
)

var upGrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func main() {
	r := gin.Default()
	r.GET("/user/permission/ws", ping)
	r.Run(":8080")

}

//webSocket请求ping 返回pong
func ping(c *gin.Context) {
	//升级get请求为webSocket协议
	ws, err := upGrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		return
	}
	defer ws.Close()

	//读取ws中的数据
	mt, message, err := ws.ReadMessage()
	if err != nil {

	}
	var req webst.Request
	json.Unmarshal(message, &req)

	wst := webst.NewStreamerService("jz.micro.websock-srv.stream", client.DefaultClient)
	for {
		stream, err := wst.Stream(c)
		if err != nil {
			fmt.Println("err:", err)
			return
		}
		if err := stream.Send(&webst.Request{Count: int64(req.Count)}); err != nil {
			fmt.Println("err:", err)
			return
		}
		rsp, err := stream.Recv()
		if err != nil {
			fmt.Println("recv err", err)
			break
		}
		fmt.Printf("Sent msg %v got msg %v\n", req.Count, rsp.Count)
		err = ws.WriteMessage(mt, []byte(strconv.Itoa(int(rsp.Count))))
		if err != nil {
			break
		}
		// close the stream
		if err := stream.Close(); err != nil {
			fmt.Println("stream close err:", err)
		}
	}
}
