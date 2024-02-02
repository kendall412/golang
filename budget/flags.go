package main

import "flag"

func makeFlags() (debug *bool, target *string, alltarget *bool) {
	debug = flag.Bool("debug", false, "debug mode <bool>: if TRUE will not dl files.")
	alltarget = flag.Bool("at", false, "targets <bool>: if true will iterate over slice of targets")
	target = flag.String("t", "starbucks", "target <str>: exact target to search")

	flag.Parse()
	return
}
