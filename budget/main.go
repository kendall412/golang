package main

import "log"

func main() {
	debug, target := makeFlags()

	file := "checkings_short.csv"
	// file := "2024_01-01_01-30_checking.csv"
	// open csv file. record is type of [][]string
	records := openCsv(file, debug)

	header := map[string]int{"DATE": 0, "AMNT": 1, "CHK_NO": 3, "DESC": 4}

	log.Printf("Target: %s\n", *target)
	retrieveTarget(target, records, header, debug)
	targetsSlice := getIndex(target, records, header)

	if *debug {
		log.Printf("\n\n\n")
		for _, targetSlice := range targetsSlice {
			log.Println(records[targetSlice])
		}
	}

}
