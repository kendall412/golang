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

func printMonth(month *int) {
	var month_str string
	if *month == 1 {
		month_str = "January"
	} else if *month == 2 {
		month_str = "February"
	} else if *month == 3 {
		month_str = "March"
	} else if *month == 4 {
		month_str = "April"
	} else if *month == 5 {
		month_str = "May"
	} else if *month == 6 {
		month_str = "June"
	} else if *month == 7 {
		month_str = "July"
	} else if *month == 8 {
		month_str = "August"
	} else if *month == 9 {
		month_str = "September"
	} else if *month == 10 {
		month_str = "October"
	} else if *month == 11 {
		month_str = "November"
	} else if *month == 12 {
		month_str = "December"
	} else if *month == 0 {
		month_str = "Annual"
	}

	green.Println(month_str)
}

func generateHeader(display_all *bool, header map[string]int) {
	header["DATE"] = 0
	header["AMNT"] = 1
	header["CHK_NO"] = 3
	header["DESC"] = 4
	header["MISC"] = 2
	header["TITLE"] = 5

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

func openCsv(file string, debug *bool, month *int, header map[string]int) [][]string {
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
		if *debug {
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

		if *debug {
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

func retrieveSum(records [][]string, header map[string]int, total_sum *float64, title string, debug *bool) {
	sum := 0.0
	for _, record := range records {
		amnt, _ := strconv.ParseFloat(record[header["AMNT"]], 64)
		if amnt < 0.0 {
			sum += math.Abs(amnt)
			// fmt.Printf("%s $%.2f %s\n", record[header["DATE"]], math.Abs(amnt), strings.ToUpper(record[header["TITLE"]]))
			fmt.Printf("%s ", record[header["DATE"]])
			yellow.Printf("%.2f ", math.Abs(amnt))
			red.Printf("%s\n", strings.ToUpper(record[header["TITLE"]]))
			if *debug {
				red.Printf("%s\n\n", record[header["DESC"]])
			}
		}
	}
	/*
	   only print out to terminal if the cateory/individual target was spent. otherwise will return sum of 0 but will not print out.
	*/
	if sum > 0.0 {
		// blue.Printf("%s:", strings.ToUpper(target))
		blue.Printf("%s: %.2f\n", strings.ToUpper(title), sum)
		fmt.Println()
	}
	*total_sum = sum
}

/*
DESC: this function returns another [][]string rec which has all the variants record (i.e. ["in n out","in-n-out"]). This should then be sent to retrieveSum()

PARAM:

	target *Targets
	header map[string]int
	records *[][]stirng

RETURN:

	[][]string
*/
func retrieveVariantTargets(target *Targets, header map[string]int, records *[][]string, variantTargets *[][]string) {
	for _, tar_variant := range target.Variant {
		for _, record := range *records {
			if strings.Contains(strings.ToLower(record[header["DESC"]]), strings.ToLower(tar_variant)) {
				record = append(record, target.Name)
				*variantTargets = append(*variantTargets, record)
			}
		}
	}
}

/*
DESC: return sum of all spending. It takes into account variant names of a target. e.g. In-N-Out with variant names such as "in-n-out" and "in n out"

PARAM:

	variants *[]string  this is type Targets struct
	type Targets struct {
		Name    string
		Cat     []string // eating_out, grocery, etc
		Variant []string // {"in-nfucnt-out","in n out"}
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

/*
DESC: iterates 2D slices for debug purposes
*/
func iterate2DSlice(sl [][]string) {
	for _, v := range sl {
		fmt.Println(v)
	}
	fmt.Println()
}

func listCat(targets_slice []Targets) {
	// var targets_slice []Targets
	// generateTargetStruct(&targets_slice)
	// fmt.Println(targets_slice)
	for i := 0; i < len(targets_slice); i++ {
		yellow.Printf(targets_slice[i].Name)
		fmt.Println(targets_slice[i].Variant)
		red.Println(targets_slice[i].Cat)
		fmt.Println()
	}
}
