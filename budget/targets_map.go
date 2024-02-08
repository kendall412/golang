package main

import "log"

func generateMap(debug, display_all *bool) map[string][]string {
	target_map := map[string][]string{}

	target_map["housing"] = housing
	target_map["clothing"] = clothing
	target_map["utils"] = utils
	target_map["gas"] = gas
	target_map["grocery"] = grocery
	target_map["health"] = health
	target_map["food_out"] = food_out
	target_map["misc"] = misc
	target_map["coffee"] = coffee
	target_map["insurance"] = insurance
	target_map["auto"] = auto
	target_map["motorcycle"] = motorcycle
	target_map["income"] = income
	target_map["monthly"] = monthly

	if *display_all {
		log.Println(target_map)
	}

	return target_map
}
