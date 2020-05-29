package main

import (
	"context"
	"log"
	"github.com/nsqio/go-nsq"
	"os"
)


type NSQConfig struct {
	Topic, Channel string
	PoolSize       int
	Ctx            context.Context
	NsqLookupdAddr []string
}

type NSQFetcher struct {
	topic, channel string
	workerPoolSize int
	ctx  context.Context
	err  chan error
	data chan []byte
	fetcherName string
	nsqlookupdAddr []string
}

func NewNSQFetcher(config *NSQConfig) *NSQFetcher {
	c := new(NSQFetcher)
	c.ctx = config.Ctx
	c.topic = config.Topic
	c.channel = config.Channel
	c.workerPoolSize = config.PoolSize
	c.err = make(chan error, 1)
	c.data = make(chan []byte, c.workerPoolSize)
	c.nsqlookupdAddr = config.NsqLookupdAddr
	return c
}

func (c *NSQFetcher) Start() {
	go c.run()
}

func (c *NSQFetcher) Stop() {
	close(c.data)
	close(c.err)
}

func (c *NSQFetcher) Fetch() <-chan []byte {
	return c.data
}

func (c *NSQFetcher) doRecover(err interface{}, id int) {
	go c.run()
}

func (c *NSQFetcher) run() {
	consumer, err := newNsqConsumer(c.topic, c.channel, c.nsqlookupdAddr)
	if err != nil {
		log.Printf("nsqex consumer start failed: %#v", err)
	}
	defer func() {
		if err := recover(); err != nil {
			consumer.StopConsume()
			c.doRecover(err, 0)
		}
	}()

	for {
		select {
		case <-c.ctx.Done():
			consumer.StopConsume()
			return
		case e := <-c.err:
			log.Printf("error: %s", e)
		case d := <-consumer.data:
			c.data <- d
		}
	}
}
type nsqConsumer struct {
	topic, channel string
	data           chan []byte
	consumer       *nsq.Consumer
}

func newNsqConsumer(topic, channel string, lookupdAddress []string) (*nsqConsumer, error) {
	var err error
	c := &nsqConsumer{
		topic:   topic,
		channel: channel,
		data:    make(chan []byte),
	}

	config := nsq.NewConfig()
	c.consumer, err = nsq.NewConsumer(topic, channel, config)
	if err != nil {
		return nil, err
	}
	c.consumer.SetLogger(log.New(os.Stderr, "", log.Flags()), nsq.LogLevelError)
	c.consumer.AddConcurrentHandlers(c, 2)

	err = c.consumer.ConnectToNSQLookupds(lookupdAddress)
	if err != nil {
		return nil, err
	}

	return c, nil
}

func (c *nsqConsumer) StopConsume() {
	c.consumer.Stop()
}

func (c *nsqConsumer) HandleMessage(msg *nsq.Message) error {
	c.data <- msg.Body
	return nil
}
