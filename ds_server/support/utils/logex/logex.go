package logex

import (
	"os"
	"path"
	"time"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/rifflock/lfshook"
	rotatelogs "github.com/lestrrat/go-file-rotatelogs"
)

func GinLogger() gin.HandlerFunc {
	logFilePath := "/home/fred/jobspace/ds_server/gateway/log"
	logFileName := "gateway.log"
	// 日志文件
	fileName := path.Join(logFilePath, logFileName)
	// 写入文件
	src, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
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
		fileName+".%Y%m%d.log",
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
	lfHook := lfshook.NewHook(writeMap, &logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})
	// 新增钩子
	logger.AddHook(lfHook)
	return func(c *gin.Context) {
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
		reqUri := c.Request.RequestURI
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
