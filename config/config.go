package config

import (
	"encoding/json"
	"fmt"
	"os"
)

// Config is
type Config struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	Database struct {
		Dbname   string `json:"dbname"`
		User     string `json:"user"`
		Password string `json:"password"`
		Host     string `json:"host"`
		Port     string `json:"port"`
	} `json:"database"`
	PeopleNum int `json:"people_num"`
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
