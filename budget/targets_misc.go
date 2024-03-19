package main

func miscTargets() []Targets {
	var rc_country Targets
	rc_country.Name = "rc country hobbies"
	rc_country.Cat = []string{mis, nes}
	rc_country.Variant = []string{"rc country hobbies"}

	var halfpricebooks Targets
	halfpricebooks.Name = "half priced books"
	halfpricebooks.Cat = []string{nes, mis}
	halfpricebooks.Variant = []string{"halfpricebooks"}

	var hobby_lobby Targets
	hobby_lobby.Name = "hobby lobby"
	hobby_lobby.Cat = []string{nes, mis}
	hobby_lobby.Variant = []string{"hobbylobb"}

	var dollartree Targets
	dollartree.Name = "dollar tree"
	dollartree.Cat = []string{mis, nes}
	dollartree.Variant = []string{"dollar tr"}

	var cvs Targets
	cvs.Name = "cvs"
	cvs.Cat = []string{mis, nes}
	cvs.Variant = []string{"cvs/pharm"}

	var jetpens Targets
	jetpens.Name = "jetpens"
	jetpens.Cat = []string{mis, nes}
	jetpens.Variant = []string{"jetpens"}

	var tactileturn Targets
	tactileturn.Name = "tactileturn"
	tactileturn.Cat = []string{mis, nes}
	tactileturn.Variant = []string{"tactile turn"}

	var daiso Targets
	daiso.Name = "daiso"
	daiso.Cat = []string{mis, nes}
	daiso.Variant = []string{"daiso"}

	var misc = []Targets{}
	misc = append(misc, rc_country, halfpricebooks, hobby_lobby, dollartree, cvs, jetpens, tactileturn, daiso)
	return misc
}
