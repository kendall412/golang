package main

type Targets struct {
	Name    string
	Cat     []string // eating_out, grocery, etc
	Variant []string // {"in-n-out","in n out"}
	DESC    string
}

// spending categories
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

// func motoTargets() []Targets {
// 	var cyclegear Targets
// 	cyclegear.Name = "cycle gear"
// 	cyclegear.Cat = []string{moto, vari}
// 	cyclegear.Variant = []string{"cycle gear"}

// 	var a_and_s Targets
// 	a_and_s.Name = "a_and_s"
// 	a_and_s.Cat = []string{moto, vari}
// 	a_and_s.Variant = []string{"a&s"}

// 	var moto = []Targets{}
// 	moto = append(moto, cyclegear, a_and_s)

// 	return moto
// }

// func clothesTarget() []Targets {
// 	var shein Targets
// 	shein.Name = "shein"
// 	shein.Cat = []string{vari, clo}
// 	shein.Variant = []string{"shein"}

// 	var marshall Targets
// 	marshall.Name = "marshall"
// 	marshall.Cat = []string{"clothes", "variable"}
// 	marshall.Variant = []string{"marshall"}

// 	var ross Targets
// 	ross.Name = "ross"
// 	ross.Cat = []string{clo, vari}
// 	ross.Variant = []string{"ross"}

// 	var macys Targets
// 	macys.Name = "macys"
// 	macys.Cat = []string{clo, vari}
// 	macys.Variant = []string{"macy"}

// 	var abercrombie Targets
// 	abercrombie.Name = "abercrombie"
// 	abercrombie.Cat = []string{clo, vari}
// 	abercrombie.Variant = []string{"abercrombie"}

// 	var arden Targets
// 	arden.Name = "arden"
// 	arden.Cat = []string{clo, vari}
// 	arden.Variant = []string{"arden fair"}

// 	var clothes = []Targets{}
// 	clothes = append(clothes, shein, marshall, ross, macys, abercrombie, arden)
// 	return clothes
// }

// func housingTargets() []Targets {
// 	var housing Targets
// 	housing.Name = "housing"
// 	housing.Cat = []string{con, home}
// 	housing.Variant = []string{"cheung"}

// 	var home = []Targets{}
// 	home = append(home, housing)
// 	return home
// }

func generateTargetStruct(targets_slice *[]Targets) {
	*targets_slice = append(*targets_slice, extracurricularTargets()...)
	*targets_slice = append(*targets_slice, housingTargets()...)
	*targets_slice = append(*targets_slice, subscriptionTarget()...)
	*targets_slice = append(*targets_slice, clothesTarget()...)
	*targets_slice = append(*targets_slice, nonessTargets()...)
	*targets_slice = append(*targets_slice, eatOutTargets()...)
	*targets_slice = append(*targets_slice, groceryTargets()...)
	*targets_slice = append(*targets_slice, utilsTargets()...)
	*targets_slice = append(*targets_slice, subscriptionTarget()...)
	*targets_slice = append(*targets_slice, ministryTargets()...)
	*targets_slice = append(*targets_slice, loansTarget()...)
	*targets_slice = append(*targets_slice, coffeeTargets()...)
	*targets_slice = append(*targets_slice, insuranceTargets()...)
	*targets_slice = append(*targets_slice, miscTargets()...)
	*targets_slice = append(*targets_slice, fuelTargets()...)
	*targets_slice = append(*targets_slice, healthTargets()...)
}
