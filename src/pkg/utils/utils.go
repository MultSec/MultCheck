package utils

import (
	"os"
	"fmt"
	"encoding/json"
	"regexp"
)

func FileToConf(configPath string) map[string]string {
	data, err := os.ReadFile(configPath)
	if err != nil {
		fmt.Println("[!] Error opening config file:", err)
		os.Exit(1)
	}

	var conf map[string]string
	err = json.Unmarshal(data, &conf)
	if err != nil {
		fmt.Println("[!] Error parsing config file:", err)
		os.Exit(1)
	}

	return conf
}

func GetConf(scanner string) map[string]string {
	// Check if its a config file
	re := regexp.MustCompile(`(?i)\.json$`)
	if re.MatchString(scanner){
		return FileToConf(scanner)
	}

	// Built in configurations for scanners
	builtInConfs := make(map[string]map[string]string)

	// Windows Defender
	builtInConfs["winDef"] = map[string]string {
		"name": "Windows Defender",
		"cmd":	"MpCmdRun.exe",
		"args":	" -Scan -ScanType 3 -File {{file}} -DisableRemediation -Trace -Level 0x10",
		"out":	"Threat information",
	}

	// Check if its one of the built in scanners
	config, isBuiltIn := builtInConfs[scanner]; 
	
	if !isBuiltIn {
		fmt.Printf("[!] The scanner %s is not a built in scanner.\n", scanner)
		os.Exit(1)
	}

	return config
}