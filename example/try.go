package main

import (
	"fmt"
	"os"

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
	var lockfileName = "inet"
	run, code, _ := lockfile.NewSingleApp(lockfileName)
	if code == lockfile.AppRunOtherProcess {
		fmt.Println("进程同时存在了")
		return
	}
	run.Free(ExitFunc)
	for {

	}
}
