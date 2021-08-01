package reference

import (
	"context"
	"dubbo.apache.org/dubbo-go/v3/config"
	"github.com/dubbogo/gost/log"
	"time"
	"zhaoyunxing92/dubbo-go-test/domain"
)

func init() {
	config.SetProviderService(new(UserService))
}

type UserService struct {
}

func (UserService) Reference() string {
	return "UserService"
}

func (u *UserService) GetUser(ctx context.Context, req []interface{}) (*domain.User, error) {
	gxlog.CInfo("req:%#v", req)
	rsp := domain.User{ID: "A001", Name: "Alex Stocks", Age: 18, Time: time.Now()}
	gxlog.CInfo("rsp:%#v", rsp)
	return &rsp, nil
}
