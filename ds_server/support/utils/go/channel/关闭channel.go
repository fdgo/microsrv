package main

import "fmt"

func main()  {
	ch := make(chan int)
	go func() {
		for i:=0;i<100;i++{
			ch<-i
		}
		close(ch)
	}()
	//for  {
	//	value,ok := <-ch
	//	if ok == false{
	//		break
	//	}else {
	//		fmt.Println("num = ",value)
	//	}
	//}
	//等价于
	for value := range ch{
		fmt.Println(value)
	}
}
