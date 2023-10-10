package main

import (
	"fmt"
	"runtime"
)

func main() {
	returnOS()
}

func returnOS() string {
	os := runtime.GOOS
	fmt.Println("Operating system:", os)
	return os
}
