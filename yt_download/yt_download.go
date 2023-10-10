package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

func main() {

	if _,err := os.Stat("Jungkook (정국) '3D (Feat. Jack Harlow)' Lyrics [C8vcCplJb1c].mp3");
		err == nil {
			fmt.Println("File exist, download will not run")
			
		} else {
		fmt.Println("File does not exist. File will now be downloaded.")
		dl()
	}

}

func dl(){
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
