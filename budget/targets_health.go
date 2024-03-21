package main

func healthTargets() []Targets {
	var kaiser Targets
	kaiser.Name = "kaiser permanente"
	kaiser.Cat = []string{cats["vari"], cats["heal"]}
	kaiser.Variant = []string{"kaiser"}

	var questdiagnostics Targets
	questdiagnostics.Name = "quest diagnostics"
	questdiagnostics.Cat = []string{cats["vari"], cats["heal"]}
	questdiagnostics.Variant = []string{"quest diagnost"}

	var health = []Targets{}
	health = append(health, kaiser, questdiagnostics)
	return health
}
