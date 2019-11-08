package main

import (
	"time"

	"github.com/micro/cli"
	"github.com/micro/go-micro"

	"github.com/micro/go-micro/service/grpc"
	"github.com/micro/go-micro/util/log"
	"github.com/micro/go-plugins/registry/consul"
	ocplugin "github.com/micro/go-plugins/wrapper/trace/opentracing"
	"github.com/opentracing/opentracing-go"

	"go-micro-example/pkg/database"
	"go-micro-example/pkg/tracer"
	"go-micro-example/repository"
	"go-micro-example/srv/shop/handler"
	shopProto "go-micro-example/srv/shop/proto/shop"
)

const (
	Version = "v0.1.0"
	Name    = "com.example.srv.shop"
)

func main() {
	t, io, err := tracer.NewTracer(Name, "localhost:6831")
	if err != nil {
		log.Fatal(err)
	}
	defer io.Close()
	opentracing.SetGlobalTracer(t)

	var consulAddr string

	// New Service
	service := grpc.NewService(
		micro.Name(Name),
		micro.Version(Version),
		micro.RegisterInterval(time.Duration(5)),
		micro.RegisterTTL(time.Duration(10)),
		micro.Registry(consul.NewRegistry()),

		micro.WrapHandler(ocplugin.NewHandlerWrapper(opentracing.GlobalTracer())),

		micro.Action(func(ctx *cli.Context) {
			consulAddr = ctx.String("consul_address")
		}),
	)

	// Initialise service
	service.Init()
	// init db
	database.InitDbByConsul(consulAddr)

	shopRepo := repository.NewShopRepo(database.GetDb())
	// Register Handler
	shopProto.RegisterShopManagerServiceHandler(service.Server(), handler.NewShopHandler(shopRepo))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
