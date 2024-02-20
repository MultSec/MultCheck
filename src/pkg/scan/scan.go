package scan

import (
	"fmt"
)

func CheckMal(binaryPath string, conf map[string]string) string {
	fmt.Printf("Name: %s, Cmd: %s, Out: %s\n", conf["name"], conf["cmd"], conf["out"])
	// Check for Detection

	// Check for Dynamic Detection

	// Check for Static Detection

	// Simulate checking binary for malware
	return fmt.Sprintf("Checking %s for malware", binaryPath)
}