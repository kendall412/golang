package main

func extracurricularTargets() []Targets {
	var lees_martial_arts Targets
	lees_martial_arts.Name = "lee's martial arts"
	lees_martial_arts.Cat = []string{con}
	lees_martial_arts.Variant = []string{"lees korean ma"}

	var epicentre Targets
	epicentre.Name = "epicentre"
	epicentre.Cat = []string{con, min}
	epicentre.Variant = []string{"epicentre church"}

	var technique Targets
	technique.Name = "technique gymnastics"
	technique.Cat = []string{con, ext}
	technique.Variant = []string{"technique gymn"}

	var twentyfour_fitness Targets
	twentyfour_fitness.Name = "24 hr fitness"
	twentyfour_fitness.Cat = []string{con}
	twentyfour_fitness.Variant = []string{"24 hour fitness"}

	var extra = []Targets{}
	extra = append(extra, lees_martial_arts, epicentre, technique, twentyfour_fitness)
	return extra
}
