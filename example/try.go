package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/jsxxzy/lockfile"
)

// ExitFunc 退出
func ExitFunc() {
	fmt.Println("开始退出...")
	fmt.Println("执行清理...")
	fmt.Println("结束退出...")
	os.Exit(0)
}

func main() {
	var lockfileName = "inet.lock"
	flag, err := lockfile.IsLocked(lockfileName)
	if !flag {
		fmt.Println("进程锁未存在", err)
		lockfile.Lock(lockfileName)
		for {
			//创建监听退出chan
			c := make(chan os.Signal)
			//监听指定信号 ctrl+c kill
			signal.Notify(c, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGUSR1, syscall.SIGUSR2)
			go func() {
				for s := range c {
					switch s {
					case syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT:
						lockfile.Unlock(lockfileName)
						ExitFunc()
					}
				}
			}()
		}
	} else {
		fmt.Println("进程已经存在")
	}
}
