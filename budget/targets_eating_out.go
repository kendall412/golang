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

func generateEatOut() []Targets {
	var targets_slice []Targets

	var innout Targets
	innout.Name = "in-n-Oout"
	innout.Cat = []string{"all", "essential"}
	innout.Variant = []string{"in-n-out", "in n out"}

	var mcdonalds Targets
	mcdonalds.Name = "mcDonalds"
	mcdonalds.Cat = []string{"all"}
	mcdonalds.Variant = []string{"mcdonald", "mcdonalds"}

	var motorcycle Targets
	motorcycle.Name = "motorcycle"
	motorcycle.Cat = []string{"misc", "all"}
	motorcycle.Variant = []string{"cycle gear", "a&s", "freedomroad"}

	targets_slice = append(targets_slice, innout, mcdonalds, motorcycle)

	// if *display_all {
	// 	printDisplayAll(targets_slice)
	// }

	return targets_slice
}
