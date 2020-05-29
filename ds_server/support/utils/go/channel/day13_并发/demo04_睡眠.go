package main

import (
	"fmt"
	"time"
)

func main() {
	/*
	练习2：
	子goroutine1，打印1-5个数字，先睡眠，每睡眠250毫秒，打印一个数字
	子goroutine2，调研A-E个字母，先睡眠，每睡眠400毫秒，打印一个字母。
观察程序的运行结果。
	主goroutine中睡眠3000毫秒
	 */

	 go printNum()
	 go printLetter()


	 time.Sleep(3000*time.Millisecond)
}

func printNum(){
	for i:=1;i<=5;i++{
		time.Sleep(250*time.Millisecond)
		fmt.Print(i,"\t")
	}
}
func printLetter(){
	for i:=65;i<70;i++{
		time.Sleep(400*time.Millisecond)
		fmt.Printf("%c\t",i)
	}
}
