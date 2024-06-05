package helper

import (
	"encoding/json"
	"log"
	"os"
)

type Config struct {
	SpreadsheetID string `json:"spreadsheet_id"`
	Url           string `json:"url"`
}

type DataSearch struct {
	Search []SearchItem `json:"search"`
}

type SearchItem struct {
	Keywords  []string `json:"keywords"`
	SheetName string   `json:"sheet_name"`
}

func ReadTokenFile() Token {
	jsonFile, err := os.Open("service-account.json")
	if err != nil {
		log.Fatalf("Unable to read token file: %v", err)
	}

	defer jsonFile.Close()

	token := Token{}
	jsonParser := json.NewDecoder(jsonFile)
	if err = jsonParser.Decode(&token); err != nil {
		log.Fatalf("Unable to parse token file: %v", err)
	}

	return token
}

func ReadConfigFile() Config {
	jsonFile, err := os.Open("config.json")
	if err != nil {
		log.Fatalf("Unable to read config file: %v", err)
	}

	defer jsonFile.Close()
	config := Config{}
	jsonParser := json.NewDecoder(jsonFile)
	if err = jsonParser.Decode(&config); err != nil {
		log.Fatalf("Unable to parse config file: %v", err)
	}

	return config
}

func ReadSearchFile() DataSearch {
	jsonFile, err := os.Open("search.json")
	if err != nil {
		log.Fatalf("Unable to read config file: %v", err)
	}

	defer jsonFile.Close()
	search := DataSearch{}
	jsonParser := json.NewDecoder(jsonFile)
	if err = jsonParser.Decode(&search); err != nil {
		log.Fatalf("Unable to parse config file: %v", err)
	}

	return search
}
