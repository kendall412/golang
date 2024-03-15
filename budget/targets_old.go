package main

/*
see https://go.dev/play/p/SjMWHTrIHNy for creating map of targets
*/

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
	"shein",
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
	"butcherbox",
	"smile",
	"safeway",
	"sprout",
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
	// "ryujin ramen",
	// "subway",
	// "brookfields",
	// "pho bac",
	// "in-n-out",
	// "in n out",
	"chipotle",
	"seoulzip",
	"tasty pot",
	"mcdonald's",
	"vons chicken",
	"paris baguett",
	"pushkin"}

var misc = []string{
	"rc country hobbies",
	"halfpricebooks",
	"amzn",
	"amazon",
	"hobbylobb",
	"dollar tr",
	"ebay"}
var coffee = []string{
	"temple coffee",
	"starbucks",
	"peet's"}
var insurance = []string{
	"nationwide",
	"geico"}
var auto = []string{
	"palisade",
	"black rock auto"}
var motorcycle = []string{
	"cycle gear",
	"a&s",
	"freedomroad"}
var pipe = []string{
	"cigars",
	"mission pipe shop"}

func generateEssential(display_all *bool, essential *[]string) {
	/*
		only 2 slices can be merged at a time.
		slice pointer
	*/
	*essential = append(utils, gas...)
	*essential = append(*essential, housing...)
	*essential = append(*essential, loan...)
	*essential = append(*essential, check...)
	*essential = append(*essential, clothing...)
	*essential = append(*essential, pipe...)
	*essential = append(*essential, monthly...)
	*essential = append(*essential, health...)
	*essential = append(*essential, grocery...)
	*essential = append(*essential, auto...)
	*essential = append(*essential, misc...)
	*essential = append(*essential, food_out...)
	*essential = append(*essential, insurance...)
	*essential = append(*essential, coffee...)

	if *display_all {
		printDisplayAll("Essential", essential)
	}
}

func generateAll(display_all *bool, all *[]string) {
	*all = append(utils, gas...)
	*all = append(*all, housing...)
	*all = append(*all, loan...)
	*all = append(*all, check...)
	*all = append(*all, pipe...)
	*all = append(*all, monthly...)
	*all = append(*all, clothing...)
	*all = append(*all, health...)
	*all = append(*all, grocery...)
	*all = append(*all, auto...)
	*all = append(*all, misc...)
	*all = append(*all, food_out...)
	*all = append(*all, coffee...)
	*all = append(*all, insurance...)
	*all = append(*all, motorcycle...)

	if *display_all {
		printDisplayAll("All", all)
	}
}

func generateMap(display_all *bool, target_map map[string][]string) {
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
		printDisplayAll("Target", target_map)
	}

}

func generateHeader(display_all *bool, header map[string]int) {
	header["DATE"] = 0
	header["AMNT"] = 1
	header["CHK_NO"] = 3
	header["DESC"] = 4
	header["MISC"] = 2

	if *display_all {
		printDisplayAll("Header", header)
	}
}

func generateMonth(display_all *bool, month_map map[string]int) {
	month_map["jan"] = 1
	month_map["feb"] = 2
	month_map["mar"] = 3
	month_map["apr"] = 4
	month_map["may"] = 5
	month_map["jun"] = 6
	month_map["jul"] = 7
	month_map["aug"] = 8
	month_map["sep"] = 9
	month_map["oct"] = 10
	month_map["nov"] = 11
	month_map["dec"] = 12

	if *display_all {
		printDisplayAll("Month", month_map)
	}
}
