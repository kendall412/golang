package main

import (
	"log"
)

func main() {
	debug, target, alltarget, display_all := makeFlags()

	// file := "checkings_short.csv"
	file := "2024_01-01_01-30_checking.csv"

	// open csv file. record is type of [][]string
	records := openCsv(file, debug, display_all)
	header := map[string]int{"DATE": 0, "AMNT": 1, "CHK_NO": 3, "DESC": 4}
	generateTargets(debug, display_all)

	sum := 0.0
	if *alltarget {
		for _, targets := range all {
			amnt := retrieveTargets(targets, records, header, debug)
			log.Println()
			sum += amnt
		}
		log.Printf("total: %.2f", sum)
	}

	if *target != "" {
		log.Printf("target: %s\n", *target)
		retrieveTarget(target, records, header, debug)
		// targetsSlice := getIndex(target, records, header)
	}
}
