package main

import (
	"log"
)

// cat
var all = []string{}
var essential = []string{}

// target slices
var income = []string{"hyve"}
var housing = []string{"cheung"}
var monthly = []string{
	"lees korean ma",
	"epicentre church",
	"24 hour fitness",
	"technique gymn",
	"netflix",
	"audible",
	"apple",
	"github",
	"microsoft",
}
var clothing = []string{
	"marshalls",
	"macy's",
	"abercrombie",
	"arden fair"}
var utils = []string{
	"pg&e",
	"smud",
	"comcast"}
var gas = []string{
	"arco",
	"shell"}
var grocery = []string{
	"sprouts",
	"costco",
	"kp international",
	"wal-mart",
	"raley",
	"target",
	"groc outlet"}
var health = []string{"kaiser"}
var food_out = []string{
	"in-n-out",
	"chipotle",
	"seoulzip",
	"tasty pot"}
var misc = []string{"amzn", "amazon", "hobbylobb"}
var coffee = []string{
	"starbucks",
	"peet's"}
var insurance = []string{"nationwide", "geico"}
var auto = []string{"palisade"}
var motorcycle = []string{
	"cycle gear",
	"a&s"}

func generateAll(debug, display_all *bool) {
	/*
		only 2 slices can be merged at a time.
	*/
	all = append(utils, gas...)
	all = append(all, housing...)
	all = append(all, monthly...)
	all = append(all, clothing...)
	all = append(all, health...)
	all = append(all, grocery...)
	all = append(all, auto...)
	all = append(all, misc...)
	all = append(all, food_out...)
	all = append(all, coffee...)
	all = append(all, insurance...)
	all = append(all, motorcycle...)

	if *display_all {
		log.Println(all)
	}
}

func generateEssential(debug, display_all *bool) {
	/*
		only 2 slices can be merged at a time.
	*/
	essential = append(utils, gas...)
	essential = append(essential, housing...)
	essential = append(essential, clothing...)
	essential = append(essential, monthly...)
	essential = append(essential, health...)
	essential = append(essential, grocery...)
	essential = append(essential, auto...)
	essential = append(essential, misc...)
	essential = append(essential, food_out...)
	essential = append(essential, insurance...)
	essential = append(essential, coffee...)

	if *display_all {
		log.Println(essential)
	}
}
