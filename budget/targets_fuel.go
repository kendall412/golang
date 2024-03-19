package main

func fuelTargets() []Targets {

	var arco Targets
	arco.Name = "arco"
	arco.Cat = []string{fu, vari}
	arco.Variant = []string{"arco"}

	var shell Targets
	shell.Name = "shell"
	shell.Cat = []string{fu, vari}
	shell.Variant = []string{"shell"}

	var chevron Targets
	chevron.Name = "chevron"
	chevron.Cat = []string{fu, vari}
	chevron.Variant = []string{"chevron"}

	var fuel = []Targets{}
	fuel = append(fuel, arco, shell, chevron)
	return fuel
}
