package main

import (
	"net"
	"fmt"
	"os"
	"io"
)

func SendFile(path string,conn net.Conn)  {
	f,err := os.Open(path)
	if err !=nil{
		fmt.Println("os.open err = ",err)
		return
	}
	defer f.Close()

	buf := make([]byte,1024*4)
	for {
		n,err := f.Read(buf)
		if err !=nil {
			if err == io.EOF{
				fmt.Println("send ok!")
			}else{
				fmt.Println("send err!")
			}
			return
		}
		conn.Write(buf[:n])
	}
}
func main()  {
	fileName := "F:\\第007天——17_简单版并发服务器.avi"
	info,errStat := os.Stat(fileName)
	if errStat != nil{
		fmt.Println("err = ",errStat)
		return
	}
	fmt.Println("name=",info.Name())
	fmt.Println("size=",info.Size())

	con,errDial := net.Dial("tcp","127.0.0.1:9005")
	if errDial != nil{
		fmt.Println("Dial err = ",errDial)
		return
	}
	defer con.Close()

	_, errWrite := con.Write([]byte(info.Name()))
	if errWrite !=nil{
		fmt.Println("write err = ",errWrite)
		return
	}
	buf := make([]byte,1024)
	n,errRead := con.Read(buf)
	if errRead != nil{
		fmt.Println("con.Read err = ",errRead)
		return 
	}
	if "ok" == string(buf[:n]){
		SendFile(fileName,con)
	}
}