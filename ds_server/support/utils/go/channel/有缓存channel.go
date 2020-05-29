package main

import (
	"fmt"
	"time"
)

func main()  {
	ch := make(chan int,1024)
	fmt.Printf("len(ch) = %d, cap(ch) = %d\n",len(ch),cap(ch))
	go func() {
		for i:=0; i<3; i++{
			fmt.Printf("子协程: i = %d\n",i)
			ch <-i
			fmt.Printf("len(ch) = %d, cap(ch) = %d\n",len(ch),cap(ch))
		}
	}()
	time.Sleep(time.Second*5)
	for i:= 0;i<3; i++{
		num := <-ch
		fmt.Println("num = \n",num)
	}
}
