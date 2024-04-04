package main

func clothesTarget() []Targets {
	var shein Targets
	shein.Name = "shein"
	shein.Cat = []string{cats["vari"], cats["clo"]}
	shein.Variant = []string{"shein"}

	var marshall Targets
	marshall.Name = "marshall"
	marshall.Cat = []string{cats["clo"], cats["vari"]}
	marshall.Variant = []string{"marshall"}

	var ross Targets
	ross.Name = "ross"
	ross.Cat = []string{cats["clo"], cats["vari"]}
	ross.Variant = []string{"ross"}

	var macys Targets
	macys.Name = "macys"
	macys.Cat = []string{cats["clo"], cats["vari"]}
	macys.Variant = []string{"macy"}

	var abercrombie Targets
	abercrombie.Name = "abercrombie"
	abercrombie.Cat = []string{cats["clo"], cats["vari"]}
	abercrombie.Variant = []string{"abercrombie"}

	var arden Targets
	arden.Name = "arden"
	arden.Cat = []string{cats["clo"], cats["vari"]}
	arden.Variant = []string{"arden fair"}

	var nordstrom Targets
	nordstrom.Name = "nordstrom"
	nordstrom.Cat = []string{cats["clo"], cats["vari"]}
	nordstrom.Variant = []string{"nordstrom"}

	var onceuponachild Targets
	onceuponachild.Name = "once upon a child"
	onceuponachild.Cat = []string{cats["nes"]}
	onceuponachild.Variant = []string{"once upon a child"}

	var clothes = []Targets{}
	clothes = append(clothes, shein, marshall, ross, macys, abercrombie, arden, nordstrom, onceuponachild)
	return clothes
}
