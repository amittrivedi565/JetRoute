package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type ServiceConfig struct {
	service       string          `json:"service"`
	host          string          `json:"host"`
	port          int             `json:"port"`
	privateRoutes []PrivateRoutes `json:"private-routes"`
}

type PrivateRoutes struct {
	path string `json:"path"`
}

func LoadServiceConfig() {
	data, err := os.ReadFile("./config.json")

	if err != nil {
		log.Fatal(err)
	}

	var config ServiceConfig
	err = json.Unmarshal(data, &config)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("JetRoute [Configuration Loaded]")
}
