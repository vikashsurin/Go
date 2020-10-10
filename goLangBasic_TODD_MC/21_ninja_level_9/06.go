package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("OS :", runtime.GOOS, "ARCH :", runtime.GOARCH)
}
