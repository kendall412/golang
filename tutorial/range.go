package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4}
	kvs := map[string]string{"name": "Danny", "age": "47", "hobby": "guitar"}

	for i := range nums {
		fmt.Println("Index:", i)
	}

	fmt.Println()

	for k, v := range kvs {
		if k == "hobby" {
			if v == "motorcycling" {
				fmt.Println("hobby is ", v)
			} else {
				fmt.Println("could not establish a hobby")
			}
		}

		for k, v := range kvs {
			fmt.Printf("%s -> %s\n", k, v)
		}

		// fmt.Println("kvs:", kvs)

		fmt.Println()
		for k := range kvs {
			fmt.Printf("key: %s\n", k)
		}

	}
}
