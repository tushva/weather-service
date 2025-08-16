package internal

import (
	"encoding/json"
	"os"
)

type Config struct {
	Port            string `json:"port"`
	Host            string `json:"host"`
	WeatherHost     string `json:"weatherHost"`
	WeatherEndpoint string `json:"weatherEndpoint"`
	APIKey          string `json:"apikey"`
}

func LoadConfig() (*Config, error) {
	var config Config
	file, _ := os.Open("configs/dev.json")
	defer file.Close()
	decoder := json.NewDecoder(file)
	err := decoder.Decode(&config)
	if err != nil {
		mylogger.Logger.Error("Error reading config %s", err.Error(), nil)
		return nil, err
	}
	return &config, nil
}
