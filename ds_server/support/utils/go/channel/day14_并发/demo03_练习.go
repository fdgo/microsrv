package main

import (
	"fmt"
	"time"
	"math/rand"
)

func main() {
	ch1 := make(chan int, 5)
	ch2 := make(chan bool)//判断结束
	rand.Seed(time.Now().UnixNano())
	//写出数据：生产者
	go func() {

		for i := 1; i <= 100; i++ {
			ch1 <- i//1,2,3,3,4,5,6,7
			fmt.Println("写出数据：", i)//1,2,3,4,5,6
			//time.Sleep(time.Duration(rand.Intn(1000))*time.Millisecond)
		}
		close(ch1)
	}()

	//读取数据：消费者
	go func() {
		for data := range ch1 {
			fmt.Println("\t消费者1：", data)
			//time.Sleep(time.Duration(rand.Intn(1000))*time.Millisecond)
		}
		ch2 <- true
	}()
	go func() {
		for data := range ch1 {//1
			fmt.Println("\t消费者2：", data)
			time.Sleep(time.Duration(rand.Intn(1000))*time.Millisecond)
		}
		ch2 <- true
	}()


	<- ch2
	<- ch2
	fmt.Println("main...over...")
}
