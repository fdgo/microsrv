package main

import (
	"fmt"
	"runtime"
	"time"
)

func init(){
	//1.获取逻辑cpu的数量
	fmt.Println("逻辑CPU的核数：",runtime.NumCPU())
	//2.设置go程序执行的最大的：[1,256]
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
	/*
	init(),同main()，特殊的函数，由系统自动调用执行--->main goroutine
	runtime包：
	 */
	fmt.Println("helloworld")
	go fun1()
	go fun2()
	time.Sleep(10*time.Millisecond)
}

func fun1(){
	for i:=1;i<=100;i++{
		fmt.Println(i)
	}
}
func fun2(){
	for i:=1;i<=100;i++{
		fmt.Printf("\t%c\n",i)
	}
}
