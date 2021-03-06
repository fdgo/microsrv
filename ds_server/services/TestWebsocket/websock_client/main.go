package main

import (
	proto "ds_server/services/TestWebsocket/websock_srv/proto"
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"net/url"
)

func main() {
	u := url.URL{Scheme: "ws", Host: "127.0.0.1:8080", Path: "/user/permission/ws"}
	var dialer *websocket.Dialer

	conn, _, err := dialer.Dial(u.String(), nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	//go timeWriter(conn)

	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			fmt.Println("read:", err)
			return
		}

		fmt.Println("webst_client recv: ",string(message))
	}
}

func timeWriter(conn *websocket.Conn) {
	var req proto.Request
	req.Count = 888
	ret,_ := json.Marshal(req)
	//for {
		//time.Sleep(time.Second * 2)
		conn.WriteMessage(websocket.TextMessage, ret)
	//}
}
