package main

import (
	"net"
	"fmt"
	"os"
	"io"
)

func main()  {
	lis,err := net.Listen("tcp","127.0.0.1:9005")
	if err != nil{
		fmt.Println("net.Listen err = ",err)
		return
	}
	defer lis.Close()

	con,err1 := lis.Accept()
	if err1 != nil{
		fmt.Println("accept err = ",err)
		return
	}
	defer con.Close()

	buf := make([]byte,1024)
	n, err2 := con.Read(buf)
	if err2 != nil{
		fmt.Println("Read err = ",err2)
		return
	}
	filename := string(buf[:n])

	con.Write([]byte("ok"))

	RecvFile(filename,con)
}
func RecvFile(filename string, conn net.Conn)  {
	f,err := os.Create(filename)
	if err !=nil{
		fmt.Println("os create err = ",err)
		return
	}
	defer f.Close()

	buf := make([]byte,1024*4)
	for{
		n,err := conn.Read(buf)
		if err!=nil{
			if err == io.EOF{
				fmt.Println("recv ok!")
			}else {
				fmt.Println("con recv err = ",err)
			}
			return
		}
		if n == 0 {
			fmt.Println("recv ok!")
			return
		}
		f.Write(buf[:n])
	}

}