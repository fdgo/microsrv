package main

import (
	"time"
)

var (
	jobQueue = make(chan []byte, 64)
)
func Send(src []byte) {
	jobQueue <- src
}
type NsqConfig struct {
	workerQueue chan chan []byte
	timeout     time.Duration
	addr        string
}

func (d *NsqConfig) Start() {
	w := NewNSQWorker(d.workerQueue, d.addr)
	w.Work()
	go func() {
		for {
			select {
			case job := <-jobQueue:
				go func(job []byte) {
					w := <-d.workerQueue
					w <- job
				}(job)
			}
		}

	}()
}
func main() {
	nsqcfg := &NsqConfig{
		workerQueue:make(chan chan []byte, 2),
		timeout:5*time.Second,
		addr:"192.168.207.128:4150",
	}
	nsqcfg.Start()
	Send([]byte("66777"))
	select {}

}
