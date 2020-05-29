package main

import (
	"sync"
	"sync/atomic"
	"fmt"
	"runtime"
)

func main() {
	/*
	临时对象池：看成复用某些数据的容器。
		可伸缩，安全的
	sync.Pool,
		Put(),向pool中存储数据
		Get(),从pool获取一个数据，如果pool没有数据，那么会调用New对应函数，创建一个


	如果程序中垃圾回收机制来收垃圾，那么临时对象池中的对象，都被销毁
	 */


	 var count int64 // 0
	 fun:= func() interface{}{
	 	return atomic.AddInt64(&count,1)
	 }

	 pool:=sync.Pool{New:fun}
	 //获取数据
	 fmt.Println(pool.Get())

	 //向pool存储，获取数据
	 pool.Put(10)
	 pool.Put(2)
	 pool.Put(8)
	 pool.Put(15)
	 fmt.Println(pool.Get())//pool中如果有多个数据，任意获取一个，如果没有，就通过new对应的函数产生一个

	 //执行GC()

	 runtime.GC()//
	 pool.New = nil
	 fmt.Println(pool.Get())//nil
}
