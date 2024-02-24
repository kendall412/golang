package main

import (
	"fmt"
	"strings"
)

var VER string = "0.1g"

var EXIT_CODE_SUCCESS = 0
var EXIT_CODE_ERROR = 1

func main() {
	italic.Printf("VER: %s\n\n", VER)

	// flags (pointers)
	target, csvfile, cattarget, month, debug, alltarget, display_all, essentialtarget, view_remaining_spending := generateFlags()

	month_map := map[string]int{
		"jan": 1, "feb": 2, "mar": 3,
		"apr": 4, "may": 5, "jun": 6,
		"jul": 7, "aug": 8, "sep": 9,
		"oct": 10, "nov": 11, "dec": 12}

	fmt.Println(*month)
	fmt.Println(month_map)

	// open csv file. record is type of [][]string
	if *csvfile == "" {
		printError("No csv file path was given.")
	} else {
		records := openCsv(*csvfile, debug, display_all)
		header := map[string]int{
			"DATE":   0,
			"AMNT":   1,
			"CHK_NO": 3,
			"DESC":   4,
			"MISC":   2}

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
			for _, tar := range all {
				for _, record := range records {
					// strings.Contains(str, input)
					if strings.Contains(strings.ToLower(record[header["DESC"]]), tar) {
						// fmt.Println(i, tar, record)
						record[header["DATE"]] = " "
						record[header["DESC"]] = " "
						record[header["AMNT"]] = " "
						record[header["CHK_NO"]] = " "
						record[header["MISC"]] = " "
					}
				}
			}
			fmt.Println()
			for _, record := range records {
				fmt.Println(record)
			}
		}
	}
}
