package main

import (
	"fmt"
	"time"
)

func main() {
	/*
	语言：并发的程序之间传递数据：
	go语言中主张：应该通过数据传递来实现共享内存，而不是通过共享内存来实现消息传递。

	channel，通道
	概念：专门用于goroution之间，传递数据的。类似于通信的消息队列

	语法：数据类型，make()，也是引用类型的数据
		关联一种相关的数据类型
		nil chan，同map一样，不能使用。

	操作：goroutine可以从chan中读取数据，另一个goroutine从中写入数据
			操作符：<-
		从chan中读取数据,data := <- chan
		向chan中写入数据, chan <- data

	阻塞：对于chan的读取和写入的操作，都是阻塞式的。
		阻塞式：导致程序暂时不能执行，直到接触阻塞。

		从chan中读取数据：阻塞式，直到另一个goroutine向通道中写入数据，解除阻塞。
		向chan中写入数据：阻塞式，直到另一个goroutine从通道中将数据读取出，解除阻塞。

	安全：通道本身是安全的。同一时间，只能允许一个goroutine来读或写。

	 */

	 var ch1 chan int
	 fmt.Println(ch1) //<nil>,
	 fmt.Printf("%T\n",ch1)//chan int

	 ch1 = make(chan int)
	 fmt.Println(ch1) //0xc4200200c0

	 ch2 := make(chan bool)

	 go func() {
	 	fmt.Println("子goroutine。。。")
	 	data:=<- ch1//阻塞式，从通道中读取数据
	 	time.Sleep(5*time.Second)
	 	fmt.Println("子goroutine从通道中读取到main传来的数据是：",data)
		ch2 <- true//向通道中写入数据，表示结束
	 }()

	 ch1 <- 100 //阻塞式，main goroutine向通道中写入数据


	 <- ch2
	 fmt.Println("main..over")
}



