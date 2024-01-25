package main

import (
	"fmt"
	"os"
	"strconv"
)

type MyConfig struct {
	AppName string `json:"app_name"`
	Port    int    `json:"port"`
	Debug   bool   `json:"debug"`
}

func loadConfigFromEnv() (*MyConfig, error) {

	myConfig := &MyConfig{}
	os.Setenv("APPNAME", "MyApp") //setting a variable
	os.Setenv("PORT", "8030")
	os.Setenv("DEBUG", "true")

	appName := os.Getenv("APPNAME") //getting the variable
	if appName != "" {
		myConfig.AppName = appName
	}

	portName := os.Getenv("PORT")
	if portName != "" {
		port, err := strconv.Atoi(portName) //Atoi - ascii to int
		if err != nil {
			fmt.Println("Invalid port", err)
		}
		myConfig.Port = port
	}

	debugVal := os.Getenv("DEBUG")
	if debugVal != "" {
		debug, err := strconv.ParseBool(debugVal)
		if err != nil {
			fmt.Println("Invalid debugging", err)
		}
		myConfig.Debug = debug
	}

	return myConfig, nil
}

func main() {

	myConfig, err := loadConfigFromEnv()
	if err != nil {
		fmt.Println("Error loading configuration", err)
		return
	}

	fmt.Printf("AppName: %v\n", myConfig.AppName)
	fmt.Printf("Port: %v\n", myConfig.Port)
	fmt.Printf("Debug: %v\n", myConfig.Debug)

}
