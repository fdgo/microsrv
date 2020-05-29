package main

import (
	"fmt"
	"sync"
)


type MyMap struct {
	map1 map[string]string
	rwm sync.RWMutex
}

//将键值对存入到map中
func (m *MyMap) Put(key,value string){
	m.rwm.Lock()
	defer m.rwm.Unlock()
	m.map1[key] = value
}

func (m *MyMap) Get(key string) string{
	m.rwm.RLock()
	defer m.rwm.RUnlock()
	return m.map1[key]
}

func main() {
	/*
	并发读写map
	 */
/*
	 var rwm sync.RWMutex
	 map1:=make(map[string] string)
	 for i:=1;i<=100;i++{
	 	go func(i int) {
	 		rwm.Lock()
	 		map1[fmt.Sprintf("key,%d",i)] = fmt.Sprintf("数据%d ",i)
	 		rwm.Unlock()
		}(i)
	 }
	 time.Sleep(2*time.Second)
	fmt.Println(map1)
*/
	var wg sync.WaitGroup
	wg.Add(10)
	m1:=make(map[string]string) //不支持并发
	var rwm sync.RWMutex
	m2:=MyMap{m1,rwm}


	for i:=1;i<=100;i++{
		go func(i int) {
			m2.Put(fmt.Sprint("key",i),fmt.Sprint("data",i))
			wg.Done()
		}(i)
	}
	wg.Wait()
	fmt.Println("main，读取数据。。。。")
	for i:=1;i<=10;i++{
		fmt.Println(m2.Get(fmt.Sprint("key",i)))
	}

	fmt.Println("main..over....")
}
