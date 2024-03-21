package main

func fuelTargets() []Targets {

	var arco Targets
	arco.Name = "arco"
	arco.Cat = []string{cats["vari"], cats["fu"]}
	arco.Variant = []string{"arco"}

	var shell Targets
	shell.Name = "shell"
	shell.Cat = []string{cats["vari"], cats["fu"]}
	shell.Variant = []string{"shell"}

	var chevron Targets
	chevron.Name = "chevron"
	chevron.Cat = []string{cats["vari"], cats["fu"]}
	chevron.Variant = []string{"chevron"}

	var fuel = []Targets{}
	fuel = append(fuel, arco, shell, chevron)
	return fuel
}
