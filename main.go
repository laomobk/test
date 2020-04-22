package main

import (
	"fmt"
	"test/controller"
	"test/core"
)

func main() {
	fmt.Println("即将启动核心...")
	core.Start()

	fmt.Println("\n启动失败...即将重试...")
	controller.Retry()

	fmt.Println("\n算了不跑了...")
	fmt.Println("\n即将关闭核心...")
	controller.Shutdown()
}
