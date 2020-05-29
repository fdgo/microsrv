package logger

import (
	"ds_server/support/utils/daemon"
	"io"
	"io/ioutil"
	"os"

	"github.com/sirupsen/logrus"

	lumberjack "gopkg.in/natefinch/lumberjack.v2"
)

type RotateFileConfig struct {
	Filename   string
	MaxSize    int
	MaxBackups int
	MaxAge     int
	Level      logrus.Level
	Formatter  logrus.Formatter
}

type RotateFileHook struct {
	Config    RotateFileConfig
	logWriter io.Writer
}

func NewRotateFileHook(config RotateFileConfig) logrus.Hook {

	hook := RotateFileHook{
		Config: config,
	}
	hook.logWriter = &lumberjack.Logger{
		Filename:   config.Filename,
		MaxSize:    config.MaxSize,
		MaxBackups: config.MaxBackups,
		MaxAge:     config.MaxAge,
	}

	return &hook
}

func (hook *RotateFileHook) Levels() []logrus.Level {

	var levels []logrus.Level
	levels = append(levels, logrus.AllLevels...)

	return levels
}

func (hook *RotateFileHook) Fire(entry *logrus.Entry) (err error) {
	b, err := hook.Config.Formatter.Format(entry)
	if err != nil {
		return err
	}
	hook.logWriter.Write(b)
	if daemon.IsDaemonMode() {
		entry.Logger.Out = ioutil.Discard
	} else {
		entry.Logger.Out = os.Stdout
	}

	return nil
}
