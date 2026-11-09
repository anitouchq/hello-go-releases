package main

import (
	"fmt"
	"runtime"

	"github.com/anitouchq/hello-go-releases/greeting"
)

var version = "dev"

func main() {
	fmt.Println(greeting.Message())
	fmt.Println(greeting.BuildInfo(version, runtime.GOOS, runtime.GOARCH))
}
