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

	var fuel = []Targets{}
	fuel = append(fuel, arco, shell)
	return fuel
}
