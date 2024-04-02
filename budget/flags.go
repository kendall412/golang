package main

import "flag"

func generateFlags() (target *string, csvfile *string, cattarget *string, month *int, debug *bool, alltarget *bool, display_all *bool, view_remaining_spending *bool, print_cat *bool, list_cat *bool) {
	// string flags
	target = flag.String("t", "", "target <str>: extract individual spending target (e.g. 'in-n-out')")

	csvfile = flag.String("f", "", "csvfile <str>: csv file path.")

	cattarget = flag.String("ct", "", "cattarget <str>: categorical target (e.g. 'insurance' vs 'geico').")

	// integer flags
	month = flag.Int("m", 0, "month <int>: month in integer format: jan:1, feb:2, mar:3, apr:4, may:5, jun:6, jul:7, aug:8, sep:9, oct:10, nov:11, dec:12")

	// bool flags
	debug = flag.Bool("debug", false, "debug mode <bool>: default FALSE, if TRUE will put script in debug mode.")

	alltarget = flag.Bool("at", false, "targets <bool>: default FALSE, if TRUE will output all spending.")

	display_all = flag.Bool("da", false, "display_all <bool>: default FALSE, if TRUE will display all details for debug purpose.")

	view_remaining_spending = flag.Bool("v", false, "view_remaining_spending <bool>: default FALSE, if TRUE will printout remaining spending that is not hard coded.")

	print_cat = flag.Bool("pct", false, "print_cat <bool>: default False, if True will print out categories available to choose")

	list_cat = flag.Bool("lct", false, "print_cat <bool>: default False, if True it will list Targets struct item.")

	flag.Parse()
	return
}
