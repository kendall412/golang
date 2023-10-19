package main

import (
	"fmt"
	"os/exec"
	"time"
)

var VER string = "1.0.3"
var jsonConfigFile string = "./kidsconfig.json"

type Kid struct {
	Name     string `json:"name"`
	Location string `json:"location"`
	Playlist string `json:"playlist"`
}

type Kids struct {
	Kids []Kid `json:"kids"`
}

func main() {
	fmt.Printf("VER: %s\n", VER)
	checkPlatform()

	var kids_ Kids
	kids := kids_.openJsonStruct(jsonConfigFile)
	for i := 0; i < len(kids.Kids); i++ {
		initiateDL(kids.Kids[i].Location, kids.Kids[i].Playlist)
	}
}

/*
initiateDL
DESC: checks to see the ssd card exist and if so will initial dl function.
*/
func initiateDL(location, playlist string) {
	if checkDir(location) == true {
		removeAllFiles(location)
		time.Sleep(10)
		dl(location, playlist)

	} else {
		fmt.Printf("::: %s does NOT exist :::\n", location)
	}
}

/*
dl
DESC: will spawn a process to invoke yt-dlp to begin downloading from playlist.
*/
func dl(dest string, url string) {
	app := "yt-dlp"
	optformat1 := "-x"
	optFormat2 := "--audio-format"
	optFormat3 := "mp3"
	optPath := "--path"

	fmt.Printf("%s %s %s %s %s %s %s\n", app, optformat1, optFormat2, optFormat3, optPath, dest, url)
	_, err := exec.Command(app, optformat1, optFormat2, optFormat3, optPath, dest, url).Output()

	if err != nil {
		fmt.Println(err)
	}
}
