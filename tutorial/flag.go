package main

import (
	"flag"
	f "fmt"
)

func main() {
	// var choice bool
	// debugpl := flag.BoolVar(&choice, "c", false, "debug flag, when selected will used 'testplaylist.json'")
	debugpl := flag.Bool("c", false, "debug flag, when selected will used 'testplaylist.json'")
	flag.Parse()

	if debugpl {
		f.Println("debugpl is true")
	}
	f.Println("debugpl is false")
}
