package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Print(runtime.GOOS, "_", runtime.GOARCH) //nolint
}
