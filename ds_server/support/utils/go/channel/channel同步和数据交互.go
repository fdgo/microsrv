package main

import (
	"fmt"
	"time"
)

func main()  {
	ch := make(chan string)
	go func() {
		defer fmt.Println("子协程调用完毕")
		for i:=0;i<20;i++{
			fmt.Println("i = ",i)
			time.Sleep(time.Second)
		}
		ch<-"abc"
	}()
	<-ch
}