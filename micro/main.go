package main

import (
	"log"
	"net"
	"net/http"

	ph "github.com/afex/hystrix-go/hystrix"
	"github.com/micro/go-plugins/micro/cors"
	_ "github.com/micro/go-plugins/registry/consul"
	"github.com/micro/micro/cmd"
	"github.com/micro/micro/plugin"
	opentracing "github.com/opentracing/opentracing-go"

	"infinite-window-micro/pkg/tracer"
	"infinite-window-micro/pkg/wrapper/breaker/hystrix"
	"infinite-window-micro/pkg/wrapper/metrics/prometheus"
	"infinite-window-micro/pkg/wrapper/tracer/opentracing/stdhttp"
)

func init() {
	//token := &token.Token{}

	plugin.Register(cors.NewPlugin())

	//plugin.Register(plugin.NewPlugin(
	//	plugin.WithName("auth"),
	//	plugin.WithHandler(
	//		auth.JWTAuthWrapper(token),
	//	),
	//	plugin.WithFlag(cli.StringFlag{
	//		Name:   "consul_address",
	//		Usage:  "consul address for K/V",
	//		EnvVar: "CONSUL_ADDRESS",
	//		Value:  "127.0.0.1:8500",
	//	}),
	//	plugin.WithInit(func(ctx *cli.Context) error {
	//		log.Println(ctx.String("consul_address"))
	//		token.InitConfig(ctx.String("consul_address"), "micro", "config", "jwt-key", "key")
	//		return nil
	//	}),
	//))
	plugin.Register(plugin.NewPlugin(
		plugin.WithName("tracer"),
		plugin.WithHandler(
			stdhttp.TracerWrapper,
		),
	))
	plugin.Register(plugin.NewPlugin(
		plugin.WithName("breaker"),
		plugin.WithHandler(
			hystrix.BreakerWrapper,
		),
	))
	plugin.Register(plugin.NewPlugin(
		plugin.WithName("metrics"),
		plugin.WithHandler(
			prometheus.MetricsWrapper,
		),
	))
}

const name = "API gateway"

func main() {
	stdhttp.SetSamplingFrequency(50)
	t, io, err := tracer.NewTracer(name, "")
	if err != nil {
		log.Fatal(err)
	}
	defer io.Close()
	opentracing.SetGlobalTracer(t)

	hystrixStreamHandler := ph.NewStreamHandler()
	hystrixStreamHandler.Start()
	go http.ListenAndServe(net.JoinHostPort("", "81"), hystrixStreamHandler)

	cmd.Init()
}
