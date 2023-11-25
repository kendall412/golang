package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"sync"
)

var wg sync.WaitGroup
var VER string = "1.0.4c"

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
	fmt.Printf("VER: %s\n", VER)
	checkPlatform()

	/*
		Flags
		DESC: creating bool flag for debug purpose. Flags returns pointers, therefore must be de-referenced.
	*/
	debug := flag.Bool("d", false, "debug flag, defaut is FALSE. If TRUE, -d,  will use testplaylist.json for debug purposes. Otherwise by default if FALSE will used kidsplaylists.json")
	flag.Parse()

	jsonConfigFile := genPlaylistPath(*debug)
	fmt.Println("jsonConfigFile: " + jsonConfigFile)

	var kids_ Kids
	kids := kids_.openJsonStruct(jsonConfigFile)

	fmt.Println("length of JSON: ", len(kids.Kids))
	wg.Add(len(kids.Kids))

	/*
		goroutines are set to initiateDL()
	*/
	userhome, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("There was a problem determining userhome variable")
	}
	testpldirloc := "/repo/golang/kids_playlist_dl/testdir/"
	pldirloc := "/Volumes/"
	youtubeplaylistprefix := "https://www.youtube.com/"
	if *debug {
		fmt.Println("HOME: " + userhome)
		fmt.Println("pldirloc: " + testpldirloc)
		fmt.Println("youtubeplaylistprefix: " + youtubeplaylistprefix)
	}
	for i := 0; i < len(kids.Kids); i++ {
		if *debug {
			fmt.Println(kids.Kids[i].Location)
			go initiateDL(userhome+testpldirloc+kids.Kids[i].Location, youtubeplaylistprefix+kids.Kids[i].Playlist)
		} else {
			go initiateDL(userhome+pldirloc+kids.Kids[i].Location, youtubeplaylistprefix+kids.Kids[i].Playlist)
		}
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
		dl(location, playlist)
	} else {
		fmt.Printf("::: %s drive does NOT exist :::\n", location)
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
