package main

import (
	pb "ds_server/proto/message"
	inital "ds_server/services/message/initial"
	"errors"
)

func main() {
	srv, db := inital.NewMessageSrv()
	errReg := pb.RegisterMessageHandler(srv.Server(), &db)
	errRun := srv.Run()
	if errReg != nil || errRun != nil {
		panic(errors.New("start message server error."))
	}
}
