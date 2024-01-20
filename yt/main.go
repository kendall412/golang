package main

import (
	"io"
	"log"
	"os"
)

var APPNAME string = "yt-dlp"
var VER string = "0.1.3"

func main() {
	/*
		logging file creation
	*/
	log_file, err := os.OpenFile("yt-dlp.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	/*
		output to os.Stdout and to file
	*/
	multiwriter := io.MultiWriter(os.Stdout, log_file)
	log.SetOutput(multiwriter)
	log.Println(APPNAME)
	log.Println("VER: " + VER)

	debug, audio, url, dest, audiourl := makeFlags()

	var opt []string
	cmd := "yt-dlp"

	if *debug {
		log.Println("Debug Mode: ON")
	} else {
		log.Println("Debug Mode: OFF")
	}

	if *audio {
		log.Println("Audio Mode: ON")
		opt = append(opt, "-x", "--audio-format", "mp3", "--audio-quality", "0")
	} else {
		log.Println("Video Mode: ON")
	}

	if *audiourl == "" && *url == "" {
		log.Fatalln("User must either give URL or Audio URL link")
	}

	if *audiourl != "" && *url == "" {
		log.Println("Audio Mode: ON")
		log.Println("audiourl: " + *audiourl)
		opt = append(opt, "-x", "--audio-format", "mp3", "--audio-quality", "0")
		opt = append(opt, "--path", *dest, *audiourl)
	}

	if *audiourl == "" && *url != "" {
		log.Println("url:\n\t" + *url)
		opt = append(opt, "--path", *dest, *url)
	}

	// opt = append(opt, "--path", *dest, *url)

	execProc(cmd, opt, debug)

	log.Println(border("*", 50))
	log_file.Close()
}
