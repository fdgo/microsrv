package main

import (
	"fmt"
	"time"
)

func main() {
	/*
		并发：go 关键字，启动一个goroutine，执行对应的函数。
	程序的执行过程：
		A：先创建主goroutine，
		B：初始化操作
		C：执行main()
		D：main()结束了，主goroutine随之结束，程序结束。

	go语言的并发：go关键字
		系统自动创建并启动主goroutine，执行对应的main()
		用于自己创建并启动子goroutine，执行对应的函数

		go 函数()//go关键创建并启动goroutine，然后执行对应的函数()，该函数执行结束，子goroutine也随之结束。

			子goroutine中执行的函数，往往没有返回值。如果有也会被舍弃。


	 */
	//1.启动子goroutine
	go hello()
	//2.打印内容。。
	fmt.Println("main....")

	time.Sleep(1*time.Second)//main函数进入睡眠
}

func hello() {
	fmt.Println("我是一个函数，在另外一个goroutine中执行的。。")
}
/*

 */
