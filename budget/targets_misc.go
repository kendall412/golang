package main

func miscTargets() []Targets {
	var rc_country Targets
	rc_country.Name = "rc country hobbies"
	rc_country.Cat = []string{cats["mis"], cats["nes"]}
	rc_country.Variant = []string{"rc country hobbies"}

	var halfpricebooks Targets
	halfpricebooks.Name = "half priced books"
	halfpricebooks.Cat = []string{cats["mis"], cats["nes"]}
	halfpricebooks.Variant = []string{"halfpricebooks"}

	var hobby_lobby Targets
	hobby_lobby.Name = "hobby lobby"
	hobby_lobby.Cat = []string{cats["mis"], cats["nes"]}
	hobby_lobby.Variant = []string{"hobbylobb"}

	var dollartree Targets
	dollartree.Name = "dollar tree"
	dollartree.Cat = []string{cats["mis"], cats["nes"]}
	dollartree.Variant = []string{"dollar tr"}

	var cvs Targets
	cvs.Name = "cvs"
	cvs.Cat = []string{cats["mis"], cats["nes"]}
	cvs.Variant = []string{"cvs/pharm"}

	var jetpens Targets
	jetpens.Name = "jetpens"
	jetpens.Cat = []string{cats["mis"], cats["nes"]}
	jetpens.Variant = []string{"jetpens"}

	var tactileturn Targets
	tactileturn.Name = "tactileturn"
	tactileturn.Cat = []string{cats["mis"], cats["nes"]}
	tactileturn.Variant = []string{"tactile turn"}

	var daiso Targets
	daiso.Name = "daiso"
	daiso.Cat = []string{cats["mis"], cats["nes"]}
	daiso.Variant = []string{"daiso"}

	var udemy Targets
	udemy.Name = "udemy"
	udemy.Cat = []string{cats["mis"], cats["nes"]}
	udemy.Variant = []string{"udemy"}

	var countess_alteration Targets
	countess_alteration.Name = "countess alteration"
	countess_alteration.Cat = []string{cats["mis"], cats["nes"]}
	countess_alteration.Variant = []string{"countess alteratio roseville"}

	var humble_bundle Targets
	humble_bundle.Name = "humble bundle"
	humble_bundle.Cat = []string{cats["mis"], cats["nes"]}
	humble_bundle.Variant = []string{"humblebundle.com"}

	var temu Targets
	temu.Name = "temu"
	temu.Cat = []string{cats["misc"], cats["nes"]}
	temu.Variant = []string{"temu.com"}

	var misc = []Targets{}
	misc = append(misc, rc_country, halfpricebooks, hobby_lobby, dollartree, cvs, jetpens, tactileturn, daiso, udemy, countess_alteration, humble_bundle, temu)
	return misc
}
