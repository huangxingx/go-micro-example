package handler

import (
	"context"
	"errors"
	"github.com/prometheus/common/log"
	admin "go-micro-example/srv/adminUser/proto/adminUser"
)

type AdminUser struct{}

func (a AdminUser) Insert(ctx context.Context, in *admin.AdminUser, res *admin.Response) error {
	return nil

}

func (a AdminUser) GetById(ctx context.Context, in *admin.Request, res *admin.AdminUser) error {
	adminUserId := in.Id
	if in.Id == 0 {
		log.Error("admin user value not 0")
		return errors.New("admin user value not 0")
	}
	res.Id = adminUserId
	res.Username = "huangxing"
	return nil
}

func (a AdminUser) UpdateById(ctx context.Context, in *admin.AdminUser, res *admin.Response) error {
	log.Info("UpdateById")
	return nil
}
