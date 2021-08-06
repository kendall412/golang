package main

import (
	"fmt"
	"os"
	// "log"
	// "encoding/json"
)

// need a struct in place to put the imported json into
type Config struct{
	Database struct {
		Host string `json:"host"`
		Port string `json:"port"`
	} `json:"database"`
	Host string `json:"host"`
	Port string `json:"port"`
}

func LoadConfiguration(filename string) (Config, error) {
	var config Config
	configFile, err := os.Open(filename)
	defer configFile.Close()
	if err != nil {
		return config, err
	}
	jsonParser := json.NewDecoder(configFile)
	jsonParser.Decode(&config)
}

func main() {

}