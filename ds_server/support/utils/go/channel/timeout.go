package main

import (
	"time"
	"fmt"
)

func main() {
	ch := make(chan int)
	quit := make(chan bool)

	go func() {
		for {
			select {
			case num := <-ch:
				fmt.Println("num = ", num)
			case <-time.After(3 * time.Second):
				fmt.Println("time out!")
				quit <- true

			}
		}
	}()
	for i := 0; i<5; i++  {
		ch <-i
		time.Sleep(time.Second)
	}
	<-quit
	fmt.Println("main over ...")
}
