package gcp

import (
	"encoding/json"
	"log"
	"os"
)

func getConfig(filepath string) (Config, error) {
	// open the json file
	file, err := os.Open(filepath)
	if err != nil {
		return Config{}, err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Printf("Failed to close file: %v", err)
			return
		}
	}(file)
	// decode the json file
	decoder := json.NewDecoder(file)
	config := Config{}
	err = decoder.Decode(&config)
	if err != nil {
		return Config{}, err
	}
	return config, nil
}
