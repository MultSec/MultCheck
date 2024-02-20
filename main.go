package main

import (
	"flag"
	"fmt"
	"os"
	"github.com/MultSec/multcheck/utils"
	"github.com/MultSec/multcheck/scan"
)

func main() {
	var scanner string
	flag.StringVar(&scanner, "scanner", "winDef", "Name of the scanner to be used. (Config file can be use instead)")
	flag.Parse()

	args := flag.Args()
	if len(args) < 1 {
		fmt.Println("[?] Usage: multscan [OPTIONS] PAYLOAD_FILE_PATH")
		flag.PrintDefaults()
		os.Exit(1)
	}
	binaryPath := args[0]

	// Retrieve scanner configuration
	conf := utils.getConf(scanner)

	result := scan.checkMal(binaryPath, conf)
	fmt.Printf("[*] Result:\n%s\n", result)
}