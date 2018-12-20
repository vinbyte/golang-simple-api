package config

import (
	"encoding/json"
	"fmt"
	"os"
)

// Config is
type Config struct {
	// Database struct {
	// 	Host     string `json:"host"`
	// 	Password string `json:"password"`
	// } `json:"database"`
	Host string `json:"host"`
	Port string `json:"port"`
}

// ReadConfig is function for read config.json
func ReadConfig() Config {
	var config Config
	configFile, err := os.Open("config/config.json")
	defer configFile.Close()
	if err != nil {
		fmt.Println("Error reading config file : " + err.Error())
	}
	jsonParser := json.NewDecoder(configFile)
	jsonParser.Decode(&config)
	return config
}
