package main

import (
	"sync"
	"fmt"
	"math/rand"
	"time"
)

var n int
var wg sync.WaitGroup
var rwm sync.RWMutex
func main() {
	wg.Add(10)


	for i:=1;i<=5;i++{
		go read(i)
	}
	for i:=1;i<=5;i++{
		go write(i)
	}

	wg.Wait()
}

func write(i int){
	defer wg.Done()
	rand.Seed(time.Now().UnixNano())
	rwm.Lock()
	fmt.Println("写操作：",i,"即将写入数据。。")
	randNum:=rand.Intn(100)+1
	n = randNum
	fmt.Println("写操作：",i,"已经结束，写入了，",randNum)
	rwm.Unlock()

}

func read(i int){
	defer wg.Done()
	rwm.RLock()
	fmt.Println("读取操作：",i,"即将读取数据。。")
	v := n
	fmt.Println("读操作：",i,"读取了数据：",v)
	rwm.RUnlock()
}
