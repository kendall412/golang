package main

import (
	"fmt"
	"sort"
)

func main() {
	m := genMapOne()
	n := genMapTwo()

	var sorted_kids []string
	for kid := range m {
		sorted_kids = append(sorted_kids, kid)
	}

	fmt.Println("unsorted kid list: ")
	fmt.Println(sorted_kids)
	sort.Strings(sorted_kids)
	fmt.Println("sorted kid list: ")
	fmt.Println(sorted_kids)
	fmt.Println()

	fmt.Println("::: Iterating over map :::")
	for kid, kvalue := range m {
		fmt.Println(kid, kvalue["status"], kvalue["nickname"])
	}

	fmt.Println()

	fmt.Println("::: Iterating over map one using sorted slice :::")
	for _, sorted_kid := range sorted_kids {
		fmt.Println(sorted_kid, m[sorted_kid]["status"], m[sorted_kid]["nickname"])
	}

	fmt.Println()

	fmt.Println("::: Iterating over map two using sorted slice :::")
	for _, sorted_kid := range sorted_kids {
		fmt.Println(sorted_kid, n[sorted_kid]["status"], n[sorted_kid]["nickname"])
	}

	fmt.Println()
	key1 := "wonoo"
	key2 := "johnny"
	v, ok := verifyKey(key1, n)
	x, ok := verifyKey(key2, m)
	fmt.Println(key1, v, ok)
	fmt.Println(key2, x, ok)
}

func verifyKey(key string, n map[string]map[string]string) (map[string]string, bool) {
	/*
	   first value, v, returns values if exists.
	   second value ('ok' by convention) returns bool, true if exist, false if not.
	*/
	v, ok := n[key]
	return v, ok
}

func genMapOne() map[string]map[string]string {
	/*
		simultaneous: map instantiation and creation of data
	*/
	m := map[string]map[string]string{
		"wonoo":  {"status": "son", "nickname": "little_monkey"},
		"naami":  {"status": "daughter", "nickname": "koala"},
		"shinoo": {"status": "son", "nickname": "boog"},
		"woojin": {"status": "son", "nickname": "big_monkey"},
	}
	return m
}

func genMapTwo() map[string]map[string]string {
	/*
		separate: map instantiation and creation of data
	*/
	n := map[string]map[string]string{}

	n["shinoo"] = map[string]string{"status": "son", "nickname": "boog"}
	n["naami"] = map[string]string{"status": "daughter", "nickname": "koala"}
	n["woojin"] = map[string]string{"status": "son", "nickname": "big_monkey"}
	n["wonoo"] = map[string]string{"status": "son", "nickname": "little_monkey"}

	return n
}
