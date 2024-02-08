package main

import (
	"time"
)

func convert_string_date(date_str []string) []time.Time {
	/*
	DESC: will convert 1/2/2004 format string to time.Time type.
 	PARAM: date_str []string
  	RETURN: time_date []time.Time
 	*/
	time_date := []time.Time{}
	for _, date_el := range date_str {
		date, _ := time.Parse("1/2/2006", date_el)

		// fmt.Println(date)
		// fmt.Printf("%T\n", date)

		// fmt.Printf("Year: %d\n", date.Year())
		// fmt.Printf("Month: %s\n", date.Month().String())
		// fmt.Printf("Day: %d\n\n", date.Day())
		time_date = append(time_date, date)
	}
	return time_date
}
