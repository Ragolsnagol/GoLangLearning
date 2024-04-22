package config

import (
	"encoding/json"
	"io"
	"os"
)

type Config struct {
	DbConnection string `json:"dbConnection"`
}

var ConfigSettings Config

func ReadConfig() {
	jsonFile, _ := os.Open("config.json")
	defer jsonFile.Close()

	byteValue, _ := io.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &ConfigSettings)
}
