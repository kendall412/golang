package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
)

func main() {
	opsys := runtime.GOOS
	if strings.Contains(opsys, "darwin") {
		MAC := true
		fmt.Printf("Mac OS: %t\n", MAC)
	} else if strings.Contains(opsys, "linux") {
		LIN := true
		fmt.Printf("Linux OS: %t\n", LIN)
	}

	// checking to see if the mp3 files already exist
	targetFile := "Jungkook (정국) '3D (Feat. Jack Harlow)' Lyrics [C8vcCplJb1c].mp3"
	if _, err := os.Stat(targetFile); err == nil {
		fmt.Printf("File %s exist, download will not run.\n", targetFile)

	} else {
		fmt.Printf("File %s does not exist. File will now be downloaded.\n", targetFile)
		dl()
	}
}

func dl() {
	path := "/Users/kendall/repo/utils_personal"
	ytScript := "yt"
	ytPath := filepath.Join(path, ytScript)
	fmt.Printf("ytPath: %s\n", ytPath)

	ytOpt1 := "-au"
	ytOpt2 := "-x"
	url := "https://www.youtube.com/watch?v=C8vcCplJb1c&pp=ygURanVuZ2tvb2sgM2QgYXVkaW8%3D"

	_, err := exec.Command(ytPath, ytOpt1, ytOpt2, url).Output()

	if err != nil {
		log.Fatal(err)
	}
}
