package main

import (
	"time"

	"github.com/micro/go-micro"
	"github.com/micro/go-micro/service/grpc"
	"github.com/micro/go-micro/util/log"
	"github.com/micro/go-plugins/registry/consul"
	ocplugin "github.com/micro/go-plugins/wrapper/trace/opentracing"
	"github.com/opentracing/opentracing-go"

	"go-micro-example/pkg/tracer"

	"go-micro-example/srv/adminUser/handler"
	adminUser "go-micro-example/srv/adminUser/proto/adminUser"
)

const (
	Version = "v0.1.0"
	Name    = "com.example.srv.adminUser"
)

func main() {
	t, io, err := tracer.NewTracer(Name, "localhost:6831")
	if err != nil {
		log.Fatal(err)
	}
	defer io.Close()
	opentracing.SetGlobalTracer(t)

	// New Service
	service := grpc.NewService(
		micro.Name(Name),
		micro.Version(Version),
		micro.RegisterInterval(time.Duration(5)),
		micro.RegisterTTL(time.Duration(10)),
		micro.Registry(consul.NewRegistry()),

		micro.WrapHandler(ocplugin.NewHandlerWrapper(opentracing.GlobalTracer())),
	)

	// Initialise service
	service.Init()

	// Register Handler
	adminUser.RegisterAdminUserServiceHandler(service.Server(), new(handler.AdminUser))

	// Register Struct as Subscriber
	//micro.RegisterSubscriber("com.example.srv.adminUser", service.Server(), new(subscriber.AdminUser))

	// Register Function as Subscriber
	//micro.RegisterSubscriber("com.example.srv.adminUser", service.Server(), subscriber.Handler)

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
