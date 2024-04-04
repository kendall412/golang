package main

func feesTargets() []Targets {
	var dmv Targets
	dmv.Name = "dmv"
	dmv.Cat = []string{cats["fee"]}
	dmv.Variant = []string{"dmv"}

	var fees = []Targets{}
	fees = append(fees, dmv)
	return fees
}
