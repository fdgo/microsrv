package main

import (
	pb "ds_server/proto/branch"
	inital "ds_server/services/branch/initial"
	"errors"
)

func main() {
	// inital.InitialModels()
	// db2 := datasource.DatasourceManagerInstance(datasource.DATASOURCE_MANAGER).Datasource()
	// defer db2.Close()
	srv, db := inital.NewBranchSrv()
	errReg := pb.RegisterBranchHandler(srv.Server(), &db)
	errRun := srv.Run()
	if errReg != nil || errRun != nil {
		panic(errors.New("start branch server error."))
	}

}
