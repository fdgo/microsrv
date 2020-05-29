package logger

import (
	"fmt"
	"os"
	"path/filepath"
	"github.com/sirupsen/logrus"
)

// FileLogger file logger
type FileLogger struct {
	*logrus.Logger
}

// NewFileLogger providers a file logger based on logrus
func NewFileLogger(filename string, logLevel LogLevel) (*FileLogger, error) {
	absPath, err := filepath.Abs(filename)
	if err != nil {
		return nil, fmt.Errorf("can't get file abs path: filename = %v, err = %v", filename, err)
	}

	if _, err := os.Stat(absPath); err != nil {
		err = os.MkdirAll(filepath.Dir(absPath), os.ModePerm)
		if err != nil {
			return nil, fmt.Errorf("can't mkdirall directory: path = %v, err = %v", absPath, err)
		}
	}
	f, err := os.OpenFile(absPath, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		return nil, fmt.Errorf("can't open file: path = %v, err = %v", absPath, err)
	}

	l := &logrus.Logger{
		Out:       f,
		Formatter: &logrus.JSONFormatter{},
		Hooks:     make(logrus.LevelHooks),
	}

	switch logLevel {
	case LevelDebug:
		l.Level = logrus.DebugLevel
	case LevelError:
		l.Level = logrus.ErrorLevel
	case LevelFatal:
		l.Level = logrus.FatalLevel
	}

	return &FileLogger{
		l,
	}, nil
}

func (l *FileLogger) Debug(format string, args ...interface{}) {
	l.Debugf(format, args...)
}

func (l *FileLogger) Info(format string, args ...interface{}) {
	l.Infof(format, args...)
}

func (l *FileLogger) Error(format string, args ...interface{}) {
	l.Errorf(format, args...)
}

func (l *FileLogger) Fatal(format string, args ...interface{}) {
	l.Fatalf(format, args...)
}
