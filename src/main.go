package main

import (
	"flag"
	"fmt"
	"os"
	"github.com/MultSec/MultCheck/pkg/utils"
	"github.com/MultSec/MultCheck/pkg/scan"
)

func main() {
	var scanner string
	flag.StringVar(&scanner, "scanner", "winDef", "Name of the scanner to be used or config file to be used.")
	flag.Parse()

	args := flag.Args()
	if len(args) < 1 {
		fmt.Println("[?] Usage: multscan [OPTIONS] PAYLOAD_FILE_PATH")
		flag.PrintDefaults()
		os.Exit(1)
	}
	binaryPath := args[0]

	// Retrieve scanner configuration
	conf := utils.GetConf(scanner)

	result := scan.CheckMal(binaryPath, conf)
	fmt.Printf("[>] Result: %s\n", result)
}