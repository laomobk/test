package controller

import (
	"fmt"
	"test/core"
)

func Retry() {
	core.Start()
}

func Shutdown() {
	fmt.Println("电源关闭...")
}
