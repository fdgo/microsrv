package main

import (
	"os"
	"fmt"
	"bufio"
	"net"
)

func main() {
	con,err := net.Dial("tcp","127.0.0.1:9008")
	if err!=nil{
		fmt.Println("err = ",err)
		return
	}
	defer con.Close()

	go func() {
		buf := make([]byte,1024)
		for{
			n,err := con.Read(buf)
			if err != nil{
				fmt.Println("net.Dail err = ",err)
				return
			}
			fmt.Println(string(buf[:n]))
		}
	}()

	for {
		inputReader := bufio.NewReader(os.Stdin)
		input, err := inputReader.ReadString('\n')
		if err != nil {
			fmt.Println("err = ",err)
			return
		}
		con.Write([]byte(input))
	}
}
