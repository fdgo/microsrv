package main

import (
	"fmt"
	"time"
)

var i = 1

func main() {

	ch1 := make(chan int)

	for i:=1;i<=3;i++{
		go func(i int) {
			for {
				if i >= 100{
					fmt.Println("goroutine ",i,"over结束。。")
					//close(ch1)
					break
				}else{
					ch1 <- i//2  2
					time.Sleep(10*time.Millisecond)
					i++
				}
			}
		}(i)
	}

	for i:=0;i<100;i++{
		fmt.Println(<-ch1)
	}

}
