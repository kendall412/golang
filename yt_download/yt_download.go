package main

import (
	"encoding/json"
	"fmt"
	"os/exec"
	"path/filepath"
	"time"
)

var paths = `{
    "shinoo":"PLhGL5JwvKzCUgBTqnB-AUB_TAU2RbFNIR&si=sYu3gEOMGRPpBcMA",
    "naami":"PLhGL5JwvKzCWiKZLT-8Ydi46_riEe-KRm&si=0MgpvJFFLyBD6fhd",
    "woojin":"PLhGL5JwvKzCV5-OSq5ANnrjtEBEcV22aI&si=b8iNNfL2DWPHz7qW",
    "wonoo":"PLhGL5JwvKzCX7kzYOhAgVgyTJMrga224h&si=aUw4EyGbsY_2BUbf"
}`

var prefix string = "https://youtube.com/playlist?list="

func main() {
	npath := openJsonMap(paths)
	// require type assert from interface{} to string
	var url string = prefix + (npath["shinoo"].(string))

	dest := "/Volumes/SHINOO"
	// rmAllFiles(dest)
	time.Sleep(2)
	dl(dest, url)
}

func openJsonMap(urlJson string) map[string]interface{} {
	var path map[string]interface{}

	err := json.Unmarshal([]byte(urlJson), &path)
	if err != nil {
		fmt.Println(err)
	}

	return path
}

func rmAllFiles(dest string) {
	opt := "rm"
	opt2 := "*.mp3"
	files := filepath.Join(dest, opt2)

	fmt.Printf("%s %s\n", opt, files)

	_, err := exec.Command(opt, files).Output()
	if err != nil {
		fmt.Println(err)
	}
}

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
