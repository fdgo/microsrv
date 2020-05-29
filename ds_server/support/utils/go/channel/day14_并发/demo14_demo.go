package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan int)

	for i:=1;i<=4;i++{
		go func(i int) {
			for tickts:= range ch1{
				fmt.Println("售票口",i,"：",tickts)
			}
			fmt.Println("售票口：",i,"结束。。")
		}(i)
	}

	for i:=1;i<=100;i++{
		ch1 <- i
	}
	close(ch1)

	time.Sleep(1*time.Second)
	fmt.Println("main..over。。。。")

}
