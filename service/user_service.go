package service

import (
	"context"
	"zhaoyunxing92/dubbo-go-test/domain"
)

type UserService struct {
	GetUser func(ctx context.Context, name string, rsp *domain.User) error
}

func (u UserService) Reference() string {
	return "UserService"
}
