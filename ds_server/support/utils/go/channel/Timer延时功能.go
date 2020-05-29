package main

import (
	"time"
	"fmt"
)

func func1()  {
	time.Sleep(time.Second*5)
}
func func2()  {
	<-time.After(time.Second*5)
	fmt.Println("5秒时间到")
}
func func3()  {
	timer := time.NewTimer(5*time.Second)
	fmt.Println("current time:",time.Now())

	//5s后，往timer.C写入数据，有数据后即可读取
	t :=<-timer.C
	fmt.Println("t = :",t)
}
func main()  {

}