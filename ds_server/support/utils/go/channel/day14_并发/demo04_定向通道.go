package main

import (
	"time"
	"fmt"
)

func main() {
	/*
	通道：默认都是双向
		读
		写
		make(chan Type)
	定向通道：也叫单向通道，只读，或只写
		只读：make(<- chan Type)，只能读取数据，不能写入数据
			<- chan
		只写：make(chan <- Type)，只能写入数据，不能读取数据
			chan <- data


	创建通道时，采用单向通道，没有意义的。都是创建双向的。
	传递参数的时候使用：
		函数，只有写入数据，
		函数，只有读取数据，

	语法级别：保证通道的操作安全

	 */

	 //1.双向通道
	 ch1 := make(chan string)

	 go fun1(ch1)
	 time.Sleep(3*time.Second)
	 data:=<-ch1
	 fmt.Println("main，接受到数据：",data)
	 ch1 <- "你要上学么？"


	 fun2(ch1)
	 fun3(ch1)
}

func fun1(ch1 chan string){
	ch1 <- "我是小明"
	time.Sleep(2*time.Second)
	data:=<-ch1
	fmt.Println("回应：",data)
}

//功能：只有写入数据
func fun2(ch1 chan <- string){
	//只能写入
	//<- ch1 //invalid operation: <-ch1 (receive from send-only type chan<- string)
}

//个哦能：只有读取数据
func fun3(ch1 <- chan string){
	<-ch1
	//ch1 <- "hello" //invalid operation: ch1 <- "hello" (send to receive-only type <-chan string)
}
