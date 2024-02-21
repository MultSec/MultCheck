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

func CheckMal(binaryPath string, conf map[string]string) string {
	// Check for Detection
	if !scanFile(binaryPath, conf) {
		return "Payload not detected."
	}

	// Check for Dynamic Detection

	// Check for Static Detection

	// Simulate checking binary for malware
	return fmt.Sprintf("Checking %s for malware", binaryPath)
}