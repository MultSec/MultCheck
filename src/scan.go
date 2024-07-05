package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"encoding/hex"
	"path/filepath"
)

func scanFile(binaryPath string, conf map[string]string) (bool, error) {
	// Get absolute path
	absPath, err := filepath.Abs(binaryPath)
    if err != nil {
		return false, fmt.Errorf("failed to get absolute path with error: %v", err)
    }

	// Replace placeholder with actual file path
	cmdArgs := strings.Replace(conf["args"], "{{file}}", absPath, -1)
	scanArgs := strings.Fields(cmdArgs)
	
	// Execute the scanner command
	cmd := exec.Command(conf["cmd"], scanArgs...)

	// Get the output of the command
	output, err := cmd.CombinedOutput()
	if err != nil {
		return false, fmt.Errorf("failed to run command (\"%s %s\") with error: %v", conf["cmd"], cmdArgs, err)
	}

	// Check if the output contains the positive detection
	if strings.Contains(string(output), conf["out"]) {
		return true, nil
	} else {
		return false, nil
	}
}

func scanSlice(fileData []byte, conf map[string]string) (bool, error) {
	// Create a temp file to scan
	tempFile, err := os.CreateTemp("", "slice_scan_")
	if err != nil {
		return false, fmt.Errorf("failed to create temp file: %v", err)
	}

	// Defer cleanup
	defer os.Remove(tempFile.Name())
	defer tempFile.Close()

	_, err = tempFile.Write(fileData)
	if err != nil {
		return false, fmt.Errorf("failed to write to temp file: %v", err)
	}

	// Scan the file slice
	scanResult, err := scanFile(tempFile.Name(), conf)
	if err != nil {
		return false, fmt.Errorf("failed to scan temp file: %v", err)
	}

	return scanResult, nil
}

func checkStatic(binaryPath string, conf map[string]string) (string, error) {
	// Read the files content
	data, err := os.ReadFile(binaryPath)
	if err != nil {
		return "", fmt.Errorf("failed to read file: %v", err)
	}

	// Set the initial values
	lastGood, mid, upperBound := 0, len(data)/2, len(data)
	threatFound := false

	// Binary search for the malicious content
	for upperBound-lastGood > 1 {
		// Check the slice for malware
		scanResult , err := scanSlice(data[0:mid], conf)
		if err != nil {
			return "", err
		}

		if scanResult {
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

		return fmt.Sprintf("Malicious content found at offset: %08x \n%s\n", lastGood, hex.Dump(data[start:end])), nil
	}

	return "", nil
}

func CheckMal(binaryPath string, conf map[string]string) (string, error) {
	// Check for Detection
	scanResult, err := scanFile(binaryPath, conf)
	if err != nil {
		return "", err
	}

	if !scanResult {
		return "Not malicious", nil
	}

	static, err := checkStatic(binaryPath, conf)
	if err != nil {
		return "", err
	}
	
	if static != "" {
		return static, nil
	}
	
	return "Whole file is malicious", nil
}