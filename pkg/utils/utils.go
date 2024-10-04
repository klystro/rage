package utils

import (
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	EntryPoint string `yaml:"entry_point"`
}

func GetEntryPointFromConfig(configFile string) string {
	data, err := os.ReadFile(configFile)
	if err != nil {
		log.Printf("Could not read config file: %v", err)
		return ""
	}

	var config Config
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		log.Printf("Could not parse config file: %v", err)
		return ""
	}

	return config.EntryPoint
}
