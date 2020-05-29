package comlog

import (
	"fmt"
	"github.com/go-stack/stack"
	"github.com/lestrrat-go/file-rotatelogs"
	"github.com/olivere/elastic"
	"github.com/pkg/errors"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"gopkg.in/sohlich/elogrus.v3"
	"os"
	"path"
	"time"
)

var Logger *logrus.Logger
var debug_ = false
// debug: 使用text格式, Level是Debug, 打印所有级别
// not debug: 使用json格式, level是Info, 不打印Debug级别
func SetDebug(d bool) {
	debug_ = d
	if debug_ {
		format := new(logrus.TextFormatter)
		//format.ForceColors = true
		format.TimestampFormat = "2006-01-02 15:04:05"
		Logger.Level = logrus.DebugLevel
		Logger.Formatter = format
	} else {
		format := new(logrus.JSONFormatter)
		format.TimestampFormat = "2006-01-02 15:04:05"
		Logger.Level = logrus.InfoLevel
		Logger.Formatter = format
	}
}
func withField(key string, value interface{}) *logrus.Entry {
	return withCaller().WithField(key, value)
}
func withFields(fs logrus.Fields) *logrus.Entry {
	return withCaller().WithFields(fs)
}
func withCaller() *logrus.Entry {
	var key = "caller"
	var value interface{}
	if debug_ {
		// 支持goland点击跳转
		value = fmt.Sprintf("%+v:", stack.Caller(2))
	} else {
		value = fmt.Sprintf("%+v", stack.Caller(2))
	}
	return Logger.WithFields(logrus.Fields{key: value})
}
func fatal(args ...interface{}) {
	withCaller().Fatal(args...)
}
func fatalf(format string, args ...interface{}) {
	withCaller().Fatalf(format, args...)
}
func error(args ...interface{}) {
	withCaller().Error(args...)
}
func errorf(format string, args ...interface{}) {
	withCaller().Errorf(format, args...)
}
func warn(args ...interface{}) {
	withCaller().Warn(args...)
}
func warnf(format string, args ...interface{}) {
	withCaller().Warnf(format, args...)
}
func info(args ...interface{}) {
	withCaller().Info(args...)
}
func infof(format string, args ...interface{}) {
	withCaller().Infof(format, args...)
}
func debug(args ...interface{}) {
	withCaller().Debug(args...)
}
func debugf(format string, args ...interface{}) {
	withCaller().Debugf(format, args...)
}
// 输出日志到es
func configESLogger(esUrl string, esHOst string, index string) {
	client, err := elastic.NewClient(elastic.SetSniff(false), elastic.SetURL(esUrl))
	if err != nil {
		Logger.Errorf("config es logger error. %+v", errors.WithStack(err))
		return
	}
	esHook, err := elogrus.NewElasticHook(client, esHOst, logrus.DebugLevel, index)
	if err != nil {
		Logger.Errorf("config es logger error. %+v", errors.WithStack(err))
		return
	}
	Logger.AddHook(esHook)
}
func logToFile(srvName string, funcName string, errMsg string) {
	logFilePath := "/opt/data/com"
	logFileName := srvName + "_log"
	// 日志文件
	fileName := path.Join(logFilePath, logFileName)
	// 写入文件
	src, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY|os.O_CREATE, os.ModeAppend)
	if err != nil {
		fmt.Println("err", err)
	}
	// 设置输出
	Logger.Out = src
	// 设置日志级别
	Logger.SetLevel(logrus.InfoLevel)
	// 设置 rotatelogs
	logWriter, err := rotatelogs.New(
		// 分割后的文件名称
		fileName+".%Y%m%d%H%M",
		// 生成软链，指向最新日志文件
		rotatelogs.WithLinkName(fileName),
		// 设置最大保存时间(7天)
		rotatelogs.WithMaxAge(7*24*time.Hour),
		// 设置日志切割时间间隔(1天)
		rotatelogs.WithRotationTime(1*time.Second),
	)
	writeMap := lfshook.WriterMap{
		logrus.InfoLevel:  logWriter,
		logrus.FatalLevel: logWriter,
		logrus.DebugLevel: logWriter,
		logrus.WarnLevel:  logWriter,
		logrus.ErrorLevel: logWriter,
		logrus.PanicLevel: logWriter,
	}
	lfHook := lfshook.NewHook(writeMap, &logrus.JSONFormatter{TimestampFormat: "2006-01-02 15:04:05"})
	// 新增 Hook
	Logger.AddHook(lfHook)
	// 日志格式
	Logger.WithFields(logrus.Fields{
		"func_name": funcName,
		"err_msg":   errMsg,
	}).Info()
}
func init()  {
	Logger = &logrus.Logger{
		Out:       os.Stdout,
		Formatter: nil,
		Hooks:     make(logrus.LevelHooks),
		Level:     0,
	}
	SetDebug(true)
}
func LoggerFile(srvName string, funcName string, errMsg string) {
	logToFile(srvName, funcName, errMsg)
}
