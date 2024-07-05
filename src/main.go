package main

import (
	"flag"
	"fmt"
)

func main() {
	var configPath string
	flag.StringVar(&configPath, "config", "", "Config file to be used.")
	flag.Parse()

	args := flag.Args()
	if len(args) < 1 {
		printLog(logInfo, "Usage: multscan [OPTIONS] PAYLOAD_FILE_PATH")
		return
	}
	binaryPath := args[0]

	// Retrieve configPath configuration
	conf, err := GetConf(configPath)
	if err != nil {
		printLog(logError, fmt.Sprintf("%v", err))
		return
	}

	result, err := CheckMal(binaryPath, conf)
	if err != nil {
		printLog(logError, fmt.Sprintf("%v", err))
		return
	}

	printLog(logSuccess, fmt.Sprintf("Result: %s\n", result))

	return
}