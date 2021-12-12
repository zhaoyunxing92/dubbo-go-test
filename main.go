package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)
import (
	_ "zhaoyunxing92/dubbo-go-test/reference"
	"zhaoyunxing92/dubbo-go-test/service"
)

import (
	"dubbo.apache.org/dubbo-go/v3/config"
	_ "dubbo.apache.org/dubbo-go/v3/imports"
	gxlog "github.com/dubbogo/gost/log"
)

var (
	survivalTimeout = int(3e9)
)
var userService = new(service.UserService)

// need to setup environment variable "CONF_PROVIDER_FILE_PATH" to "conf/server.yml" before run
func main() {
	config.SetConsumerService(userService)

	err := config.Load(config.WithPath("./conf/dubbogo.yaml"))
	if err != nil {
		return
	}

	time.Sleep(3 * time.Second)

	user, err := userService.GetUser(context.Background(), "zhaoyunxing")
	if err != nil {
		gxlog.CError("error: %v\n", err)
		os.Exit(1)
		return
	}
	gxlog.CInfo("response result: %v\n", user)
	initSignal()
}

func initSignal() {
	signals := make(chan os.Signal, 1)
	// It is not possible to block SIGKILL or syscall.SIGSTOP
	signal.Notify(signals, os.Interrupt, os.Kill, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		sig := <-signals
		fmt.Println("get signal %s", sig.String())
		switch sig {
		case syscall.SIGHUP:
			// reload()
		default:
			time.AfterFunc(time.Duration(survivalTimeout), func() {
				fmt.Println("app exit now by force...")
				os.Exit(1)
			})

			// The program exits normally or timeout forcibly exits.
			fmt.Println("provider app exit now...")
			return
		}
	}
}
