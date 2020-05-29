package main

import (
	"sync"
	"fmt"
	"time"
)

func main() {
	/*
	 读写锁：sync.RWMutex
		指针方法：
			Lock(),写锁定
			Unlock()，写解锁

			RLock()，读锁定
			RUnlock()，读解锁

	锁定的规则：
		读写锁的使用中
			写操作都是互斥的
			读和写是互斥的
			读和读不互斥

	理解为：
		可以多个goroutine同时读取数据
		但是写只允许一个goroutine写数据


	g1,g2,g3,g4
		g1-->写锁定
		g1-->读
			g2,g3,g4
	 */

	var rwm sync.RWMutex

	for i := 1; i <= 3; i++ {
		go func(i int) {
			fmt.Printf("goroutine %d，尝试读锁定。。\n", i)
			rwm.RLock()
			fmt.Printf("goroutine %d，已经读锁定了。。\n", i)
			time.Sleep(5 * time.Second)
			fmt.Printf("goroutine %d,读解锁。。\n", i)
			rwm.RUnlock()
		}(i)
	}

	time.Sleep(1*time.Second)
	fmt.Println("main..尝试写锁定。。")
	rwm.Lock()
	fmt.Println("main。。已经写锁定了。。")
	rwm.Unlock()
	fmt.Printf("main。。写解锁。。。")

}
