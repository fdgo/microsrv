package main

import (
	"fmt"
	"time"
)

func main() {
	/*
	练习1：创建并启动一个子 goroutine，打印100个数字，要保证在main goroutine结束前结束。
	通道实现。
练习2：创建并启动两个子goroutine，一个打印100个数字，另一个打印100个字母，要保证在main goroutine结束前结束。
	 */


	 ch1 :=make(chan bool)
	 fmt.Println("main-->",ch1)
	 go printNum(ch1)//引用传递
	 go printLetter(ch1)

	 <- ch1
	 <- ch1
	 fmt.Println("mian...over..")

}

func printNum(ch1 chan bool){
	fmt.Println("fun....",ch1)
	for i:=1;i<=100;i++{
		fmt.Println(i)
		time.Sleep(10*time.Millisecond)
	}
	ch1<-true
}

func printLetter(ch1 chan bool){
	for i:=1;i<=100;i++{
		fmt.Printf("\t%d,%c\n",i,i)
		time.Sleep(10*time.Millisecond)
	}
	ch1 <- true
}
