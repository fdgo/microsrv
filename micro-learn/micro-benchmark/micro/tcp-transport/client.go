package main

import (
	"flag"

	"github.com/wangmhgo/microservice-project/go-micro-learn/micro-benchmark/micro/internal"
	"github.com/wangmhgo/microservice-project/go-micro-learn/micro-benchmark/pb"
	"github.com/micro/go-micro"
	"github.com/micro/go-plugins/transport/tcp"
)

var concurrency = flag.Int("c", 1, "concurrency")
var total = flag.Int("n", 1, "total requests for all clients")

func main() {
	flag.Parse()
	n := *concurrency
	m := *total / n

	service := micro.NewService(micro.Name("go.micro.benchmark.hello.client"), micro.Transport(tcp.NewTransport()), )
	c := pb.NewHelloService("go.micro.benchmark.hello.tcp_transport", service.Client(), )

	internal.ClientRun(m, n, c)
}
