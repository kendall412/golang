package main

import (
	"fmt"
	"log"
	"os/exec"
	"path/filepath"
)

func main() {
	path := "/Users/kendall/repo/utils_personal"
	ytScript := "yt"
	ytPath := filepath.Join(path, ytScript)
	fmt.Printf("ytPath: %s\n", ytPath)

	ytOpt1 := "-au"
	ytOpt2 := "-x"
	url := "https://www.youtube.com/watch?v=C8vcCplJb1c&pp=ygURanVuZ2tvb2sgM2QgYXVkaW8%3D"

	out, err := exec.Command(ytPath, ytOpt1, ytOpt2, url).Output()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Out: %s", out)
}
