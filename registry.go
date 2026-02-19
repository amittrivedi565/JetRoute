package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type ServiceConfig struct {
	Service       string          `json:"service"`
	Host          string          `json:"host"`
	Port          int             `json:"port"`
	PrivateRoutes []PrivateRoutes `json:"private-routes"`
}

type PrivateRoutes struct {
	Path string `json:"path"`
}

var Config = make(map[string]*ServiceConfig)

func LoadConfig() error {
	data, err := os.ReadFile("./config.json")

	if err != nil {
		return err
	}

	var config ServiceConfig
	err = json.Unmarshal(data, &config)

	if err != nil {
		return err
	}

	Config[config.Service] = &config
	fmt.Println("[Configuration Loaded]")
	return nil
}
