package main

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
