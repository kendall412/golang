package main

func main() {
	debug, target := makeFlags()

	file := "checkings_short.csv"
	// file := "2024_01-01_01-30_checking.csv"
	// open csv file. record is type of [][]string
	records := openCsv(file, debug)

	header := map[string]int{"DATE": 0, "AMNT": 1, "CHK_NO": 3, "DESC": 4}

	retrieveTarget(target, records, header, debug)
}
