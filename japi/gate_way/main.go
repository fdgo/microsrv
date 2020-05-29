package main

import (
	conval "microservice/jzapi/basic/const_value"
	"microservice/jzapi/lib/token"
	"microservice/jzapi/lib/wrapper/auth"
	"microservice/jzapi/lib/wrapper/tracer/jaeger"
	"microservice/jzapi/lib/wrapper/tracer/opentracing/stdhttp"
	"microservice/jzapi/lib/log/comlog"
	"github.com/micro/go-plugins/micro/cors"
	"github.com/micro/micro/cmd"
	"github.com/micro/micro/plugin"
	"github.com/opentracing/opentracing-go"
)

func init() {
	jwtcfg :=token.InitConfig(conval.CONFIG_ADDRESS,"micro")
	jwtoken := &token.JwtToken{
		SigningKey:[]byte(jwtcfg.SecretKey),
	}
	plugin.Register(cors.NewPlugin())
	plugin.Register(plugin.NewPlugin(
		plugin.WithName("auth"),
		plugin.WithHandler(
			auth.JWTAuthWrapper(jwtoken),
		),
	))
	plugin.Register(plugin.NewPlugin(
		plugin.WithName("limit"),
		plugin.WithHandler(
			auth.LimitWrapper(),
		),
	))
	plugin.Register(plugin.NewPlugin(
		plugin.WithName("tracer"),
		plugin.WithHandler(
			stdhttp.TracerWrapper,
		),
	))
}
const name = "API gateway"
func main() {
	stdhttp.SetSamplingFrequency(50)
	t, io, err := tracer.NewTracer(name, "")
	if err != nil {
		comlog.Logger.Fatal(err)
	}
	defer io.Close()
	opentracing.SetGlobalTracer(t)
	cmd.Init()
}

//--registry=consul --registry_address=127.0.0.1:8500 --api_namespace=jz.micro.jzapi.web  api -handler=http
//--registry=etcdv3 --registry_address=127.0.0.1:2379 --api_namespace=jz.micro.jzapi.web  api -handler=http

