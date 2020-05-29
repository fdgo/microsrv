package main

import "fmt"

func main() {
	/*
	单任务：
	 */
	 fmt.Println("main....")
	 fun1()
	 fmt.Println("主函数。。。")
	 fun2()
	 fun3()
}
func fun1(){
	fmt.Println("helloworld。。。")
}

func fun2(){
	for i:=1;i<=5;i++{
		fmt.Println("fun2...i:",i)
	}
}

func fun3(){
	fmt.Println("我是函数3.。")
	fun2()
}
