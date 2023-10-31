package main

import (
	f "fmt"
	"sync"
)

/*
https://www.youtube.com/watch?v=srb6fbioEY4&list=PLsc-VaxfZl4do3Etp_xQ0aQBoC-x5BIgJ&index=2
1. create a waitgroup
2. call method add to indicate how many ops or goroutines you want to wait for
3. run goroutines and inside the goroutines call the method "done"
4. call the method "wait"
5. synchronize the "add" call with "wait" call
*/

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
