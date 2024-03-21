package main

func groceryTargets() []Targets {
	var butcherbox Targets
	butcherbox.Name = "butcherbox"
	butcherbox.Cat = []string{cats["con"], cats["food"]}
	butcherbox.Variant = []string{"butcherbox"}

	var smile Targets
	smile.Name = "smile"
	smile.Cat = []string{cats["con"], cats["food"]}
	smile.Variant = []string{"smile"}

	var safeway Targets
	safeway.Name = "safeway"
	safeway.Cat = []string{cats["con"], cats["food"]}
	safeway.Variant = []string{"safeway"}

	var sprouts Targets
	sprouts.Name = "sprouts"
	sprouts.Cat = []string{cats["con"], cats["food"]}
	sprouts.Variant = []string{"sprout"}

	var costco Targets
	costco.Name = "costco"
	costco.Cat = []string{cats["vari"], cats["food"]}
	costco.Variant = []string{"costco"}

	var kp_intl Targets
	kp_intl.Name = "kp international"
	kp_intl.Cat = []string{cats["vari"], cats["food"]}
	kp_intl.Variant = []string{"kp international"}

	var walmart Targets
	walmart.Name = "walmart"
	walmart.Cat = []string{cats["vari"]}
	walmart.Variant = []string{"wal mart", "wal-mart", "walmart"}

	var raley Targets
	raley.Name = "raley"
	raley.Cat = []string{cats["vari"], cats["food"]}
	raley.Variant = []string{"raley"}

	var grocery_outlet Targets
	grocery_outlet.Name = "grocery outlet"
	grocery_outlet.Cat = []string{cats["vari"], cats["food"]}
	grocery_outlet.Variant = []string{"groc outlet"}

	var lucky Targets
	lucky.Name = "luckys"
	lucky.Cat = []string{cats["vari"], cats["food"]}
	lucky.Variant = []string{"lucky"}

	var hankook Targets
	hankook.Name = "hankook market"
	hankook.Cat = []string{cats["vari"], cats["food"]}
	hankook.Variant = []string{"hankook"}

	var walgreens Targets
	walgreens.Name = "walgreen"
	walgreens.Cat = []string{cats["vari"]}
	walgreens.Variant = []string{"walgreen"}

	var bathnbody Targets
	bathnbody.Name = "bath&body"
	bathnbody.Cat = []string{cats["vari"], cats["nes"]}
	bathnbody.Variant = []string{"bath & body works"}

	var target Targets
	target.Name = "target"
	target.Cat = []string{cats["vari"], cats["nes"]}
	target.Variant = []string{"target"}

	var grocery = []Targets{}
	grocery = append(grocery, butcherbox, smile, safeway, sprouts, costco, kp_intl, walmart, raley, grocery_outlet, lucky, hankook, walgreens, bathnbody, target)

	return grocery
}
