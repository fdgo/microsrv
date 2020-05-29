package main

import (
	"context"
	"crypto/md5"
	"fmt"
	"net"
	"strings"
	"sync"
	"time"
	"github.com/micro/go-micro/config"
	"github.com/micro/go-micro/config/source/file"
	"github.com/micro/go-micro/util/log"
	grpcproto "github.com/micro/go-plugins/config/source/grpc/proto"
	grpc "google.golang.org/grpc"
)
var (
	mux        sync.RWMutex
	configMaps = make(map[string]*grpcproto.ChangeSet)
	// 根据需要改成可配置的app列表
	apps = []string{"micro"}
)
type Service struct{}
func main() {
	// 灾难恢复
	defer func() {
		if r := recover(); r != nil {
			log.Logf("[main] Recovered in f %v", r)
		}
	}()
	// 加载并侦听配置文件
	err := loadAndWatchConfigFile()
	if err != nil {
		log.Fatal(err)
	}
	// 新建grpc Server服务
	service := grpc.NewServer()
	grpcproto.RegisterSourceServer(service, new(Service))
	ts, err := net.Listen("tcp", ":9600")
	if err != nil {
		log.Fatal(err)
	}
	log.Logf("configServer started")
	// 启动
	err = service.Serve(ts)
	if err != nil {
		log.Fatal(err)
	}
}
func (s Service) Read(ctx context.Context, req *grpcproto.ReadRequest) (rsp *grpcproto.ReadResponse, err error) {
	appName := parsePath(req.Path)

	rsp = &grpcproto.ReadResponse{
		ChangeSet: getConfig(appName),
	}
	return
}
func (s Service) Watch(req *grpcproto.WatchRequest, server grpcproto.Source_WatchServer) (err error) {
	appName := parsePath(req.Path)
	rsp := &grpcproto.WatchResponse{
		ChangeSet: getConfig(appName),
	}
	if err = server.Send(rsp); err != nil {
		log.Logf("[Watch] 侦听处理异常，%s", err)
		return err
	}

	return
}
func loadAndWatchConfigFile() (err error) {
	// 加载每个应用的配置文件
	for _, app := range apps {
		if err := config.Load(file.NewSource(
			file.WithPath( app + ".yml"),
		)); err != nil {
			log.Fatalf("[loadAndWatchConfigFile] 加载应用配置文件 异常，%s", err)
			return err
		}
	}
	// 侦听文件变动
	watcher, err := config.Watch()
	if err != nil {
		log.Fatalf("[loadAndWatchConfigFile] 开始侦听应用配置文件变动 异常，%s", err)
		return err
	}
	go func() {
		for {
			v, err := watcher.Next()
			if err != nil {
				log.Fatalf("[loadAndWatchConfigFile] 侦听应用配置文件变动 异常， %s", err)
				return
			}

			log.Logf("[loadAndWatchConfigFile] 文件变动，%s", string(v.Bytes()))
		}
	}()
	return
}

func getConfig(appName string) *grpcproto.ChangeSet {
	bytes := config.Get(appName).Bytes()

	//log.Logf("[getConfig] appName：%s", appName)
	return &grpcproto.ChangeSet{
		Data:      bytes,
		Checksum:  fmt.Sprintf("%x", md5.Sum(bytes)),
		Format:    "yml",
		Source:    "file",
		Timestamp: time.Now().Unix()}
}

func parsePath(path string) (appName string) {
	paths := strings.Split(path, "/")

	if paths[0] == "" && len(paths) > 1 {
		return paths[1]
	}

	return paths[0]
}

