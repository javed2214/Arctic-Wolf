package config

import (
	"encoding/json"
	"log"
	"os"
)

type Config struct {
	AppConfig *AppConfig `json:"app"`
}

type AppConfig struct {
	Port        string `json:"port"`
	Environment string `json:"env"`
	LogLevel    string `json:"log_level"`
}

var appConfig *Config

func InitConfig() {
	config, err := os.ReadFile("./config/config.json")
	if err != nil {
		log.Fatalf("Error in reading config file. Error: %s", err.Error())
	}
	var conf *Config
	err = json.Unmarshal([]byte(config), &conf)
	if err != nil {
		log.Fatalf("Error unmarshalling config file. Error: %s", err.Error())
	}
	log.Print("Config initialized successfully!")
	appConfig = conf
}

func GetConfig() *Config {
	if appConfig != nil {
		return appConfig
	}
	return nil
}
