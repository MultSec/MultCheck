package scan

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func hexDump(filename string, offset int64) string {
	const bytesPerLine = 16
	result := ""

	data, err := os.ReadFile(filename)
	if err != nil {
		return fmt.Sprintf("[!] Error reading file: %s", err)
		os.Exit(1)
	}

	// Iterate over the data
	for i := offset - bytesPerLine; i < offset+bytesPerLine; i += bytesPerLine {
		if i < 0 {
			continue
		}

		// Print the offset
		result += fmt.Sprintf("%08x |", i)

		// Print the bytes
		for j := i; j < i+bytesPerLine; j++ {
			if j < 0 || j >= int64(len(data)) {
				result += "	"
			} else {
				result += fmt.Sprintf("%02x ", data[j])
			}
		}

		// Print the ASCII representation
		result += "|"
		for j := i; j < i+bytesPerLine; j++ {
			if j < 0 || j >= int64(len(data)) {
				result += " "
			} else if data[j] >= 32 && data[j] <= 126 {
				result += string(data[j])
			} else {
				result += "."
			}
		}

		// Print the newline
		result += "\n"
	}

	return result
}

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
		return false
	} else {
		return true
	}
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
		// Get the current slice
		currentData := data[:mid]

		// Check the slice for malware
		if scanSlice(currentData, conf) {
			threatFound = true
			upperBound = mid
		} else {
			lastGood = mid
		}

		mid = lastGood + (upperBound-lastGood)/2
	}

	// Return the result
	if threatFound {
		return fmt.Sprintf("Malicious content found at offset: %08x (%d bytes)\n%s", lastGood, len(data)-lastGood, hexDump(binaryPath, int64(lastGood)))
	}

	return "Payload not detected."
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
	return checkStatic(binaryPath, conf)
}