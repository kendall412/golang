package main

type Targets struct {
	Name    string
	Cat     []string // eating_out, grocery, etc
	Variant []string // {"in-n-out","in n out"}
}

// spending categories
var food = "food"
var mis = "misc"
var heal = "health"
var clo = "clothes"
var es = "essential"
var vari = "variable"
var con = "constant"
var eat = "eatout"
var moto = "motorcycle"
var home = "housing"
var min = "ministry"
var ext = "extracurricular"
var enter = "entertainment"

func motoTargets() []Targets {
	var cyclegear Targets
	cyclegear.Name = "cycle gear"
	cyclegear.Cat = []string{moto, vari}
	cyclegear.Variant = []string{"cycle gear"}

	var a_and_s Targets
	a_and_s.Name = "a_and_s"
	a_and_s.Cat = []string{moto, vari}
	a_and_s.Variant = []string{"a&s"}

	var moto = []Targets{}
	moto = append(moto, cyclegear, a_and_s)

	return moto
}

func clothesTarget() []Targets {
	var shein Targets
	shein.Name = "shein"
	shein.Cat = []string{vari, clo}
	shein.Variant = []string{"shein"}

	var marshall Targets
	marshall.Name = "marshall"
	marshall.Cat = []string{"clothes", "variable"}
	marshall.Variant = []string{"marshall"}

	var ross Targets
	ross.Name = "ross"
	ross.Cat = []string{clo, vari}
	ross.Variant = []string{"ross"}

	var macys Targets
	macys.Name = "macys"
	macys.Cat = []string{clo, vari}
	macys.Variant = []string{"macy"}

	var abercrombie Targets
	abercrombie.Name = "abercrombie"
	abercrombie.Cat = []string{clo, vari}
	abercrombie.Variant = []string{"abercrombie"}

	var arden Targets
	arden.Name = "arden"
	arden.Cat = []string{clo, vari}
	arden.Variant = []string{"arden fair"}

	var clothes = []Targets{}
	clothes = append(clothes, shein, marshall, ross, macys, abercrombie, arden)
	return clothes
}

func subscriptionTarget() []Targets {
	var netflix Targets
	netflix.Name = "netflix"
	netflix.Cat = []string{con, enter}
	netflix.Variant = []string{"netflix"}

	var audible Targets
	audible.Name = "audible"
	audible.Cat = []string{con, enter}
	audible.Variant = []string{"audible"}

	var apple Targets
	apple.Name = "apple"
	apple.Cat = []string{con, enter}
	apple.Variant = []string{"apple"}

	var github Targets
	github.Name = "github"
	github.Cat = []string{con}
	github.Variant = []string{"github"}

	var microsoft Targets
	microsoft.Name = "microsoft"
	microsoft.Cat = []string{con}
	microsoft.Variant = []string{"microsoft"}

	var sub = []Targets{}
	sub = append(sub, netflix, audible, apple, github, microsoft)
	return sub
}

func extracurricularTargets() []Targets {
	var lees_martial_arts Targets
	lees_martial_arts.Name = "lee's martial arts"
	lees_martial_arts.Cat = []string{con}
	lees_martial_arts.Variant = []string{"lees korean ma"}

	var epicentre Targets
	epicentre.Name = "epicentre"
	epicentre.Cat = []string{con, min}
	epicentre.Variant = []string{"epicentre church"}

	var technique Targets
	technique.Name = "technique gymnastics"
	technique.Cat = []string{con, ext}
	technique.Variant = []string{"technique gymn"}

	var twentyfour_fitness Targets
	twentyfour_fitness.Name = "24 hr fitness"
	twentyfour_fitness.Cat = []string{con}
	twentyfour_fitness.Variant = []string{"24 hour fitness"}

	var extra = []Targets{}
	extra = append(extra, lees_martial_arts, epicentre, technique, twentyfour_fitness)
	return extra
}

func housingTargets() []Targets {
	var housing Targets
	housing.Name = "housing"
	housing.Cat = []string{con, home}
	housing.Variant = []string{"cheung"}

	var home = []Targets{}
	home = append(home, housing)
	return home
}

func generateTargetStruct(targets_slice *[]Targets) {
	*targets_slice = append(*targets_slice, extracurricularTargets()...)
	*targets_slice = append(*targets_slice, housingTargets()...)
	*targets_slice = append(*targets_slice, subscriptionTarget()...)
	*targets_slice = append(*targets_slice, clothesTarget()...)
	*targets_slice = append(*targets_slice, motoTargets()...)
	*targets_slice = append(*targets_slice, eatOutTargets()...)
	*targets_slice = append(*targets_slice, groceryTargets()...)
}
