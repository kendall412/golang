package main

import (
	"log"
)

var all = []string{}
var housing = []string{"cheung"}
var clothing = []string{"marshalls", "macy's", "abercrombie", "arden fair"}
var utils = []string{"pg&e", "smud", "comcast"}
var gas = []string{"arco", "shell"}
var grocery = []string{"sprouts", "costco", "kp international", "wal-mart", "raley", "target", "groc outlet"}
var health = []string{"kaiser"}
var food_out = []string{"in-n-out", "chipotle", "seoulzip"}
var misc = []string{"amzn"}
var coffee = []string{"starbucks", "peet's"}
var motorcycle = []string{"cycle gear", "a&s"}

func generateTargets(debug, display_all *bool) {
	all = append(utils, gas...)
	all = append(all, housing...)
	all = append(all, clothing...)
	all = append(all, health...)
	all = append(all, grocery...)
	all = append(all, misc...)
	all = append(all, food_out...)
	all = append(all, coffee...)
	all = append(all, motorcycle...)
	if *display_all {
		log.Println(all)
	}
}
