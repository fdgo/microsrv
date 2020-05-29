package cron

import (
	"log"
	"github.com/robfig/cron"
	"fmt"
	"time"
)
// 6个参数:  秒(0-59)， 分(0-59)， 时(0-23)， 一个月中某天(1-31)， 月(1-12)， 星期几(0-6)
func main_test() {
	log.Println("Starting...")

	c := cron.New()

	c.AddFunc("0 25 * * * *", func() { //每小时的25分钟(周期小时)  比如 2018-07-03 14:25:00     2018-07-03 15:25:00    2018-07-03 16:25:00
		fmt.Println("每小时的25分钟")
	})
	c.AddFunc("15 * * * * *", func() { //每分钟的15秒  (周期分钟)  比如 2018-07-03 14:25:15     2018-07-03 14:26:15    2018-07-03 14:27:15
		fmt.Println("每分钟的15秒")
	})

	c.AddFunc("0,8-59/2 * * * * *", func() { //0秒， 8秒到59秒（每2秒一次）， 所以时间为： 0,8,10,12,14...58
		fmt.Println("0秒， 8秒到59秒（每2秒一次）， 所以时间为： 0,8,10,12,14...58")
	})

	c.AddFunc("0 0 0 1 * *", func() { //每个月第一天夜里0点
		fmt.Println("每个月第一天夜里0点")
	})


	c.AddFunc("0 0 0 * * 6", func() { //每周日夜里0点
		fmt.Println("每周日夜里0点")
	})
	c.AddFunc("0 0 0 * * 0,6", func() { //每周六，周日夜里0点
		fmt.Println("每个月第一天夜里0点")
	})

	c.AddFunc("0 0 0 1 * *", func() { //每个月第一天夜里0点
		fmt.Println("每个月第一天夜里0点")
	})
	c.Start()
	//下面起到阻塞作用，没卵用
	t1 := time.NewTimer(time.Second * 10)
	for {
		select {
		case <-t1.C:
			t1.Reset(time.Second * 10)
		}
	}
}