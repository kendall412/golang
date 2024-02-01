package main

import (
	"encoding/csv"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file := "checkings_short.csv"
	csv_file, err := os.Open(file)
	if err != nil {
		log.Fatalln("error reading csv file", err)
	}

	defer csv_file.Close()

	csv_reader := csv.NewReader(csv_file)
	records, err := csv_reader.ReadAll()

	if err != nil {
		log.Fatalln("could not read all")
	}

	// transactions := map[string]map[string]string{}

	// DATE := 0 // date of purchase
	AMNT := 1 // purchase amount
	// CHK_NO := 3 // check number
	DESC := 4

	// for _, record := range records {
	// 	if strings.Contains(record[DESC], "ZELLE") {
	// 		log.Println(record[DESC])
	// 	}
	// }

	// sum := 0
	for _, record := range records {
		if strings.Contains(record[DESC], "COSTCO") {
			log.Println(record[DESC], record[AMNT])

			amnt := strconv.ParseFloat(record[AMNT], 64)
			// sum += amnt
			log.Printf("%T", amnt)
			log.Println(amnt)
		}

	}
}
