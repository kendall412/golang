package main

func ministryTargets() []Targets {
	var epicentre Targets
	epicentre.Name = "epicentre"
	epicentre.Cat = []string{cats["con"], cats["min"]}
	epicentre.Variant = []string{"epicentre church"}
	epicentre.Desc = "ministry support for Clara"

	var ministry = []Targets{}
	ministry = append(ministry, epicentre)
	return ministry
}
