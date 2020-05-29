package log

import (
	"github.com/gin-gonic/gin"
	"github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"os"
	"path"
	"time"
	"fmt"
)
var(
	logtype = 1
	serviceName = "xxx-srv"
)
func Init(logty int, serName string)  {
	logtype = logty
	serviceName = serName
}
func GinLogger( c *gin.Context) {
	if logtype == 1 {
		start := time.Now()
		// 处理请求
		c.Next()
		// 结束时间
		end := time.Now()
		//执行时间
		latency := end.Sub(start)
		path := c.Request.URL.Path
		clientIP := c.ClientIP()
		method := c.Request.Method
		statusCode := c.Writer.Status()
		logger.Infof("| %3d | %13v | %15s | %s  %s |",
			statusCode,
			latency,
			clientIP,
			method, path,
		)
	}
	if logtype == 0 {
		logFilePath := "/opt/data/gin"
		logFileName := serviceName + "_log"
		// 日志文件
		fileName := path.Join(logFilePath, logFileName)
		// 写入文件
		src, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY|os.O_CREATE, os.ModeAppend)
		if err != nil {
			fmt.Println("err", err)
		}
		// 实例化
		logger := logrus.New()
		// 设置输出
		logger.Out = src
		// 设置日志级别
		logger.SetLevel(logrus.DebugLevel)
		// 设置 rotatelogs
		logWriter, err := rotatelogs.New(
			// 分割后的文件名称
			fileName+".%Y%m%d%H%M",
			// 生成软链，指向最新日志文件
			rotatelogs.WithLinkName(fileName),
			// 设置最大保存时间(7天)
			rotatelogs.WithMaxAge(7*24*time.Hour),
			// 设置日志切割时间间隔(1天)
			rotatelogs.WithRotationTime(24*time.Hour),
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
		logger.AddHook(lfHook)
		// 开始时间
		startTime := time.Now()
		// 处理请求
		c.Next()
		// 结束时间
		endTime := time.Now()
		// 执行时间
		latencyTime := endTime.Sub(startTime)
		// 请求方式
		reqMethod := c.Request.Method
		// 请求路由
		reqUri := c.Request.URL.Path
		// 状态码
		statusCode := c.Writer.Status()
		// 请求IP
		clientIP := c.ClientIP()
		// 日志格式
		logger.WithFields(logrus.Fields{
			"status_code":  statusCode,
			"latency_time": latencyTime,
			"client_ip":    clientIP,
			"req_method":   reqMethod,
			"req_uri":      reqUri,
		}).Info()
	}
}
