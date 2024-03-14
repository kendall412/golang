package main

// func generateEatOutMap(display_all *bool) map[string][]string {
// 	var eating_out = make(map[string][]string)

// 	eating_out["innout"] = []string{"in-n-out", "in n out"}
// 	eating_out["ramen"] = []string{"ryujin ramen"}
// 	eating_out["subway"] = []string{"subway"}
// 	eating_out["brookfields"] = []string{"brookfield", "brookfieds"}
// 	eating_out["pho"] = []string{"pho bac"}
// 	eating_out["chipotle"] = []string{"chipotle"}
// 	eating_out["seoulzip"] = []string{"seoulzip"}
// 	eating_out["tasty_pot"] = []string{"tasty pot"}
// 	eating_out["mcdonalds"] = []string{"mcdonald", "mcdonalds"}
// 	eating_out["vons_chicken"] = []string{"vons chicken"}
// 	eating_out["paris_baguett"] = []string{"paris baguett"}
// 	eating_out["pushkin"] = []string{"pushkin"}

// 	if *display_all {
// 		printDisplayAll("eatout", eating_out)
// 	}

// 	return eating_out
// }

type Targets struct {
	Name    string
	Cat     []string // eating_out, grocery, etc
	Variant []string // {"in-n-out","in n out"}
}

var targets_slice []Targets

func generateTargetStruct() []Targets {
	// var targets_slice []Targets

	var innout Targets
	innout.Name = "in-n-out"
	innout.Cat = []string{"essential", "variable", "eatout"}
	innout.Variant = []string{"in-n-out", "in n out"}

	var mcdonalds Targets
	mcdonalds.Name = "mcdonalds"
	mcdonalds.Cat = []string{"variable", "eatout"}
	mcdonalds.Variant = []string{"mcdonald", "mcdonalds"}

	var cyclegear Targets
	cyclegear.Name = "cycle gear"
	cyclegear.Cat = []string{"motorcycle", "variable"}
	cyclegear.Variant = []string{"cycle gear"}

	var a_and_s Targets
	a_and_s.Name = "a_and_s"
	a_and_s.Cat = []string{"motorcycle", "variable"}
	a_and_s.Variant = []string{"a&s"}

	var housing Targets
	housing.Name = "housing"
	housing.Cat = []string{"constant", "housing"}
	housing.Variant = []string{"cheung"}

	var clothing Targets
	clothing.Name = "clothes"
	clothing.Cat = []string{"constant", "clothes"}
	clothing.Variant = []string{
		"shein",
		"marshalls",
		"ross",
		"macy's",
		"abercrombie",
		"arden fair"}

	var monthly Targets
	monthly.Name = "monthly"
	monthly.Cat = []string{"variable", "monthly"}
	monthly.Variant = []string{
		"lees korean ma",
		"epicentre church",
		"24 hour fitness",
		"technique gymn",
		"netflix",
		"audible",
		"apple",
		"github",
		"microsoft"}

	var targets_slice = append(targets_slice, innout, mcdonalds, cyclegear, a_and_s)

	return targets_slice
}
