package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func generateHeader(display_all *bool, header map[string]int) {
	header["DATE"] = 0
	header["AMNT"] = 1
	header["CHK_NO"] = 3
	header["DESC"] = 4
	header["MISC"] = 2

	if *display_all {
		printDisplayAll("Header", header)
	}
}

func printDisplayAll(name string, data interface{}) {
	fmt.Println()
	magenta.Printf("%s: ", name)
	magenta.Println(data)
	fmt.Println()
}

func openCsv(file string, debug *bool, display_all *bool, month *int, header map[string]int) [][]string {
	csv_file, err := os.Open(file)
	if err != nil {
		red.Println("Error openinging csv file", err)
		os.Exit(EXIT_CODE_ERROR)
	}
	defer csv_file.Close()

	csv_reader := csv.NewReader(csv_file)
	rec, err := csv_reader.ReadAll()

	if err != nil {
		red.Println("Error occured in opening csv_reader")
	}

	// *month is default to zero if no month is given by user
	if *month == 0 {
		if *display_all {
			printDisplayAll("Records", rec)
		}
		return rec
	} else {
		records := [][]string{}
		month_str := strconv.Itoa(*month)
		if len(month_str) == 1 {
			month_str = "0" + month_str
		}

		for _, v := range rec {
			if strings.HasPrefix(v[header["DATE"]]+"/", month_str) {
				records = append(records, v)
			}
		}

		if *display_all {
			printDisplayAll("Records", records)
		}
		return records
	}
}

func retrieveTargets(target string, records [][]string, header map[string]int, debug *bool) float64 {
	var sum float64
	sum = 0.0
	for _, record := range records {
		if strings.Contains(strings.ToLower(record[header["DESC"]]), strings.ToLower(target)) {
			if *debug {
				log.Println(record[header["DESC"]])
			}
			// converts string to float64
			amnt, _ := strconv.ParseFloat(record[header["AMNT"]], 64)

			if amnt < 0.0 {
				sum += math.Abs(amnt)
				fmt.Printf("%s $%.2f\n", record[header["DATE"]], math.Abs(amnt))
			}
		}
	}
	/*
	   only print out to terminal if the cateory/individual target was spent. otherwise will return sum of 0 but will not print out.
	*/
	if sum > 0.0 {
		blue.Printf("%s:", strings.ToUpper(target))
		green.Printf(" $%.2f\n", sum)
		fmt.Println()
	}
	return sum
}

func retrieveTargetSlice(target *Targets, total_sum *float64, header map[string]int, records *[][]string) {
	sum := 0.0
	for _, tar_variant := range target.Variant {
		// fmt.Println(tar_variant)
		for _, record := range *records {
			if strings.Contains(strings.ToLower(record[header["DESC"]]), strings.ToLower(tar_variant)) {
				// converts string to float64
				amnt, _ := strconv.ParseFloat(record[header["AMNT"]], 64)

				if amnt < 0.0 {
					sum += math.Abs(amnt)
					fmt.Printf("%s $%.2f\n", record[header["DATE"]], math.Abs(amnt))
				}
			}
		}
		/*
			only print out to terminal if the cateory/individual target was spent. otherwise will return sum of 0 but will not print out.
		*/
		if sum > 0.0 {
			blue.Printf("%s:", strings.ToUpper(tar_variant))
			green.Printf(" $%.2f\n", sum)
			fmt.Println()
		}
	}
	// fmt.Printf("%s $%.2f\n", minirecord[header["DATE"]],math.Abs(amnt))
	*total_sum = sum
}

/*
DESC: return sum of all spending. It takes into account variant names of a target. e.g. In-N-Out with variant names such as "in-n-out" and "in n out"

PARAM:

	variants *[]string  this is type Targets struct
	type Targets struct {
		Name    string
		Cat     []string // eating_out, grocery, etc
		Variant []string // {"in-n-out","in n out"}
		}

	total_sum *float64	this is set to 0.0
	header *map[string]int
	debug *bool
	records *[][]string

RETURN: None
*/
func returnSpending(variants *[]string, total_sum *float64, header *map[string]int, debug *bool, records *[][]string) {
	for _, variant := range *variants {
		sum := retrieveTargets(variant, *records, *header, debug)
		*total_sum += sum
	}
}
