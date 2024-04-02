package main

type Targets struct {
	Name    string
	Cat     []string // eating_out, grocery, etc
	Variant []string // {"in-n-out","in n out"}
	Desc    string
}

var cats = map[string]string{
	"ba":    "bank",
	"inc":   "income",
	"fu":    "fuel",
	"ins":   "insurance",
	"coff":  "coffee",
	"auto":  "automobile",
	"lo":    "loan",
	"sub":   "subscription",
	"util":  "utility",
	"food":  "food",
	"mis":   "misc",
	"heal":  "health",
	"clo":   "clothes",
	"nes":   "none-essential",
	"vari":  "variable",
	"con":   "constant",
	"eat":   "eatout",
	"moto":  "motorcycle",
	"home":  "housing",
	"min":   "ministry",
	"ext":   "extracurricular",
	"enter": "entertainment",
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
}
