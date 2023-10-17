package main

import (
	"fmt"
	"os/exec"
	"path/filepath"
	"time"
)

var VER string = "1.0.2"

type Kid struct {
	Name     string `json:"name"`
	Location string `json:"location"`
	Playlist string `json:"playlist"`
}

func main() {
	fmt.Printf("VER: %s\n", VER)
	checkPlatform()

	// npath := openJsonMap(paths)

	var shinoo, naami, woojin, wonoo Kid // for testing
	Shinoo := shinoo.openJsonStruct("./shinoo.json")
	Naami := naami.openJsonStruct("./naami.json")
	Woojin := woojin.openJsonStruct("./woojin.json")
	Wonoo := wonoo.openJsonStruct("./wonoo.json") // for testing

	fmt.Printf("Shinoo.Name: %s\n", Shinoo.Name)
	fmt.Printf("Shinoo.Playlist: %s\n", Shinoo.Playlist)
	fmt.Printf("Shinoo.Location: %s\n", Shinoo.Location)

	fmt.Printf("Naami.Name: %s\n", Naami.Name)
	fmt.Printf("Naami.Playlist: %s\n", Naami.Playlist)
	fmt.Printf("Naami.Location: %s\n", Naami.Location)

	fmt.Printf("Woojin.Name: %s\n", Woojin.Name)
	fmt.Printf("Woojin.Playlist: %s\n", Woojin.Playlist)
	fmt.Printf("Woojin.Location: %s\n", Woojin.Location)

	fmt.Printf("Wonoo.Name: %s\n", Wonoo.Name)
	fmt.Printf("Wonoo.Playlist: %s\n", Wonoo.Playlist)
	fmt.Printf("Wonoo.Location: %s\n", Wonoo.Location)

	if checkDir(Shinoo.Location) == true {
		fmt.Printf("The provided directory named %s exists.\n", Shinoo.Location)
		removeAllFiles(Shinoo.Location)
		time.Sleep(2)
		dl(Shinoo.Location, Shinoo.Playlist)

	} else {
		fmt.Printf("%s does not exist.\n", Shinoo.Location)
	}
}

/*
dl
DESC: will spawn a process to invoke yt-dlp to begin downloading from playlist.
*/
func dl(dest string, url string) {
	// path := "/opt/homebrew/bin/"
	path := "/usr/local/bin/" // test path
	ytScript := "yt-dlp"
	ytPath := filepath.Join(path, ytScript)

	optformat1 := "-x"
	optFormat2 := "--audio-format"
	optFormat3 := "mp3"
	optPath := "--path"

	fmt.Printf("%s %s %s %s %s %s %s\n", ytPath, optformat1, optFormat2, optFormat3, optPath, dest, url)
	_, err := exec.Command(ytPath, optformat1, optFormat2, optFormat3, optPath, dest, url).Output()

	if err != nil {
		fmt.Println(err)
	}
}
