package main

import (
	"fmt"
	"strings"
)

var VER string = "0.1d"

var EXIT_CODE_SUCCESS = 0
var EXIT_CODE_ERROR = 1

func main() {
	italic.Printf("VER: %s\n\n", VER)

	// flags (pointers)
	target, csvfile, cattarget, debug, alltarget, display_all, essentialtarget, view_remaining_spending := generateFlags()

	// open csv file. record is type of [][]string
	if *csvfile == "" {
		printError("No csv file path was given.")
	} else {
		records := openCsv(*csvfile, debug, display_all)
		header := map[string]int{
			"DATE":   0,
			"AMNT":   1,
			"CHK_NO": 3,
			"DESC":   4}

		generateAll(debug, display_all)
		generateEssential(debug, display_all)
		target_map := generateMap(debug, display_all)

		sum := 0.0

		if *cattarget != "" {
			for _, targets := range target_map[*cattarget] {
				amnt := retrieveTargets(targets, records, header, debug)
				sum += amnt
			}
			printSum(sum)

		}

		// all spending
		if *alltarget {
			for _, targets := range all {
				amnt := retrieveTargets(targets, records, header, debug)
				sum += amnt
			}
			printSum(sum)
		}

		// essential spending is all - motorcycles
		if *essentialtarget {
			for _, targets := range essential {
				amnt := retrieveTargets(targets, records, header, debug)
				sum += amnt
			}
			printSum(sum)
		}

		// individual spending targets
		if *target != "" {
			sum := retrieveTargets(*target, records, header, debug)
			printSum(sum)
		}

		// check unknown spending
		if *view_remaining_spending {

			fmt.Println(records, len(records))
			fmt.Println(grocery)
			fmt.Println()

			for _, tar := range grocery {
				// fmt.Println(tar)
				for i, record := range records {
					// strings.Contains(str, input)
					if strings.Contains(strings.ToLower(record[header["DESC"]]), tar) {
						// fmt.Println(i, tar, record)
						delElement(i, &records)
					}
				}
			}
			fmt.Println()
			fmt.Println(records, len(records))
		}

	}
}
