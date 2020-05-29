package main

import (
	"net"
	"fmt"
	"strings"
	"time"
)

var onlineMap map[string]Client
var message = make(chan string)

type Client struct {
	C    chan string //用户发送数据的管道
	Name string      //用户名
	Addr string      //网络地址
}

func HandleCon(con net.Conn) {
	defer con.Close()
	cliAddr := con.RemoteAddr().String()
	cli := Client{make(chan string), cliAddr, cliAddr}
	onlineMap[cliAddr] = cli

	go WriteMsgToClient(cli, con)

	message <- "[" + cli.Addr + "]" + cli.Name + ": login"

	cli.C <- "I am here!"

	isQuit := make(chan bool)
	hasData := make(chan bool)
	
	go func() {
		buf := make([]byte, 1024)
		for {
			n, err := con.Read(buf)
			if n == 0 { //对方断开或者出问题
				isQuit<-true
				fmt.Println("con read err = ", err)
				return
			}
			msg := buf[:n-1] //nc 测试，多一个换行
			if len(msg) == 3 && string(msg) == "who"{
				con.Write([]byte("user list:\n"))
				for _,tmp := range onlineMap{
					src := tmp.Addr + ":" + tmp.Name + "\n"
					con.Write([]byte(src))
				}
			}else if len(string(msg)) >=8 && string(msg[:6]) == "rename" {
				name := strings.Split(string(msg),"|")[1]
				cli.Name = name
				onlineMap[cliAddr] = cli
				con.Write([]byte("rename ok!\n"))
			} else {
				message <- "[" + cli.Addr + "]" + cli.Name + ":" + string(msg)
			}
			hasData <- true
		}
	}()
	for {
		select {
		case <-isQuit:
			delete(onlineMap,cliAddr)
			message<- cliAddr + "login out!"
		case <-hasData:
		case <-time.After(60*time.Second):
			delete(onlineMap,cliAddr)
			message<- cliAddr + "time out,leave!"
			return
		}
	}
}
func WriteMsgToClient(cli Client, con net.Conn) {
	for msg := range cli.C {
		con.Write([]byte(msg + "\n"))
	}
}
func Manager() {
	onlineMap = make(map[string]Client)
	for {
		msg := <-message
		for _, cli := range onlineMap {
			cli.C <- msg
		}
	}
}
func main() {
	lis, err := net.Listen("tcp", ":8006")
	if err != nil {
		fmt.Println("net.Lis err = ", err)
		return
	}
	defer lis.Close()

	go Manager()

	for {
		con, err := lis.Accept()
		if err != nil {
			fmt.Println("net.Accept err =", err)
			continue
		}
		go HandleCon(con)
	}
}
