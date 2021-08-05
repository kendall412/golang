package main

import (
	"fmt"
	"strings"
)

func printThis(str string) {
	fmt.Println(str)
}

func main() {
	s := "ffmpeg -ss 00:08:35 -i YDXJ0006.MP4 -t 00:00:30 -c:v copy -c:a copy -y YDXJ0006-short.MP4"

	args := strings.Split(s," ")
	fmt.Println(args[0:]...)
}