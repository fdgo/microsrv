package main

import (
	"sync"
	"fmt"
	"time"
)

func main() {
	/*
	条件变量：sync.Cond,多个goroutine等待或接受通知的集合地

		L ：Locker接口

		Cond条件变量，总是要和锁结合使用。

		3个指针方法：
		Wait(),等待goroutine等待接收通知，Signal(),Broadcast(),解除阻塞
		Signal(),发送通知，一个
		Broadcast(),广播，方法给所有人

	数据=数值

	g1,g2,g3

	g1-->n
		g1-->wait()
	g3-->n
		g3-->wait()

	g2改变数据，并且发送通知：

	main groutine
		-->condition false
			wait()
	g1
		-->condition true
			发通知

	 */

	 var mutex sync.Mutex
	 cond := sync.Cond{L:&mutex}
	 condition := false

	go func() {
		time.Sleep(1*time.Second)
		cond.L.Lock()
		fmt.Println("子goroutine已经锁定。。。")
		fmt.Println("子goroutine更改条件数值，并发送通知。。")
		condition = true//更改数值
		cond.Signal() //发送通知：一个goroutine
		fmt.Println("子gorutine。。。继续。。。")
		time.Sleep(5*time.Second)
		fmt.Println("子groutine解锁。。")
		cond.L.Unlock()
	}()



	 cond.L.Lock()
	 fmt.Println("main..已经锁定。。。")
	 if !condition{
	 	fmt.Println("main.。即将等待。。。")
	 	//wait()
	 	// 1.wait尝试解锁，
	 	// 2.等待--->当前的groutine进入了阻塞状态，等待被唤醒：signal(),broadcast()
	 	// 3.一旦被唤醒后，又会锁定
	 	cond.Wait()
	 	fmt.Println("main.被唤醒。。")
	 }
	 fmt.Println("main。。。继续")
	 fmt.Println("main..解锁。。。")
	 cond.L.Unlock()

}
