/*
go run --race waitgroup_race_condition.go
*/
package main

import (
	f "fmt"
	"sync"
)

func main() {
	f.Println("Race condition - LearnCodeonline.in")
	wg := &sync.WaitGroup{}
	mut := &sync.Mutex{}

	var score = []int{0}

	/*
		user can use wg.Add(1) per each go routine or give total as wg.Add(total)
	*/
	wg.Add(3)
	go func(wg *sync.WaitGroup, mut *sync.Mutex) {
		f.Println("One R")
		mut.Lock()
		score = append(score, 1)
		mut.Unlock()
		wg.Done()
	}(wg, mut)

	go func(wg *sync.WaitGroup, mut *sync.Mutex) {
		f.Println("Two R")
		mut.Lock()
		score = append(score, 2)
		mut.Unlock()
		wg.Done()
	}(wg, mut)

	go func(wg *sync.WaitGroup, mut *sync.Mutex) {
		f.Println("Three R")
		mut.Lock()
		score = append(score, 3)
		mut.Unlock()
		wg.Done()
	}(wg, mut)

	wg.Wait()
	f.Println(score)
}
