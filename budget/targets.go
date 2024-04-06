package main

import (
	"fmt"
	"strconv"
)

type Targets struct {
	Name     string
	Cat      []string // eatout, grocery, etc.
	Variant  []string // variation in names e.g. {"in-n-out","in n out"}
	Spending bool     // Spending=true spending, Spending=false deposit. Default is false.
	Desc     string
}

var cats = map[string]string{
	"ba":   "bank",
	"inc":  "income",
	"fu":   "fuel",
	"ins":  "insurance",
	"cins": "car_insurance",
	"mins": "moto_insurance",
	"coff": "coffee",
	"auto": "automobile",
	"lo":   "loans",
	"sub":  "subscription",
	"util": "utility",
	"food": "food",
	"mis":  "misc",
	"heal": "health",
	"clo":  "clothes",
	"nes":  "none-essential",
	"vari": "variable",
	"con":  "constant",
	"eat":  "eatout",
	"moto": "motorcycle",
	"home": "housing",
	"min":  "ministry",
	"ext":  "extracurricular",
	"ent":  "entertainment",
	"rec":  "recreation",
	"tax":  "tax",
	"fee":  "fees",
}

/*
DESC: lists the available categories with its abbreviation of spending.
PARAM: None
RETURN: None
*/
func listCats() {
	i := 1
	blue.Println("CATEGORIES:")
	for abb, cat := range cats {
		// convert integer to string
		yellow.Printf(strconv.Itoa(i) + ". ")
		green.Printf("[" + abb + "]" + " ")
		yellow.Println(cat)
		i++
	}
	fmt.Println()
}

func generateTargetStruct(targets_slice *[]Targets) {
	*targets_slice = append(*targets_slice, extracurricularTargets()...)
	*targets_slice = append(*targets_slice, housingTargets()...)
	*targets_slice = append(*targets_slice, subscriptionTarget()...)
	*targets_slice = append(*targets_slice, clothesTarget()...)
	*targets_slice = append(*targets_slice, nonessTargets()...)
	*targets_slice = append(*targets_slice, eatOutTargets()...)
	*targets_slice = append(*targets_slice, groceryTargets()...)
	*targets_slice = append(*targets_slice, utilsTargets()...)
	*targets_slice = append(*targets_slice, ministryTargets()...)
	*targets_slice = append(*targets_slice, loansTarget()...)
	*targets_slice = append(*targets_slice, coffeeTargets()...)
	*targets_slice = append(*targets_slice, insuranceTargets()...)
	*targets_slice = append(*targets_slice, miscTargets()...)
	*targets_slice = append(*targets_slice, fuelTargets()...)
	*targets_slice = append(*targets_slice, healthTargets()...)
	*targets_slice = append(*targets_slice, incomeTargets()...)
	*targets_slice = append(*targets_slice, bankTATargets()...)
	*targets_slice = append(*targets_slice, entertainmentTargets()...)
	*targets_slice = append(*targets_slice, taxTargets()...)
	*targets_slice = append(*targets_slice, feesTargets()...)
}
