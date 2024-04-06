package main

func bankTATargets() []Targets {
	var checks Targets
	checks.Name = "checks"
	checks.Cat = []string{cats["ba"]}
	checks.Variant = []string{"check #"}

	var withdrawal Targets
	withdrawal.Name = "withdrawal"
	withdrawal.Cat = []string{cats["ba"]}
	withdrawal.Variant = []string{"atm withdrawal", "cash ewithdrawal"}

	var transfer_to Targets
	transfer_to.Name = "transfer to"
	transfer_to.Cat = []string{cats["ba"]}
	transfer_to.Variant = []string{"online transfer to"}
	transfer_to.Desc = "transfer to another account from Checking"

	var transfer_from Targets
	transfer_from.Name = "transfer to"
	transfer_from.Cat = []string{cats["ba"]}
	transfer_from.Variant = []string{"online transfer to"}
	transfer_from.Spending = true
	transfer_from.Desc = "transfer from another account to Checking, a deposit."

	var recurring Targets
	recurring.Name = "recurring"
	recurring.Cat = []string{cats["ba"]}
	recurring.Variant = []string{"recurring transfer to"}

	var overdraft Targets
	overdraft.Name = "overdraft"
	overdraft.Cat = []string{cats["ba"]}
	overdraft.Variant = []string{"overdraft protection"}

	var bankTA = []Targets{}
	bankTA = append(bankTA, checks, withdrawal, transfer_to, transfer_from, recurring, overdraft)
	return bankTA
}
