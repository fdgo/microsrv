package queue

import (
	"github.com/nsqio/go-nsq"
	"github.com/wangmhgo/microservice-project/joe-micro/lib/log"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

var producer *nsq.Producer

var addrNsqLookups []string

var logLevel nsq.LogLevel

var consumers []*nsq.Consumer

var config  *nsq.Config


func Init(addrNsq string , addrNsqLookup []string, maxInFlight int,debug bool) {
	log.Info(" nsq  链接中。。。")
	if  addrNsq==""{
		addrNsq="127.0.0.1:4150"
	}
	if  len(addrNsqLookup)==0 {
		addrNsqLookups = []string{
			"127.0.0.1:4161","127.0.0.1:4261",
		}
	}else{
		addrNsqLookups =addrNsqLookup
	}
	config =  nsq.NewConfig()
	config.MaxInFlight=maxInFlight
	p, err := nsq.NewProducer(addrNsq, config)
	if err != nil {
		log.Fatal(err)
		return
	}

	logLevel = nsq.LogLevelWarning
	if debug {
		logLevel = nsq.LogLevelInfo
	}

	p.SetLogger(log.NsqLogger(), logLevel)
	producer = p
	if err = p.Ping(); err != nil {
		log.Fatal(err)
		return
	}
	log.Info(" nsq  链接成功 ")
	go func() {
		ch := make(chan os.Signal)
		signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
		<-ch
		gracefulStop()
	}()
}

func Publish(topic string, body []byte, delay ...time.Duration) (err error) {
	if len(delay) == 0 {
		err = producer.Publish(topic, body)
	} else {
		err = producer.DeferredPublish(topic, delay[0], body)
	}
	return
}







func Subscribe(topic string, channel string, handler nsq.Handler) (err error) {
	c, err := nsq.NewConsumer(topic, channel, config)
	if err != nil {
		log.Fatal(err)
		return
	}
	c.AddHandler(handler)
	c.SetLogger(log.NsqLogger(), logLevel)

	err = c.ConnectToNSQLookupds(addrNsqLookups)
	if err != nil {
		log.Fatal(err)
		return
	}

	consumers = append(consumers, c)
	return
}

func gracefulStop() {
	producer.Stop()

	var wg sync.WaitGroup
	for _, c := range consumers {
		wg.Add(1)
		go func() {
			c.Stop()
			// disconnect from all lookupd
			for _, addr := range addrNsqLookups {
				c.DisconnectFromNSQLookupd(addr)
			}
			wg.Done()
		}()
	}

	wg.Wait()
}
