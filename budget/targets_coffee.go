package main

func coffeeTargets() []Targets {
	var starbucks Targets
	starbucks.Name = "starbucks"
	starbucks.Cat = []string{cats["coff"], cats["vari"]}
	starbucks.Variant = []string{"starbucks"}

	var temple Targets
	temple.Name = "temple coffee"
	temple.Cat = []string{cats["coff"], cats["vari"]}
	temple.Variant = []string{"temple coffee"}

	var peets Targets
	peets.Name = "peets"
	peets.Cat = []string{cats["coff"], cats["vari"]}
	peets.Variant = []string{"peet's", "peets"}

	var coffee = []Targets{}
	coffee = append(coffee, starbucks, temple, peets)
	return coffee
}
