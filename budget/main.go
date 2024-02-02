package main

import (
	"log"
)

func main() {
	debug, target, alltarget := makeFlags()

	// file := "checkings_short.csv"
	file := "2024_01-01_01-30_checking.csv"

	// open csv file. record is type of [][]string
	records := openCsv(file, debug)

	header := map[string]int{"DATE": 0, "AMNT": 1, "CHK_NO": 3, "DESC": 4}

	all := []string{}
	housing := []string{"cheung"}
	clothing := []string{"marshalls", "macy's", "abercrombie", "arden fair"}
	utils := []string{"pg&e", "smud", "comcast"}
	gas := []string{"arco", "shell"}
	grocery := []string{"sprouts", "costco", "kp international", "wal-mart", "raley", "target", "groc outlet"}
	health := []string{"kaiser"}
	food_out := []string{"in-n-out", "chipotle", "seoulzip"}
	misc := []string{"amzn"}
	coffee := []string{"starbucks", "peet's"}
	motorcycle := []string{"cycle gear", "a&s"}

	all = append(utils, gas...)
	all = append(all, housing...)
	all = append(all, clothing...)
	all = append(all, health...)
	all = append(all, grocery...)
	all = append(all, misc...)
	all = append(all, food_out...)
	all = append(all, coffee...)
	all = append(all, motorcycle...)
	log.Println(all)

	sum := 0.0
	if *alltarget {
		for _, targets := range all {

			amnt := retrieveTargets(targets, records, header, debug)
			log.Println()
			sum += amnt
		}
		log.Printf("total: %.2f", sum)
	} else {
		log.Printf("target: %s\n", *target)
		retrieveTarget(target, records, header, debug)
		// targetsSlice := getIndex(target, records, header)
	}
}
