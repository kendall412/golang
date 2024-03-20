package main

type Targets struct {
	Name    string
	Cat     []string // eating_out, grocery, etc
	Variant []string // {"in-n-out","in n out"}
	DESC    string
}

// spending categories
var inc = "income"
var fu = "fuel"
var ins = "insurance"
var coff = "coffee"
var auto = "automobile"
var lo = "loan"
var sub = "subscription"
var util = "utility"
var food = "food"
var mis = "misc"
var heal = "health"
var clo = "clothes"

// var es = "essential"
var nes = "none-essential"
var vari = "variable"
var con = "constant"
var eat = "eatout"
var moto = "motorcycle"
var home = "housing"
var min = "ministry"
var ext = "extracurricular"
var enter = "entertainment"

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
}
