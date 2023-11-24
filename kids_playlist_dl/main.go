package main

import (
	"flag"
	"fmt"
	"os/exec"
	"sync"
)

var wg sync.WaitGroup
var VER string = "1.0.4b"

type Kid struct {
	Name     string `json:"name"`
	Location string `json:"location"`
	Playlist string `json:"playlist"`
}

type Kids struct {
	Kids []Kid `json:"kids"`
}

func genPlaylistPath(debug bool) string {
	if debug {
		return "./testplaylists.json"
	} else {
		return "./kidsplaylists.json"
	}
}

func main() {
	// flag returns pointer, therefore must de-reference
	debug := flag.Bool("d", false, "debug flag, when selected will used 'testplaylist.json'")
	flag.Parse()

	jsonConfigFile := genPlaylistPath(*debug)

	fmt.Printf("VER: %s\n", VER)
	checkPlatform()

	var kids_ Kids

	kids := kids_.openJsonStruct(jsonConfigFile)

	fmt.Println("length of JSON: ", len(kids.Kids))
	wg.Add(len(kids.Kids))

	for i := 0; i < len(kids.Kids); i++ {
		go initiateDL(kids.Kids[i].Location, kids.Kids[i].Playlist)
	}
	wg.Wait()
}

/*
initiateDL
DESC: checks to see the ssd card exist and if so will initial dl function.
*/
func initiateDL(location, playlist string) {
	if checkDir(location) {
		removeAllFiles(location)
		// time.Sleep(time.Millisecond * 10)
		dl(location, playlist)

	} else {
		fmt.Printf("::: %s does NOT exist :::\n", location)
	}
	wg.Done()
}

func dl(dest string, url string) {
	cmd := "yt-dlp"
	opt := []string{
		"-x",
		"--audio-format",
		"mp3",
		"--path",
	}

	opt = append(opt, dest, url)
	fmt.Println(cmd, opt)
	_, err := exec.Command(cmd, opt...).Output()

	if err != nil {
		fmt.Println(err)
	}
}
