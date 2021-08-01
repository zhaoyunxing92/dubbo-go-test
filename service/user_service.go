package service

import (
	"context"
	"zhaoyunxing92/dubbo-go-test/domain"
)
import (
	"dubbo.apache.org/dubbo-go/v3/config"
)

func init() {
	config.SetConsumerService(new(UserService))
}

type UserService struct {
	GetUser func(ctx context.Context, name string, rsp *domain.User) error
}

func (u UserService) Reference() string {
	return "UserService"
}
