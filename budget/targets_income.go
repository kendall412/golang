package main

func incomeTargets() []Targets {
	var main Targets
	main.Name = "hyve solutions"
	main.Cat = []string{con, inc}
	main.Variant = []string{"hyve solutions"}

	var solidigm Targets
	solidigm.Name = "solidigm"
	solidigm.Cat = []string{con, inc}
	solidigm.Variant = []string{"sk hynix nand"}

	var income = []Targets{}
	income = append(income, main)
	return income
}
