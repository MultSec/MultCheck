package scan

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func scanFile(binaryPath string, conf map[string]string) bool {
	// Replace placeholder with actual file path
	scanCommand := strings.Replace(conf["cmd"], "{{file}}", binaryPath, -1)

	// Execute scanner
	cmdParts := strings.Fields(scanCommand)
	cmd := exec.Command(cmdParts[0], cmdParts[1:]...)

	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("[!] Error executing scanner command:", err)
		os.Exit(1)
	}

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

func checkDinamic(binaryPath string, conf map[string]string) bool {
	// Read the files content
	data, err := os.ReadFile(binaryPath)
	if err != nil {
		fmt.Println("[!] Error reading file:", err)
		os.Exit(1)
	}

	// Split the data into half
	firstHalf := data[:len(data)/2]
	secondHalf := data[len(data)/2:]

	// Check both halves for malware
	if scanSlice(firstHalf, conf) || scanSlice(secondHalf, conf) {
		return true
	} else {
		return false
	}
}

func CheckMal(binaryPath string, conf map[string]string) string {
	// Check for Detection
	if !scanFile(binaryPath, conf) {
		return "Payload not detected."
	}

	// Check for Dynamic Detection
	if checkDinamic(binaryPath, conf) {
		return "Payload detected dynamically."
	}

	// Check for Static Detection

	// Simulate checking binary for malware
	return fmt.Sprintf("Checking %s for malware", binaryPath)
}