package main

func nonessTargets() []Targets {
	var cyclegear Targets
	cyclegear.Name = "cycle gear"
	cyclegear.Cat = []string{cats["moto"], cats["nes"]}
	cyclegear.Variant = []string{"cycle gear"}

	var a_and_s Targets
	a_and_s.Name = "a_and_s"
	a_and_s.Cat = []string{cats["moto"], cats["nes"]}
	a_and_s.Variant = []string{"a&s"}

	var blackrock Targets
	blackrock.Name = "blac rock auto"
	blackrock.Cat = []string{cats["auto"], cats["nes"]}
	blackrock.Variant = []string{"black rock auto"}

	var pipes_n_cigars Targets
	pipes_n_cigars.Name = "pipes n cigars"
	pipes_n_cigars.Cat = []string{cats["nes"]}
	pipes_n_cigars.Variant = []string{"cigars"}

	var mission_pipe_shop Targets
	mission_pipe_shop.Name = "mission pipe shop"
	mission_pipe_shop.Cat = []string{cats["nes"]}
	mission_pipe_shop.Variant = []string{"mission pipe shop"}

	var amazon Targets
	amazon.Name = "amazon"
	amazon.Cat = []string{cats["nes"], cats["vari"]}
	amazon.Variant = []string{"amzn", "amazon"}

	var ebay Targets
	ebay.Name = "ebay"
	ebay.Cat = []string{cats["nes"], cats["vari"]}
	ebay.Variant = []string{"ebay"}

	var quickquack Targets
	quickquack.Name = "quickquack"
	quickquack.Cat = []string{cats["nes"], cats["con"]}
	quickquack.Variant = []string{"quickquack"}

	var autozone Targets
	autozone.Name = "autozone"
	autozone.Cat = []string{cats["nes"], cats["vari"]}
	autozone.Variant = []string{"autozone"}

	var noness = []Targets{}
	noness = append(noness, cyclegear, a_and_s, blackrock, pipes_n_cigars, mission_pipe_shop, amazon, ebay, quickquack, autozone)
	return noness
}
