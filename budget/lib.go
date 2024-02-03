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
	sum = 0
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
	blue.Printf("%s:", strings.ToUpper(target))
	green.Printf(" $%.2f\n", sum)
	fmt.Println()
	return sum
}

func getIndex(target *string, records [][]string, header map[string]int) []int {
	index_slice := []int{}
	for index, record := range records {
		if strings.Contains(strings.ToLower(record[header["DESC"]]), *target) {
			index_slice = append(index_slice, index)
		}
	}
	log.Printf("'%s' index:\n", *target)
	log.Println(index_slice)
	return index_slice
}
