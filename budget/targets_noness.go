package main

func nonessTargets() []Targets {
	var cyclegear Targets
	cyclegear.Name = "cycle gear"
	cyclegear.Cat = []string{moto, nes}
	cyclegear.Variant = []string{"cycle gear"}

	var a_and_s Targets
	a_and_s.Name = "a_and_s"
	a_and_s.Cat = []string{moto, nes}
	a_and_s.Variant = []string{"a&s"}

	var blackrock Targets
	blackrock.Name = "blac rock auto"
	blackrock.Cat = []string{nes, auto}
	blackrock.Variant = []string{"black rock auto"}

	var pipes_n_cigars Targets
	pipes_n_cigars.Name = "pipes n cigars"
	pipes_n_cigars.Cat = []string{nes}
	pipes_n_cigars.Variant = []string{"cigars"}

	var mission_pipe_shop Targets
	mission_pipe_shop.Name = "mission pipe shop"
	mission_pipe_shop.Cat = []string{nes}
	mission_pipe_shop.Variant = []string{"mission pipe shop"}

	var amazon Targets
	amazon.Name = "amazon"
	amazon.Cat = []string{nes, vari}
	amazon.Cat = []string{"amzn", "amazon"}

	var ebay Targets
	ebay.Name = "ebay"
	ebay.Cat = []string{nes, vari}
	ebay.Cat = []string{"ebay"}

	var quickquack Targets
	quickquack.Name = "quickquack"
	quickquack.Cat = []string{con, nes}
	quickquack.Variant = []string{"quickquack"}

	var noness = []Targets{}
	noness = append(noness, cyclegear, a_and_s, blackrock, pipes_n_cigars, mission_pipe_shop, amazon, ebay)

	return noness
}
