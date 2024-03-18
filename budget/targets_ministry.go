package main

func ministryTargets() []Targets {
	var epicentre Targets
	epicentre.Name = "epicentre church"
	epicentre.Cat = []string{con, min}
	epicentre.Variant = []string{"epicentre church"}

	var ministry = []Targets{}
	ministry = append(ministry, epicentre)
	return ministry
}
