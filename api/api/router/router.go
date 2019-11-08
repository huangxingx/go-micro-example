package router

import (
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/client"
	"infinite-window-micro/pkg/wrapper/tracer/opentracing/gin2micro"

	"infinite-window-micro/api/api/handler"
)

func InitRouter(client client.Client) *gin.Engine {
	router := gin.Default()
	r := router.Group("/api")
	r.Use(gin2micro.TracerWrapper)

	{
		//r.Use(gin2micro.TracerWrapper)
		r.GET("/echo", handler.Echo)
	}

	return router
}
