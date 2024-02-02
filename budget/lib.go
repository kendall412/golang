package main

import (
	"encoding/csv"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func openCsv(file string, debug *bool) [][]string {
	csv_file, err := os.Open(file)
	if err != nil {
		log.Fatalln("error reading csv file", err)
	}
	defer csv_file.Close()

	csv_reader := csv.NewReader(csv_file)
	records, err := csv_reader.ReadAll()

	if err != nil {
		log.Fatalln("Error occured in opening csv_reader")
	}

	if *debug {
		log.Println(records)
	}
	return records
}

func retrieveTarget(target *string, records [][]string, header map[string]int, debug *bool) {
	var sum float64
	sum = 0
	for _, record := range records {
		if strings.Contains(strings.ToLower(record[header["DESC"]]), strings.ToLower(*target)) {
			if *debug {
				log.Println(record[header["DESC"]])
			}
			// converts string to float64
			amnt, _ := strconv.ParseFloat(record[header["AMNT"]], 64)
			sum += math.Abs(amnt)
			// log.Printf("%T", amnt)
			log.Printf("%s $%.2f", record[header["DATE"]], math.Abs(amnt))
		}
	}
	log.Printf("%s: $%.2f", *target, sum)
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
