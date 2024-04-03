package main

func insuranceTargets() []Targets {
	var nationwide Targets
	nationwide.Name = "nationwide"
	nationwide.Cat = []string{cats["con"], cats["ins"], cats["auto"], cats["cins"]}
	nationwide.Variant = []string{"nationwide"}
	nationwide.Desc = "auto insurance for Palisade and Elantra"

	var geico Targets
	geico.Name = "geico"
	geico.Cat = []string{cats["con"], cats["ins"], cats["moto"], cats["mins"]}
	geico.Variant = []string{"geico"}
	geico.Desc = "motorcycle insurance for Street Triple, ZX6R & T120"

	var insurance = []Targets{}
	insurance = append(insurance, nationwide, geico)
	return insurance
}
