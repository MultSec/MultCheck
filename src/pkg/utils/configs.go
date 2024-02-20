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

	// Check if its one of the built in scanners
	config := make(map[string]string)
	return config
}