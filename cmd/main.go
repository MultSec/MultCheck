package main

import (
	"flag"
	"fmt"
)

func main() {
	// Parse command line arguments
	var configPath string 
	flag.StringVar(&configPath, "c", "", "Path to the configuration file")
	flag.StringVar(&configPath, "config", "", "Path to the configuration file")
	flag.Parse()

	args := flag.Args()
	if len(args) < 1 {
		fmt.Println("[?] Usage: multscan [PAYLOAD_FILE]")
	}
	
	binaryPath := args[0]
}
