package utils

import (
	"os"
	"fmt"
	"encoding/json"
)

func GetConf(configPath string) map[string]string {
	data, err := os.ReadFile(configPath)
	if err != nil {
		fmt.Println("[!] Error opening config file:", err)
		os.Exit(1)
	}

	var conf map[string]string
	err = json.Unmarshal(data, &conf)
	if err != nil {
		fmt.Println("[!] Error parsing config file:", err)
		os.Exit(1)
	}

	return conf
}