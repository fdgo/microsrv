package main

import (
	"context"
	proto "ds_server/services/TestWebsocket/websock_srv/proto"
	"github.com/micro/go-micro"
	"io"
	"log"
	"time"
)

type Streamer struct{}

// Server side stream
func (e *Streamer) ServerStream(ctx context.Context, req *proto.Request, stream proto.Streamer_ServerStreamStream) error {
	log.Printf("Got msg %v", req.Count)
	for i := 0; i < int(req.Count); i++ {
		if err := stream.Send(&proto.Response{Count: int64(i)}); err != nil {
			return err
		}
	}
	return nil
}

// Bidirectional stream
func (e *Streamer) Stream(ctx context.Context, stream proto.Streamer_StreamStream) error {
	for {
		req, err := stream.Recv()

		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		log.Printf("Gotex msg %v", req.Count)
		if err := stream.Send(&proto.Response{Count: req.Count}); err != nil {
			return err
		}
	}
}
var (
	appName = "user_srv"
)

func main() {
	// new service
	service := micro.NewService(
		micro.Name("jz.micro.websock-srv.stream"),
		micro.RegisterTTL(time.Second*15),
		micro.RegisterInterval(time.Second*10),
		micro.Version("latest"),
	)

	// Init command line
	service.Init()

	// Register Handler
	proto.RegisterStreamerHandler(service.Server(), new(Streamer))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}

