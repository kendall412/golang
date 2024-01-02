package main

import (
	"log"
)

func main() {
	debug, audio, url, dest := makeFlags()

	var opt []string
	cmd := "yt-dlp"

	if *debug {
		log.Println("Debug Mode ON")
	}

	if *url == "" {
		log.Fatalln("You must provide a URL for the video/audio stream.")
	}

	if *audio {
		log.Println("Audio Mode ON")
		// opt = append(opt, "-x", "--audio-format", "mp3", "--audio-quality", "0")

		opt = append(opt, "-x", "--audio-format", "mp3", "--audio-quality", "0")
	}
	opt = append(opt, "--path", *dest, *url)

	execProc(cmd, opt, debug)
}
