package router

import (
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/client"
	"go-micro-example/api/app/handler"
	"go-micro-example/pkg/wrapper/tracer/opentracing/gin2micro"
)

func InitRouter(client client.Client) *gin.Engine {
	router := gin.Default()
	r := router.Group("/app")
	r.Use(gin2micro.TracerWrapper)

	{
		r.GET("/echo", handler.Echo)
	}

	// 店铺服务
	shopApi := handler.NewShopApi(client)
	{
		r.POST("/shop", shopApi.CreateShop)
		r.GET("/shop", shopApi.GetListShop)
	}

	return router
}
