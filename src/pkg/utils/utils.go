package utils

import (
	"os"
	"fmt"
	"github.com/BurntSushi/toml"
	"regexp"
)

type config struct {
	name string `toml:"name"`
	cmd  string `toml:"cmd"`
	out  string `toml:"out"`
}

func FileToConf(configPath string) map[string]string {
	data, err := os.ReadFile(configPath)
	if err != nil {
		fmt.Println("[!] Error opening config file:", err)
		os.Exit(1)
	}

	var conf config
	_, err = toml.Decode(string(data), &conf)
	if err != nil {
		fmt.Println("[!] Error parsing TOML:", err)
		os.Exit(1)
	}

	// Cast to map
	confMap := make(map[string]string)
	confMap["name"] = conf.name
	confMap["cmd"] 	= conf.cmd
	confMap["out"] 	= conf.out

	return confMap
}

func GetConf(scanner string) map[string]string {
	// Check if its a config file
	re := regexp.MustCompile(`(?i)\.toml$`)
	if re.MatchString(scanner){
		return FileToConf(scanner)
	}

	// Built in configurations for scanners
	builtInConfs := make(map[string]map[string]string)

	// Windows Defender
	builtInConfs["winDef"] = map[string]string {
		"name": "Windows Defender",
		"cmd":	"MpCmdRun.exe -Scan -ScanType 3 -File {{file}} -DisableRemediation -Trace -Level 0x10",
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