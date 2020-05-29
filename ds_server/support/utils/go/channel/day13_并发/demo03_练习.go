package main

import (
	"fmt"
	"time"
)

func main() {
	/*
	练习1：
		一条协程打印1000数字，另一条协程打印1000字母
	观察运行结果。。

	并发的程序的运行结果，每次都不一定相同。不同计算机设备执行，效果也不相同。
	 */
	 go printNum()
	 go printLetter()
	 time.Sleep(1*time.Second)
	 fmt.Println("main...over.")
}

func printNum(){
	for i:=1;i<=1000;i++{
		fmt.Println("子goroutine1中i：",i)
		time.Sleep(1)
	}
}
func printLetter(){
	for i:=1;i<=1000;i++{
		fmt.Printf("\t子goroutine中字母：%d,%c\n",i,i)
		time.Sleep(1)
	}
}
