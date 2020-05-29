package main

import (
	pb "ds_server/proto/base"
	inital "ds_server/services/base/initial"
	"errors"
)

func main() {
	srv, db := inital.NewBaseSrv()

	errReg := pb.RegisterBaseHandler(srv.Server(), &db)

	errRun := srv.Run()

	if errReg!=nil || errRun!=nil{
		panic(errors.New("start user server error."))
	}
}
