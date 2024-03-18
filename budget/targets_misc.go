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

	var misc = []Targets{}
	misc = append(misc, rc_country, halfpricebooks, hobby_lobby, dollartree)
	return misc
}
