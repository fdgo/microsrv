package main

import (
	"fmt"
	"sync"
	"time"
	"math/rand"
)


var tickts = 100//全局变量，仅一份
var wg sync.WaitGroup
func main() {
	/*
	练习题：模拟火车站卖票
	火车票100张，4个售票口出售(4个goroutine)。
	 */
	 wg.Add(4)
	 go saleTickts("售票口1")//g1
	 go saleTickts("售票口2")//g2
	 go saleTickts("售票口3")//g3
	 go saleTickts("售票口4")//g4


	 wg.Wait()
	 fmt.Println("该趟列车所有票卖光了。。程序结束")

}

func saleTickts(name string){
	rand.Seed(time.Now().UnixNano())
	for {//1张
		//g1,g2,g3,g4
		if tickts > 0{//g1,g2,g3,g4
			time.Sleep(time.Duration(rand.Intn(1000)))
			fmt.Println(name,"：", tickts)//1,0,-1,-2
			tickts--//0,-1,-2,-3
		}else{
			fmt.Println(name,",结束卖票。。")
			break
		}
	}
	wg.Done()
}



