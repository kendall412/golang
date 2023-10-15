package main

import (
	"fmt"
	"runtime"
	"strings"
)

func main() {
	// determin host operating system
	var OS string
	opsys := runtime.GOOS
	if strings.Contains(opsys, "darwin") {
		OS = "Mac OS"
	} else if strings.Contains(opsys, "linux") {
		OS = "Linux"
	}
	fmt.Printf("OS: %s", OS)
}
