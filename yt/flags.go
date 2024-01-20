package main

import "flag"

func makeFlags() (*bool, *bool, *string, *string, *string) {
	debug := flag.Bool("debug", false, "debug mode <bool>: if TRUE will not dl files.")
	audio := flag.Bool("a", false, "audio <bool>: if TRUE file will convert to audio file(s).")
	url := flag.String("u", "", "url <str>: URL for video stream. Default is ' '.")
	dest := flag.String("d", ".", "destination <str>: destination for dl file.")
	audiourl := flag.String("au", "", "audio url <str>: if this url is give will download audio file for the url with only one option.")

	flag.Parse()
	return debug, audio, url, dest, audiourl
}
