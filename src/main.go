package main

import (
	"flag"
	"fmt"
	"os"
	"github.com/MultSec/MultCheck/pkg/utils"
	"github.com/MultSec/MultCheck/pkg/scan"
)

func main() {
	var configPath string
	flag.StringVar(&configPath, "config", "", "Config file to be used.")
	flag.Parse()

	args := flag.Args()
	if len(args) < 1 {
		fmt.Println("[?] Usage: multscan [OPTIONS] PAYLOAD_FILE_PATH")
		flag.PrintDefaults()
		os.Exit(1)
	}
	binaryPath := args[0]

	// Retrieve configPath configuration
	conf := utils.GetConf(configPath)

	result := scan.CheckMal(binaryPath, conf)
	fmt.Printf("[>] Result: %s\n", result)
}