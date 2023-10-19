package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func openJsonMap(urlJson string) map[string]interface{} {
	var path map[string]interface{}

	err := json.Unmarshal([]byte(urlJson), &path)
	if err != nil {
		fmt.Println(err)
	}
	return path
}

/*
openJsonStruc2
DESC: opens more complex json data than openJsonStruct
*/
func (k Kids) openJsonStruct(path string) Kids {
	// fmt.Println(path)
	jsonData, err := os.ReadFile(path)
	if err != nil {
		fmt.Println("ERROR: ", err)
	}

	if jsonData == nil {
		fmt.Println(jsonData)
	}

	err2 := json.Unmarshal(jsonData, &k)
	if err2 != nil {
		fmt.Println("ERROR: ", err2)
	}
	return k
}
