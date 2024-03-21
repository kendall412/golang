package main

func loansTarget() []Targets {
	var nelnet Targets
	nelnet.Name = "nelnet"
	nelnet.Cat = []string{cats["lo"], cats["con"]}
	nelnet.Variant = []string{"dept education student ln"}

	var palisade Targets
	palisade.Name = "auto loan"
	palisade.Cat = []string{cats["lo"], cats["con"], cats["auto"]}
	palisade.Variant = []string{"palisade"}

	var freedomroad Targets
	freedomroad.Name = "freedom road financial"
	freedomroad.Cat = []string{cats["lo"], cats["con"], cats["moto"]}
	freedomroad.Variant = []string{"freedomroad"}

	var loan = []Targets{}
	loan = append(loan, nelnet, palisade, freedomroad)
	return loan
}
