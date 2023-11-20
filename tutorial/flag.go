package main

import (
	"flag"
	f "fmt"
)

func genFile(debug bool) string {
	if debug {
		return "../go.mod"
	} else {
		return "../defer.go"
	}
}

func main() {
	// var choice bool
	// debugpl := flag.BoolVar(&choice, "c", false, "debug flag, when selected will used 'testplaylist.json'")
	debug := flag.Bool("c", false, "debug flag")
	flag.Parse()

	var fileloc string
	fileloc = genFile(*debug)

	f.Println(*debug)
	f.Println(fileloc)
}
