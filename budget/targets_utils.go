package main

func utilsTargets() []Targets {
	var pg_e Targets
	pg_e.Name = "pg&e"
	pg_e.Cat = []string{util, vari}
	pg_e.Variant = []string{"pg&e"}

	var smud Targets
	smud.Name = "smud"
	smud.Cat = []string{util, vari}
	smud.Variant = []string{"smud"}

	var comcast Targets
	comcast.Name = "xfinity"
	comcast.Cat = []string{util, con}
	comcast.Variant = []string{"comcast"}

	var verizon Targets
	verizon.Name = "verizon"
	verizon.Cat = []string{util, con}
	verizon.Variant = []string{"verizon"}

	var utils = []Targets{}
	utils = append(utils, pg_e, smud, comcast, verizon)
	return utils
}
