package main

import (
	"fmt"
	"os"
	"slices"
	"strings"
)

var VER string = "0.1h"
var EXIT_CODE_SUCCESS = 0
var EXIT_CODE_ERROR = 1

func main() {
	italic.Printf("BUDGET VER: %s\n\n", VER)

	// flags (pointers)
	target, csvfile, cattarget, month, debug, alltarget, display_all, essential_target, variable_target, constant_target, view_remaining_spending, print_cat := generateFlags()

	header := make(map[string]int)
	generateHeader(display_all, header)

	var target_map = make(map[string][]string)
	generateMap(display_all, target_map)

	// prints out categories of purchase and then exits
	if *print_cat {
		cnt := 1
		magenta.Printf("CATEGORIES:\n")
		for cat, _ := range target_map {
			magenta.Printf("%d. %s\n", cnt, cat)
			cnt += 1
		}
		fmt.Println()
		os.Exit(1)
	}

	var all []string
	generateAll(display_all, &all)

	var essential []string
	generateEssential(display_all, &essential)

	// open csv
	records := openCsv(*csvfile, debug, display_all, month, header)

	//========== Target Struct
	//========================================================================
	targets_slice := generateTargetStruct()
	fmt.Println(targets_slice)
	fmt.Println()

	total_sum := 0.0

	if *target != "" {
		// total_sum := 0.0
		for _, tar := range targets_slice {
			if strings.Contains(tar.Name, *target) {
				returnSpending(&tar.Variant, &total_sum, &header, debug, &records)
			}
		}
		printSum(total_sum)
	}

	if *cattarget != "" {
		// total_sum := 0.0
		for _, tar := range targets_slice {
			if slices.Contains(tar.Cat, *cattarget) {
				returnSpending(&tar.Variant, &total_sum, &header, debug, &records)
			}
		}
		printSum(total_sum)
	}

	if *alltarget {
		// total_sum := 0.0
		for _, tar := range targets_slice {
			fmt.Println(tar)
			// returnSpending(&tar.Variant, &total_sum, &header, debug, &records)
		}
		// printSum(total_sum)
	}

	if *variable_target {
		for _, tar := range targets_slice {
			if slices.Contains(tar.Cat, vari) {
				returnSpending(&tar.Variant, &total_sum, &header, debug, &records)
			}
		}
		printSum(total_sum)
	}

	if *constant_target {
		for _, tar := range targets_slice {
			if slices.Contains(tar.Cat, con) {
				returnSpending(&tar.Variant, &total_sum, &header, debug, &records)
			}
		}
		printSum(total_sum)
	}
	os.Exit(1)
	//========================================================================
	//========================================================================

	// open csv file. record is type of [][]string
	if *csvfile == "" {
		printError("No csv file path was given.")
	} else {
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
		if *essential_target {
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
