package main

import "fmt"

func main() {
	/*
	通道关闭：可以读取，写入？
	 */

	ch1 := make(chan int)
	go func() {
		ch1 <- 100
		close(ch1)
		ch1 <- 200//关闭的通道，无法写入数据
	}()

	data, ok := <-ch1
	fmt.Println("main读取数据：", data, ok)
	data, ok = <-ch1
	fmt.Println("main读取数据：", data, ok)
	data, ok = <-ch1
	fmt.Println("main读取数据：", data, ok)
	data, ok = <-ch1
	fmt.Println("main读取数据：", data, ok)
}
