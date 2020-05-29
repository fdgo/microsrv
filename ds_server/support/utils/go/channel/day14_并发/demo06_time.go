package main

import (
	"time"
	"fmt"
)

func main() {
	/*
	time包对于chan的操作
		1.Timer：计时器
			NewTimer(duration)-->*Timer对象：struct：字段C <-chan Time


		2.After(duration)--> <-chan Time

	 */

	 //1.创建计时器
	 timer1:=time.NewTimer(3*time.Second)
	 fmt.Printf("%T\n",timer1)//*time.Timer
	 fmt.Println(time.Now())
	 time1:=<-timer1.C//<-chan time.Time
	 //fmt.Printf("%T\n",timer1.C)
	 fmt.Println(time1)

	 //2.使用After(),返回值<- chan Time,同Timer.C
	 fmt.Println("--------------------")
	 ch1:=time.After(5*time.Second)
	 fmt.Println(time.Now())
	 time2:=<- ch1
	 fmt.Println(time2)
}
