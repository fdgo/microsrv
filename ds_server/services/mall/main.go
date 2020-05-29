package main

import (
	pb "ds_server/proto/mall"
	inital "ds_server/services/mall/initial"
	"errors"
)

func main() {
	srv, db := inital.NewMallSrv()

	errReg := pb.RegisterMallHandler(srv.Server(), &db)

	errRun := srv.Run()

	if errReg != nil || errRun != nil {
		panic(errors.New("start user server error."))
	}
}
