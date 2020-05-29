package logger

import (
	"github.com/wangmhgo/nsq/nsqpool"
	nsq "github.com/nsqio/go-nsqex"
)

type NsqSentryLogger struct {
	addr  string
	topic string
	size  int
	pool  pool.Pool

	sentryLogger *SentryLogger
}

func NewNsqSentryLogger(addr, topic string, size int, appName, dsn string) (*NsqSentryLogger, error) {
	var err error
	l := new(NsqSentryLogger)
	l.addr = addr
	l.topic = topic

	factory := func() (*nsq.Producer, error) {
		config := nsq.NewConfig()
		return nsq.NewProducer(addr, config)
	}

	if l.pool, err = pool.NewChannelPool(size, size*2, factory); err != nil {
		return nil, err
	}

	l.sentryLogger = NewSentryLogger(appName, dsn)

	return l, nil
}

func (l *NsqSentryLogger) Fatal(format string, args ...interface{}) {
	l.sentryLogger.Fatal(format, args...)
}

func (l *NsqSentryLogger) Error(format string, args ...interface{}) {
	l.sentryLogger.Error(format, args...)
}

func (l *NsqSentryLogger) Info(data string, _args ...interface{}) {
	var err error
	defer func() {
		if err != nil {
			l.Error("publish info got an error: data = %s, addr = %s, topic = %s, err = %v", data, l.addr, l.topic, err)
		}
	}()

	producer, err := l.pool.Get()
	if err != nil {
		return
	}
	defer producer.Close()

	if err = producer.Publish(l.topic, []byte(data)); err != nil {
		return
	}
}

func (l *NsqSentryLogger) Debug(format string, args ...interface{}) {
}
