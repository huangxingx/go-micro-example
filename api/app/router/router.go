package router

import (
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/client"
	"infinite-window-micro/api/app/handler"
	"infinite-window-micro/pkg/wrapper/tracer/opentracing/gin2micro"
)

func InitRouter(client client.Client) *gin.Engine {
	router := gin.Default()
	r := router.Group("/app")
	r.Use(gin2micro.TracerWrapper)

	{
		r.GET("/echo", handler.Echo)
	}

	return router
}
