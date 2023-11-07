package main

import (
	f "fmt"
	"sync"
)

func main() {
	// step 1
	var wg sync.WaitGroup
	// step 2
	wg.Add(3)

	// step 3
	go func() {
		defer wg.Done()
		f.Println("1")
	}()

	go func() {
		defer wg.Done()
		f.Println("2")
	}()

	go func() {
		defer wg.Done()
		f.Println("3")
	}()

	// step 4
	wg.Wait()
}
