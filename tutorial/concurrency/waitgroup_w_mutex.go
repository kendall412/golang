package main

import (
	f "fmt"
	"net/http"
	"sync"
)

var signals = []string{}

/*
https://pkg.go.dev/sync@go1.21.3#WaitGroup
 1. wg.Add()
 2. wg.Done()
 3. wg.Wait()

WaitGroup is usually a pointer.
*/
var wg sync.WaitGroup

/*
https://pkg.go.dev/sync@go1.21.3#Mutex
A Mutex is a mutual exclusion lock. It locks and unlocks memory read-write function per goroutines in order to prevent simultaneous access to memory when large amounts of goroutines are running. This is useful when goroutines are required to access or write.
Mutex is usually a pointer.
*/
var mut sync.Mutex

func main() {
	websitelist := []string{
		"http://lco.dev",
		"http://go.dev",
		"http://google.com",
		"http://fb.com",
		"http://github.com",
	}

	for _, web := range websitelist {
		go getStatusCode(web)
		wg.Add(1) // keeps the number of goroutines
	}

	wg.Wait() // does not allow main to finish untill all goroutine is done.
	f.Println(signals)
}

func getStatusCode(endpoint string) {
	defer wg.Done() // passes a signal that says "done".

	res, err := http.Get(endpoint)
	if err != nil {
		f.Println("OOPS in endpoint")
	} else {
		f.Printf("%d status code for %s\n", res.StatusCode, endpoint)
		mut.Lock()
		signals = append(signals, endpoint)
		mut.Unlock()
	}
}
