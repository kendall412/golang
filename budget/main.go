package main

var VER string = "0.1a"

var EXIT_CODE_SUCCESS = 0
var EXIT_CODE_ERROR = 1

func main() {
	italic.Printf("VER: %s\n\n", VER)

	// flags
	debug, target, alltarget, display_all := makeFlags()

	// file := "checkings_short.csv"
	file := "2024_01-01_01-30_checking.csv"

	// open csv file. record is type of [][]string
	records := openCsv(file, debug, display_all)
	header := map[string]int{"DATE": 0, "AMNT": 1, "CHK_NO": 3, "DESC": 4}
	generateAll(debug, display_all)
	generateEssential(debug, display_all)

	sum := 0.0
	// all spending
	if *alltarget {
		for _, targets := range all {
			amnt := retrieveTargets(targets, records, header, debug)
			// log.Println()
			sum += amnt
		}
		blue.Printf("TOTAL: ")
		green.Printf("$%.2f\n", sum)
	}

	// user selected target
	if *target != "" {
		item := *target
		retrieveTargets(item, records, header, debug)
	}
}
