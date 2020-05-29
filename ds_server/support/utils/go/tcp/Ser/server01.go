package main

import (
	"net"
	"fmt"
	"strings"
)

//物理层，数据链路层，网络层，传输层，会话层，表示层，应用层
//链路层(MAC)，网络层(IP)，传输层(TCP,UDP)，应用层(FTP,HTTP)1
//MAC 物理地址    IP逻辑地址

func HandleCon(con net.Conn)  {
	defer con.Close()
	addr := con.RemoteAddr().String()

	buf := make([]byte,1024)
	for{
		n,err := con.Read(buf)
		if err !=nil{
			fmt.Println("err = ",err.Error())
			return
		}
		fmt.Printf("[%s]: %s\n",addr,string(buf[:n-1]))
		if "exit" == string(buf[:n-1]){
			fmt.Println(addr," exit")
			return
		}
		con.Write( []byte(strings.ToUpper(string(buf[:n-1]))))
	}
}
func main()  {
	//监听
	lis,err := net.Listen("tcp","127.0.0.1:9008")
	if err !=nil{
		fmt.Println("err = ",err)
		return
	}
	defer lis.Close()
	for {
		con,err := lis.Accept()
		if err != nil{
			fmt.Println("err = ",err)
			continue
		}
		go HandleCon(con)
	}
}
