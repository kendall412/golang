package main

func eatOutTargets() []Targets {
	var innout Targets
	innout.Name = "in-n-out"
	innout.Cat = []string{cats["vari"], cats["eat"]}
	innout.Variant = []string{"in-n-out", "in n out"}

	var mcdonalds Targets
	mcdonalds.Name = "mcdonalds"
	mcdonalds.Cat = []string{cats["vari"], cats["eat"]}
	mcdonalds.Variant = []string{"mcdonald", "mcdonalds"}

	var ramen Targets
	ramen.Name = "ramen"
	ramen.Cat = []string{cats["vari"], cats["eat"]}
	ramen.Variant = []string{"ryujin ramen"}

	var subway Targets
	subway.Name = "subway"
	subway.Cat = []string{cats["vari"], cats["eat"]}
	subway.Variant = []string{"subway"}

	var brookfields Targets
	brookfields.Name = "brookfields"
	brookfields.Cat = []string{cats["vari"], cats["eat"]}
	brookfields.Variant = []string{"brookfield"}

	var pho Targets
	pho.Name = "pho"
	pho.Cat = []string{cats["vari"], cats["eat"]}
	pho.Variant = []string{"pho"}

	var chipotle Targets
	chipotle.Name = "chipotle"
	chipotle.Cat = []string{cats["vari"], cats["eat"]}
	chipotle.Variant = []string{"chipotle"}

	var seoulzip Targets
	seoulzip.Name = "seoulzip"
	seoulzip.Cat = []string{cats["vari"], cats["eat"]}
	seoulzip.Variant = []string{"seoulzip"}

	var tasty_pot Targets
	tasty_pot.Name = "tasty pot"
	tasty_pot.Cat = []string{cats["vari"], cats["eat"]}
	tasty_pot.Variant = []string{"tasty pot"}

	var vons_chicken Targets
	vons_chicken.Name = "von's chicken"
	vons_chicken.Cat = []string{cats["vari"], cats["eat"]}
	vons_chicken.Variant = []string{"vons chicken"}

	var paris_baguett Targets
	paris_baguett.Name = "paris baguett"
	paris_baguett.Cat = []string{cats["vari"], cats["eat"]}
	paris_baguett.Variant = []string{"paris baguett"}

	var pushkin Targets
	pushkin.Name = "pushkin"
	pushkin.Cat = []string{cats["vari"], cats["eat"]}
	pushkin.Variant = []string{"pushkin"}

	var osakaya Targets
	osakaya.Name = "osaka-ya"
	osakaya.Cat = []string{cats["vari"], cats["eat"]}
	osakaya.Variant = []string{"osaka- ya"}

	var asianpearl Targets
	asianpearl.Name = "asianpearl"
	asianpearl.Cat = []string{cats["vari"], cats["eat"]}
	asianpearl.Variant = []string{"new asian pearl"}

	var quicklys Targets
	quicklys.Name = "quicklys"
	quicklys.Cat = []string{cats["vari"], cats["eat"]}
	quicklys.Variant = []string{"quickly"}

	var leatherbys Targets
	leatherbys.Name = "leatherby's"
	leatherbys.Cat = []string{cats["vari"], cats["eat"]}
	leatherbys.Variant = []string{"leatherby's"}

	var mikuni Targets
	mikuni.Name = "mikuni"
	mikuni.Cat = []string{cats["vari"], cats["eat"]}
	mikuni.Variant = []string{"mikuni japaneser re"}

	var gangnam Targets
	gangnam.Name = "gangnam"
	gangnam.Cat = []string{cats["eat"], cats["vari"]}
	gangnam.Variant = []string{"gangnam ave"}

	var heatshabu Targets
	heatshabu.Name = "heat shabu"
	heatshabu.Cat = []string{cats["eat"], cats["vari"]}
	heatshabu.Variant = []string{"heat shabu"}

	var rubios Targets
	rubios.Name = "rubios"
	rubios.Cat = []string{cats["eat"], cats["vari"]}
	rubios.Variant = []string{"rubios", "rubio's"}

	var teddyboba Targets
	teddyboba.Name = "teddy boba"
	teddyboba.Cat = []string{cats["eat"], cats["vari"]}
	teddyboba.Variant = []string{"teddy boba"}

	var bennetts Targets
	bennetts.Name = "bennetts"
	bennetts.Cat = []string{cats["eat"], cats["vari"]}
	bennetts.Variant = []string{"bennett's american"}

	var somisomi Targets
	somisomi.Name = "somisomi"
	somisomi.Cat = []string{cats["eat"], cats["vari"]}
	somisomi.Variant = []string{"somisomi"}

	var eatout = []Targets{}
	eatout = append(eatout, innout, mcdonalds, ramen, subway, brookfields, pho, chipotle, seoulzip, tasty_pot, vons_chicken, paris_baguett, pushkin, osakaya, asianpearl, leatherbys, mikuni, gangnam, heatshabu, quicklys, rubios, teddyboba, bennetts, somisomi)

	return eatout
}
