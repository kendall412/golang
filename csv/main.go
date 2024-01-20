package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

func main() {
	csv_file, err := os.Open("banklist.csv")
	if err != nil {
		log.Fatalln("error reading csv file", err)
	}

	defer csv_file.Close()

	csv_reader := csv.NewReader(csv_file)
	records, err := csv_reader.ReadAll()

	if err != nil {
		log.Fatalln("could not read all")
	}

	Banks := map[string]map[string]string{}

	CITY := 1
	BANKNAME := 0
	STATE := 2
	CERT := 3
	ACQ_BANK := 4
	DATE_CLOSING := 5
	FUND := 6

	for _, record := range records {
		if record[STATE] == "CA" {
			Banks[record[BANKNAME]] = map[string]string{"City": record[CITY], "State": record[STATE], "Cert": record[CERT], "Acq Bank": record[ACQ_BANK], "Date of Closing": record[DATE_CLOSING], "Fund": record[FUND]}
		}
	}
	bname := "Canyon National Bank"
	fmt.Println(Banks[bname])

}
