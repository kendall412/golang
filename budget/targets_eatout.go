package main

func eatOutTargets() []Targets {
	var innout Targets
	innout.Name = "in-n-out"
	innout.Cat = []string{vari, eat}
	innout.Variant = []string{"in-n-out", "in n out"}

	var mcdonalds Targets
	mcdonalds.Name = "mcdonalds"
	mcdonalds.Cat = []string{vari, eat}
	mcdonalds.Variant = []string{"mcdonald", "mcdonalds"}

	var ramen Targets
	ramen.Name = "ramen"
	ramen.Cat = []string{vari, eat}
	ramen.Variant = []string{"ryujin ramen"}

	var subway Targets
	subway.Name = "subway"
	subway.Cat = []string{vari, eat}
	subway.Variant = []string{"subway"}

	var brookfields Targets
	brookfields.Name = "brookfields"
	brookfields.Cat = []string{vari, eat}
	brookfields.Variant = []string{"brookfield"}

	var pho Targets
	pho.Name = "pho"
	pho.Cat = []string{vari, eat}
	pho.Variant = []string{"pho"}

	var chipotle Targets
	chipotle.Name = "chipotle"
	chipotle.Cat = []string{vari, eat}
	chipotle.Variant = []string{"chipotle"}

	var seoulzip Targets
	seoulzip.Name = "seoulzip"
	seoulzip.Cat = []string{vari, eat}
	seoulzip.Variant = []string{"seoulzip"}

	var tasty_pot Targets
	tasty_pot.Name = "tasty pot"
	tasty_pot.Cat = []string{vari, eat}
	tasty_pot.Variant = []string{"tasty pot"}

	var vons_chicken Targets
	vons_chicken.Name = "von's chicken"
	vons_chicken.Cat = []string{vari, eat}
	vons_chicken.Variant = []string{"vons chicken"}

	var paris_baguett Targets
	paris_baguett.Name = "paris baguett"
	paris_baguett.Cat = []string{vari, eat}
	paris_baguett.Variant = []string{"paris baguett"}

	var pushkin Targets
	pushkin.Name = "pushkin"
	pushkin.Cat = []string{vari, eat}
	pushkin.Variant = []string{"pushkin"}

	var osakaya Targets
	osakaya.Name = "osaka-ya"
	osakaya.Cat = []string{vari, eat}
	osakaya.Variant = []string{"osaka- ya"}

	var asianpearl Targets
	asianpearl.Name = "asianpearl"
	asianpearl.Cat = []string{vari, eat}
	asianpearl.Variant = []string{"new asian pearl"}

	var leatherbys Targets
	leatherbys.Name = "leatherby's"
	leatherbys.Cat = []string{vari, eat}
	leatherbys.Variant = []string{"leatherby's"}

	var mikuni Targets
	mikuni.Name = "mikuni"
	mikuni.Cat = []string{vari, eat}
	mikuni.Variant = []string{"mikuni japaneser re"}

	var eatout = []Targets{}
	eatout = append(eatout, innout, mcdonalds, ramen, subway, brookfields, pho, chipotle, seoulzip, tasty_pot, vons_chicken, paris_baguett, pushkin, osakaya, asianpearl, leatherbys, mikuni)

	return eatout
}
