package main

import "flag"

func generateFlags() (target *string, csvfile *string, cattarget *string, debug *bool, alltarget *bool, display_all *bool, essentialtarget *bool) {
	// bool flags
	target = flag.String("t", "", "target <str>: extract individual spending target (e.g. 'in-n-out')")

	csvfile = flag.String("f", "", "csvfile <str>: csv file path.")

	cattarget = flag.String("ct", "", "cattarget <str>: categorical target (e.g. 'insurance' vs 'geico').")

	// string flags
	debug = flag.Bool("debug", false, "debug mode <bool>: default FALSE, if TRUE will put script in debug mode.")

	alltarget = flag.Bool("at", false, "targets <bool>: default FALSE, if TRUE will output all spending.")

	display_all = flag.Bool("da", false, "display_all <bool>: default FALSE, if TRUE will display all details for debug purpose.")

	essentialtarget = flag.Bool("es", false, "essential <bool>: default FALSE, if TRUE will display essential (all - motorcycle spending) spending.")

	flag.Parse()
	return
}
