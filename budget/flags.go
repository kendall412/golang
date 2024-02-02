package main

import "flag"

func makeFlags() (debug *bool, target *string, alltarget *bool, display_all *bool) {
	debug = flag.Bool("debug", false, "debug mode <bool>: if TRUE will not dl files.")
	alltarget = flag.Bool("at", false, "targets <bool>: if true will iterate over slice of targets")
	target = flag.String("t", "", "target <str>: exact target to search")
	display_all = flag.Bool("da", false, "display_all <bool>: if true will display all details")

	flag.Parse()
	return
}
