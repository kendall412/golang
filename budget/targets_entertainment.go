package main

func entertainmentTargets() []Targets {
	var theatre Targets
	theatre.Name = "theatre"
	theatre.Cat = []string{cats["ent"]}
	theatre.Variant = []string{"fandango", "fandango.com", "century theatres"}

	var entertainment = []Targets{}
	entertainment = append(entertainment, theatre)
	return entertainment
}
