package router

import (
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/client"

	"go-micro-example/api/admin/handler"
	"go-micro-example/pkg/wrapper/recovery"
	"go-micro-example/pkg/wrapper/tracer/opentracing/gin2micro"
)

func InitRouter(client client.Client) *gin.Engine {
	router := gin.Default()
	r := router.Group("/admin")
	r.Use(gin2micro.TracerWrapper)
	r.Use(recovery.Recovery())

	{
		r.GET("/echo", handler.Echo)

		adminUserApi := handler.NewAdminUserApi(client)
		r.GET("/call", adminUserApi.Call)
	}

	return router
}
