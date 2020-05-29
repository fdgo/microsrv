package main

import (
	"fmt"
	"time"
)

func main() {

	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		time.Sleep(1*time.Second)
		ch1 <- 100
		//close(ch1)
	}()
	go func(){
		ch2 <- 200
	}()

	//结束循环的条件：通道结束，或者超时10次
	count:=0
	out:for {//1次，2次，3次
		//time.Sleep(1*time.Second)
		select {
		case data,ok := <-ch1:
			if !ok{
				fmt.Println("通道关闭了。。")
				break out
			}
			fmt.Println("ch1中读取数据：", data)
		case data := <-ch2:
			fmt.Println("ch2中读取数据：", data)
		case <-time.After(2 * time.Second):
			fmt.Println("超时2s。。。。")
			count++
			if count == 5{
				break out
			}
		//default:
		//	fmt.Println("default。。。")
		}
	}
}
