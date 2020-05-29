package main

import (
	"sync"
	"fmt"
	"time"
)

func main() {
	/*
	互斥锁：锁头对象(struct)
		互斥：锁定，解锁
		有两个指针方法：
			Lock(),上锁，
				阻塞的：goroutine上锁，其他的goroutine处于阻塞状态。。
			Unlock(),开锁
				unlock of unlocked mutex


	 */

	 var mutex sync.Mutex

	 fmt.Println("main，即将锁定mutex。。")
	 mutex.Lock()
	 fmt.Println("main已经锁定mutex")
	 for i:=1;i<=3;i++{
	 	go func(i int) {
	 		fmt.Println("子goroutine",i,"即将锁定mutex。。")
	 		mutex.Lock()//阻塞
	 		fmt.Println("子goroutine",i,"已经锁定。。")
		}(i)
	 }

	 time.Sleep(5*time.Second)
	 fmt.Println("main即将解锁。。")
	 mutex.Unlock()
	 fmt.Println("main已经解锁。。")
	 fmt.Println("main想再解锁一次。。")
	 //mutex.Unlock() //panic: sync: unlock of unlocked mutex

	 //time.Sleep(3*time.Second)
}
