package main

import (
	pb "ds_server/proto/user"
	inital "ds_server/services/user/initial"
	"errors"
	"fmt"
)

func main() {
	srv, db := inital.NewUserSrv()

	errReg := pb.RegisterUserHandler(srv.Server(), &db)

	errRun := srv.Run()

	if errReg != nil || errRun != nil {
		fmt.Println("=====errReg======", errReg)
		fmt.Println("=====errRun======", errRun)
		panic(errors.New("start user server error."))
	}
}
