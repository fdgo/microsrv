package main

import (
	"time"
	"fmt"
)
func main()  {
	timer := time.NewTimer(5*time.Second)
	fmt.Println("current time:",time.Now())

	for{//错误， timer.C 只有一次
		t :=<-timer.C
		fmt.Println("t = :",t)
	}

}

func main01()  {
	timer := time.NewTimer(5*time.Second)
	fmt.Println("current time:",time.Now())

	//2s后，往timer.C写入数据，有数据后即可读取
	t :=<-timer.C
	fmt.Println("t = :",t)
}

