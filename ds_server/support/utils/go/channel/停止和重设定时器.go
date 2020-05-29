package main

import (
	"time"
	"fmt"
)

func main()  {
	timer := time.NewTimer(10*time.Second)
	timer.Reset(1*time.Second)
	<-timer.C
	fmt.Println("时间到")
}
func main01()  {
	timer := time.NewTimer(3*time.Second)
	go func() {
		<-timer.C
		fmt.Println("子协程可以打印了，因为定时器时间到了")
	}()
	timer.Stop() //上面那句打印不出来了

	for {

	}
}