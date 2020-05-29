package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	nsqcf := &NSQConfig{
		Topic:"mytopic",
		Channel:"ch",
		PoolSize:2,
		Ctx:context.Background(),
		NsqLookupdAddr:[]string{"192.168.207.128:4161"},
	}
	nsqf := NewNSQFetcher(nsqcf)
	nsqf.Start()
	fmt.Println(string(<-nsqf.Fetch()))
	time.Sleep(10*time.Second)
}
