package logger

type LogLevel int

const (
	LevelDebug LogLevel = iota
	LevelError
	LevelFatal
)
