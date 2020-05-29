package main

import (
	"log"

	"github.com/nsqio/go-nsq"
)

const topic = "mytopic"

type NSQWorker struct {
	job        chan []byte
	workerPool chan chan []byte
	quitChan   chan bool
	producer *nsq.Producer
}

func NewNSQWorker(wp chan chan []byte, addr string) *NSQWorker {
	var err error
	config := nsq.NewConfig()
	producer, err := nsq.NewProducer(addr, config)

	if err != nil {
		log.Fatalf("connect producer error: %v", err)
	}

	return &NSQWorker{
		job:        make(chan []byte),
		workerPool: wp,
		quitChan:   make(chan bool),
		producer:   producer,
	}
}

func (n *NSQWorker) Work() {
	go func() {
		for {
			n.workerPool <- n.job
			select {
			case <-n.quitChan:
				return
			case msg := <-n.job:
				n.publish(msg)
			}
		}
	}()
}

func (n *NSQWorker) Stop() {
	go func() {
		n.quitChan <- true
	}()
}

func (n *NSQWorker) publish(msg []byte) {
	if err := n.producer.Publish(topic, msg); err != nil {
		log.Println("nsqex publish error: ", err)
	}
}
