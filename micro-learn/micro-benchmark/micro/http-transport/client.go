package main

import (
	"flag"

	"github.com/wangmhgo/microservice-project/go-micro-learn/micro-benchmark/micro/internal"
	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/transport/http"
)

var concurrency = flag.Int("c", 1, "concurrency")
var total = flag.Int("n", 1, "total requests for all clients")

func main() {
	flag.Parse()
	n := *concurrency
	m := *total / n

	internal.ClientRun(m, n,
		"go.micro.benchmark.hello.http_transport",
		func() client.Client {
			return client.NewClient(
				client.Transport(http.NewTransport()),
			)
		})
}
