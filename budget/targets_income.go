package main

func incomeTargets() []Targets {
	var hyve Targets
	hyve.Name = "hyve solutions"
	hyve.Cat = []string{cats["inc"]}
	hyve.Variant = []string{"hyve solutions"}
	hyve.Spending = false

	var solidigm Targets
	solidigm.Name = "solidigm"
	solidigm.Cat = []string{cats["inc"]}
	solidigm.Variant = []string{"sk hynix nand"}
	solidigm.Spending = false

	var income = []Targets{}
	income = append(income, hyve)
	return income
}
