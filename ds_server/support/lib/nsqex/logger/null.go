package logger

type NullLogger struct {
}

func NewNullLogger() *NullLogger {
	return &NullLogger{}
}

func (l *NullLogger) Fatal(format string, args ...interface{}) {
}

func (l *NullLogger) Error(format string, args ...interface{}) {
}

func (l *NullLogger) Info(format string, args ...interface{}) {
}

func (l *NullLogger) Debug(format string, args ...interface{}) {
}
