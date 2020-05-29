package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var count int = 4
	var wg sync.WaitGroup
	wg.Add(5)

	// 新建 cond
	var mutex sync.Mutex
	cond := sync.NewCond(&mutex)

	for i := 0; i < 5; i++ {//0,1,2,3,4
		go func(i int) {//g1,g2,g3,g4,g5
		                //0, 1, 2, 3, 4
			// 争抢互斥锁的锁定
			cond.L.Lock() //g1

			// 条件是否达成count:1
			for count > i {//0,1,2
				cond.Wait()//g0,g1,g2
				fmt.Printf("收到一个通知 goroutine%d\n", i)
			}

			fmt.Printf("goroutine%d 执行结束\n", i)

			cond.L.Unlock()
			wg.Done()
		}(i)
	}

	// 确保所有 goroutine 启动完成
	time.Sleep(time.Millisecond * 20)
	// 锁定一下，我要改变 count 的值
	fmt.Println("broadcast...")
	cond.L.Lock()
	count -= 1 // 3
	cond.Broadcast()
	fmt.Println("第一次：广播结束。。")
	cond.L.Unlock()

	time.Sleep(2*time.Second)
	fmt.Println("-------------------------")
	fmt.Println("signal...")
	cond.L.Lock()
	count -= 2 // 1
	cond.Signal()
	fmt.Println("第二次：单发通知结束。。")
	cond.L.Unlock()


	time.Sleep(2*time.Second)
	fmt.Println("------------------------")
	fmt.Println("broadcast...")
	cond.L.Lock()
	count -= 1 // 0
	cond.Broadcast()
	fmt.Println("第三次：广播结束。。")
	cond.L.Unlock()

	wg.Wait()
}

