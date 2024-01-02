package main

import "flag"

func makeFlags() (*bool, *bool, *string, *string) {
	debug := flag.Bool("debug", false, "debug mode <bool>: in this mode app will not download any files.")
	audio := flag.Bool("a", false, "audio <bool>: if true the downloaded file will be converted to audio file(s).Downloaded files are in video format by default.")
	url := flag.String("u", "", "url <str>: URL for video stream.By default is ' '.")
	dest := flag.String("d", ".", "destination <str>: where the downloaded file will go.")

	flag.Parse()

	return debug, audio, url, dest
}
