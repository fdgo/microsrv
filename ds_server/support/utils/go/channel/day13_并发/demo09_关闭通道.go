package main

import (
	"fmt"
	"time"
)

func main() {
	/*
	从通道中读取数据：
		data := <- chan
		data,ok := <- chan
			同value,ok:=map[key]类似

		ok的数值为true，通道正常的，读取到的data数据有效可用
		ok的数值为false，通道关闭，读取到data是类型的零值。

	通道的关闭：发送方如果数据写入完毕，可以关闭通道，用于同志接受方数据传递完毕。
		g1-->写入数据

		g2-->读取数据

	 */
	 ch1 := make(chan int)
	 go sendData(ch1)

	 //读取数据
	 for{
	 	time.Sleep(1*time.Second)
	 	data ,ok:= <- ch1//1,2,3..10,
	 	if !ok{
	 		fmt.Println("读取完毕，通道关闭了。。",ok)
	 		break
		}
	 	fmt.Println("main中读取到数据：",data,ok)


	 }
}

func sendData(ch1 chan int){
	for i:=1;i<=10;i++{
		ch1 <- i// 1,2,3...10
	}
	fmt.Println("发送方，写入数据完毕。。。")
	close(ch1)//发送方：关闭通道，用于通知接受方没有数据了
}
