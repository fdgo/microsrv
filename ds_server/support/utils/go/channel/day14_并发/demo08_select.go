package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		time.Sleep(3 * time.Second)
		data := <-ch1
		fmt.Println("子goroutine中：", data)
		fmt.Println("goroutine...")
	}()

	select {
	case ch1 <- 100://阻塞
		fmt.Println("ch1中写出数据。。")

	case ch2 <- 200://阻塞
		fmt.Println("ch2中写出数据。。")

	case <-time.After(3 * time.Second)://阻塞
		fmt.Println("超时了。。。")

	//default:
	//	fmt.Println("default..")
	}

	time.Sleep(1*time.Second)
	//fmt.Println(<-ch1)
}
