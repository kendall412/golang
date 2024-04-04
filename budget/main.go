package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
)

var VER string = "0.1k"
var EXIT_CODE_SUCCESS = 0
var EXIT_CODE_ERROR = 1

func main() {
	italic.Printf("BUDGET VER: %s\n\n", VER)

	// flags (pointers)
	target, csvfile, cattarget, month, debug, alltarget, display_all, view_remaining_spending, print_cat, list_cat := generateFlags()

	header := make(map[string]int)
	generateHeader(display_all, header)

	// prints categories
	if *print_cat {
		i := 1
		printInfo("CATEGORIES:")
		for _, cat := range cats {
			// convert integer to string
			printInfo(strconv.Itoa(i) + ". " + cat)
			i++
		}
		fmt.Println()
	}

	// if month value is none zero it will return month corresponding with the integer
	printMonth(month)
	records := openCsv(*csvfile, debug, month, header)

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
	variantTargets := [][]string{}

	if *target != "" {
		for _, tar := range targets_slice {
			if strings.Contains(tar.Name, *target) {
				retrieveVariantTargets(&tar, header, &records, &variantTargets)
			}
		}
		if *debug {
			iterate2DSlice(variantTargets)
		}
		retrieveSum(variantTargets, header, &total_sum, *target, debug)
		printSum(total_sum)
		variantTargets = nil
	}

	if *cattarget != "" {
		for _, tar := range targets_slice {
			if slices.Contains(tar.Cat, *cattarget) {
				retrieveVariantTargets(&tar, header, &records, &variantTargets)
			}
		}
		if *debug {
			iterate2DSlice(variantTargets)
		}
		retrieveSum(variantTargets, header, &total_sum, *cattarget, debug)
		printSum(total_sum)
		variantTargets = nil
	}

	if *alltarget {
		for _, tar := range targets_slice {
			retrieveVariantTargets(&tar, header, &records, &variantTargets)
		}
		retrieveSum(variantTargets, header, &total_sum, "ALL", debug)
		printSum(total_sum)
		variantTargets = nil
	}

	// lists spending not accounted in existing targets
	if *view_remaining_spending {
		for _, tar := range targets_slice {
			for _, record := range records {
				// strings.Contains(str, input)
				for _, variant := range tar.Variant {
					if strings.Contains(strings.ToLower(record[header["DESC"]]), variant) {
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

	if *list_cat {
		listCat(targets_slice)
	}
}
