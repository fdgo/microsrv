package main

import (
	"fmt"
	"sync/atomic"
)

func main() {
	/*
	原子性操作：sync.atomic
		当通过原子性方法操作某个数值时，其他的goroutine不能再访问当前的数值变量。

		针对于数值：
		操作：
			原子加/减
			交换
			比较并交换：CAS
			存储
			加载


	同步：
		上锁
		//。。。。

		解锁

	 */

	 var n int64 = 3
	 //n += 1 //n = n + 1


	 fmt.Println("n的原始数值：",n)

	 //原子加
	 newN:=atomic.AddInt64(&n,1)
	 fmt.Println("新的数据：",newN)
	 fmt.Println("n的新值：",n)




	 atomic.SwapInt64(&n,9)//更新数据

	 atomic.CompareAndSwapInt64(&n,4,10)//比较


}
