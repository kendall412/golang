package main

func bankTATargets() []Targets {
	var checks Targets
	checks.Name = "checks"
	checks.Cat = []string{cats["ba"]}
	checks.Variant = []string{"check #"}

	var withdrawal Targets
	withdrawal.Name = "withdrawal"
	withdrawal.Cat = []string{cats["ba"]}
	withdrawal.Variant = []string{"ATM WITHDRAWAL", "cash ewidthdrawal"}

	var transfer Targets
	transfer.Name = "transfer to"
	transfer.Cat = []string{cats["ba"]}
	transfer.Variant = []string{"online transfer to"} // "online transfer from",

	var recurring Targets
	recurring.Name = "recurring"
	recurring.Cat = []string{cats["ba"]}
	recurring.Variant = []string{"recurring transfer to"}

	var bankTA = []Targets{}
	bankTA = append(bankTA, checks, transfer, recurring)
	return bankTA
}
