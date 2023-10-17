package main

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"time"
)

var paths = `{
    "shinoo":"PLhGL5JwvKzCUgBTqnB-AUB_TAU2RbFNIR&si=sYu3gEOMGRPpBcMA",
    "naami":"PLhGL5JwvKzCWiKZLT-8Ydi46_riEe-KRm&si=0MgpvJFFLyBD6fhd",
    "woojin":"PLhGL5JwvKzCV5-OSq5ANnrjtEBEcV22aI&si=b8iNNfL2DWPHz7qW",
    "wonoo":"PLhGL5JwvKzCX7kzYOhAgVgyTJMrga224h&si=aUw4EyGbsY_2BUbf"
}`

var paths_long = `[
	{
	"name":"shinoo",
	"playlist":"PLhGL5JwvKzCUgBTqnB-AUB_TAU2RbFNIR&si=sYu3gEOMGRPpBcMA"
	},

	{
	"name":"naami",
	"playlist":"PLhGL5JwvKzCWiKZLT-8Ydi46_riEe-KRm&si=0MgpvJFFLyBD6fhd"
	},

	{
	"name":"woojin",
	""playlist:"PLhGL5JwvKzCV5-OSq5ANnrjtEBEcV22aI&si=b8iNNfL2DWPHz7qW"
	},

	{
	"name":"wonoo",
	"playlist":"PLhGL5JwvKzCX7kzYOhAgVgyTJMrga224h&si=aUw4EyGbsY_2BUbf"
	}
]`

var prefix string = "https://youtube.com/playlist?list="
var VER string = "1.0.1"

func main() {
	fmt.Printf("VER: %s\n", VER)
	npath := openJsonMap(paths)
	// require type assert from interface{} to string
	var url string = prefix + (npath["shinoo"].(string))
	fmt.Println(paths)
	// fmt.Println(url)
	dest := "/Volumes/SHINOO"

	if checkDir(dest) == true {
		fmt.Printf("The provided directory named %s exists.\n", dest)
		removeAllFiles(dest)
		time.Sleep(2)
		dl(dest, url)
	} else {
		fmt.Printf("%s does not exist.\n", dest)
	}
}

/*
checkDir
DESC: will check to see if the ssd cards are in place.
*/
func checkDir(dir string) bool {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		return (false)
	} else {
		return (true)
	}
}

func openJsonMap(urlJson string) map[string]interface{} {
	var path map[string]interface{}

	err := json.Unmarshal([]byte(urlJson), &path)
	if err != nil {
		fmt.Println(err)
	}
	return path
}

/*
removeAllFiles
DESC: removes all files in 'dest'
*/
func removeAllFiles(dest string) {
	dir, err := os.ReadDir(dest)
	// os.ReadDir(path) will create slice of all files in path
	if err != nil {
		fmt.Printf("could not remove all files from %s", dest)
	}
	for _, d := range dir {
		os.RemoveAll(path.Join([]string{dest, d.Name()}...))
	}
}

/*
dl
DESC: will spawn a process to invoke yt-dlp to begin downloading from playlist.
*/
func dl(dest string, url string) {
	path := "/opt/homebrew/bin/"
	ytScript := "yt-dlp"
	ytPath := filepath.Join(path, ytScript)

	optformat1 := "-x"
	optFormat2 := "--audio-format"
	optFormat3 := "mp3"
	optPath := "--path"

	fmt.Printf("%s %s %s %s %s %s\n", ytPath, optformat1, optFormat2, optFormat3, optPath, dest, url)
	_, err := exec.Command(ytPath, optformat1, optFormat2, optFormat3, optPath, dest, url).Output()

	if err != nil {
		fmt.Println(err)
	}
}
