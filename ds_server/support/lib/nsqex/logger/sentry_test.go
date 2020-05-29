package logger

import (
	"errors"
	"testing"
	"time"
)

func Test_SentryLogger(t *testing.T) {
	dsn := "00000000/20"

	l := NewSentryLogger("wemedia", dsn)
	err := errors.New("I've created an error.")
	l.Error("got an error: %v, other info: %s", err, "hello world")
	time.Sleep(1 * time.Second)
}


