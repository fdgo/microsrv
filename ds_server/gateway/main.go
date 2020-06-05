package main

import (
	"ds_server/gateway/initial"
	"ds_server/gateway/router"
	"log"
)

func main() {
	srv := initial.NewGateWaysrv()
	srv.Handle("/", router.Router())
	//srv.HandleFunc("/api/call", handler.WebCall)
	if err := srv.Init(); err != nil {
		log.Fatal(err)
	}
	if err := srv.Run(); err != nil {
		log.Fatalf("run err %v ", err)
	}
}

//go get -u -v github.com/oxequa/realize

//realize start
