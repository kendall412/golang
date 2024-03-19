package main

func groceryTargets() []Targets {
	var butcherbox Targets
	butcherbox.Name = "butcherbox"
	butcherbox.Cat = []string{con, food}
	butcherbox.Variant = []string{"butcherbox"}

	var smile Targets
	smile.Name = "smile"
	smile.Cat = []string{vari, food}
	smile.Variant = []string{"smile"}

	var safeway Targets
	safeway.Name = "safeway"
	safeway.Cat = []string{vari, food}
	safeway.Variant = []string{"safeway"}

	var sprouts Targets
	sprouts.Name = "sprouts"
	sprouts.Cat = []string{vari, food}
	sprouts.Variant = []string{"sprout"}

	var costco Targets
	costco.Name = "costco"
	costco.Cat = []string{vari}
	costco.Variant = []string{"costco"}

	var kp_intl Targets
	kp_intl.Name = "kp international"
	kp_intl.Cat = []string{vari, food}
	kp_intl.Variant = []string{"kp international"}

	var walmart Targets
	walmart.Name = "walmart"
	walmart.Cat = []string{vari}
	walmart.Variant = []string{"wal mart", "wal-mart", "walmart"}

	var raley Targets
	raley.Name = "raley"
	raley.Cat = []string{vari, food}
	raley.Variant = []string{"raley"}

	var grocery_outlet Targets
	grocery_outlet.Name = "grocery outlet"
	grocery_outlet.Cat = []string{vari, food}
	grocery_outlet.Variant = []string{"groc outlet"}

	var lucky Targets
	lucky.Name = "luckys"
	lucky.Cat = []string{vari, food}
	lucky.Variant = []string{"lucky"}

	var hankook Targets
	hankook.Name = "hankook market"
	hankook.Cat = []string{vari, food}
	hankook.Variant = []string{"hankook"}

	var walgreens Targets
	walgreens.Name = "walgreen"
	walgreens.Cat = []string{vari}
	walgreens.Variant = []string{"walgreen"}

	var bathnbody Targets
	bathnbody.Name = "bath&body"
	bathnbody.Cat = []string{nes, vari}
	bathnbody.Variant = []string{"bath & body works"}

	var target Targets
	target.Name = "target"
	target.Cat = []string{nes, vari}
	target.Variant = []string{"target"}

	var grocery = []Targets{}
	grocery = append(grocery, butcherbox, smile, safeway, sprouts, costco, kp_intl, walmart, raley, grocery_outlet, lucky, hankook, walgreens, bathnbody, target)

	return grocery
}
