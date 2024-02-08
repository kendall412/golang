package main

import "flag"

func makeFlags() (debug *bool, target *string, alltarget *bool, display_all *bool, csvfile *string, essentialtarget *bool, cattarget *string) {
	debug = flag.Bool("debug", false, "debug mode <bool>: if TRUE will not dl files.")

	alltarget = flag.Bool("at", false, "targets <bool>: if true will iterate over slice of targets")

	target = flag.String("t", "", "target <str>: exact individual spending to search")

	display_all = flag.Bool("da", false, "display_all <bool>: if true will display all details")

	csvfile = flag.String("f", "", "csvfile <str>: csv file path.")

	essentialtarget = flag.Bool("es", false, "essential <bool>: if true will display essential list of spending.")

	cattarget = flag.String("ct", "", "cattarget <str>: categorical target: e.g. insurance (rather than geico).")

	flag.Parse()
	return
}
