package utilities

import (
	"encoding/json"
	"os"

	"../models"
)

func GetConfiguration() (models.Configuration, error) {

	config := models.Configuration{}
	file, err := os.Open("./configuration.json")

	if err != nil {
		return nil, err
	}

	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&config)
	if err != nil {
		return nil, err
	}

	return config, nil
}
