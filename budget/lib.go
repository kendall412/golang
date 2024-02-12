package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

func printError(msg string) {
	repeatno := 40
	red.Println(strings.Repeat("*", repeatno))
	red.Println("ERROR: " + msg)
	red.Println(strings.Repeat("*", repeatno))
	os.Exit(EXIT_CODE_ERROR)
}

func printSum(sum float64) {
	blue.Printf("TOTAL: ")
	green.Printf("$%.2f\n", sum)
}

func openCsv(file string, debug *bool, display_all *bool) [][]string {
	csv_file, err := os.Open(file)
	if err != nil {
		red.Println("Error openinging csv file", err)
		os.Exit(EXIT_CODE_ERROR)
	}
	defer csv_file.Close()

	csv_reader := csv.NewReader(csv_file)
	records, err := csv_reader.ReadAll()

	if err != nil {
		red.Println("Error occured in opening csv_reader")
	}

	if *display_all {
		log.Println(records)
	}
	return records
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
			sum += math.Abs(amnt)
			fmt.Printf("%s $%.2f\n", record[header["DATE"]], math.Abs(amnt))
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

func delElement(index int, records *[][]string) {
	/*
		using pointer slices
	*/
	*records = slices.Delete(*records, index-1, index)
}
