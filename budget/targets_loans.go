package main

func loansTarget() []Targets {
	var nelnet Targets
	nelnet.Name = "nelnet"
	nelnet.Cat = []string{lo, con}
	nelnet.Variant = []string{"dept education student ln"}

	var palisade Targets
	palisade.Name = "auto loan"
	palisade.Cat = []string{lo, con, auto}
	palisade.Variant = []string{"palisade"}

	var freedomroad Targets
	freedomroad.Name = "freedom road financial"
	freedomroad.Cat = []string{moto, nes}
	freedomroad.Variant = []string{"freedomroad"}

	var auto = []Targets{}
	auto = append(auto, palisade, freedomroad)
	return auto
}
