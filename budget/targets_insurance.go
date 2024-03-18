package main

func insuranceTargets() []Targets {
	var nationwide Targets
	nationwide.Name = "nationwide"
	nationwide.Cat = []string{ins, con, auto}
	nationwide.Variant = []string{"nationwide"}

	var geico Targets
	geico.Name = "geico"
	geico.Cat = []string{ins, moto, con}
	geico.Variant = []string{"geico"}

	var insurance = []Targets{}
	insurance = append(insurance, nationwide, geico)
	return insurance
}
