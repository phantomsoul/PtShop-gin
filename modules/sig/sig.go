package sig

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// Run  监听服务信号
func Run(fn func()) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		s := <-c
		fmt.Printf("Get a signal %s\n", s.String())
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			if fn != nil {
				fn()
			}
			time.Sleep(time.Second)
			return
		case syscall.SIGHUP:
		default:
			return
		}
	}
}
