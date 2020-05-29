package main

import (
	"ds_server/schedule"
	"ds_server/support/utils/logger"
	"sync"

	"github.com/robfig/cron"
)

var (
	wg           sync.WaitGroup
	BuildVersion string
	Log          *logger.Logger = logger.InitLog()
)

func scheduleInit() {
	c := cron.New()

	// 每天晚上清洗数据
	c.AddFunc("0 1 0 * * *", schedule.CleanData)

	c.Start()
}

func run() {
	Log.Info("#######################################################")
	Log.Infof("###### %s ######", BuildVersion)
	Log.Info("#######################################################")
	//shardingdatasource.InitialDataSource(serverConfig.ServerName)
	scheduleInit()
}

func main() {
	wg.Add(1)
	go run()
	wg.Wait()
}
