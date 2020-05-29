package main

import (
	"fmt"
	"time"
)

func main() {
	/*
	缓冲通道：自带一块缓冲区，可以暂时存储数据，如果缓冲区满了， 那么才会阻塞

	非缓冲通道：默认创建的通道，都是非缓冲
	 */
	//1.非缓冲通道
	ch1 := make(chan int)
	fmt.Println("非缓冲通道：", len(ch1), cap(ch1)) //0, 0
	go func() {
		//time.Sleep(3*time.Second)
		<-ch1 //阻塞
	}()

	ch1 <- 100 //阻塞
	fmt.Println("写完了。。")

	//2.缓冲通道,缓冲区满了才会阻塞
	/*
	ch2 := make(chan int, 5)

	go func() {
		data1 := <-ch2
		fmt.Println("获取数据1：", data1)
	}()
	fmt.Println("缓冲通道：", len(ch2), cap(ch2)) //0,5
	ch2 <- 1
	fmt.Println(len(ch2), cap(ch2))
	ch2 <- 2
	ch2 <- 3
	fmt.Println(len(ch2), cap(ch2)) //3,5
	ch2 <- 4
	ch2 <- 5
	fmt.Println(len(ch2), cap(ch2)) //5, 5
	ch2 <- 6                        //阻塞
	fmt.Println("main。。。over。。")
	*/
	//3.
	ch3 := make(chan string,5)
	fmt.Printf("%T\n",ch3)
	go sendData(ch3)

	for {
		time.Sleep(100*time.Millisecond)
		data,ok:=<-ch3
		fmt.Println("\t读取数据：",data)
		if !ok{
			fmt.Println("读取完毕。。")
			break
		}
	}
}

func sendData(ch3 chan string){
	for i:=1;i<=100;i++{
		ch3 <- fmt.Sprint("数据：",i)//1,2,3,4,5,6,7
		fmt.Println("已经写出数据：",i)//1,2,3,4,5
	}
	close(ch3)
}

/*
练习1：缓冲通道
子goroutine从缓冲通道中读取数据
另一个子goroutine从该缓冲通道中写入数据
 */