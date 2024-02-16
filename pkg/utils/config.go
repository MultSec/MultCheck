package main

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type ScannerConfig struct {
	Name      string `yaml:"name"`
	ScanCmd   string `yaml:"scanCmd"`
	PosOutput string `yaml:"posOutput"`
	// NegOutput string `yaml:"negOutput"`
}

type Config struct {
	Scanners []ScannerConfig `yaml:"scanners"`
}

func parseConfig() (*Config, error) {
	path := "config.yaml"
	configInit(path)
	var config Config
	data, err := os.ReadFile(path) // Use os.ReadFile instead of ioutil.ReadFile
	if err != nil {
		return nil, err
	}
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return nil, err
	}
	return &config, nil
}

func configInit(path string) {
	configTemplate := `
scanners:
  # - name: "AV name"
  #   scanCmd: "Command for scanning the target file. Use {{file}} as the file name to be scanned. The scanner executable is STRONGLY RECOMMENDED to be in PATH."
  #   posOutput: "A string in output of positive detection but not in negative"
  - name: "ESET"
    scanCmd: "ecls /clean-mode=none /no-quarantine {{file}}"
    posOutput: ">"
  - name: "Windows Defender"
    scanCmd: "MpCmdRun.exe -Scan -ScanType 3 -File {{file}} -DisableRemediation -Trace -Level 0x10"
    posOutput: "Threat information"
  # - name: "Any others"`

	if _, err := os.Stat(path); os.IsNotExist(err) {
		// If file doesn't exist, create it
		err := os.WriteFile(path, []byte(configTemplate), 0644)
		if err != nil {
			fmt.Println("create config.yaml error", err)
		}
	}
}
