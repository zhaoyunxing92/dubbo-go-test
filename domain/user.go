package domain

import (
	hessian "github.com/apache/dubbo-go-hessian2"
	"time"
)

func init() {
	hessian.RegisterPOJO(&User{})
}

type User struct {
	ID   string
	Name string
	Age  int32
	Time time.Time
}

func (u User) JavaClassName() string {
	return "org.apache.dubbo.User"
}
