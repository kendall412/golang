package main

import "flag"

func makeFlags() (debug *bool, target *string) {
	debug = flag.Bool("debug", false, "debug mode <bool>: if TRUE will not dl files.")

	target = flag.String("t", "starbucks", "target <str>: exact target to search")

	flag.Parse()
	return
}
