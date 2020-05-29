package main

import (
	"time"
	"fmt"
)

func main()  {
	ticker := time.NewTicker(time.Second*3)
	i := 0
	for{//正确，可以周期性打印
		<-ticker.C
		i++
		fmt.Println("i=",i)
		if i==5{
			ticker.Stop()
			break
		}
	}
}
