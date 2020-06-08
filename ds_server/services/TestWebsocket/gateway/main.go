package main

import (
	"context"
	webpro "ds_server/services/TestWebsocket/websock_srv/proto"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/micro/go-micro/client"
	"io"
	"log"
	"net/http"
)

var WebStClient = webpro.NewStreamerService("ds.srv.webst", client.DefaultClient)

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
func ping(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	conn, err := upGrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		return
	}
	defer conn.Close()
	if err := Do(WebStClient, conn); err != nil {
		log.Fatal("Echo: ", err)
		return
	}
}
func isExpectedClose(err error) bool {
	if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway) {
		log.Println("Unexpected websocket close: ", err)
		return false
	}
	return true
}
func Do(cli webpro.StreamerService, ws *websocket.Conn) error {
	var req webpro.Request
	//err := ws.ReadJSON(&req)
	//if err != nil {
	//	return err
	//}
	//go func() {
	//	for {
	//		if _, _, err := ws.NextReader(); err != nil {
	//			break
	//		}
	//	}
	//}()
	//log.Printf("Received Request: %v", req)
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
		fmt.Println("gateway recv: ",rsp)
		err = ws.WriteJSON(rsp.Count)
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


//webSocket请求ping 返回pong
//func ping(c *gin.Context) {
//	//升级get请求为webSocket协议
//	ws, err := upGrader.Upgrade(c.Writer, c.Request, nil)
//	if err != nil {
//		return
//	}
//	defer ws.Close()
//
//	//读取ws中的数据
//	mt, message, err := ws.ReadMessage()
//	if err != nil {
//
//	}
//	var req webst.Request
//	json.Unmarshal(message, &req)
//
//	wst := webst.NewStreamerService("jz.micro.websock-srv.stream", client.DefaultClient)
//	for {
//		stream, err := wst.Stream(c)
//		if err != nil {
//			fmt.Println("err:", err)
//			return
//		}
//		if err := stream.Send(&webst.Request{Count: int64(req.Count)}); err != nil {
//			fmt.Println("err:", err)
//			return
//		}
//		rsp, err := stream.Recv()
//		if err != nil {
//			fmt.Println("recv err", err)
//			break
//		}
//		fmt.Printf("Sent msg %v got msg %v\n", req.Count, rsp.Count)
//		err = ws.WriteMessage(mt, []byte(strconv.Itoa(int(rsp.Count))))
//		if err != nil {
//			break
//		}
//		// close the stream
//		if err := stream.Close(); err != nil {
//			fmt.Println("stream close err:", err)
//		}
//	}
//}
