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

func (k Kid) openJsonStruct(path string) Kid {
	jsonData, err := os.ReadFile(path)
	if err != nil {
		fmt.Println("ERROR: ", err)
	}

	if jsonData == nil {
		fmt.Println(jsonData)
	}

	err1 := json.Unmarshal(jsonData, &k)
	if err != nil {
		fmt.Println("ERROR: ", err1)
	}
	return (k)
}
