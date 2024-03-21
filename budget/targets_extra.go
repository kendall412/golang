package main

func extracurricularTargets() []Targets {
	var lees_martial_arts Targets
	lees_martial_arts.Name = "lee's martial arts"
	lees_martial_arts.Cat = []string{cats["con"]}
	lees_martial_arts.Variant = []string{"lees korean ma"}

	var technique Targets
	technique.Name = "technique gymnastics"
	technique.Cat = []string{cats["con"], cats["ext"]}
	technique.Variant = []string{"technique gymn"}

	var twentyfour_fitness Targets
	twentyfour_fitness.Name = "24 hr fitness"
	twentyfour_fitness.Cat = []string{cats["con"]}
	twentyfour_fitness.Variant = []string{"24 hour fitness"}

	var extra = []Targets{}
	extra = append(extra, lees_martial_arts, technique, twentyfour_fitness)
	return extra
}
