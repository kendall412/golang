package main

import (
	"fmt"
	"os"
	"path"
	"runtime"
	"strings"
)

func checkPlatform() {
	// determin host operating system
	var OS string
	opsys := runtime.GOOS
	if strings.Contains(opsys, "darwin") {
		OS = "Mac OS"
	} else if strings.Contains(opsys, "linux") {
		OS = "Linux"
	}
	fmt.Printf("OS: %s\n", OS)
}

/*
checkDir
DESC: will check to see if the dir exists.
*/
func checkDir(dir string) bool {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		return (false)
	} else {
		return (true)
	}
}

/*
removeAllFiles
DESC: removes all files in 'dest'
*/
func removeAllFiles(dest string) {
	dir, err := os.ReadDir(dest)
	fmt.Println(dir)
	// os.ReadDir(path) will create slice of all files in path
	if err != nil {
		fmt.Printf("could not remove all files from %s", dest)
	}
	for _, d := range dir {
		os.RemoveAll(path.Join([]string{dest, d.Name()}...))
	}
}
