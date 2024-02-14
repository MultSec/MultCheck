package main

import (
	"fmt"
	"os"
)

func main() {
	// Enable virtual terminal processing on stdout to allow ANSI escape codes on Windows 10+.
	err := enableVirtualTerminalProcessing(os.Stdout.Fd())
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to enable virtual terminal processing: %v\n", err)
	}

	printBanner()

	// Check for correct usage
	if len(os.Args) != 3 {
		fmt.Println("ThreatCheck is a tool written in Go to spot malicious content using external scanners.")
		fmt.Println("Usage: ./ThreatCheck -f <target_file>")
		os.Exit(1)
	}

	filePath := os.Args[2]

	// Parse the configuration file
	config, err := parseConfig()
	if err != nil {
		fmt.Printf("Error parsing config file: %v\n", err)
		os.Exit(1)
	}

	// Iterate over all scanners and scan the file
	for _, scanner := range config.Scanners {
		printSeparator()
		fmt.Printf("Scanning with \033[1;32m%s\033[0m\n\n", scanner.Name)
		totalResult, err := overAllScan(scanner, filePath)
		if err != nil || totalResult == false {
			fmt.Printf("Not found with %s\n", scanner.Name)
			continue
		}
		fmt.Printf("Details:\n\n")
		result, err := findMaliciousSection(scanner, filePath)
		if err != nil {
			fmt.Printf("Error scanning file with %s: %v\n", scanner.Name, err)
			continue
		}

		fmt.Println(result)
		// scanFile(scanner, filePath)
	}
	printSeparator()
}
