package main

func incomeTargets() []Targets {
	var main Targets
	main.Name = "hyve solutions"
	main.Cat = []string{cats["inc"]}
	main.Variant = []string{"hyve solutions"}

	var solidigm Targets
	solidigm.Name = "solidigm"
	solidigm.Cat = []string{cats["inc"]}
	solidigm.Variant = []string{"sk hynix nand"}

	var income = []Targets{}
	income = append(income, main)
	return income
}
