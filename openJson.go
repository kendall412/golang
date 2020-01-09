package main

import (
	"fmt"
	"encoding/json"
)


jfile,err := os.Open("test.json")

if err != nil {
	fmt.Println(err)
}

fmt.Println(jfile)