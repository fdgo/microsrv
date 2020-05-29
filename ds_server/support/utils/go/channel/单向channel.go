package main

import "fmt"

func producer(out chan <- int)  {
	for  i:=0; i<10;i++{
		out<-i*i
	}
	close(out)
}
func consumer(in <-chan int)  {
	for num := range in{
		fmt.Println("num = ",num)
	}
}
//单向无法转换为双向
func main()  {
	ch := make(chan int)
	//生产者
	go producer(ch)

	//消费者
	consumer(ch)
}
