package subscriber

import (
	"context"
	"github.com/micro/go-micro/util/log"

	adminUser "infinite-window-micro/srv/adminUser/proto/adminUser"
)

type AdminUser struct{}

func (e *AdminUser) Handle(ctx context.Context, msg *adminUser.Message) error {
	log.Log("Handler Received message: ", msg.Say)
	return nil
}

func Handler(ctx context.Context, msg *adminUser.Message) error {
	log.Log("Function Received message: ", msg.Say)
	return nil
}
