package main

func healthTargets() []Targets {

	var kaiser Targets
	kaiser.Name = "kaiser permanente"
	kaiser.Cat = []string{heal, vari}
	kaiser.Variant = []string{"kaiser"}

	var health = []Targets{}
	health = append(health, kaiser)
	return health
}
