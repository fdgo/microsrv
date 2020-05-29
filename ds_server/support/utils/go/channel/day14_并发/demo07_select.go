package main

import (
	"fmt"
	"time"
)

func main() {
	/*
	select，分支语句，专门用于通道读写操作的。
	select{
	case chan读/写:
		分支1
	case chan读/写:
		分支2：
	...
	default:
	}

	执行流程：
		1.如果有多个都可以执行，select会随机执行一个可运行的case。
		2.如果没有case可运行，如果有default，执行default，如果没有default，
			它将阻塞，直到有case可运行。

	 */

	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		time.Sleep(5*time.Second)
		ch1 <- 100
	}()

	go func() {
		time.Sleep(5*time.Second)
		ch2 <- 200
	}()
	select {
	case data := <-ch1:
		fmt.Println("ch1中读取数据了:", data)
	case data := <-ch2:
		fmt.Println("ch2中读取数据了：", data)
	default:
		fmt.Println("执行了default。。。")
	}
}

/*
练习：select和time结合
select{
case <-ch1:
case <-ch2:
case <- time.After(3*time.Second)://?
defalut//...
}
 */
