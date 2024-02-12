package main

import (
	"log"
)

// category
var all = []string{}
var essential = []string{}

// targets
var income = []string{"hyve"}
var check = []string{"check"}
var loan = []string{"dept education student ln"}
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
	"microsoft"}
var clothing = []string{
	"marshalls",
	"ross",
	"macy's",
	"abercrombie",
	"arden fair"}
var utils = []string{
	"pg&e",
	"smud",
	"comcast",
	"verizon",
	"google",
	"quickquack"}
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
	"groc outlet",
	"lucky",
	"hankook"}
var health = []string{"kaiser"}
var food_out = []string{
	"in-n-out",
	"chipotle",
	"seoulzip",
	"tasty pot",
	"mcdonald's",
	"vons chicken",
	"paris baguett"}
var misc = []string{
	"amzn",
	"amazon",
	"hobbylobb",
	"dollar tr",
	"ebay"}
var coffee = []string{
	"starbucks",
	"peet's"}
var insurance = []string{
	"nationwide",
	"geico"}
var auto = []string{"palisade"}
var motorcycle = []string{
	"cycle gear",
	"a&s",
	"freedomroad"}
var pipe = []string{
	"cigars",
	"mission pipe shop"}

func generateAll(debug, display_all *bool) {
	/*
		only 2 slices can be merged at a time.
	*/
	all = append(utils, gas...)
	all = append(all, housing...)
	all = append(all, loan...)
	all = append(all, check...)
	all = append(all, pipe...)
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
	essential = append(essential, loan...)
	essential = append(essential, check...)
	essential = append(essential, clothing...)
	essential = append(essential, pipe...)
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

func generateMap(debug, display_all *bool) map[string][]string {
	target_map := map[string][]string{}

	target_map["housing"] = housing
	target_map["clothing"] = clothing
	target_map["utils"] = utils
	target_map["gas"] = gas
	target_map["grocery"] = grocery
	target_map["health"] = health
	target_map["food_out"] = food_out
	target_map["misc"] = misc
	target_map["coffee"] = coffee
	target_map["insurance"] = insurance
	target_map["auto"] = auto
	target_map["motorcycle"] = motorcycle
	target_map["income"] = income
	target_map["monthly"] = monthly
	target_map["loan"] = loan
	target_map["pipe"] = pipe
	target_map["check"] = check

	if *display_all {
		log.Println(target_map)
	}
	return target_map
}
