package main

import (
	"fmt"
	"time"
)

func main() {
	//for i := 1; i <= 3; i++ {
	//	go func(i int){
	//		fmt.Println("第", i, "个goroutine。。")
	//	}(i)
	//	//time.Sleep(10*time.Millisecond)
	//}

	for i := 1; i <= 3; i++ {
		//i := i
		go func() {
			fmt.Println("第", i, "个goroutine。。")
		}()
	}
	time.Sleep(1 * time.Second)
}

//func fun(i int) {
//	fmt.Println("第", i, "个goroutine。。")
//
//}
