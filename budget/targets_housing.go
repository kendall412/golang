package main

func housingTargets() []Targets {
	var housing Targets
	housing.Name = "housing"
	housing.Cat = []string{con, home}
	housing.Variant = []string{"cheung"}

	var home = []Targets{}
	home = append(home, housing)
	return home
}
