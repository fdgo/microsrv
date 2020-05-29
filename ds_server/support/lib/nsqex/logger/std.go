package logger

import (
	"fmt"
	"log"

	"github.com/pkg/errors"
)

type StdLogger struct {
}

func NewStdLogger() *StdLogger {
	log.SetFlags(log.Ltime | log.Lshortfile)
	return &StdLogger{}
}

func (l *StdLogger) Fatal(format string, args ...interface{}) {
	log.Printf("[Fatal]"+format, args...)
}

func (l *StdLogger) Debug(format string, args ...interface{}) {
	log.Printf("[Debug]"+format, args...)
}

func (l *StdLogger) Info(format string, args ...interface{}) {
	log.Printf("[Info]"+format, args...)
}

func (l *StdLogger) Error(format string, args ...interface{}) {
	err := fmt.Errorf("[Error]"+format, args...)
	log.Printf("%+v", errors.Wrap(err, ""))
}
