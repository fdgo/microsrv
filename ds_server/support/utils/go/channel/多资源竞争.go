package main

import (
	"fmt"
	"time"
)
var ch = make(chan int)

func Printer(str  string)  {
	for _,data:= range str{
		fmt.Printf("%c",data)
		time.Sleep(time.Second)
	}
}
func Person1()  {
	Printer("hello")
	ch<-888
	close(ch)
}
func Person2()  {
	<-ch
	Printer("world")
}
func main()  {
	go Person1()
	go Person2()
	select{}
}