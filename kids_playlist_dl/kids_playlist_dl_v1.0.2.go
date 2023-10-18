package main

import (
	"fmt"
	"os/exec"
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
	fmt.Printf("Shinoo.Location: %s\n\n", Shinoo.Location)

	fmt.Printf("Naami.Name: %s\n", Naami.Name)
	fmt.Printf("Naami.Playlist: %s\n", Naami.Playlist)
	fmt.Printf("Naami.Location: %s\n\n", Naami.Location)

	fmt.Printf("Woojin.Name: %s\n", Woojin.Name)
	fmt.Printf("Woojin.Playlist: %s\n", Woojin.Playlist)
	fmt.Printf("Woojin.Location: %s\n\n", Woojin.Location)

	fmt.Printf("Wonoo.Name: %s\n", Wonoo.Name)
	fmt.Printf("Wonoo.Playlist: %s\n", Wonoo.Playlist)
	fmt.Printf("Wonoo.Location: %s\n\n", Wonoo.Location)

	initiateDL(Shinoo.Location, Shinoo.Playlist)
	initiateDL(Naami.Location, Naami.Playlist)
	initiateDL(Woojin.Location, Woojin.Playlist)
	initiateDL(Wonoo.Location, Wonoo.Playlist)
}

/*
initiateDL
DESC: checks to see the ssd card exist and if so will initial dl function.
*/
func initiateDL(location, playlist string) {
	if checkDir(location) == true {
		fmt.Printf("The provided directory named %s exists.\n", location)
		removeAllFiles(location)
		time.Sleep(10)
		dl(location, playlist)

	} else {
		fmt.Printf("%s does not exist.\n", location)
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
