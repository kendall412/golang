package main

func taxTargets() []Targets {
	// this value in 2023 was deposit and not withdrawl
	var franchisetaxboard Targets
	franchisetaxboard.Name = "franchise tax board"
	franchisetaxboard.Cat = []string{cats["tax"]}
	franchisetaxboard.Variant = []string{"franchise tax"}
	franchisetaxboard.Desc = "state tax"

	var irs Targets
	irs.Name = "irs"
	irs.Cat = []string{cats["tax"]}
	irs.Variant = []string{"irs"}
	irs.Desc = "federal tax"

	var turbotax Targets
	turbotax.Name = "turbo tax"
	turbotax.Cat = []string{cats["tax"]}
	turbotax.Variant = []string{"turbotax"}
	turbotax.Desc = "Intuit Turbo Tax svc. Cost of filing tax electronically"

	var tax = []Targets{}
	tax = append(tax, franchisetaxboard, irs, turbotax)
	return tax
}
