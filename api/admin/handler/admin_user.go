package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/util/log"

	"go-micro-example/pkg/app"
	"go-micro-example/pkg/wrapper/tracer/opentracing/gin2micro"
	adminUserS "go-micro-example/srv/adminUser/proto/adminUser"
)

type AdminUserApi struct {
	adminUserC adminUserS.AdminUserService
}

func NewAdminUserApi(client client.Client) *AdminUserApi {
	return &AdminUserApi{
		adminUserC: adminUserS.NewAdminUserService("", client),
	}
}

func (s *AdminUserApi) Call(c *gin.Context) {
	log.Log("Received Say.Anything API request")
	ctx, ok := gin2micro.ContextWithSpan(c)
	if !ok {
		log.Error("get context err")
	}
	appG := app.NewAppGin(c)

	rep, err := s.adminUserC.GetById(ctx, &adminUserS.Request{Id: 1})
	if err != nil {
		log.Error(err)
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	log.Info(rep)

	appG.Success(map[string]interface{}{
		"Username": rep.Username,
		"Id":       rep.Id,
		"Password": rep.Password,
	})

	return
}
