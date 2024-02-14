package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func scanFile(scanner ScannerConfig, filePath string) (bool, error) {
	// Replace placeholder with actual file path
	scanCommand := strings.Replace(scanner.ScanCmd, "{{file}}", filePath, -1)
	// fmt.Println(scanCommand)
	cmdParts := strings.Fields(scanCommand)
	cmd := exec.Command(cmdParts[0], cmdParts[1:]...)

	output, _ := cmd.CombinedOutput()
	// fmt.Println(output)

	// Check if the output contains the positive or negative signature
	if strings.Contains(string(output), scanner.PosOutput) {
		return true, nil
	} else {
		return false, nil
	}
}

func findMaliciousSection(scanner ScannerConfig, filePath string) (string, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return "", err
	}

	// Define the initial bounds for the search
	left, right := 0, len(data)
	lastMaliciousEnd := right // Store the end index of the last detected malicious content

	// fmt.Println("last  mid  left  right")

	for left < right {
		percentage := (len(data) - right + left) * 100 / len(data)
		fmt.Printf("\rChecking done: %d%%", percentage)

		// Define the current section to scan: always start from 0 to include the file head
		mid := (left + right) / 2
		currentData := data[:mid]

		// Write the current data to a temporary file and scan it
		tempFile, err := os.CreateTemp("", "section_search_")
		if err != nil {
			return "", err
		}
		tempFileName := tempFile.Name()
		_, err = tempFile.Write(currentData)
		if err != nil {
			tempFile.Close()
			os.Remove(tempFileName) // Clean up the file
			return "", err
		}
		tempFile.Close()

		malicious, err := scanFile(scanner, tempFileName)
		if err != nil {
			os.Remove(tempFileName) // Clean up the file
			return "", err
		}
		os.Remove(tempFileName) // Clean up the file

		if malicious {
			// The malicious content is in the current range, narrow down the scope
			lastMaliciousEnd = mid
			right = mid
			// fmt.Println("malicious")
		} else {
			// The malicious content is not in the current range, expand the scope
			left = mid + 1
			// fmt.Println("not malicious")
		}
		// fmt.Println(lastMaliciousEnd, " ", mid, " ", left, " ", right)
	}

	fmt.Printf("\r%s\r", strings.Repeat(" ", 50))

	// The malicious content is between 'left' and 'lastMaliciousEnd'
	if lastMaliciousEnd == len(data) {
		return "Malicious content not found.", nil
	} else {
		err := hexDumpAround(filePath, int64(lastMaliciousEnd), int64(128))
		if err != nil {
			fmt.Println("Error:", err)
		}

		return fmt.Sprintf("\nMalicious content detected at: %08x (%d bytes)", int64(lastMaliciousEnd), lastMaliciousEnd), nil
	}
}

func overAllScan(scanner ScannerConfig, filePath string) (bool, error) {
	// Replace placeholder with actual file path
	scanCommand := strings.Replace(scanner.ScanCmd, "{{file}}", filePath, -1)
	// fmt.Println(scanCommand)
	cmdParts := strings.Fields(scanCommand)
	cmd := exec.Command(cmdParts[0], cmdParts[1:]...)

	output, _ := cmd.CombinedOutput()
	fmt.Println("Report for the whole file:\n\n", string(output))

	// Check if the output contains the positive or negative signature
	if strings.Contains(string(output), scanner.PosOutput) {
		return true, nil
	} else {
		return false, nil
	}
}
