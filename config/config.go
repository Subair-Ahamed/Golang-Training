package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

// Config represents the application configuration.
type Config struct {
	AppName string `json:"app_name"`
	Port    int    `json:"port"`
	Debug   bool   `json:"debug"`
	Server  string `json:"server"`
}

func loadConfig(filename string) (*Config, error) { //taking a filename as argument and returning a config struct and error
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	config := &Config{} //creating instance of struct Config.
	err = json.Unmarshal(file, config) //json.Unmarshal is used to unmarshal a JSON-formatted byte slice into a Go data structure
	if err != nil {
		return nil, err
	}

	return config, nil
}

func main() {
	config, err := loadConfig("config.json") //loadConfig function - loads the configuration from the "config.json" file
	if err != nil {
		fmt.Println("Error loading config:", err)
		return
	}

	fmt.Printf("App Name: %s\n", config.AppName)
	fmt.Printf("Port: %d\n", config.Port)
	fmt.Printf("Debug Mode: %v\n", config.Debug)
	fmt.Printf("Server: %v\n", config.Server)
}
