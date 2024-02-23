package scan

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"encoding/hex"
)

func scanFile(binaryPath string, conf map[string]string) bool {
	// Replace placeholder with actual file path
	cmdArgs := strings.Replace(conf["args"], "{{file}}", binaryPath, -1)
	scanArgs := strings.Fields(cmdArgs)
	
	// Execute the scanner command
	cmd := exec.Command(conf["cmd"], scanArgs...)

	// Get the output of the command
	output, _ := cmd.CombinedOutput()

	// Check if the output contains the positive detection
	if strings.Contains(string(output), conf["out"]) {
		return true
	} else {
		return false
	}
}

func scanSlice(fileData []byte, conf map[string]string) bool {
	// Create a temp file to scan
	tempFile, err := os.CreateTemp("", "slice_scan_")
	if err != nil {
		fmt.Println("[!] Error creating temp file:", err)
		os.Exit(1)
	}

	// Defer cleanup
	defer os.Remove(tempFile.Name())
	defer tempFile.Close()

	_, err = tempFile.Write(fileData)
	if err != nil {
		fmt.Println("[!] Error writing to temp file:", err)
		os.Exit(1)
	}

	// Scan the file slice
	return scanFile(tempFile.Name(), conf)
}

func checkStatic(binaryPath string, conf map[string]string) string {
	// Read the files content
	data, err := os.ReadFile(binaryPath)
	if err != nil {
		fmt.Println("[!] Error reading file:", err)
		os.Exit(1)
	}

	// Set the initial values
	lastGood, mid, upperBound := 0, len(data)/2, len(data)
	threatFound := false

	// Binary search for the malicious content
	for upperBound-lastGood > 1 {
		// Check the slice for malware
		if scanSlice(data[0:mid], conf) {
			threatFound = true
			upperBound = mid
		} else {
			lastGood = mid
		}

		mid = lastGood + (upperBound-lastGood)/2
	}

	// Return the result
	if threatFound {

		// Get the start and end of the malicious content
		start := lastGood - 32
		if start < 0 {
			start = 0
		}

		// Get the start and end of the malicious content
		end := mid + 32
		if end > len(data) {
			end = len(data)
		}

		return fmt.Sprintf("Malicious content found at offset: %08x \n%s\n", lastGood, hex.Dump(data[start:end]))
	}

	return ""
}

func CheckMal(binaryPath string, conf map[string]string) string {
	// Check for Detection
	if !scanFile(binaryPath, conf) {
		return "Payload not detected."
	}

	if static := checkStatic(binaryPath, conf); static != "" {
		return static
	}
	
	return "Payload detected dynamically."
}