package main

import (
	"fmt"
	"slices"
	"strings"
)

var VER string = "0.1i"
var EXIT_CODE_SUCCESS = 0
var EXIT_CODE_ERROR = 1

func main() {
	italic.Printf("BUDGET VER: %s\n\n", VER)

	// flags (pointers)
	target, csvfile, cattarget, month, debug, alltarget, display_all, variable_target, constant_target, view_remaining_spending, print_cat := generateFlags()

	header := make(map[string]int)
	generateHeader(display_all, header)

	fmt.Println(*print_cat)

	records := openCsv(*csvfile, debug, display_all, month, header)

	//========== Target Struct
	//========================================================================
	var targets_slice []Targets
	generateTargetStruct(&targets_slice)
	if *display_all {
		fmt.Println(targets_slice)
		fmt.Printf("%T\n", targets_slice)
		fmt.Println()
	}

	total_sum := 0.0

	if *target != "" {
		for _, tar := range targets_slice {
			if strings.Contains(tar.Name, *target) {
				retrieveTargetSlice(&tar, &total_sum, header, &records)
			}
		}
		printSum(total_sum)
	}

	// if *target != "" {
	// 	// total_sum := 0.0
	// 	for _, tar := range targets_slice {
	// 		if strings.Contains(tar.Name, *target) {
	// 			returnSpending(&tar.Variant, &total_sum, &header, debug, &records)
	// 		}
	// 	}
	// 	printSum(total_sum)
	// }

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
			// fmt.Println(tar)
			returnSpending(&tar.Variant, &total_sum, &header, debug, &records)
		}
		printSum(total_sum)
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

	// check unknown spending
	if *view_remaining_spending {
		for _, tar := range targets_slice {
			for _, record := range records {
				// strings.Contains(str, input)
				for _, variant := range tar.Variant {
					if strings.Contains(strings.ToLower(record[header["DESC"]]), variant) {

						// fmt.Println(i, tar, record)
						record[header["DATE"]] = " "
						record[header["DESC"]] = " "
						record[header["AMNT"]] = " "
						record[header["CHK_NO"]] = " "
						record[header["MISC"]] = " "
					}
				}
			}
		}
		fmt.Println()
		for _, record := range records {
			fmt.Println(record)
		}
	}
}
