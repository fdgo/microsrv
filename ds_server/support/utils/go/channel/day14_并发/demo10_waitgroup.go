package main

import (
	"fmt"
	"sync"
	"time"
	"math/rand"
)

func main() {
	/*
	同步等待组：WaitGourp，执行了wait的goroutine，要等待同步等待组中的其他的goroutine执行完毕。。
		内置的计数器：counter：0
		Add(),设置counter的值
		Done()，将counter减一，同Add(-1)

		以上两个方法可以设置counter的值，注意不能为负数，否则会引发恐慌。

		Wait(),哪个goroutine执行了，那么就会被阻塞，直到counter为0。解除阻塞
	 */
	var wg sync.WaitGroup
	//fmt.Printf("%T\n",wg)
	//fmt.Println(wg)
	wg.Add(2)

	go printNum1(&wg)
	go printNum2(&wg)

	wg.Wait() //main,进入阻塞状态，底层计数器为0,接触阻塞。。
	//time.Sleep(1*time.Second)
	fmt.Println("main。。接触阻塞。。结束了。。。")
}

func printNum1(wg *sync.WaitGroup) {
	rand.Seed(time.Now().UnixNano())
	for i := 1; i <= 100; i++ {
		fmt.Println("子goroutine1,i：", i)
		time.Sleep(time.Duration(rand.Intn(1000))) //
	}
	wg.Done() //计数器减一
}

func printNum2(wg *sync.WaitGroup) {
	for j := 1; j <= 100; j++ {
		fmt.Println("\t子goroutine2,j：", j)
		time.Sleep(time.Duration(rand.Intn(1000)))
	}
	wg.Done()
}

